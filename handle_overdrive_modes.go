package main

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type OverdriveMode struct {
	ID			int32					`json:"id"`
	Name		string					`json:"name"`
	Description	string					`json:"description"`
	Effect		string					`json:"effect"`
	Type		string					`json:"type"`
	FillRate	*float32				`json:"fill_rate,omitempty"`
	Actions		[]NamedAPIResource		`json:"actions"`
}



type NamedApiResourceList struct {
	Count		int					`json:"count"`
	Next		*string				`json:"next"`
	Previous	*string				`json:"previous"`
	Results		[]NamedAPIResource	`json:"results"`
}


type NamedAPIResource struct {
	Name			string		`json:"name"`
	Version			*int32		`json:"version,omitempty"`
	Specification	*string		`json:"specification,omitempty"`
	URL				string		`json:"url"`
}


func (cfg *apiConfig)handleOverdriveModes(w http.ResponseWriter, r *http.Request) {
	path := strings.TrimPrefix(r.URL.Path, "/api/overdrive-modes/")
    segments := strings.Split(path, "/")

	// /api/overdrive-modes
	if path == "" {
		cfg.handleOverdriveModesRetrieve(w, r)
		return
	}
	
	switch len(segments) {
	case 1:
		fmt.Printf("%s should trigger the single resource endpoint\n", r.URL.Path)
	default:
		fmt.Printf("%s is in the wrong format and should give an error\n", r.URL.Path)
		respondWithError(w, http.StatusBadRequest, `wrong format. usage: /api/overdrive-modes/{name or id}`, nil)
		return
	}
}


func (cfg *apiConfig) handleOverdriveModesRetrieve(w http.ResponseWriter, r *http.Request) {
	var dbODModes []database.OverdriveMode
	var err error
	modeType := r.URL.Query().Get("type")

	if modeType != "" {
		// need to add type lookup validation
		dbODModes, err = cfg.db.GetOverdriveModesByType(r.Context(), database.OverdriveModeType(modeType))
		if err != nil {
			respondWithError(w, http.StatusBadRequest, "No valid overdrive mode type provided. See /api/overdrive-modes for valid values", err)
			return
		}
	} else {
		dbODModes, err = cfg.db.GetOverdriveModes(context.Background())
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve overdrive modes", err)
			return
		}
	}

	var resources []NamedAPIResource

	for _, dbMode := range dbODModes {
		overdriveMode := NamedAPIResource{
			Name: dbMode.Name,
			URL: cfg.createURL("overdrive-modes", dbMode.ID),
		}

		resources = append(resources, overdriveMode)
	}

	list := NamedApiResourceList{
		Count: len(resources),
		Results: resources,
	}

	respondWithJSON(w, http.StatusOK, list)
}