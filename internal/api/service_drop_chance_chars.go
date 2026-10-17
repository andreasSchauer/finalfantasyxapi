package api

import (
	"slices"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

// put the char specific field in their own struct and make that a pointer
// that's better than having three pointer fields, one of them a boolean
type CharacterChances struct {
	EligibleChars int32                `json:"eligible_chars"`
	AnyFinBlow    float64              `json:"any_fin_blow"`
	AnyNoFinBlow  float64              `json:"any_no_fin_blow"`
	SpecificChar  *SpecificCharChances `json:"specific_character,omitempty"`
}

func (c CharacterChances) Percent() CharacterChances {
	c.AnyFinBlow = h.DecimalToPercent(c.AnyFinBlow)
	c.AnyNoFinBlow = h.DecimalToPercent(c.AnyNoFinBlow)

	if c.SpecificChar != nil {
		newObj := c.SpecificChar.Percent()
		c.SpecificChar = &newObj
	}

	return c
}

func (c CharacterChances) Round(n int32) CharacterChances {
	c.AnyFinBlow = h.FloatRound(c.AnyFinBlow, n)
	c.AnyNoFinBlow = h.FloatRound(c.AnyNoFinBlow, n)

	if c.SpecificChar != nil {
		newObj := c.SpecificChar.Round(n)
		c.SpecificChar = &newObj
	}

	return c
}

type SpecificCharChances struct {
	IsEligible bool    `json:"is_eligible"`
	FinBlow    float64 `json:"fin_blow"`
	NoFinBlow  float64 `json:"no_fin_blow"`
}

func (c SpecificCharChances) Percent() SpecificCharChances {
	c.FinBlow = h.DecimalToPercent(c.FinBlow)
	c.NoFinBlow = h.DecimalToPercent(c.NoFinBlow)

	return c
}

func (c SpecificCharChances) Round(n int32) SpecificCharChances {
	c.FinBlow = h.FloatRound(c.FinBlow, n)
	c.NoFinBlow = h.FloatRound(c.NoFinBlow, n)

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
			NoFinBlow:  baseRandomChance,
			FinBlow:    finBlowChance,
			IsEligible: charIsEligible(char, wantedAbilities, monAbilities, params.LenientAbilities),
		}

		chances.SpecificChar = &specificChar
	}

	for _, charID := range params.PartyMembers {
		char, _ := seeding.GetResourceByID(charID, cfg.l.CharactersID)

		if charIsEligible(char, wantedAbilities, monAbilities, params.LenientAbilities) {
			eligibleChars++
		}
	}

	chances.AnyNoFinBlow = float64(eligibleChars) * baseRandomChance
	chances.EligibleChars = eligibleChars

	if eligibleChars > 0 {
		chances.AnyFinBlow = chances.AnyNoFinBlow + 0.25
	}

	// if we don't care about who gets the equipment and all characters in the party are eligible
	// then it doesn't matter who scores the finishing blow
	// therefore, it's always a guaranteed character match
	if eligibleChars == h.Len32(params.PartyMembers) {
		chances.AnyFinBlow = 1
		chances.AnyNoFinBlow = 1
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
	return forcedSlotExists(monAbilities) &&
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
