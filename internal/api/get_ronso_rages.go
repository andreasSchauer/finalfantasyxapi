package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func getRonsoRage(r *http.Request, i handlerInput[seedtypes.RonsoRage, apitypes.RonsoRage, apiresource.NamedAPIResource], id int32) (apitypes.RonsoRage, error) {
	ronsoRage, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return apitypes.RonsoRage{}, err
	}

	monsters, err := E.Monsters.dbGetRelations(r.Context(), ronsoRage, config.Cfg.DB.GetRonsoRageMonsterIDs)
	if err != nil {
		return apitypes.RonsoRage{}, err
	}

	response := apitypes.RonsoRage{
		ID:          ronsoRage.ID,
		Name:        ronsoRage.Name,
		Description: ronsoRage.Description,
		Effect:      ronsoRage.Effect,
		Overdrive:   nameToNamedAPIResource(E.Overdrives, ronsoRage.Name, nil),
		Monsters:    monsters,
	}

	return response, nil
}

func retrieveRonsoRages(r *http.Request, i handlerInput[seedtypes.RonsoRage, apitypes.RonsoRage, apiresource.NamedAPIResource]) ([]int32, error) {
	return i.queryParamsVerifyRetrieve(r)
}
