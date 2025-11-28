package main

import (
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
	Actions		[]OverdriveModeAction	`json:"actions"`
}

type OverdriveModeAction struct {
	User		NamedAPIResource	`json:"user"`
	Amount		int32				`json:"amount"`
	
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
		// /api/overdrive-modes/{name or id}
		segment := segments[0]
		cfg.handleOverdriveModeGet(w, r, segment)
		return
	default:
		respondWithError(w, http.StatusBadRequest, `Wrong format. Usage: /api/overdrive-modes/{name or id}`, nil)
		return
	}
}


func (cfg *apiConfig) handleOverdriveModeGet(w http.ResponseWriter, r *http.Request, segment string) {
	id, err := parseSingleSegmentResource(segment, cfg.l.OverdriveModes)
	if err != nil {
		if httpErr, ok := err.(httpError); ok {
        respondWithError(w, httpErr.code, httpErr.msg, httpErr.err)
		return
		}
	}

	dbMode, err := cfg.db.GetOverdriveMode(r.Context(), id)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't get Overdrive Mode", err)
		return
	}

	actions, err := cfg.getOverdriveModeActions(r, dbMode.ID)
	if err != nil {
		if httpErr, ok := err.(httpError); ok {
			respondWithError(w, httpErr.code, httpErr.msg, httpErr.err)
			return
		}
	}

	response := OverdriveMode{
		ID: 			dbMode.ID,
		Name: 			dbMode.Name,
		Description: 	dbMode.Description,
		Effect: 		dbMode.Effect,
		Type: 			string(dbMode.Type),
		FillRate: 		anyToFloat32Ptr(dbMode.FillRate),
		Actions: 		actions,
	}

	respondWithJSON(w, http.StatusOK, response)
}


func (cfg *apiConfig) getOverdriveModeActions(r *http.Request, id int32) ([]OverdriveModeAction, error) {
	dbActions, err := cfg.db.GetOverdriveModeActions(r.Context(), id)
	if err != nil {
		return nil, NewHTTPError(http.StatusInternalServerError, "Couldn't get Overdrive Mode Actions", err)
	}

	actions := []OverdriveModeAction{}

	for _, dbAction := range dbActions {
		user := NamedAPIResource{
			Name: 	dbAction.Character.String,
			URL: 	cfg.createURL("characters", dbAction.UserID),
		}

		action := OverdriveModeAction{
			User: 	user,
			Amount: dbAction.Amount,
		}

		actions = append(actions, action)
	}

	return actions, nil
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
		dbODModes, err = cfg.db.GetOverdriveModes(r.Context())
		if err != nil {
			respondWithError(w, http.StatusInternalServerError, "Couldn't retrieve overdrive modes", err)
			return
		}
	}

	// I swear I can generalize this, if I can find a way to get to the name and id fields
	// could create an interface, but that is a last resort
	var resources []NamedAPIResource

	for _, dbMode := range dbODModes {
		overdriveMode := NamedAPIResource{
			Name: 	dbMode.Name,
			URL: 	cfg.createURL("overdrive-modes", dbMode.ID),
		}

		resources = append(resources, overdriveMode)
	}

	resourceList := NamedApiResourceList{
		Count: 		len(resources),
		Results: 	resources,
	}

	respondWithJSON(w, http.StatusOK, resourceList)
}