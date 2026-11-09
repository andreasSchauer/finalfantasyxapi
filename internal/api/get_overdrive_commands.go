package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func getOverdriveCommand(r *http.Request, i handlerinput.Main[seedtypes.OverdriveCommand, apitypes.OverdriveCommand, apiresource.Named], id int32) (apitypes.OverdriveCommand, error) {
	command, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.OverdriveCommand{}, err
	}

	overdrives, err := apilookups.E.Overdrives.DBGetRelations(r.Context(), command, handlerinput.ToIntManyNull(config.Cfg.DB.GetOverdriveCommandOverdriveIDs))
	if err != nil {
		return apitypes.OverdriveCommand{}, err
	}

	response := apitypes.OverdriveCommand{
		ID:          command.ID,
		Name:        command.Name,
		Description: command.Description,
		Rank:        command.Rank,
		User:        handlerinput.NameToNamedAPIResource(apilookups.E.CharacterClasses, command.User, nil),
		Topmenu:     handlerinput.NamePtrToNamedAPIResPtr(apilookups.E.Topmenus, command.Topmenu, nil),
		OpenSubmenu: handlerinput.NameToNamedAPIResource(apilookups.E.Submenus, command.OpenSubmenu, nil),
		Overdrives:  overdrives,
	}

	return response, nil
}

func retrieveOverdriveCommands(r *http.Request, i handlerinput.Main[seedtypes.OverdriveCommand, apitypes.OverdriveCommand, apiresource.Named]) ([]int32, error) {
	return i.VerifyQueriesRetrieve(r)
}
