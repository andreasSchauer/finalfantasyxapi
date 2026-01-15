package main

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Area struct {
	ID                int32                `json:"id"`
	Name              string               `json:"name"`
	Version           *int32               `json:"version,omitempty"`
	Specification     *string              `json:"specification,omitempty"`
	ParentLocation    NamedAPIResource     `json:"parent_location"`
	ParentSublocation NamedAPIResource     `json:"parent_sublocation"`
	StoryOnly         bool                 `json:"story_only"`
	HasSaveSphere     bool                 `json:"has_save_sphere"`
	AirshipDropOff    bool                 `json:"airship_drop_off"`
	HasCompSphere     bool                 `json:"has_comp_sphere"`
	CanRideChocobo    bool                 `json:"can_ride_chocobo"`
	ConnectedAreas    []AreaConnection     `json:"connected_areas"`
	Characters        []NamedAPIResource   `json:"characters"`
	Aeons             []NamedAPIResource   `json:"aeons"`
	Shops             []UnnamedAPIResource `json:"shops"`
	Treasures         []UnnamedAPIResource `json:"treasures"`
	Monsters          []NamedAPIResource   `json:"monsters"`
	Formations        []UnnamedAPIResource `json:"formations"`
	Sidequest         *NamedAPIResource    `json:"sidequest"`
	Music             *LocationMusic       `json:"music"`
	FMVs              []NamedAPIResource   `json:"fmvs"`
}

func (cfg *Config) HandleAreas(w http.ResponseWriter, r *http.Request) {
	i := cfg.e.areas

	segments := getPathSegments(r.URL.Path, i.endpoint)

	switch len(segments) {
	case 0:
		handleEndpointList(w, r, i)
		return

	case 1:
		handleEndpointIDOnly(cfg, w, r, i, segments)
		return

	case 2:
		handleEndpointSubsections(w, i, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: '/api/%s/{id}', or '/api/%s/{id}/{subsection}'. supported subsections: %s.", i.endpoint, i.endpoint, h.GetMapKeyStr(i.subsections)), nil)
		return
	}
}

func (cfg *Config) getArea(r *http.Request, id int32) (Area, error) {
	i := cfg.e.areas

	area, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return Area{}, err
	}

	rel, err := cfg.getAreaRelationships(r, area)
	if err != nil {
		return Area{}, err
	}

	response := Area{
		ID:                area.ID,
		Name:              area.Name,
		Version:           area.Version,
		Specification:     area.Specification,
		ParentLocation:    nameToNamedAPIResource(cfg, cfg.e.locations, area.SubLocation.Location.Name, nil),
		ParentSublocation: nameToNamedAPIResource(cfg, cfg.e.sublocations, area.SubLocation.Name, nil),
		StoryOnly:         area.StoryOnly,
		HasSaveSphere:     area.HasSaveSphere,
		AirshipDropOff:    area.AirshipDropOff,
		HasCompSphere:     area.HasCompilationSphere,
		CanRideChocobo:    area.CanRideChocobo,
		ConnectedAreas:    rel.ConnectedAreas,
		Characters:        rel.Characters,
		Aeons:             rel.Aeons,
		Shops:             rel.Shops,
		Treasures:         rel.Treasures,
		Monsters:          rel.Monsters,
		Formations:        rel.Formations,
		Sidequest:         rel.Sidequest,
		Music:             rel.Music,
		FMVs:              rel.FMVs,
	}

	return response, nil
}

func (cfg *Config) retrieveAreas(r *http.Request) (LocationApiResourceList, error) {
	i := cfg.e.areas

	resources, err := retrieveLocationAPIResources(cfg, r, i)
	if err != nil {
		return LocationApiResourceList{}, err
	}

	filteredLists := []filteredResList[LocationAPIResource]{
		frl(idQueryLocBased(cfg, r, i, resources, "location", len(cfg.l.Locations), cfg.db.GetLocationAreaIDs)),
		frl(idQueryLocBased(cfg, r, i, resources, "sublocation", len(cfg.l.Sublocations), cfg.db.GetSublocationAreaIDs)),
		frl(idQueryWrapperLocBased(r, i, resources, "item", len(cfg.l.Items), cfg.queryAreasByItemMethod)),
		frl(idQueryWrapperLocBased(r, i, resources, "key-item", len(cfg.l.KeyItems), cfg.getAreasByKeyItem)),
		frl(boolQueryLocBased(cfg, r, i, resources, "story-based", cfg.db.GetAreaIDsStoryOnly)),
		frl(boolQueryLocBased(cfg, r, i, resources, "save-sphere", cfg.db.GetAreaIDsWithSaveSphere)),
		frl(boolQueryLocBased(cfg, r, i, resources, "comp-sphere", cfg.db.GetAreaIDsWithCompSphere)),
		frl(boolQueryLocBased(cfg, r, i, resources, "airship", cfg.db.GetAreaIDsWithDropOff)),
		frl(boolQueryLocBased(cfg, r, i, resources, "chocobo", cfg.db.GetAreaIDsChocobo)),
		frl(boolAccumulatorLocBased(cfg, r, i, resources, "characters", cfg.db.GetAreaIDsCharacters)),
		frl(boolAccumulatorLocBased(cfg, r, i, resources, "aeons", cfg.db.GetAreaIDsAeons)),
		frl(boolAccumulatorLocBased(cfg, r, i, resources, "monsters", cfg.db.GetAreaIDsMonsters)),
		frl(boolAccumulatorLocBased(cfg, r, i, resources, "boss-fights", cfg.db.GetAreaIDsBosses)),
		frl(boolAccumulatorLocBased(cfg, r, i, resources, "shops", cfg.db.GetAreaIDsShops)),
		frl(boolAccumulatorLocBased(cfg, r, i, resources, "treasures", cfg.db.GetAreaIDsTreasures)),
		frl(boolAccumulatorLocBased(cfg, r, i, resources, "sidequests", cfg.db.GetAreaIDsSidequests)),
		frl(boolAccumulatorLocBased(cfg, r, i, resources, "fmvs", cfg.db.GetAreaIDsFMVs)),
	}

	return filterLocationAPIResources(cfg, r, i, resources, filteredLists)
}
