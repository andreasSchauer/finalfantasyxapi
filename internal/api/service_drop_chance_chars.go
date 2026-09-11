package api

import (
	"slices"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type CharacterChances struct {
	EligibleChars    int32    `json:"eligible_chars"`
	AnyCharFinBlow   float64  `json:"any_char_fin_blow"`
	AnyCharNoFinBlow float64  `json:"any_char_no_fin_blow"`
	CharFinBlow      *float64 `json:"char_fin_blow,omitempty"`
	CharNoFinBlow    *float64 `json:"char_no_fin_blow,omitempty"`
}

func (c CharacterChances) Percent() CharacterChances {
	c.AnyCharFinBlow = h.DecimalToPercent(c.AnyCharFinBlow)
	c.AnyCharNoFinBlow = h.DecimalToPercent(c.AnyCharNoFinBlow)
	c.CharFinBlow = h.DecimalPtrToPercent(c.CharFinBlow)
	c.CharNoFinBlow = h.DecimalPtrToPercent(c.CharNoFinBlow)

	return c
}

func calcCharRandomChances(cfg *Config, params DropChanceParams, wantedAbilities []seeding.AutoAbility, monAbilities []seeding.EquipmentDrop) CharacterChances {
	var eligibleChars int32
	baseRandomChance := 0.75 / h.FloatLen(params.PartyMembers)
	var chances CharacterChances

	if params.Character != nil {
		char, _ := seeding.GetResourceByID(*params.Character, cfg.l.CharactersID)

		if charLearnsReqAbilities(char, wantedAbilities, monAbilities) {
			finBlowChance := baseRandomChance + 0.25
			chances.CharNoFinBlow = &baseRandomChance
			chances.CharFinBlow = &finBlowChance
		}
	}

	for _, charID := range params.PartyMembers {
		char, _ := seeding.GetResourceByID(charID, cfg.l.CharactersID)

		if charLearnsReqAbilities(char, wantedAbilities, monAbilities) {
			eligibleChars++
		}
	}

	chances.AnyCharNoFinBlow = float64(eligibleChars) * baseRandomChance
	chances.EligibleChars = eligibleChars

	if eligibleChars > 0 {
		chances.AnyCharFinBlow = chances.AnyCharNoFinBlow + 0.25
	}

	return chances
}

func charLearnsReqAbilities(char seeding.Character, wantedAbilities []seeding.AutoAbility, monAbilities []seeding.EquipmentDrop) bool {
	for _, ability := range wantedAbilities {
		for _, monAbility := range monAbilities {
			if ability.Name == monAbility.Ability && !charGetsAbility(monAbility, &char) {
				return false
			}
		}
	}

	return true
}

func charGetsAbility(ability seeding.EquipmentDrop, charPtr *seeding.Character) bool {
	if charPtr == nil {
		return true
	}

	chars := ability.Characters

	if len(chars) == 0 {
		return true
	}

	return slices.Contains(chars, charPtr.Name)
}

func getCharPtr(cfg *Config, params DropChanceParams) *seeding.Character {
	if params.Character == nil {
		return nil
	}

	char, _ := seeding.GetResourceByID(*params.Character, cfg.l.CharactersID)
	return &char
}
