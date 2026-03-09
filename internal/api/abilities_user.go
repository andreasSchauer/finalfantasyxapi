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

type biReplacement struct {
	Range          *int32
	ShatterRate    *int32
	Accuracy       *Accuracy
	DamageConstant *int32
}


func applyUser(cfg *Config, r *http.Request, ability userAbility, queryName string, queryLookup map[string]QueryType) ([]BattleInteraction, error) {
	queryParam := queryLookup[queryName]
	queryParamBomb := queryLookup["bomb_wpn"]

	resType, unitStr, err := parseResTypeQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return ability.getBattleInteractions(), nil
	}
	if err != nil {
		return nil, err
	}

	unitName, repl, err := findUnit(cfg, r, unitStr, resType, queryParam, queryParamBomb)
	if err != nil {
		return nil, err
	}

	err = verifyAbilityUsage(cfg, ability, unitName, resType, queryName)
	if err != nil {
		return nil, err
	}

	battleInteractions := applyBiReplacement(ability.getBattleInteractions(), repl)

	return battleInteractions, nil
}

func findUnit(cfg *Config, r *http.Request, unitStr, resType string, queryParamUser, queryParamBomb QueryType) (string, biReplacement, error) {
	var repl biReplacement
	var unitName string

	bombWpn, err := parseBooleanQuery(r, queryParamBomb)
	if err != nil && !errors.Is(err, errEmptyQuery) {
		return "", biReplacement{}, err
	}

	switch resType {
	case "character":
		id, err := parseQueryNamedVal(unitStr, resType, queryParamUser, cfg.l.Characters)
		if err != nil {
			return "", biReplacement{}, err
		}
		character, _ := seeding.GetResourceByID(id, cfg.l.CharactersID)
		unitName = character.Name

		repl.Range = &character.PhysAtkRange

		if bombWpn {
			repl.DamageConstant = h.GetInt32Ptr(18)
		}

	case "aeon":
		id, err := parseQueryNamedVal(unitStr, resType, queryParamUser, cfg.l.Aeons)
		if err != nil {
			return "", biReplacement{}, err
		}
		aeon, _ := seeding.GetResourceByID(id, cfg.l.AeonsID)
		unitName = aeon.Name

		repl.Range = aeon.PhysAtkRange
		repl.ShatterRate = aeon.PhysAtkShatterRate
		repl.Accuracy = convertObjPtr(cfg, aeon.PhysAtkAccuracy, convertAccuracy)
		repl.DamageConstant = aeon.PhysAtkDmgConstant
	}

	return unitName, repl, nil
}



func verifyAbilityUsage(cfg *Config, ability userAbility, unitName, resType, queryName string) error {
	if !ability.canUseAbility(cfg, unitName) {
		return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid input for parameter '%s': %s '%s' can't learn %s", queryName, resType, unitName, ability), nil)
	}

	return nil
}


func applyBiReplacement(battleInteractions []BattleInteraction, repl biReplacement) []BattleInteraction {
	for i, battleInteraction := range battleInteractions {
		if !battleInteraction.BasedOnPhysAttack {
			continue
		}

		if repl.Range != nil {
			battleInteraction.Range = repl.Range
		}

		if repl.ShatterRate != nil {
			battleInteraction.ShatterRate = *repl.ShatterRate
		}

		if repl.Accuracy != nil {
			battleInteraction.Accuracy = *repl.Accuracy
		}

		if repl.DamageConstant != nil {
			battleInteraction.Damage.DamageCalc[0].DamageConstant = *repl.DamageConstant
		}

		battleInteractions[i] = battleInteraction
	}

	return battleInteractions
}