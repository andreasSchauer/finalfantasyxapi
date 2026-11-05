package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func getTopmenu(r *http.Request, i handlerInput[seedtypes.Topmenu, Topmenu, apiresource.NamedAPIResource], id int32) (Topmenu, error) {
	topmenu, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Topmenu{}, err
	}

	rel, err := getTopmenuRelationships(r, topmenu)
	if err != nil {
		return Topmenu{}, err
	}

	response := Topmenu{
		ID:                topmenu.ID,
		Name:              topmenu.Name,
		Submenus:          rel.Submenus,
		OverdriveCommands: rel.OverdriveCommands,
		Overdrives:        rel.Overdrives,
		AeonCommands:      rel.AeonCommands,
		Abilities:         rel.Abilities,
	}

	return response, nil
}

func retrieveTopmenus(r *http.Request, i handlerInput[seedtypes.Topmenu, Topmenu, apiresource.NamedAPIResource]) ([]int32, error) {
	return i.queryParamsVerifyRetrieve(r)
}
