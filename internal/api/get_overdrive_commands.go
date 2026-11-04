package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func (srv *Server) getOverdriveCommand(r *http.Request, i handlerInput[seedtypes.OverdriveCommand, OverdriveCommand, NamedAPIResource], id int32) (OverdriveCommand, error) {
	command, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return OverdriveCommand{}, err
	}

	overdrives, err := E.Overdrives.dbGetRelations(srv, r.Context(), command, ToIntManyNull(srv.Cfg.DB.GetOverdriveCommandOverdriveIDs))
	if err != nil {
		return OverdriveCommand{}, err
	}

	response := OverdriveCommand{
		ID:          command.ID,
		Name:        command.Name,
		Description: command.Description,
		Rank:        command.Rank,
		User:        nameToNamedAPIResource(srv, E.CharacterClasses, command.User, nil),
		Topmenu:     namePtrToNamedAPIResPtr(srv, E.Topmenus, command.Topmenu, nil),
		OpenSubmenu: nameToNamedAPIResource(srv, E.Submenus, command.OpenSubmenu, nil),
		Overdrives:  overdrives,
	}

	return response, nil
}

func (srv *Server) retrieveOverdriveCommands(r *http.Request, i handlerInput[seedtypes.OverdriveCommand, OverdriveCommand, NamedAPIResource]) ([]int32, error) {
	return i.queryParamsVerifyRetrieve(r)
}
