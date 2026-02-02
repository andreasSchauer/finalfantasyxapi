package main

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type MonsterFormation struct {
	ID              int32            			`json:"id"`
	Version         *int32           			`json:"version,omitempty"`
	Category       	string             			`json:"category"`
	IsForcedAmbush 	bool               			`json:"is_forced_ambush"`
	CanEscape      	bool               			`json:"can_escape"`
	Notes          	*string            			`json:"notes,omitempty"`
	BossMusic      	*FormationBossSong 			`json:"boss_music,omitempty"`
	Monsters		[]MonsterAmount				`json:"monsters"`
	Areas			[]EncounterLocation			`json:"areas"`
	TriggerCommands	[]FormationTriggerCommand	`json:"trigger_commands"`
}


type EncounterLocation struct {
	Area  			LocationAPIResource `json:"area"`
	Specification 	*string 			`json:"specification"`
}

func createEncounterLocations(cfg *Config, items []seeding.EncounterLocation) []EncounterLocation {
	els := []EncounterLocation{}

	for _, item := range items {
		el := EncounterLocation{
			Area: 			locAreaToLocationAPIResource(cfg, cfg.e.areas, item.LocationArea),
			Specification: 	item.Specification,
		}
		els = append(els, el)
	}

	return els
}

type FormationTriggerCommand struct {
	Ability			NamedAPIResource	`json:"ability"`
	Condition       *string  			`json:"condition"`
	UseAmount       *int32   			`json:"use_amount"`
	Users           []NamedAPIResource 	`json:"users"`
}


func createFormationTriggerCommands (cfg *Config, items []seeding.FormationTriggerCommand) []FormationTriggerCommand {
	triggerCommands := []FormationTriggerCommand{}

	for _, item := range items {
		tc := FormationTriggerCommand{
			Ability: 	nameToNamedAPIResource(cfg, cfg.e.triggerCommands, item.Name, item.Version),
			Condition: 	item.Condition,
			UseAmount: 	item.UseAmount,
			Users: 		namesToNamedAPIResources(cfg, cfg.e.characterClasses, item.Users),
		}
		triggerCommands = append(triggerCommands, tc)
	}

	return triggerCommands
}




type FormationBossSong struct {
	Song             NamedAPIResource 	`json:"song"`
	CelebrateVictory bool   			`json:"celebrate_victory"`
}


func getFormationBossMusic(cfg *Config, bossMusic *seeding.FormationBossSong) *FormationBossSong {
	if bossMusic == nil {
		return nil
	}

	bossSong := FormationBossSong{
		Song: nameToNamedAPIResource(cfg, cfg.e.songs, bossMusic.Song, nil),
		CelebrateVictory: bossMusic.CelebrateVictory,
	}

	return &bossSong
}



func (cfg *Config) HandleMonsterFormations(w http.ResponseWriter, r *http.Request) {
	i := cfg.e.monsterFormations

	segments := getPathSegments(r.URL.Path, i.endpoint)

	switch len(segments) {
	case 0:
		handleEndpointList(w, r, i)
		return

	case 1:
		handleEndpointIDOnly(cfg, w, r, i, segments)
		return

	case 2:
		handleEndpointSubsections(cfg, w, r, i, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: '/api/%s/{id}', or '/api/%s/{id}/{subsection}'. supported subsections: %s.", i.endpoint, i.endpoint, h.GetMapKeyStr(i.subsections)), nil)
		return
	}
}

func (cfg *Config) getMonsterFormation(r *http.Request, i handlerInput[seeding.MonsterFormation, MonsterFormation, UnnamedAPIResource, UnnamedApiResourceList], id int32) (MonsterFormation, error) {
	formation, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return MonsterFormation{}, err
	}

	response := MonsterFormation{
		ID:                	formation.ID,
		Version: 			formation.Version,
		Category: 			formation.FormationData.Category,
		IsForcedAmbush: 	formation.FormationData.IsForcedAmbush,
		CanEscape: 			formation.FormationData.CanEscape,
		Notes: 				formation.FormationData.Notes,
		BossMusic: 			getFormationBossMusic(cfg, formation.FormationData.BossMusic),
		Monsters: 			createMonsterAmounts(cfg, formation.Monsters),
		Areas: 				createEncounterLocations(cfg, formation.EncounterLocations),
		TriggerCommands: 	createFormationTriggerCommands(cfg, formation.TriggerCommands),
	}

	return response, nil
}


func (cfg *Config) retrieveMonsterFormations(r *http.Request, i handlerInput[seeding.MonsterFormation, MonsterFormation, UnnamedAPIResource, UnnamedApiResourceList]) (UnnamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return UnnamedApiResourceList{}, err
	}

	filteredLists := []filteredResList[UnnamedAPIResource]{
		frl(idOnlyQuery(cfg, r, i, resources, "monster", len(cfg.l.Monsters), cfg.db.GetMonsterFormationIDsByMonster)),
		frl(idOnlyQuery(cfg, r, i, resources, "location", len(cfg.l.Locations), cfg.db.GetMonsterFormationIDsByLocation)),
		frl(idOnlyQuery(cfg, r, i, resources, "sublocation", len(cfg.l.Sublocations), cfg.db.GetMonsterFormationIDsBySublocation)),
		frl(idOnlyQuery(cfg, r, i, resources, "area", len(cfg.l.Areas), cfg.db.GetMonsterFormationIDsByArea)),
		frl(boolQuery(cfg, r, i, resources, "ambush", cfg.db.GetMonsterFormationIDsByForcedAmbush)),
		frl(typeQuery(cfg, r, i, cfg.t.MonsterFormationCategory, resources, "category", cfg.db.GetMonsterFormationIDsByCategory)),
	}

	return filterAPIResources(cfg, r, i, resources, filteredLists)
}
