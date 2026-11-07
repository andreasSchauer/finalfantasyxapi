package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func getRonsoRage(r *http.Request, i handlerinput.HandlerInput[seedtypes.RonsoRage, apitypes.RonsoRage, apiresource.NamedAPIResource], id int32) (apitypes.RonsoRage, error) {
	ronsoRage, err := i.QueryParamsVerifyGet(r, id)
	if err != nil {
		return apitypes.RonsoRage{}, err
	}

	monsters, err := E.Monsters.DBGetRelations(r.Context(), ronsoRage, config.Cfg.DB.GetRonsoRageMonsterIDs)
	if err != nil {
		return apitypes.RonsoRage{}, err
	}

	response := apitypes.RonsoRage{
		ID:          ronsoRage.ID,
		Name:        ronsoRage.Name,
		Description: ronsoRage.Description,
		Effect:      ronsoRage.Effect,
		Overdrive:   handlerinput.NameToNamedAPIResource(E.Overdrives, ronsoRage.Name, nil),
		Monsters:    monsters,
	}

	return response, nil
}

func retrieveRonsoRages(r *http.Request, i handlerinput.HandlerInput[seedtypes.RonsoRage, apitypes.RonsoRage, apiresource.NamedAPIResource]) ([]int32, error) {
	return i.QueryParamsVerifyRetrieve(r)
}
