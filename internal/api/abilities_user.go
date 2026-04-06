package api

import (
	"errors"
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type userAbility interface {
	canUseAbility(*Config, string) bool
	getBattleInteractions() []BattleInteraction
	error
}

type unitRepl struct {
	unit     seeding.PlayerUnit
	bombWpn  bool
	replVals biReplacement
}

type biReplacement struct {
	Range          *int32
	ShatterRate    *int32
	Accuracy       *Accuracy
	DamageConstant *int32
}

func applyUser[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], ability userAbility, queryName string) ([]BattleInteraction, error) {
	repl, err := getUnitRepl(cfg, r, i, queryName)
	if errors.Is(err, errEmptyQuery) {
		return ability.getBattleInteractions(), nil
	}
	if err != nil {
		return nil, err
	}

	err = verifyAbilityUsage(cfg, ability, repl, queryName)
	if err != nil {
		return nil, err
	}

	battleInteractions := applyBiReplacement(ability.getBattleInteractions(), repl.replVals)

	return battleInteractions, nil
}

func getUnitRepl[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], queryName string) (unitRepl, error) {
	queryParamUser := i.queryLookup[queryName]
	queryParamBomb := i.queryLookup["bomb_wpn"]

	unitID, err := parseNameIdQuery(r, queryParamUser, cfg.e.playerUnits.resourceType, cfg.e.playerUnits.objLookup)
	if err != nil {
		return unitRepl{}, err
	}
	unit, _ := seeding.GetResourceByID(unitID, cfg.l.PlayerUnitsID)

	bombWpn, err := parseBooleanQuery(r, queryParamBomb)
	if errIsNotEmptyQuery(err) {
		return unitRepl{}, err
	}

	repl := unitRepl{
		unit:     unit,
		bombWpn:  bombWpn,
		replVals: biReplacement{},
	}

	switch repl.unit.Type {
	case "character":
		repl, err = populateReplCharacter(cfg, repl, queryParamUser)
		if err != nil {
			return unitRepl{}, err
		}

	case "aeon":
		repl, err = populateReplAeon(cfg, repl, queryParamUser)
		if err != nil {
			return unitRepl{}, err
		}
	}

	return repl, nil
}

func populateReplCharacter(cfg *Config, repl unitRepl, queryParamUser QueryParam) (unitRepl, error) {
	id, err := checkQueryNameID(repl.unit.Name, string(repl.unit.Type), queryParamUser, cfg.l.Characters)
	if err != nil {
		return unitRepl{}, err
	}
	character, _ := seeding.GetResourceByID(id, cfg.l.CharactersID)

	repl.replVals.Range = &character.PhysAtkRange

	if repl.bombWpn {
		repl.replVals.DamageConstant = h.GetInt32Ptr(18)
	}

	return repl, nil
}

func populateReplAeon(cfg *Config, repl unitRepl, queryParamUser QueryParam) (unitRepl, error) {
	id, err := checkQueryNameID(repl.unit.Name, string(repl.unit.Type), queryParamUser, cfg.l.Aeons)
	if err != nil {
		return unitRepl{}, err
	}
	aeon, _ := seeding.GetResourceByID(id, cfg.l.AeonsID)

	repl.replVals.Range = aeon.PhysAtkRange
	repl.replVals.ShatterRate = aeon.PhysAtkShatterRate
	repl.replVals.Accuracy = convertObjPtr(cfg, aeon.PhysAtkAccuracy, convertAccuracy)
	repl.replVals.DamageConstant = aeon.PhysAtkDmgConstant

	return repl, nil
}

func verifyAbilityUsage(cfg *Config, ability userAbility, repl unitRepl, queryName string) error {
	if !ability.canUseAbility(cfg, repl.unit.Name) {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid input for parameter '%s': %s '%s' can't learn %s", queryName, cfg.e.playerUnits.resourceType, repl.unit.Name, ability), nil)
	}

	return nil
}

func applyBiReplacement(battleInteractions []BattleInteraction, replVals biReplacement) []BattleInteraction {
	for i, battleInteraction := range battleInteractions {
		if !battleInteraction.BasedOnUserAttack {
			continue
		}

		if replVals.Range != nil {
			battleInteraction.Range = replVals.Range
		}

		if replVals.ShatterRate != nil {
			battleInteraction.ShatterRate = *replVals.ShatterRate
		}

		if replVals.Accuracy != nil {
			battleInteraction.Accuracy = *replVals.Accuracy
		}

		if replVals.DamageConstant != nil {
			battleInteraction.Damage.DamageCalc[0].DamageConstant = *replVals.DamageConstant
		}

		battleInteractions[i] = battleInteraction
	}

	return battleInteractions
}
