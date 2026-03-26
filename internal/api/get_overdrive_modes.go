package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getOverdriveMode(r *http.Request, i handlerInput[seeding.OverdriveMode, OverdriveMode, NamedAPIResource, NamedApiResourceList], id int32) (OverdriveMode, error) {
	mode, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return OverdriveMode{}, err
	}

	response := OverdriveMode{
		ID:          mode.ID,
		Name:        mode.Name,
		Description: mode.Description,
		Effect:      mode.Effect,
		Type:        mode.Type,
		FillRate:    mode.FillRate,
		Actions:     nameAmtsToResAmts(cfg, cfg.e.characters, mode.ActionsToLearn),
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
