package main

import (
	"errors"
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
	i := cfg.e.overdriveModes

	mode, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return OverdriveMode{}, err
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
		actionAmount := nameToResourceAmount(cfg, cfg.e.characters, a.User, nil, a.Amount, newActionAmount)
		actions = append(actions, actionAmount)
	}

	return actions
}


func (cfg *Config) retrieveOverdriveModes(r *http.Request) (NamedApiResourceList, error) {
	i := cfg.e.overdriveModes

	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	filteredLists := []filteredResList[NamedAPIResource]{
		frl(cfg.getOverdriveModesType(r, resources, i)),
	}

	return filterAPIResources(cfg, r, i, resources, filteredLists)
}


// can be generalized, now that I use id lookups
// the extra things I need to pass are the db func, a conversion func, the query param name, the typeLookup, and maybe the err msg
// though I can also return a predefined error from errors.go and write the message under the function call
// need a fourth type, next to handlerInput types, called enumType, or E
// query func signature: func (context.Context, E) ([]int32, error) (put actually maybe I can define them in config_type_lookup.go)
// conversion func signature: func (enumName string) E
// sharedResources is exclusive to this function as it's the only one for this endpoint
// normally, they return after creating the resources
func (cfg *Config) getOverdriveModesType(r *http.Request, inputModes []NamedAPIResource, i handlerInput[seeding.OverdriveMode, OverdriveMode, NamedAPIResource, NamedApiResourceList]) ([]NamedAPIResource, error) {
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

	resources := idsToAPIResources(cfg, i, dbIDs)
	sharedResources := getSharedResources(inputModes, resources)

	return sharedResources, nil
}