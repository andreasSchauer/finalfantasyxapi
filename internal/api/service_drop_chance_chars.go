package api

import (
	"slices"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

type CharChanceParams struct {
	monMaxSlots     int32
	wantedAbilities []seedtypes.AutoAbility
	monAbilities    []seedtypes.EquipmentDrop
}

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

func (p CharChanceParams) calcChances(params DropChanceParams) CharacterChances {
	eligibleChars := []string{}
	baseRandomChance := 0.75 / h.FloatLen(params.PartyMembers)
	var chances CharacterChances

	if params.CharacterID != nil {
		finBlowChance := baseRandomChance + 0.25
		char, _ := lookup.GetResourceByID(*params.CharacterID, lookup.L.CharactersID)

		specificChar := SpecificCharChances{
			NoFinBlow:  baseRandomChance,
			FinBlow:    finBlowChance,
			IsEligible: p.charIsEligible(params, char),
		}

		chances.SpecificChar = &specificChar
	}

	for _, charID := range params.PartyMembers {
		char, _ := lookup.GetResourceByID(charID, lookup.L.CharactersID)

		if p.charIsEligible(params, char) {
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

func (p CharChanceParams) charIsEligible(params DropChanceParams, char seedtypes.Character) bool {
	if p.forcedAbilityEliminatesChar(params, char) {
		return false
	}

	for _, ability := range p.wantedAbilities {
		for _, monAbility := range p.monAbilities {
			if ability.Name == monAbility.Ability && !charGetsAbility(monAbility, &char) {
				return false
			}
		}
	}

	return true
}

func (p CharChanceParams) forcedAbilityEliminatesChar(params DropChanceParams, char seedtypes.Character) bool {
	// the character gets the forced ability, but the ability wasn't specified in auto_abilities
	charGetsUnwantedForced := p.charGetsUnwantedForced(char)
	noLooseSlotLeft := p.noLooseSlotLeft(params) // there's no loose slot left to fill in with an optional ability

	if charGetsUnwantedForced { // if this is true, the char gets the forced ability, but it is not desired
		return !params.OptionalAbilities || // optionalAbilities is false, so the forced ability can't be on it
			noLooseSlotLeft // optionalAbilities is true, but requirements leave no space for the forced ability
	}

	return false
}

func (p CharChanceParams) charGetsUnwantedForced(char seedtypes.Character) bool {
	return p.forcedSlotExists() &&
		!p.forcedSlotWanted() &&
		charGetsAbility(p.monAbilities[0], &char)
}

func (p CharChanceParams) noLooseSlotLeft(params DropChanceParams) bool {
	return params.GetRequiredSlots() == p.monMaxSlots
}

func (p CharChanceParams) forcedSlotExists() bool {
	return p.monAbilities[0].IsForced
}

func (p CharChanceParams) forcedSlotWanted() bool {
	if !p.forcedSlotExists() {
		return false
	}

	for _, ability := range p.wantedAbilities {
		if ability.Name == p.monAbilities[0].Ability {
			return true
		}
	}

	return false
}

func charGetsAbility(ability seedtypes.EquipmentDrop, charPtr *seedtypes.Character) bool {
	chars := ability.Characters

	if charPtr == nil || len(chars) == 0 {
		return true
	}

	return slices.Contains(chars, charPtr.Name)
}
