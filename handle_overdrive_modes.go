package main

import (
	"net/http"

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

	modeType, err := newNamedAPIResourceFromType(cfg, cfg.e.overdriveModeType.endpoint, mode.Type, cfg.t.OverdriveModeType)
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
		frl(typeQuery(cfg, r, i, cfg.t.OverdriveModeType, resources, "type", cfg.db.GetOverdriveModeIDsByType)),
	}

	return filterAPIResources(cfg, r, i, resources, filteredLists)
}