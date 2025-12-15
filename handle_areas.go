package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


type Area struct {
	ID					int32					`json:"id"`
	Name				string					`json:"name"`
	Version				*int32					`json:"version,omitempty"`
	Specification		*string					`json:"specification,omitempty"`
	ParentLocation		NamedAPIResource		`json:"parent_location"`
	ParentSublocation	NamedAPIResource		`json:"parent_sublocation"`
	StoryOnly			bool					`json:"story_only"`
	HasSaveSphere		bool					`json:"has_save_sphere"`
	AirshipDropOff		bool					`json:"airship_drop_off"`
	HasCompSphere		bool					`json:"has_comp_sphere"`
	CanRideChocobo		bool					`json:"can_ride_chocobo"`
	ConnectedAreas		[]AreaConnection		`json:"connected_areas"`
}


type AreaConnection struct {
	Area			LocationAPIResource			`json:"area"`
	ConnectionType	NamedAPIResource			`json:"connection_type"`
	StoryOnly		bool						`json:"story_only"`
	Notes			*string						`json:"notes,omitempty"`
}


func (cfg *apiConfig) handleAreas(w http.ResponseWriter, r *http.Request) {
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
		idStr := segments[0]
		
		id, err := strconv.Atoi(idStr)
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


func (cfg *apiConfig) getArea(r *http.Request, id int32) (Area, error) {
	dbArea, err := cfg.db.GetArea(r.Context(), id)
	if err != nil {
		return Area{}, newHTTPError(http.StatusNotFound, "Couldn't get Area. Area with this ID doesn't exist.", err)
	}

	connections, err := cfg.getConnectedAreas(r, dbArea)
	if err != nil {
		return Area{}, err
	}

	location  := cfg.newNamedAPIResourceSimple("locations", h.NullInt32ToVal(dbArea.LocationID), h.NullStringToVal(dbArea.Location))

	sublocation := cfg.newNamedAPIResourceSimple("sublocations", dbArea.SublocationID, h.NullStringToVal(dbArea.Sublocation))

	area := Area{
		ID: 				dbArea.ID,
		Name: 				dbArea.Name,
		Version: 			h.NullInt32ToPtr(dbArea.Version),
		Specification: 		h.NullStringToPtr(dbArea.Specification),
		ParentLocation: 	location,
		ParentSublocation: 	sublocation,
		StoryOnly: 			dbArea.StoryOnly,
		HasSaveSphere: 		dbArea.HasSaveSphere,
		AirshipDropOff: 	dbArea.AirshipDropOff,
		HasCompSphere: 		dbArea.HasCompilationSphere,
		CanRideChocobo: 	dbArea.CanRideChocobo,
		ConnectedAreas: 	connections,
	}

	return area, nil
}


func (cfg *apiConfig) getConnectedAreas(r *http.Request, area database.GetAreaRow) ([]AreaConnection, error) {
	locArea := newLocationArea(h.NullStringToVal(area.Location), h.NullStringToVal(area.Sublocation), area.Name, h.NullInt32ToPtr(area.Version))
	dbConnAreas, err := cfg.db.GetAreaConnections(r.Context(), area.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("Couldn't retrieve connected areas of %s", locArea.Error()), err)
	}

	connectedAreas := []AreaConnection{}

	for _, dbConnArea := range dbConnAreas {
		locArea := newLocationArea(h.NullStringToVal(dbConnArea.Location), h.NullStringToVal(dbConnArea.Sublocation), h.NullStringToVal(dbConnArea.Area), h.NullInt32ToPtr(dbConnArea.Version))

		connType, err := cfg.newNamedAPIResourceFromType("connection-type", string(dbConnArea.ConnectionType), cfg.t.AreaConnectionType)
		if err != nil {
			return nil, err
		}

		connection := AreaConnection{
			Area: 			cfg.newLocationBasedAPIResource(locArea),
			ConnectionType: connType,
			StoryOnly: 		dbConnArea.StoryOnly,
			Notes: 			h.NullStringToPtr(dbConnArea.Notes),
		}

		connectedAreas = append(connectedAreas, connection)
	}

	return connectedAreas, nil
}


func (cfg *apiConfig) retrieveAreas(r *http.Request) (LocationdApiResourceList, error) {
	dbAreas, err := cfg.db.GetAreas(r.Context())
	if err != nil {
		return LocationdApiResourceList{}, newHTTPError(http.StatusInternalServerError, "Couldn't retrieve areas", err)
	}

	resources := createLocationBasedAPIResources(cfg, dbAreas, func(area database.GetAreasRow) (string, string, string, *int32) {
		return h.NullStringToVal(area.Location), h.NullStringToVal(area.Sublocation), area.Name, h.NullInt32ToPtr(area.Version)
	})

	resourceList, err := cfg.newLocationAPIResourceList(r, resources)
	if err != nil {
		return LocationdApiResourceList{}, err
	}

	return resourceList, nil
}