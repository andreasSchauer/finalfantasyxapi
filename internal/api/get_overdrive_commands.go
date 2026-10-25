package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getOverdriveCommand(r *http.Request, i handlerInput[seeding.OverdriveCommand, OverdriveCommand, NamedAPIResource, NamedApiResourceList], id int32) (OverdriveCommand, error) {
	command, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return OverdriveCommand{}, err
	}

	overdrives, err := getResourcesDbItem(cfg, r.Context(), cfg.E.Overdrives, command, ToIntManyNull(cfg.db.GetOverdriveCommandOverdriveIDs))
	if err != nil {
		return OverdriveCommand{}, err
	}

	response := OverdriveCommand{
		ID:          command.ID,
		Name:        command.Name,
		Description: command.Description,
		Rank:        command.Rank,
		User:        nameToNamedAPIResource(cfg, cfg.E.CharacterClasses, command.User, nil),
		Topmenu:     namePtrToNamedAPIResPtr(cfg, cfg.E.Topmenus, command.Topmenu, nil),
		OpenSubmenu: nameToNamedAPIResource(cfg, cfg.E.Submenus, command.OpenSubmenu, nil),
		Overdrives:  overdrives,
	}

	return response, nil
}

func (cfg *Config) retrieveOverdriveCommands(r *http.Request, i handlerInput[seeding.OverdriveCommand, OverdriveCommand, NamedAPIResource, NamedApiResourceList]) ([]int32, error) {
	return verifyParamsAndRetrieve(r, i)
}
