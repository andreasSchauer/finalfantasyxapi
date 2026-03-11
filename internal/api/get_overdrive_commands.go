package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getOverdriveCommand(r *http.Request, i handlerInput[seeding.OverdriveCommand, OverdriveCommand, NamedAPIResource, NamedApiResourceList], id int32) (OverdriveCommand, error) {
	command, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return OverdriveCommand{}, err
	}
	
	overdrives, err := getResourcesDB(cfg, r, cfg.e.overdrives, command, cfg.db.GetOverdriveCommandOverdriveIDs)
	if err != nil {
		return OverdriveCommand{}, err
	}

	response := OverdriveCommand{
		ID:             command.ID,
		Name:           command.Name,
		Description:    command.Description,
		Rank: 			command.Rank,
		User: 			nameToNamedAPIResource(cfg, cfg.e.characterClasses, command.Name, nil),
		Topmenu: 		namePtrToNamedAPIResPtr(cfg, cfg.e.topmenus, command.Topmenu, nil),
		OpenSubmenu: 	nameToNamedAPIResource(cfg, cfg.e.submenus, command.OpenSubmenu, nil),
		Overdrives: 	overdrives,
	}

	return response, nil
}

func (cfg *Config) retrieveOverdriveCommands(r *http.Request, i handlerInput[seeding.OverdriveCommand, OverdriveCommand, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return i.resToListFunc(cfg, r, resources)
}
