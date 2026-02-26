package api

import (
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func applyPlayerAbilityUser(cfg *Config, r *http.Request, ability PlayerAbility, queryName string) (PlayerAbility, error) {
	queryParam := cfg.q.playerAbilities[queryName]

	resType, idStr, err := parseResTypeQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return ability, nil
	}
	if err != nil {
		return PlayerAbility{}, err
	}

	var repl biReplacement
	var unitName string

	switch resType {
	case "character":
		parseRes, err := parseID(idStr, resType, len(cfg.l.Characters))
		if err != nil {
			return PlayerAbility{}, err
		}
		character, _ := seeding.GetResourceByID(parseRes.ID, cfg.l.CharactersID)
		unitName = character.Name

		repl.Range = &character.PhysAtkRange

	case "aeon":
		parseRes, err := parseID(idStr, resType, len(cfg.l.Aeons))
		if err != nil {
			return PlayerAbility{}, err
		}
		aeon, _ := seeding.GetResourceByID(parseRes.ID, cfg.l.AeonsID)
		unitName = aeon.Name

		repl.Range = aeon.PhysAtkRange
		repl.ShatterRate = aeon.PhysAtkShatterRate
		repl.Accuracy = convertObjPtr(cfg, aeon.PhysAtkAccuracy, convertAccuracy)
		repl.DamageConstant = aeon.PhysAtkDmgConstant
	}

	err = checkUserInCharClass(cfg, ability, unitName, resType, queryName)
	if err != nil {
		return PlayerAbility{}, err
	}

	ability.BattleInteractions = replaceBattleInteractions(ability.BattleInteractions, repl)
	return ability, nil
}

func replaceBattleInteractions(battleInteractions []BattleInteraction, repl biReplacement) []BattleInteraction {
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

type biReplacement struct {
	Range			*int32
	ShatterRate		*int32
	Accuracy		*Accuracy
	DamageConstant	*int32
}


func checkUserInCharClass(cfg *Config, ability PlayerAbility, unitName, resType, queryName string) error {
	for _, class := range ability.LearnedBy {
		classLookup, _ := seeding.GetResourceByID(class.ID, cfg.l.CharClassesID)

		if slices.Contains(classLookup.Members, unitName) {
			return nil
		}
	}

	return newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid input for parameter '%s': %s '%s' can't learn player ability '%s'", queryName, resType, unitName, nameToString(ability.Name, ability.Version, nil)), nil)
}