package main

import (
	"errors"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type OverdriveMode struct {
	ID          int32          		`json:"id"`
	Name        string         		`json:"name"`
	Description string         		`json:"description"`
	Effect      string         		`json:"effect"`
	Type        NamedAPIResource    `json:"type"`
	FillRate    *float32       		`json:"fill_rate,omitempty"`
	Actions     []ActionAmount 		`json:"actions"`
}

type ActionAmount struct {
	User   NamedAPIResource `json:"user"`
	Amount int32            `json:"amount"`
}

func (a ActionAmount) GetAPIResource() IsAPIResource {
	return a.User
}

func (a ActionAmount) GetName() string {
	return a.User.Name
}

func (a ActionAmount) GetVal() int32 {
	return a.Amount
}

func (cfg *Config) HandleOverdriveModes(w http.ResponseWriter, r *http.Request) {
	i := cfg.e.overdriveModes

	segments := getPathSegments(r.URL.Path, i.endpoint)

	switch len(segments) {
	case 0:
		handleEndpointList(w, r, i)
		return

	case 1:
		handleEndpointNameOrID(cfg, w, r, i, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, `Wrong format. Usage: /api/overdrive-modes/{name or id}`, nil)
		return
	}
}

func (cfg *Config) getOverdriveMode(r *http.Request, id int32) (OverdriveMode, error) {
	endpoint := cfg.e.overdriveModes.endpoint
	
	err := verifyQueryParams(r, endpoint, &id, cfg.q.overdriveModes)
	if err != nil {
		return OverdriveMode{}, err
	}

	dbMode, err := cfg.db.GetOverdriveMode(r.Context(), id)
	if err != nil {
		return OverdriveMode{}, newHTTPError(http.StatusNotFound, "Couldn't get Overdrive Mode. Overdrive mode with this ID doesn't exist.", err)
	}

	modeType, err := cfg.newNamedAPIResourceFromType(cfg.e.overdriveModeType.endpoint, string(dbMode.Type), cfg.t.OverdriveModeType)
	if err != nil {
		return OverdriveMode{}, err
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
		Type:        modeType,
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
			User:   cfg.newNamedAPIResourceSimple(cfg.e.characters.endpoint, dbAction.UserID, dbAction.Character.String),
			Amount: dbAction.Amount,
		}

		actions = append(actions, action)
	}

	return actions, nil
}

func (cfg *Config) retrieveOverdriveModes(r *http.Request) (NamedApiResourceList, error) {
	endpoint := cfg.e.overdriveModes.endpoint

	err := verifyQueryParams(r, endpoint, nil, cfg.q.overdriveModes)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	dbODModes, err := cfg.db.GetOverdriveModes(r.Context())
	if err != nil {
		return NamedApiResourceList{}, newHTTPError(http.StatusInternalServerError, "Couldn't retrieve overdrive modes", err)
	}

	resources := createNamedAPIResourcesSimple(cfg, dbODModes, endpoint, func(mode database.OverdriveMode) (int32, string) {
		return mode.ID, mode.Name
	})

	resources, err = cfg.getOverdriveModesType(r, endpoint, resources)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	resourceList, err := cfg.newNamedAPIResourceList(r, resources)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return resourceList, nil
}

func (cfg *Config) getOverdriveModesType(r *http.Request, endpoint string, inputModes []NamedAPIResource) ([]NamedAPIResource, error) {
	queryParam := cfg.q.overdriveModes["type"]
	enum, err := parseTypeQuery(r, queryParam, cfg.t.OverdriveModeType)
	if errors.Is(err, errEmptyQuery) {
		return inputModes, nil
	}
	if err != nil {
		return nil, err
	}

	modeType := database.OverdriveModeType(enum.Name)

	dbODModes, err := cfg.db.GetOverdriveModesByType(r.Context(), modeType)
	if err != nil {
		return nil, newHTTPError(http.StatusBadRequest, "No valid overdrive mode type provided. See /api/overdrive-modes for valid values", err)
	}

	resources := createNamedAPIResourcesSimple(cfg, dbODModes, endpoint, func(mode database.OverdriveMode) (int32, string) {
		return mode.ID, mode.Name
	})

	sharedResources := getSharedResources(inputModes, resources)

	return sharedResources, nil
}
