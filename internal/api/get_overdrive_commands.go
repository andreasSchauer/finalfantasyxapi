package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func getOverdriveCommand(r *http.Request, i handlerInput[seedtypes.OverdriveCommand, apitypes.OverdriveCommand, apiresource.NamedAPIResource], id int32) (apitypes.OverdriveCommand, error) {
	command, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return apitypes.OverdriveCommand{}, err
	}

	overdrives, err := E.Overdrives.dbGetRelations(r.Context(), command, ToIntManyNull(config.Cfg.DB.GetOverdriveCommandOverdriveIDs))
	if err != nil {
		return apitypes.OverdriveCommand{}, err
	}

	response := apitypes.OverdriveCommand{
		ID:          command.ID,
		Name:        command.Name,
		Description: command.Description,
		Rank:        command.Rank,
		User:        nameToNamedAPIResource(E.CharacterClasses, command.User, nil),
		Topmenu:     namePtrToNamedAPIResPtr(E.Topmenus, command.Topmenu, nil),
		OpenSubmenu: nameToNamedAPIResource(E.Submenus, command.OpenSubmenu, nil),
		Overdrives:  overdrives,
	}

	return response, nil
}

func retrieveOverdriveCommands(r *http.Request, i handlerInput[seedtypes.OverdriveCommand, apitypes.OverdriveCommand, apiresource.NamedAPIResource]) ([]int32, error) {
	return i.queryParamsVerifyRetrieve(r)
}
