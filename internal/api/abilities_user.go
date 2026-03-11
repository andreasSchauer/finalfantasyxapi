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
	resType 	string
	unitName	string
	bombWpn		bool
	replVals	biReplacement
}

type biReplacement struct {
	Range          *int32
	ShatterRate    *int32
	Accuracy       *Accuracy
	DamageConstant *int32
}


func applyUser(cfg *Config, r *http.Request, ability userAbility, queryName string, queryLookup map[string]QueryType) ([]BattleInteraction, error) {
	repl, err := getUnitRepl(cfg, r, queryName, queryLookup)
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


func getUnitRepl(cfg *Config, r *http.Request, queryName string, queryLookup map[string]QueryType) (unitRepl, error) {
	queryParamUser := queryLookup[queryName]
	queryParamBomb := queryLookup["bomb_wpn"]

	resType, unitStr, err := parseResTypeQuery(r, queryParamUser)
	if err != nil {
		return unitRepl{}, err
	}
	
	bombWpn, err := parseBooleanQuery(r, queryParamBomb)
	if err != nil && !errors.Is(err, errEmptyQuery) {
		return unitRepl{}, err
	}

	repl := unitRepl{
		resType: 	resType,
		bombWpn: 	bombWpn,
		replVals: 	biReplacement{},
	}

	switch repl.resType {
	case "character":
		repl, err = populateReplCharacter(cfg, unitStr, repl, queryParamUser)
		if err != nil {
			return unitRepl{}, err
		}

	case "aeon":
		repl, err = populateReplAeon(cfg, unitStr, repl, queryParamUser)
		if err != nil {
			return unitRepl{}, err
		}
	}

	return repl, nil
}


func populateReplCharacter(cfg *Config, unitStr string, repl unitRepl, queryParamUser QueryType) (unitRepl, error) {
	id, err := parseQueryNamedVal(unitStr, repl.resType, queryParamUser, cfg.l.Characters)
	if err != nil {
		return unitRepl{}, err
	}
	character, _ := seeding.GetResourceByID(id, cfg.l.CharactersID)
	repl.unitName = character.Name

	repl.replVals.Range = &character.PhysAtkRange

	if repl.bombWpn {
		repl.replVals.DamageConstant = h.GetInt32Ptr(18)
	}

	return repl, nil
}


func populateReplAeon(cfg *Config, unitStr string, repl unitRepl, queryParamUser QueryType) (unitRepl, error) {
	id, err := parseQueryNamedVal(unitStr, repl.resType, queryParamUser, cfg.l.Aeons)
	if err != nil {
		return unitRepl{}, err
	}
	aeon, _ := seeding.GetResourceByID(id, cfg.l.AeonsID)
	repl.unitName = aeon.Name

	repl.replVals.Range = aeon.PhysAtkRange
	repl.replVals.ShatterRate = aeon.PhysAtkShatterRate
	repl.replVals.Accuracy = convertObjPtr(cfg, aeon.PhysAtkAccuracy, convertAccuracy)
	repl.replVals.DamageConstant = aeon.PhysAtkDmgConstant

	return repl, nil
}



func verifyAbilityUsage(cfg *Config, ability userAbility, repl unitRepl, queryName string) error {
	if !ability.canUseAbility(cfg, repl.unitName) {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid input for parameter '%s': %s '%s' can't learn %s", queryName, repl.resType, repl.unitName, ability), nil)
	}

	return nil
}


func applyBiReplacement(battleInteractions []BattleInteraction, replVals biReplacement) []BattleInteraction {
	for i, battleInteraction := range battleInteractions {
		if !battleInteraction.BasedOnPhysAttack {
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