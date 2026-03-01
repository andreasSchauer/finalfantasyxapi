package api

import (
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)


type biReplacement struct {
	Range			*int32
	ShatterRate		*int32
	Accuracy		*Accuracy
	DamageConstant	*int32
}

func applyPlayerAbilityUser(cfg *Config, r *http.Request, ability PlayerAbility, queryName string) (PlayerAbility, error) {
	queryParam := cfg.q.playerAbilities[queryName]

	resType, unitStr, err := parseResTypeQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return ability, nil
	}
	if err != nil {
		return PlayerAbility{}, err
	}
	
	unitName, repl, err := findUnit(cfg, unitStr, resType, queryParam)
	if err != nil {
		return PlayerAbility{}, err
	}

	err = canUsePlayerAbility(cfg, ability, unitName, resType, queryName)
	if err != nil {
		return PlayerAbility{}, err
	}

	ability.BattleInteractions = replaceBattleInteractionVals(ability.BattleInteractions, repl)

	return ability, nil
}

func applyOtherAbilityUser(cfg *Config, r *http.Request, ability OtherAbility, queryName string) (OtherAbility, error) {
	queryParam := cfg.q.otherAbilities[queryName]

	resType, unitStr, err := parseResTypeQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return ability, nil
	}
	if err != nil {
		return OtherAbility{}, err
	}
	
	unitName, repl, err := findUnit(cfg, unitStr, resType, queryParam)
	if err != nil {
		return OtherAbility{}, err
	}

	err = canUseOtherAbility(cfg, ability, unitName, resType, queryName)
	if err != nil {
		return OtherAbility{}, err
	}

	ability.BattleInteractions = replaceBattleInteractionVals(ability.BattleInteractions, repl)

	return ability, nil
}

func applyTriggerCommandUser(cfg *Config, r *http.Request, ability TriggerCommand, queryName string) (TriggerCommand, error) {
	queryParam := cfg.q.triggerCommands[queryName]

	resType, unitStr, err := parseResTypeQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return ability, nil
	}
	if err != nil {
		return TriggerCommand{}, err
	}
	
	unitName, repl, err := findUnit(cfg, unitStr, resType, queryParam)
	if err != nil {
		return TriggerCommand{}, err
	}

	err = canUseTriggerCommand(cfg, ability, unitName, resType, queryName)
	if err != nil {
		return TriggerCommand{}, err
	}

	ability.BattleInteractions = replaceBattleInteractionVals(ability.BattleInteractions, repl)

	return ability, nil
}


func findUnit(cfg *Config, unitStr, resType string, queryParam QueryType) (string, biReplacement, error) {
	var repl biReplacement
	var unitName string

	switch resType {
	case "character":
		id, err := parseQueryNamedVal(unitStr, resType, queryParam, cfg.l.Characters)
		if err != nil {
			return "", biReplacement{}, err
		}
		character, _ := seeding.GetResourceByID(id, cfg.l.CharactersID)
		unitName = character.Name

		repl.Range = &character.PhysAtkRange

	case "aeon":
		id, err := parseQueryNamedVal(unitStr, resType, queryParam, cfg.l.Aeons)
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


func replaceBattleInteractionVals(battleInteractions []BattleInteraction, repl biReplacement) []BattleInteraction {
	for i, battleInteraction := range battleInteractions {
		if !battleInteraction.BasedOnPhysAttack {
			continue
		}

		if repl.Range != nil {
			battleInteraction.Range = repl.Range
		}

		if repl.ShatterRate != nil {
			battleInteraction.ShatterRate = repl.ShatterRate
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


func canUsePlayerAbility(cfg *Config, ability PlayerAbility, unitName, resType, queryName string) error {
	for _, class := range ability.LearnedBy {
		classLookup, _ := seeding.GetResourceByID(class.ID, cfg.l.CharClassesID)

		if slices.Contains(classLookup.Members, unitName) {
			return nil
		}
	}

	return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid input for parameter '%s': %s '%s' can't learn player ability '%s'", queryName, resType, unitName, nameToString(ability.Name, ability.Version, nil)), nil)
}


func canUseOtherAbility(cfg *Config, ability OtherAbility, unitName, resType, queryName string) error {
	for _, class := range ability.LearnedBy {
		classLookup, _ := seeding.GetResourceByID(class.ID, cfg.l.CharClassesID)

		if slices.Contains(classLookup.Members, unitName) {
			return nil
		}
	}

	return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid input for parameter '%s': %s '%s' can't learn other ability '%s'", queryName, resType, unitName, nameToString(ability.Name, ability.Version, nil)), nil)
}

func canUseTriggerCommand(cfg *Config, ability TriggerCommand, unitName, resType, queryName string) error {
	for _, class := range ability.UsedBy {
		classLookup, _ := seeding.GetResourceByID(class.ID, cfg.l.CharClassesID)

		if slices.Contains(classLookup.Members, unitName) {
			return nil
		}
	}

	return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid input for parameter '%s': %s '%s' can't learn trigger command '%s'", queryName, resType, unitName, nameToString(ability.Name, ability.Version, nil)), nil)
}