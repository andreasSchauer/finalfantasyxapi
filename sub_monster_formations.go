package main

import (
	"fmt"
	"net/http"
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
);

type MonsterFormationSub struct {
	ID        		int32           `json:"id"`
	URL       		string          `json:"url"`
	Category		string			`json:"category"`
	IsForcedAmbush	bool			`json:"is_forced_ambush"`
	Monsters		[]string		`json:"monsters"`
	Areas			[]string		`json:"areas"`
}


func (m MonsterFormationSub) GetSectionName() string {
	return "monster-formations"
}

func (m MonsterFormationSub) GetURL() string {
	return m.URL
}

func idToLocAreaString(cfg *Config, areaID int32) string {
	area, _ := seeding.GetResourceByID(areaID, cfg.l.AreasID)
	areaString := nameVersionToString(area.Name, area.Version, area.Specification)
	return fmt.Sprintf("%s - %s - %s", area.SubLocation.Location.Name, area.SubLocation.Name, areaString)
}

func locAreaString(cfg *Config, locArea seeding.LocationArea) string {
	area, _ := seeding.GetResource(locArea, cfg.l.Areas)
	areaString := nameVersionToString(area.Name, area.Version, area.Specification)
	return fmt.Sprintf("%s - %s - %s", locArea.Location, locArea.Sublocation, areaString)
}

func encounterLocationsStrings(cfg *Config, items []seeding.EncounterLocation) []string {
	strings := []string{}

	for _, item := range items {
		s := locAreaString(cfg, item.LocationArea)

		if !slices.Contains(strings, s) {
			strings = append(strings, s)
		}
	}

	return strings
}

func getLocAreaStrings(cfg *Config, items []seeding.LocationArea) []string {
	strings := []string{}

	for _, item := range items {
		s := locAreaString(cfg, item)
		strings = append(strings, s)
	}

	return strings
}


func monsterAmountString(cfg *Config, ma seeding.MonsterAmount) string {
	key := seeding.LookupObject{
		Name: ma.MonsterName,
		Version: ma.Version,
	}
	mon, _ := seeding.GetResource(key, cfg.l.Monsters)
	monsterString := nameVersionToString(mon.Name, mon.Version, mon.Specification)
	return nameAmountString(monsterString, ma.Amount)
}

func getMonsterAmountStrings(cfg *Config, items []seeding.MonsterAmount) []string {
	strings := []string{}

	for _, item := range items {
		s := monsterAmountString(cfg, item)
		strings = append(strings, s)
	}

	return strings
}


func handleMonsterFormationsSection(cfg *Config, r *http.Request, dbIDs []int32) ([]SubResource, error) {
	i := cfg.e.monsterFormations
	formations := []MonsterFormationSub{}

	for _, formationID := range dbIDs {
		formation, _ := seeding.GetResourceByID(formationID, i.objLookupID)

		formationSub := MonsterFormationSub{
			ID:	formation.ID,
			URL: createResourceURL(cfg, i.endpoint, formationID),
			Category: formation.FormationData.Category,
			IsForcedAmbush: formation.FormationData.IsForcedAmbush,
			Monsters: getMonsterAmountStrings(cfg, formation.Monsters),
			Areas: encounterLocationsStrings(cfg, formation.EncounterLocations),
		}

		formations = append(formations, formationSub)
	}

	return toSubResourceSlice(formations), nil
}