package main

import (
	"net/http"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type OverdriveMode struct {
	ID          int32          `json:"id"`
	Name        string         `json:"name"`
	Description string         `json:"description"`
	Effect      string         `json:"effect"`
	Type        string         `json:"type"`
	FillRate    *float32       `json:"fill_rate,omitempty"`
	Actions     []ActionAmount `json:"actions"`
}

type ActionAmount struct {
	User   NamedAPIResource `json:"user"`
	Amount int32            `json:"amount"`
}

func (cfg *apiConfig) handleOverdriveModes(w http.ResponseWriter, r *http.Request) {
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
		input, err := parseSingleSegmentResource(segment, cfg.l.OverdriveModes)
		if handleHTTPError(w, err) {
			return
		}

		cfg.handleOverdriveModeGet(w, r, input)
		return
	default:
		respondWithError(w, http.StatusBadRequest, `Wrong format. Usage: /api/overdrive-modes/{name or id}`, nil)
		return
	}
}

func (cfg *apiConfig) handleOverdriveModeGet(w http.ResponseWriter, r *http.Request, input parseResponse) {
	dbMode, err := cfg.db.GetOverdriveMode(r.Context(), input.ID)
	if err != nil {
		respondWithError(w, http.StatusNotFound, "Couldn't get Overdrive Mode. Overdrive mode with this ID doesn't exist.", err)
		return
	}

	actions, err := cfg.getOverdriveModeActions(r, dbMode.ID)
	if handleHTTPError(w, err) {
		return
	}

	response := OverdriveMode{
		ID:          dbMode.ID,
		Name:        dbMode.Name,
		Description: dbMode.Description,
		Effect:      dbMode.Effect,
		Type:        string(dbMode.Type),
		FillRate:    anyToFloat32Ptr(dbMode.FillRate),
		Actions:     actions,
	}

	respondWithJSON(w, http.StatusOK, response)
}

func (cfg *apiConfig) getOverdriveModeActions(r *http.Request, id int32) ([]ActionAmount, error) {
	dbActions, err := cfg.db.GetOverdriveModeActions(r.Context(), id)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "Couldn't get Overdrive Mode Actions", err)
	}

	actions := []ActionAmount{}

	for _, dbAction := range dbActions {
		action := ActionAmount{
			User:   cfg.newNamedAPIResourceSimple("characters", dbAction.UserID, dbAction.Character.String),
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

	// this can potentially be made generic, but I'm not sure yet
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

	resources := createNamedAPIResourcesSimple(cfg, dbODModes, "overdrive-modes", func(mode database.OverdriveMode) (int32, string) {
		return mode.ID, mode.Name
	})

	resourceList, err := cfg.newNamedAPIResourceList(r, resources)
	if handleHTTPError(w, err) {
		return
	}

	respondWithJSON(w, http.StatusOK, resourceList)
}
