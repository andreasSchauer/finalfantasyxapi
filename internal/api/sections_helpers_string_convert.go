package api

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func idToLocAreaString(cfg *Config, areaID int32) string {
	area, _ := seeding.GetResourceByID(areaID, cfg.L.AreasID)
	return areaToLocAreaString(area)
}

func locAreaString(cfg *Config, locArea seedtypes.LocationArea) string {
	area, _ := seeding.GetResource(locArea, cfg.L.Areas)
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

func locAreaStrings[T seedtypes.HasLocArea](cfg *Config, items []T) []string {
	strings := []string{}
	seen := make(map[string]bool)

	for _, item := range items {
		s := locAreaString(cfg, item.GetLocationArea())

		if !seen[s] {
			seen[s] = true
			strings = append(strings, s)
		}
	}

	return strings
}

func monsterAmountString(cfg *Config, ma seedtypes.MonsterAmount) string {
	mon, _ := seeding.GetResource(ma, cfg.L.Monsters)
	return h.NameAmountString(mon.Name, mon.Version, mon.Specification, ma.Amount)
}

func idToMonsterSimpleString(cfg *Config, monID int32) string {
	mon, _ := seeding.GetResourceByID(monID, cfg.L.MonstersID)
	return h.NameToString(mon.Name, mon.Version, mon.Specification)
}

func monsterAutoAbilityString(_ *Config, drop seedtypes.EquipmentDrop) string {
	if len(drop.Characters) == 0 {
		return drop.Ability
	}

	formattedChars := h.StringSliceToListString(drop.Characters)
	return fmt.Sprintf("%s (%s)", drop.Ability, formattedChars)
}

func abilityRefString(cfg *Config, ref seedtypes.AbilityReference) string {
	ability, _ := seeding.GetResource(ref, cfg.L.Abilities)
	return h.NameToString(ability.Name, ability.Version, ability.Specification)
}
