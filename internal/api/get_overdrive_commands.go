package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func (cfg *Config) getOverdriveCommand(r *http.Request, i handlerInput[seedtypes.OverdriveCommand, OverdriveCommand, NamedAPIResource], id int32) (OverdriveCommand, error) {
	command, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return OverdriveCommand{}, err
	}

	overdrives, err := E.Overdrives.dbGetRelations(cfg, r.Context(), command, ToIntManyNull(cfg.db.GetOverdriveCommandOverdriveIDs))
	if err != nil {
		return OverdriveCommand{}, err
	}

	response := OverdriveCommand{
		ID:          command.ID,
		Name:        command.Name,
		Description: command.Description,
		Rank:        command.Rank,
		User:        nameToNamedAPIResource(cfg, E.CharacterClasses, command.User, nil),
		Topmenu:     namePtrToNamedAPIResPtr(cfg, E.Topmenus, command.Topmenu, nil),
		OpenSubmenu: nameToNamedAPIResource(cfg, E.Submenus, command.OpenSubmenu, nil),
		Overdrives:  overdrives,
	}

	return response, nil
}

func (cfg *Config) retrieveOverdriveCommands(r *http.Request, i handlerInput[seedtypes.OverdriveCommand, OverdriveCommand, NamedAPIResource]) ([]int32, error) {
	return i.queryParamsVerifyRetrieve(r)
}
