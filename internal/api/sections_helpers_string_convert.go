package api

import (
	"fmt"
	"slices"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func idToLocAreaString(cfg *Config, areaID int32) string {
	area, _ := seeding.GetResourceByID(areaID, cfg.l.AreasID)
	return areaToLocAreaString(area)
}

func locAreaString(cfg *Config, locArea seeding.LocationArea) string {
	area, _ := seeding.GetResource(locArea, cfg.l.Areas)
	return areaToLocAreaString(area)
}

func areaToLocAreaString(area seeding.Area) string {
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


func locAreaStrings[T seeding.HasLocArea](cfg *Config, items []T) []string {
	strings := []string{}

	for _, item := range items {
		s := locAreaString(cfg, item.GetLocationArea())

		if !slices.Contains(strings, s) {
			strings = append(strings, s)
		}
	}

	return strings
}

func monsterAmountString(cfg *Config, ma seeding.MonsterAmount) string {
	key := seeding.LookupObject{
		Name:    ma.MonsterName,
		Version: ma.Version,
	}
	mon, _ := seeding.GetResource(key, cfg.l.Monsters)
	return h.NameAmountString(mon.Name, mon.Version, mon.Specification, ma.Amount)
}

func idToMonsterSimpleString(cfg *Config, monID int32) string {
	mon, _ := seeding.GetResourceByID(monID, cfg.l.MonstersID)
	return h.NameToString(mon.Name, mon.Version, mon.Specification)
}

func monsterAutoAbilityString(_ *Config, drop seeding.EquipmentDrop) string {
	if len(drop.Characters) == 0 {
		return drop.Ability
	}

	formattedChars := h.StringSliceToListString(drop.Characters)
	return fmt.Sprintf("%s (%s)", drop.Ability, formattedChars)
}

func abilityRefString(cfg *Config, ref seeding.AbilityReference) string {
	ability, _ := seeding.GetResource(ref, cfg.l.Abilities)
	return h.NameToString(ability.Name, ability.Version, ability.Specification)
}
