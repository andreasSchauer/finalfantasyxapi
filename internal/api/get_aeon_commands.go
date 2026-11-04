package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func (s *Server) getAeonCommand(r *http.Request, i handlerInput[seedtypes.AeonCommand, AeonCommand, NamedAPIResource], id int32) (AeonCommand, error) {
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
		User:              nameToNamedAPIResource(s, E.CharacterClasses, userName, nil),
		Topmenu:           namePtrToNamedAPIResPtr(s, E.Topmenus, command.Topmenu, nil),
		OpenSubmenu:       namePtrToNamedAPIResPtr(s, E.Submenus, command.OpenSubmenu, nil),
		PossibleAbilities: convertObjSlice(s, command.PossibleAbilities, convertPossibleAbilityList),
	}

	return response, nil
}

func (s *Server) retrieveAeonCommands(r *http.Request, i handlerInput[seedtypes.AeonCommand, AeonCommand, NamedAPIResource]) ([]int32, error) {
	return i.queryParamsVerifyRetrieve(r)
}
