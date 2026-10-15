package api

import (
	"slices"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

// put the char specific field in their own struct and make that a pointer
// that's better than having three pointer fields, one of them a boolean
type CharacterChances struct {
	EligibleChars    int32    				`json:"eligible_chars"`
	AnyCharFinBlow   float64  				`json:"any_char_fin_blow"`
	AnyCharNoFinBlow float64  				`json:"any_char_no_fin_blow"`
	SpecificChar	 *SpecificCharChances	`json:"specific_character,omitempty"`
}

func (c CharacterChances) Percent() CharacterChances {
	c.AnyCharFinBlow = h.DecimalToPercent(c.AnyCharFinBlow)
	c.AnyCharNoFinBlow = h.DecimalToPercent(c.AnyCharNoFinBlow)

	if c.SpecificChar != nil {
		newObj := c.SpecificChar.Percent()
		c.SpecificChar = &newObj
	}
	
	return c
}

func (c CharacterChances) Round(n int32) CharacterChances {
	c.AnyCharFinBlow = h.FloatRound(c.AnyCharFinBlow, n)
	c.AnyCharNoFinBlow = h.FloatRound(c.AnyCharNoFinBlow, n)

	if c.SpecificChar != nil {
		newObj := c.SpecificChar.Round(n)
		c.SpecificChar = &newObj
	}

	return c
}

type SpecificCharChances struct {
	CharIsEligible	 bool	  `json:"char_is_eligible"`
	CharFinBlow      float64 `json:"char_fin_blow"`
	CharNoFinBlow    float64 `json:"char_no_fin_blow"`
}

func (c SpecificCharChances) Percent() SpecificCharChances {
	c.CharFinBlow = h.DecimalToPercent(c.CharFinBlow)
	c.CharNoFinBlow = h.DecimalToPercent(c.CharNoFinBlow)

	return c
}

func (c SpecificCharChances) Round(n int32) SpecificCharChances {
	c.CharFinBlow = h.FloatRound(c.CharFinBlow, n)
	c.CharNoFinBlow = h.FloatRound(c.CharNoFinBlow, n)

	return c
}


func calcCharChances(cfg *Config, params DropChanceParams, wantedAbilities []seeding.AutoAbility, monAbilities []seeding.EquipmentDrop) CharacterChances {
	var eligibleChars int32
	baseRandomChance := 0.75 / h.FloatLen(params.PartyMembers)
	var chances CharacterChances

	if params.CharacterID != nil {
		finBlowChance := baseRandomChance + 0.25
		char, _ := seeding.GetResourceByID(*params.CharacterID, cfg.l.CharactersID)

		specificChar := SpecificCharChances{
			CharNoFinBlow: baseRandomChance,
			CharFinBlow: finBlowChance,
			CharIsEligible: charIsEligible(char, wantedAbilities, monAbilities, params.LenientAbilities),
		}

		chances.SpecificChar = &specificChar
	}

	for _, charID := range params.PartyMembers {
		char, _ := seeding.GetResourceByID(charID, cfg.l.CharactersID)

		if charIsEligible(char, wantedAbilities, monAbilities, params.LenientAbilities) {
			eligibleChars++
		}
	}

	chances.AnyCharNoFinBlow = float64(eligibleChars) * baseRandomChance
	chances.EligibleChars = eligibleChars

	if eligibleChars > 0 {
		chances.AnyCharFinBlow = chances.AnyCharNoFinBlow + 0.25
	}

	// if we don't care about who gets the equipment and all characters in the party are eligible
	// then it doesn't matter who scores the finishing blow
	// therefore, it's always a guaranteed character match
	if eligibleChars == h.Len32(params.PartyMembers) {
		chances.AnyCharFinBlow = 1
		chances.AnyCharNoFinBlow = 1
	}

	return chances
}

func charIsEligible(char seeding.Character, wantedAbilities []seeding.AutoAbility, monAbilities []seeding.EquipmentDrop, lenientAbilities bool) bool {
	if charGetsForcedButNotWanted(char, wantedAbilities, monAbilities, lenientAbilities) {
		return false
	}

	for _, ability := range wantedAbilities {
		for _, monAbility := range monAbilities {
			if ability.Name == monAbility.Ability && !charGetsAbility(monAbility, &char) {
				return false
			}
		}
	}

	return true
}

func charGetsForcedButNotWanted(char seeding.Character, wantedAbilities []seeding.AutoAbility, monAbilities []seeding.EquipmentDrop, lenientAbilities bool) bool {
	return 	forcedSlotExists(monAbilities) &&
			!forcedSlotWanted(wantedAbilities, monAbilities) && 
			charGetsAbility(monAbilities[0], &char) &&
			!lenientAbilities
}

func forcedSlotExists(monAbilities []seeding.EquipmentDrop) bool {
	return monAbilities[0].IsForced
}

func forcedSlotWanted(wantedAbilities []seeding.AutoAbility, monAbilities []seeding.EquipmentDrop) bool {
	if !forcedSlotExists(monAbilities) {
		return false
	}

	for _, ability := range wantedAbilities {
		if ability.Name == monAbilities[0].Ability {
			return true
		}
	}

	return false
}

func charGetsAbility(ability seeding.EquipmentDrop, charPtr *seeding.Character) bool {
	chars := ability.Characters

	if charPtr == nil || len(chars) == 0 {
		return true
	}

	return slices.Contains(chars, charPtr.Name)
}
