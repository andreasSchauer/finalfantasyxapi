package api

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func getAeonCommand(r *http.Request, i handlerinput.Main[seedtypes.AeonCommand, apitypes.AeonCommand, apiresource.Named], id int32) (apitypes.AeonCommand, error) {
	command, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.AeonCommand{}, err
	}

	var userName string

	if len(command.PossibleAbilities) == 3 {
		userName = "magus-sisters"
	} else {
		userName = command.PossibleAbilities[0].User
	}

	response := apitypes.AeonCommand{
		ID:                command.ID,
		Name:              command.Name,
		Description:       command.Description,
		Effect:            command.Effect,
		Cursor:            command.Cursor,
		User:              handlerinput.NameToNamedAPIResource(apilookups.E.CharacterClasses, userName, nil),
		Topmenu:           handlerinput.NamePtrToNamedAPIResPtr(apilookups.E.Topmenus, command.Topmenu, nil),
		OpenSubmenu:       handlerinput.NamePtrToNamedAPIResPtr(apilookups.E.Submenus, command.OpenSubmenu, nil),
		PossibleAbilities: ah.ConvertObjSlice(command.PossibleAbilities, ConvertPossibleAbilityList),
	}

	return response, nil
}

func retrieveAeonCommands(r *http.Request, i handlerinput.Main[seedtypes.AeonCommand, apitypes.AeonCommand, apiresource.Named]) ([]int32, error) {
	return i.VerifyQueriesRetrieve(r)
}
