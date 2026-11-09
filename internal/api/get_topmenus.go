package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/assemble"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func getTopmenu(r *http.Request, i handlerinput.Main[seedtypes.Topmenu, apitypes.Topmenu, apiresource.Named], id int32) (apitypes.Topmenu, error) {
	topmenu, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Topmenu{}, err
	}

	rel, err := assemble.GetTopmenuRelationships(r, topmenu)
	if err != nil {
		return apitypes.Topmenu{}, err
	}

	response := apitypes.Topmenu{
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

func retrieveTopmenus(r *http.Request, i handlerinput.Main[seedtypes.Topmenu, apitypes.Topmenu, apiresource.Named]) ([]int32, error) {
	return i.VerifyQueriesRetrieve(r)
}
