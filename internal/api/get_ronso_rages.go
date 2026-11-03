package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func (cfg *Config) getRonsoRage(r *http.Request, i handlerInput[seedtypes.RonsoRage, RonsoRage, NamedAPIResource], id int32) (RonsoRage, error) {
	ronsoRage, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return RonsoRage{}, err
	}

	monsters, err := cfg.E.Monsters.dbGetRelations(cfg, r.Context(), ronsoRage, cfg.db.GetRonsoRageMonsterIDs)
	if err != nil {
		return RonsoRage{}, err
	}

	response := RonsoRage{
		ID:          ronsoRage.ID,
		Name:        ronsoRage.Name,
		Description: ronsoRage.Description,
		Effect:      ronsoRage.Effect,
		Overdrive:   nameToNamedAPIResource(cfg, cfg.E.Overdrives, ronsoRage.Name, nil),
		Monsters:    monsters,
	}

	return response, nil
}

func (cfg *Config) retrieveRonsoRages(r *http.Request, i handlerInput[seedtypes.RonsoRage, RonsoRage, NamedAPIResource]) ([]int32, error) {
	return i.queryParamsVerifyRetrieve(r)
}
