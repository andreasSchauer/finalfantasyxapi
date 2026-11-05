package api

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api_helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func getAeonCommand(r *http.Request, i handlerInput[seedtypes.AeonCommand, AeonCommand, NamedAPIResource], id int32) (AeonCommand, error) {
	command, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return AeonCommand{}, err
	}

	var userName string

	if len(command.PossibleAbilities) == 3 {
		userName = "magus-sisters"
	} else {
		userName = command.PossibleAbilities[0].User
	}

	response := AeonCommand{
		ID:                command.ID,
		Name:              command.Name,
		Description:       command.Description,
		Effect:            command.Effect,
		Cursor:            command.Cursor,
		User:              nameToNamedAPIResource(E.CharacterClasses, userName, nil),
		Topmenu:           namePtrToNamedAPIResPtr(E.Topmenus, command.Topmenu, nil),
		OpenSubmenu:       namePtrToNamedAPIResPtr(E.Submenus, command.OpenSubmenu, nil),
		PossibleAbilities: ah.ConvertObjSlice(command.PossibleAbilities, convertPossibleAbilityList),
	}

	return response, nil
}

func retrieveAeonCommands(r *http.Request, i handlerInput[seedtypes.AeonCommand, AeonCommand, NamedAPIResource]) ([]int32, error) {
	return i.queryParamsVerifyRetrieve(r)
}
