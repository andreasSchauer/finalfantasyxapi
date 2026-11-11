package get

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func OverdriveCommand(r *http.Request, i handlerinput.Main[seedtypes.OverdriveCommand, apitypes.OverdriveCommand, apiresource.Named], id int32) (apitypes.OverdriveCommand, error) {
	command, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.OverdriveCommand{}, err
	}

	rel, err := getOverdriveCommandRelationships(r, command)
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
		Overdrives:  rel.Overdrives,
	}

	return response, nil
}

func getOverdriveCommandRelationships(r *http.Request, command seedtypes.OverdriveCommand) (apitypes.OverdriveCommand, error) {
	var rel apitypes.OverdriveCommand
	var err error

	rel.Overdrives, err = apilookups.E.Overdrives.DBGetRelations(r.Context(), command, handlerinput.ToIntManyNull(config.Cfg.DB.GetOverdriveCommandOverdriveIDs))
	if err != nil {
		return apitypes.OverdriveCommand{}, err
	}

	return rel, nil
}
