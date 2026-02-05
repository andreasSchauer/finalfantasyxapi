package main

import (
	"fmt"
	"slices"
	"strconv"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)



func nameToString(name string, version *int32, spec *string) string {
	var verStr string
	var specStr string

	if version != nil {
		intVer := int(*version)
		verStr = fmt.Sprintf(" %s", strconv.Itoa(intVer))
	}

	if spec != nil {
		specStr = fmt.Sprintf(" (%s)", *spec)
	}

	return name + verStr + specStr
}

func nameAmountString(name string, version *int32, spec *string, amount int32) string {
	nameStr := nameToString(name, version, spec)
	return fmt.Sprintf("%s x%d", nameStr, amount)
}


func idToLocAreaString(cfg *Config, areaID int32) string {
	area, _ := seeding.GetResourceByID(areaID, cfg.l.AreasID)
	areaString := nameToString(area.Name, area.Version, area.Specification)
	return fmt.Sprintf("%s - %s - %s", area.SubLocation.Location.Name, area.SubLocation.Name, areaString)
}

func locAreaString(cfg *Config, locArea seeding.LocationArea) string {
	area, _ := seeding.GetResource(locArea, cfg.l.Areas)
	areaString := nameToString(area.Name, area.Version, area.Specification)
	return fmt.Sprintf("%s - %s - %s", locArea.Location, locArea.Sublocation, areaString)
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
	return nameAmountString(mon.Name, mon.Version, mon.Specification, ma.Amount)
}

func idToMonsterSubString(cfg *Config, monID int32) string {
	mon, _ := seeding.GetResourceByID(monID, cfg.l.MonstersID)
	return nameToString(mon.Name, mon.Version, mon.Specification)
}

func monsterAutoAbilityString(_ *Config, drop seeding.EquipmentDrop) string {
	if len(drop.Characters) == 0 {
		return drop.Ability
	}

	formattedChars := h.StringSliceToListString(drop.Characters)
	return fmt.Sprintf("%s (%s)", drop.Ability, formattedChars)
}