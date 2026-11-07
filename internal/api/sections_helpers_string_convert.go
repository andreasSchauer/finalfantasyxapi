package api

import (
	"cmp"
	"fmt"
	"slices"
	"strings"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func idToLocAreaString(areaID int32) string {
	area, _ := lookup.GetResourceByID(areaID, lookup.L.AreasID)
	return areaToLocAreaString(area)
}

func locAreaString(locArea seedtypes.LocationArea) string {
	area, _ := lookup.GetResource(locArea, lookup.L.Areas)
	return areaToLocAreaString(area)
}

func areaToLocAreaString(area seedtypes.Area) string {
	locArea := area.GetLocationArea()
	locAreaString := h.NameToString(area.Name, area.Version, area.Specification)

	if area.Name != locArea.Sublocation {
		locAreaString = fmt.Sprintf("%s - %s", locArea.Sublocation, locAreaString)
	}

	if locArea.Sublocation != locArea.Location {
		locAreaString = fmt.Sprintf("%s - %s", locArea.Location, locAreaString)
	}

	return locAreaString
}

func locAreaStrings[T seedtypes.HasLocArea](items []T) []string {
	strings := []string{}
	seen := make(map[string]bool)

	for _, item := range items {
		s := locAreaString(item.GetLocationArea())

		if !seen[s] {
			seen[s] = true
			strings = append(strings, s)
		}
	}

	return strings
}

func monsterAmountString(ma seedtypes.MonsterAmount) string {
	mon, _ := lookup.GetResource(ma, lookup.L.Monsters)
	return h.NameAmountString(mon.Name, mon.Version, mon.Specification, ma.Amount)
}

func idToMonsterSimpleString(monID int32) string {
	mon, _ := lookup.GetResourceByID(monID, lookup.L.MonstersID)
	return h.NameToString(mon.Name, mon.Version, mon.Specification)
}

func monsterAutoAbilityString(drop seedtypes.EquipmentDrop) string {
	if len(drop.Characters) == 0 {
		return drop.Ability
	}

	formattedChars := h.StringSliceToListString(drop.Characters)
	return fmt.Sprintf("%s (%s)", drop.Ability, formattedChars)
}

func abilityRefString(ref seedtypes.AbilityReference) string {
	ability, _ := lookup.GetResource(ref, lookup.L.Abilities)
	return h.NameToString(ability.Name, ability.Version, ability.Specification)
}
