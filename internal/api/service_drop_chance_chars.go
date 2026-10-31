package api

import (
	"slices"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)


type CharacterChances struct {
	EligibleChars []string             `json:"eligible_chars"`
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

func calcCharChances(cfg *Config, params DropChanceParams, wantedAbilities []seeding.AutoAbility, monAbilities []seeding.EquipmentDrop, monMaxSlots int32) CharacterChances {
	eligibleChars := []string{}
	baseRandomChance := 0.75 / h.FloatLen(params.PartyMembers)
	var chances CharacterChances

	if params.CharacterID != nil {
		finBlowChance := baseRandomChance + 0.25
		char, _ := seeding.GetResourceByID(*params.CharacterID, cfg.L.CharactersID)

		specificChar := SpecificCharChances{
			NoFinBlow:  baseRandomChance,
			FinBlow:    finBlowChance,
			IsEligible: charIsEligible(params, char, monMaxSlots, wantedAbilities, monAbilities),
		}

		chances.SpecificChar = &specificChar
	}

	for _, charID := range params.PartyMembers {
		char, _ := seeding.GetResourceByID(charID, cfg.L.CharactersID)

		if charIsEligible(params, char, monMaxSlots, wantedAbilities, monAbilities) {
			eligibleChars = append(eligibleChars, char.Name)
		}
	}
	eligibleCharsAmt := len(eligibleChars)

	chances.AnyNoFinBlow = float64(eligibleCharsAmt) * baseRandomChance
	chances.EligibleChars = eligibleChars

	if eligibleCharsAmt > 0 {
		chances.AnyFinBlow = chances.AnyNoFinBlow + 0.25
	}

	// if we don't care about who gets the equipment and all characters in the party are eligible
	// then it doesn't matter who scores the finishing blow
	// therefore, it's always a guaranteed character match
	if eligibleCharsAmt == len(params.PartyMembers) {
		chances.AnyFinBlow = 1
		chances.AnyNoFinBlow = 1
	}

	return chances
}

func charIsEligible(params DropChanceParams, char seeding.Character, monMaxSlots int32, wantedAbilities []seeding.AutoAbility, monAbilities []seeding.EquipmentDrop) bool {
	if forcedAbilityEliminatesChar(params, char, monMaxSlots, wantedAbilities, monAbilities) {
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

func forcedAbilityEliminatesChar(params DropChanceParams, char seeding.Character, monMaxSlots int32, wantedAbilities []seeding.AutoAbility, monAbilities []seeding.EquipmentDrop) bool {
	// the character gets the forced ability, but the ability wasn't specified in auto_abilities
	charGetsUnwantedForced := charGetsUnwantedForced(char, wantedAbilities, monAbilities)
	noLooseSlotLeft := noLooseSlotLeft(params, monMaxSlots) // there's no loose slot left to fill in with an optional ability

	if charGetsUnwantedForced { // if this is true, the char gets the forced ability, but it is not desired
		return !params.OptionalAbilities || // optionalAbilities is false, so the forced ability can't be on it
			noLooseSlotLeft // optionalAbilities is true, but requirements leave no space for the forced ability
	}

	return false
}

func charGetsUnwantedForced(char seeding.Character, wantedAbilities []seeding.AutoAbility, monAbilities []seeding.EquipmentDrop) bool {
	return forcedSlotExists(monAbilities) &&
		!forcedSlotWanted(wantedAbilities, monAbilities) &&
		charGetsAbility(monAbilities[0], &char)
}

func noLooseSlotLeft(params DropChanceParams, monMaxSlots int32) bool {
	return params.GetRequiredSlots() == monMaxSlots
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
