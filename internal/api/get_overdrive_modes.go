package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)


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
