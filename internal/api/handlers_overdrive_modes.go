package api

import (
	"fmt"
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

func convertActionAmount(res NamedAPIResource, amount int32) ActionAmount {
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

func (a ActionAmount) GetVersion() *int32 {
	return nil
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

	case 2:
		handleEndpointSubsections(cfg, w, r, i, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", getUsageString(i)), nil)
		return
	}
}

func (cfg *Config) getOverdriveMode(r *http.Request, i handlerInput[seeding.OverdriveMode, OverdriveMode, NamedAPIResource, NamedApiResourceList], id int32) (OverdriveMode, error) {
	mode, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return OverdriveMode{}, err
	}

	modeType, err := newNamedAPIResourceFromType(cfg, cfg.e.overdriveModeType.endpoint, mode.Type, cfg.t.OverdriveModeType)
	if err != nil {
		return OverdriveMode{}, err
	}
	actions := namesToResourceAmounts(cfg, cfg.e.characters, mode.ActionsToLearn, convertActionAmount)

	response := OverdriveMode{
		ID:          mode.ID,
		Name:        mode.Name,
		Description: mode.Description,
		Effect:      mode.Effect,
		Type:        modeType,
		FillRate:    mode.FillRate,
		Actions:     actions,
	}

	return response, nil
}

func (cfg *Config) retrieveOverdriveModes(r *http.Request, i handlerInput[seeding.OverdriveMode, OverdriveMode, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[NamedAPIResource]{
		frl(typeQuery(cfg, r, i, cfg.t.OverdriveModeType, resources, "type", cfg.db.GetOverdriveModeIDsByType)),
	})
}
