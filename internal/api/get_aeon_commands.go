package api

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func getAeonCommand(r *http.Request, i handlerinput.HandlerInput[seedtypes.AeonCommand, apitypes.AeonCommand, apiresource.NamedAPIResource], id int32) (apitypes.AeonCommand, error) {
	command, err := i.QueryParamsVerifyGet(r, id)
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
		User:              handlerinput.NameToNamedAPIResource(E.CharacterClasses, userName, nil),
		Topmenu:           handlerinput.NamePtrToNamedAPIResPtr(E.Topmenus, command.Topmenu, nil),
		OpenSubmenu:       handlerinput.NamePtrToNamedAPIResPtr(E.Submenus, command.OpenSubmenu, nil),
		PossibleAbilities: ah.ConvertObjSlice(command.PossibleAbilities, convertPossibleAbilityList),
	}

	return response, nil
}

func retrieveAeonCommands(r *http.Request, i handlerinput.HandlerInput[seedtypes.AeonCommand, apitypes.AeonCommand, apiresource.NamedAPIResource]) ([]int32, error) {
	return i.QueryParamsVerifyRetrieve(r)
}
