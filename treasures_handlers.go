package main

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)


type Treasure struct {
	ID              int32
	Version         int32				`json:"-"`
	Area          	LocationAPIResource	`json:"area"`
	TreasureType    NamedAPIResource    `json:"treasure_type"`
	LootType        NamedAPIResource    `json:"loot_type"`
	IsPostAirship   bool            	`json:"is_post_airship"`
	IsAnimaTreasure bool            	`json:"is_anima_treasure"`
	Notes           *string         	`json:"notes"`
	GilAmount       *int32          	`json:"gil_amount,omitempty"`
	Items           []ItemAmount    	`json:"items,omitempty"`
	Equipment       *FoundEquipment 	`json:"equipment,omitempty"`
}

type FoundEquipment struct {
	EquipmentName    NamedAPIResource   `json:"name"`
	Abilities        []NamedAPIResource `json:"abilities"`
	EmptySlotsAmount int32    			`json:"empty_slots_amount"`
}

func (cfg *Config) HandleTreasures(w http.ResponseWriter, r *http.Request) {
	i := cfg.e.treasures

	segments := getPathSegments(r.URL.Path, i.endpoint)

	switch len(segments) {
	case 0:
		handleEndpointList(w, r, i)
		return

	case 1:
		handleEndpointIDOnly(cfg, w, r, i, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: '/api/%s/{id}'.", i.endpoint), nil)
		return
	}
}

func (cfg *Config) getTreasure(r *http.Request, i handlerInput[seeding.Treasure, Treasure, UnnamedAPIResource, UnnamedApiResourceList], id int32) (Treasure, error) {
	treasure, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return Treasure{}, err
	}

	response := Treasure{
		ID:                	treasure.ID,
		
	}

	return response, nil
}


func (cfg *Config) retrieveTreasures(r *http.Request, i handlerInput[seeding.Treasure, Treasure, UnnamedAPIResource, UnnamedApiResourceList]) (UnnamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return UnnamedApiResourceList{}, err
	}

	filteredLists := []filteredResList[UnnamedAPIResource]{
		frl(idOnlyQuery(cfg, r, i, resources, "location", len(cfg.l.Locations), cfg.db.GetLocationTreasureIDs)),
		frl(idOnlyQuery(cfg, r, i, resources, "sublocation", len(cfg.l.Sublocations), cfg.db.GetSublocationTreasureIDs)),
		frl(idOnlyQuery(cfg, r, i, resources, "area", len(cfg.l.Areas), cfg.db.GetAreaTreasureIDs)),
		frl(boolQuery(cfg, r, i, resources, "anima", cfg.db.GetTreasureIDsByIsAnimaTreasure)),
		frl(boolQuery(cfg, r, i, resources, "airship", cfg.db.GetTreasureIDsByIsPostAirship)),
		frl(typeQuery(cfg, r, i, cfg.t.LootType, resources, "loot_type", cfg.db.GetTreasureIDsByLootType)),
		frl(typeQuery(cfg, r, i, cfg.t.TreasureType, resources, "treasure_type", cfg.db.GetTreasureIDsByTreasureType)),
	}

	return filterAPIResources(cfg, r, i, resources, filteredLists)
}