package main

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type Treasure struct {
	ID              int32
	Area            AreaAPIResource  `json:"area"`
	IsPostAirship   bool             `json:"is_post_airship"`
	IsAnimaTreasure bool             `json:"is_anima_treasure"`
	Notes           *string          `json:"notes,omitempty"`
	TreasureType    NamedAPIResource `json:"treasure_type"`
	LootType        NamedAPIResource `json:"loot_type"`
	GilAmount       *int32           `json:"gil_amount,omitempty"`
	Items           []ItemAmount     `json:"items,omitempty"`
	Equipment       *FoundEquipment  `json:"equipment,omitempty"`
}

type FoundEquipment struct {
	EquipmentName    NamedAPIResource   `json:"name"`
	Abilities        []NamedAPIResource `json:"abilities"`
	EmptySlotsAmount int32              `json:"empty_slots_amount"`
}

func convertFoundEquipment(cfg *Config, fe seeding.FoundEquipment) FoundEquipment {
	return FoundEquipment{
		EquipmentName:    nameToNamedAPIResource(cfg, cfg.e.equipment, fe.Name, nil),
		Abilities:        namesToNamedAPIResources(cfg, cfg.e.autoAbilities, fe.Abilities),
		EmptySlotsAmount: fe.EmptySlotsAmount,
	}
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

	case 2:
		handleEndpointSubsections(cfg, w, r, i, segments)
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

	treasureType, _ := newNamedAPIResourceFromType(cfg, cfg.e.treasureType.endpoint, treasure.TreasureType, cfg.t.TreasureType)
	lootType, _ := newNamedAPIResourceFromType(cfg, cfg.e.lootType.endpoint, treasure.LootType, cfg.t.LootType)

	response := Treasure{
		ID:              treasure.ID,
		Area:            idToAreaAPIResource(cfg, cfg.e.areas, treasure.AreaID),
		IsPostAirship:   treasure.IsPostAirship,
		IsAnimaTreasure: treasure.IsAnimaTreasure,
		Notes:           treasure.Notes,
		TreasureType:    treasureType,
		LootType:        lootType,
		GilAmount:       treasure.GilAmount,
		Items:           convertObjSlice(cfg, treasure.Items, convertItemAmount),
		Equipment:       convertObjPtr(cfg, treasure.Equipment, convertFoundEquipment),
	}

	return response, nil
}

func (cfg *Config) retrieveTreasures(r *http.Request, i handlerInput[seeding.Treasure, Treasure, UnnamedAPIResource, UnnamedApiResourceList]) (UnnamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return UnnamedApiResourceList{}, err
	}

	filteredLists := []filteredResList[UnnamedAPIResource]{
		frl(idQuery(cfg, r, i, resources, "location", len(cfg.l.Locations), cfg.db.GetLocationTreasureIDs)),
		frl(idQuery(cfg, r, i, resources, "sublocation", len(cfg.l.Sublocations), cfg.db.GetSublocationTreasureIDs)),
		frl(idQuery(cfg, r, i, resources, "area", len(cfg.l.Areas), cfg.db.GetAreaTreasureIDs)),
		frl(boolQuery(cfg, r, i, resources, "anima", cfg.db.GetTreasureIDsByIsAnimaTreasure)),
		frl(boolQuery(cfg, r, i, resources, "airship", cfg.db.GetTreasureIDsByIsPostAirship)),
		frl(typeQuery(cfg, r, i, cfg.t.LootType, resources, "loot_type", cfg.db.GetTreasureIDsByLootType)),
		frl(typeQuery(cfg, r, i, cfg.t.TreasureType, resources, "treasure_type", cfg.db.GetTreasureIDsByTreasureType)),
	}

	return filterAPIResources(cfg, r, i, resources, filteredLists)
}
