package main

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type Area struct {
	ID                int32            `json:"id"`
	Name              string           `json:"name"`
	Version           *int32           `json:"version,omitempty"`
	Specification     *string          `json:"specification,omitempty"`
	ParentLocation    NamedAPIResource `json:"parent_location"`
	ParentSublocation NamedAPIResource `json:"parent_sublocation"`
	StoryOnly         bool             `json:"story_only"`
	HasSaveSphere     bool             `json:"has_save_sphere"`
	AirshipDropOff    bool             `json:"airship_drop_off"`
	HasCompSphere     bool             `json:"has_comp_sphere"`
	CanRideChocobo    bool             `json:"can_ride_chocobo"`
	ConnectedAreas    []AreaConnection `json:"connected_areas"`
	LocRel
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
		handleEndpointSubsections(cfg, w, r, i, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: '/api/%s/{id}', or '/api/%s/{id}/{subsection}'. supported subsections: %s.", i.endpoint, i.endpoint, h.GetMapKeyStr(i.subsections)), nil)
		return
	}
}

func (cfg *Config) getArea(r *http.Request, i handlerInput[seeding.Area, Area, LocationAPIResource, LocationApiResourceList], id int32) (Area, error) {
	area, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return Area{}, err
	}

	connections, err := getAreaConnectedAreas(cfg, area)
	if err != nil {
		return Area{}, err
	}

	rel, err := getAreaRelationships(cfg, r, area)
	if err != nil {
		return Area{}, err
	}

	response := Area{
		ID:                area.ID,
		Name:              area.Name,
		Version:           area.Version,
		Specification:     area.Specification,
		ParentLocation:    nameToNamedAPIResource(cfg, cfg.e.locations, area.Sublocation.Location.Name, nil),
		ParentSublocation: nameToNamedAPIResource(cfg, cfg.e.sublocations, area.Sublocation.Name, nil),
		StoryOnly:         area.StoryOnly,
		HasSaveSphere:     area.HasSaveSphere,
		AirshipDropOff:    area.AirshipDropOff,
		HasCompSphere:     area.HasCompilationSphere,
		CanRideChocobo:    area.CanRideChocobo,
		ConnectedAreas:    connections,
		LocRel:            rel,
	}

	return response, nil
}

func (cfg *Config) retrieveAreas(r *http.Request, i handlerInput[seeding.Area, Area, LocationAPIResource, LocationApiResourceList]) (LocationApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return LocationApiResourceList{}, err
	}

	filteredLists := []filteredResList[LocationAPIResource]{
		frl(idOnlyQuery(cfg, r, i, resources, "location", len(cfg.l.Locations), cfg.db.GetLocationAreaIDs)),
		frl(idOnlyQuery(cfg, r, i, resources, "sublocation", len(cfg.l.Sublocations), cfg.db.GetSublocationAreaIDs)),
		frl(idOnlyQueryWrapper(cfg, r, i, resources, "item", len(cfg.l.Items), getAreasByItem)),
		frl(idOnlyQueryWrapper(cfg, r, i, resources, "key_item", len(cfg.l.KeyItems), getAreasByKeyItem)),
		frl(boolQuery(cfg, r, i, resources, "story_based", cfg.db.GetAreaIDsStoryOnly)),
		frl(boolQuery(cfg, r, i, resources, "save_sphere", cfg.db.GetAreaIDsWithSaveSphere)),
		frl(boolQuery(cfg, r, i, resources, "comp_sphere", cfg.db.GetAreaIDsWithCompSphere)),
		frl(boolQuery(cfg, r, i, resources, "airship", cfg.db.GetAreaIDsWithDropOff)),
		frl(boolQuery(cfg, r, i, resources, "chocobo", cfg.db.GetAreaIDsChocobo)),
		frl(boolQuery2(cfg, r, i, resources, "characters", cfg.db.GetAreaIDsWithCharacters)),
		frl(boolQuery2(cfg, r, i, resources, "aeons", cfg.db.GetAreaIDsWithAeons)),
		frl(boolQuery2(cfg, r, i, resources, "monsters", cfg.db.GetAreaIDsWithMonsters)),
		frl(boolQuery2(cfg, r, i, resources, "boss_fights", cfg.db.GetAreaIDsWithBosses)),
		frl(boolQuery2(cfg, r, i, resources, "shops", cfg.db.GetAreaIDsWithShops)),
		frl(boolQuery2(cfg, r, i, resources, "treasures", cfg.db.GetAreaIDsWithTreasures)),
		frl(boolQuery2(cfg, r, i, resources, "sidequests", cfg.db.GetAreaIDsWithSidequests)),
		frl(boolQuery2(cfg, r, i, resources, "fmvs", cfg.db.GetAreaIDsWithFMVs)),
	}

	return filterAPIResources(cfg, r, i, resources, filteredLists)
}
