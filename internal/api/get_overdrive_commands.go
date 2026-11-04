package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func (s *Server) getOverdriveCommand(r *http.Request, i handlerInput[seedtypes.OverdriveCommand, OverdriveCommand, NamedAPIResource], id int32) (OverdriveCommand, error) {
	command, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return OverdriveCommand{}, err
	}

	overdrives, err := E.Overdrives.dbGetRelations(s, r.Context(), command, ToIntManyNull(s.Cfg.DB.GetOverdriveCommandOverdriveIDs))
	if err != nil {
		return OverdriveCommand{}, err
	}

	response := OverdriveCommand{
		ID:          command.ID,
		Name:        command.Name,
		Description: command.Description,
		Rank:        command.Rank,
		User:        nameToNamedAPIResource(s, E.CharacterClasses, command.User, nil),
		Topmenu:     namePtrToNamedAPIResPtr(s, E.Topmenus, command.Topmenu, nil),
		OpenSubmenu: nameToNamedAPIResource(s, E.Submenus, command.OpenSubmenu, nil),
		Overdrives:  overdrives,
	}

	return response, nil
}

func (s *Server) retrieveOverdriveCommands(r *http.Request, i handlerInput[seedtypes.OverdriveCommand, OverdriveCommand, NamedAPIResource]) ([]int32, error) {
	return i.queryParamsVerifyRetrieve(r)
}
