package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type OverdriveMode struct {
	ID          int32            `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Effect      string           `json:"effect"`
	Type        NamedAPIResource `json:"type"`
	FillRate    *float32         `json:"fill_rate,omitempty"`
	Actions     []ActionAmount   `json:"actions"`
}

type ActionAmount struct {
	User   NamedAPIResource `json:"user"`
	Amount int32            `json:"amount"`
}

func newActionAmount(res NamedAPIResource, amount int32) ActionAmount {
	return ActionAmount{
		User:   res,
		Amount: amount,
	}
}

func (a ActionAmount) GetAPIResource() APIResource {
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
		respondWithError(w, http.StatusBadRequest, `wrong format. usage: '/api/overdrive-modes/{name or id}'.`, nil)
		return
	}
}


func (cfg *Config) getOverdriveMode(r *http.Request, id int32) (OverdriveMode, error) {
	endpoint := cfg.e.overdriveModes.endpoint

	err := verifyQueryParams(r, endpoint, &id, cfg.q.overdriveModes)
	if err != nil {
		return OverdriveMode{}, err
	}

	mode, err := seeding.GetResourceByID(id, cfg.l.OverdriveModesID)
	if err != nil {
		return OverdriveMode{}, newHTTPError(http.StatusNotFound, fmt.Sprintf("overdrive mode with id '%d' doesn't exist.", id), err)
	}

	modeType, err := cfg.newNamedAPIResourceFromType(cfg.e.overdriveModeType.endpoint, mode.Type, cfg.t.OverdriveModeType)
	if err != nil {
		return OverdriveMode{}, err
	}

	response := OverdriveMode{
		ID:          mode.ID,
		Name:        mode.Name,
		Description: mode.Description,
		Effect:      mode.Effect,
		Type:        modeType,
		FillRate:    mode.FillRate,
		Actions:     cfg.getOverdriveModeActions(mode),
	}

	return response, nil
}


func (cfg *Config) getOverdriveModeActions(mode seeding.OverdriveMode) []ActionAmount {
	actions := []ActionAmount{}

	for _, a := range mode.ActionsToLearn {
		actionAmount := nameToResourceAmount(cfg, cfg.e.characters.endpoint, a.User, nil, a.Amount, cfg.l.Characters, cfg.l.CharactersID, newActionAmount)

		actions = append(actions, actionAmount)
	}

	return actions
}


// this can potentially be generalized, if I use a loop for the filter funcs
// and all of the funcs keep the signature of (*http.Request, []APIResource, handlerInput) ([]APIResource, error)
func (cfg *Config) retrieveOverdriveModes(r *http.Request) (NamedApiResourceList, error) {
	i := cfg.e.overdriveModes

	err := verifyQueryParams(r, i.endpoint, nil, cfg.q.overdriveModes)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	dbIDs, err := cfg.db.GetOverdriveModeIDs(r.Context())
	if err != nil {
		return NamedApiResourceList{}, newHTTPError(http.StatusInternalServerError, "couldn't retrieve overdrive modes.", err)
	}

	resources := idsToNamedAPIResources(cfg, i.endpoint, dbIDs, i.objLookupID)

	resources, err = cfg.getOverdriveModesType(r, resources, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	resourceList, err := newNamedAPIResourceList(cfg, r, i, resources)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return resourceList, nil
}


// even this might be able to be generalized, now that I use id lookups
// the extra things I need to pass are the db func, the query param name, the typeLookup, and maybe the err msg
// query func signature: [T any](context.Context, T) ([]int32, error)
// potential problem might be the type casting of enum.Name
// for all of this I would probably make a struct as input
func (cfg *Config) getOverdriveModesType(r *http.Request, inputModes []NamedAPIResource, i handlerInput[seeding.OverdriveMode, OverdriveMode, NamedApiResourceList]) ([]NamedAPIResource, error) {
	queryParam := i.queryLookup["type"]
	enum, err := parseTypeQuery(r, queryParam, cfg.t.OverdriveModeType)
	if errors.Is(err, errEmptyQuery) {
		return inputModes, nil
	}
	if err != nil {
		return nil, err
	}

	modeType := database.OverdriveModeType(enum.Name)

	dbIDs, err := cfg.db.GetOverdriveModeIDsByType(r.Context(), modeType)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, "couldn't get overdrive modes by type.", err)
	}

	resources := idsToNamedAPIResources(cfg, i.endpoint, dbIDs, i.objLookupID)
	sharedResources := getSharedResources(inputModes, resources)

	return sharedResources, nil
}
