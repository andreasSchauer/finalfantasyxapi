package main

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
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

	err := verifyQueryParams(r, i, &id)
	if err != nil {
		return Area{}, err
	}

	dbArea, err := cfg.db.GetArea(r.Context(), id)
	if err != nil {
		return Area{}, newHTTPError(http.StatusNotFound, fmt.Sprintf("area with id '%d' doesn't exist.", id), err)
	}

	rel, err := cfg.getAreaRelationships(r, dbArea)
	if err != nil {
		return Area{}, err
	}

	location := cfg.newNamedAPIResourceSimple("locations", dbArea.LocationID, dbArea.Location)
	sublocation := cfg.newNamedAPIResourceSimple("sublocations", dbArea.SublocationID, dbArea.Sublocation)

	area := Area{
		ID:                dbArea.ID,
		Name:              dbArea.Name,
		Version:           h.NullInt32ToPtr(dbArea.Version),
		Specification:     h.NullStringToPtr(dbArea.Specification),
		ParentLocation:    location,
		ParentSublocation: sublocation,
		StoryOnly:         dbArea.StoryOnly,
		HasSaveSphere:     dbArea.HasSaveSphere,
		AirshipDropOff:    dbArea.AirshipDropOff,
		HasCompSphere:     dbArea.HasCompilationSphere,
		CanRideChocobo:    dbArea.CanRideChocobo,
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

	return area, nil
}


func (cfg *Config) retrieveAreas(r *http.Request) (LocationApiResourceList, error) {
	i := cfg.e.areas
	
	err := verifyQueryParams(r, i, nil)
	if err != nil {
		return LocationApiResourceList{}, err
	}

	dbAreas, err := cfg.db.GetAreas(r.Context())
	if err != nil {
		return LocationApiResourceList{}, newHTTPError(http.StatusInternalServerError, "couldn't retrieve areas.", err)
	}

	resources := createLocationBasedAPIResources(cfg, dbAreas, func(area database.GetAreasRow) (string, string, string, *int32) {
		return area.Location, area.Sublocation, area.Name, h.NullInt32ToPtr(area.Version)
	})

	filterFuncs := []func(*http.Request, []LocationAPIResource) ([]LocationAPIResource, error){
		cfg.getAreasLocation,
		cfg.getAreasSublocation,
		cfg.getAreasItem,
		cfg.getAreasKeyItem,
		cfg.getAreasStoryBased,
		cfg.getAreasSaveSphere,
		cfg.getAreasCompSphere,
		cfg.getAreasDropOff,
		cfg.getAreasChocobo,
		cfg.getAreasCharacters,
		cfg.getAreasAeons,
		cfg.getAreasMonsters,
		cfg.getAreasBosses,
		cfg.getAreasShops,
		cfg.getAreasTreasures,
		cfg.getAreasSidequests,
		cfg.getAreasFMVs,
	}

	for _, function := range filterFuncs {
		filteredResources, err := function(r, resources)
		if err != nil {
			return LocationApiResourceList{}, err
		}

		resources = getSharedResources(resources, filteredResources)
	}

	resourceList, err := newLocationAPIResourceList(cfg, r, i, resources)
	if err != nil {
		return LocationApiResourceList{}, err
	}

	return resourceList, nil
}
