package main

import (
	"fmt"
	"net/http"

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

func (cfg *Config) HandleOverdriveModes(w http.ResponseWriter, r *http.Request) {
	segments := getPathSegments(r.URL.Path, "overdrive-modes")

	switch len(segments) {
	case 0:
		// /api/overdrive-modes
		resourceList, err := cfg.retrieveOverdriveModes(r)
		if handleHTTPError(w, err) {
			return
		}
		respondWithJSON(w, http.StatusOK, resourceList)
		return

	case 1:
		// /api/overdrive-modes/{name or id}
		segment := segments[0]
		input, err := parseSingleSegmentResource("overdrive-mode", segment, cfg.l.OverdriveModes)
		if handleHTTPError(w, err) {
			return
		}

		overdriveMode, err := cfg.getOverdriveMode(r, input.ID)
		if handleHTTPError(w, err) {
			return
		}
		respondWithJSON(w, http.StatusOK, overdriveMode)
		return

	default:
		respondWithError(w, http.StatusBadRequest, `Wrong format. Usage: /api/overdrive-modes/{name or id}`, nil)
		return
	}
}

func (cfg *Config) getOverdriveMode(r *http.Request, id int32) (OverdriveMode, error) {
	dbMode, err := cfg.db.GetOverdriveMode(r.Context(), id)
	if err != nil {
		return OverdriveMode{}, newHTTPError(http.StatusNotFound, "Couldn't get Overdrive Mode. Overdrive mode with this ID doesn't exist.", err)
	}

	actions, err := cfg.getOverdriveModeActions(r, dbMode.ID)
	if err != nil {
		return OverdriveMode{}, err
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

	return response, nil
}

func (cfg *Config) getOverdriveModeActions(r *http.Request, id int32) ([]ActionAmount, error) {
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

func (cfg *Config) retrieveOverdriveModes(r *http.Request) (NamedApiResourceList, error) {
	dbODModes, err := cfg.db.GetOverdriveModes(r.Context())
	if err != nil {
		return NamedApiResourceList{}, newHTTPError(http.StatusInternalServerError, "Couldn't retrieve overdrive modes", err)
	}

	resources := createNamedAPIResourcesSimple(cfg, dbODModes, "overdrive-modes", func(mode database.OverdriveMode) (int32, string) {
		return mode.ID, mode.Name
	})

	resources, err = cfg.getOverdriveModesType(r, resources)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	resourceList, err := cfg.newNamedAPIResourceList(r, resources)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return resourceList, nil
}

func (cfg *Config) getOverdriveModesType(r *http.Request, inputModes []NamedAPIResource) ([]NamedAPIResource, error) {
	query := r.URL.Query().Get("type")

	if query == "" {
		return inputModes, nil
	}

	enum, err := GetEnumType(query, cfg.t.OverdriveModeType)
	if err != nil {
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value: %s, use /api/overdrive-mode-type to see valid values.", query), err)
	}

	modeType := database.OverdriveModeType(enum.Name)

	dbODModes, err := cfg.db.GetOverdriveModesByType(r.Context(), modeType)
	if err != nil {
		return nil, newHTTPError(http.StatusBadRequest, "No valid overdrive mode type provided. See /api/overdrive-modes for valid values", err)
	}

	resources := createNamedAPIResourcesSimple(cfg, dbODModes, "overdrive-modes", func(mode database.OverdriveMode) (int32, string) {
		return mode.ID, mode.Name
	})

	sharedResources := getSharedResources(inputModes, resources)

	return sharedResources, nil
}
