package main

import (
	"fmt"
	"net/http"
	"strconv"

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
	segments := getPathSegments(r.URL.Path, "areas")

	// this whole thing can probably be generalized
	switch len(segments) {
	case 0:
		// /api/areas
		resourceList, err := cfg.retrieveAreas(r)
		if handleHTTPError(w, err) {
			return
		}
		respondWithJSON(w, http.StatusOK, resourceList)
		return
	case 1:
		// /api/areas/{id}
		segment := segments[0]

		if segment == "parameters" {
			parameterList, err := cfg.getQueryParamList(r, cfg.q.areas)
			if handleHTTPError(w, err) {
				return
			}

			respondWithJSON(w, http.StatusOK, parameterList)
			return
		}

		id, err := strconv.Atoi(segment)
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "Wrong format. Usage: /api/areas/{id}.", err)
			return
		}

		area, err := cfg.getArea(r, int32(id))
		if handleHTTPError(w, err) {
			return
		}

		respondWithJSON(w, http.StatusOK, area)
		return

	case 2:
		// /api/areas/{id}/{subSection}
		// areaID := segments[0]
		subSection := segments[1]

		// to generalize the switch cases, I could use a map[string]WhateverList and trigger its function or return the error, if the key is not in there
		switch subSection {
		case "connected":
			fmt.Println(segments)
			fmt.Println("this should trigger /api/areas/{id}/connected")
			return
		case "monsters":
			fmt.Println(segments)
			fmt.Println("this should trigger /api/areas/{id}/monsters")
			return
		case "monster-formations":
			fmt.Println(segments)
			fmt.Println("this should trigger /api/areas/{id}/monster-formations")
			return
		case "shops":
			fmt.Println(segments)
			fmt.Println("this should trigger /api/areas/{id}/shops")
			return
		case "treasures":
			fmt.Println(segments)
			fmt.Println("this should trigger /api/areas/{id}/treasures")
			return
		default:
			fmt.Println(segments)
			fmt.Println("this should trigger an error: this sub section is not supported. Supported sub-sections: connected, monsters, monster-formations, shops, treasures.")
			return
		}

	default:
		respondWithError(w, http.StatusBadRequest, `Wrong format. Usage: /api/areas/{id}, or /api/areas/{id}/{sub-section}. Supported sub-sections: connected, monsters, monster-formations, shops, treasures.`, nil)
		return
	}
}

func (cfg *Config) getArea(r *http.Request, id int32) (Area, error) {
	err := verifyQueryParams(r, "areas", &id, cfg.q.areas)
	if err != nil {
		return Area{}, err
	}

	dbArea, err := cfg.db.GetArea(r.Context(), id)
	if err != nil {
		return Area{}, newHTTPError(http.StatusNotFound, "Couldn't get Area. Area with this ID doesn't exist.", err)
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
	err := verifyQueryParams(r, "areas", nil, cfg.q.areas)
	if err != nil {
		return LocationApiResourceList{}, err
	}

	dbAreas, err := cfg.db.GetAreas(r.Context())
	if err != nil {
		return LocationApiResourceList{}, newHTTPError(http.StatusInternalServerError, "Couldn't retrieve areas", err)
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

	resourceList, err := cfg.newLocationAPIResourceList(r, resources)
	if err != nil {
		return LocationApiResourceList{}, err
	}

	return resourceList, nil
}
