package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getTopmenu(r *http.Request, i handlerInput[seeding.Topmenu, Topmenu, NamedAPIResource, NamedApiResourceList], id int32) (Topmenu, error) {
	topmenu, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return Topmenu{}, err
	}

	submenus, err := getResourcesDbItem(cfg, r, cfg.e.submenus, topmenu, cfg.db.GetTopmenuSubmenuIDs)
	if err != nil {
		return Topmenu{}, err
	}

	overdriveCommands, err := getResourcesDbItem(cfg, r, cfg.e.overdriveCommands, topmenu, cfg.db.GetTopmenuOverdriveCommandIDs)
	if err != nil {
		return Topmenu{}, err
	}

	overdrives, err := getResourcesDbItem(cfg, r, cfg.e.overdrives, topmenu, cfg.db.GetTopmenuOverdriveIDs)
	if err != nil {
		return Topmenu{}, err
	}

	aeonCommands, err := getResourcesDbItem(cfg, r, cfg.e.aeonCommands, topmenu, cfg.db.GetTopmenuAeonCommandIDs)
	if err != nil {
		return Topmenu{}, err
	}

	abilities, err := getResourcesDbItem(cfg, r, cfg.e.abilities, topmenu, NullToIntMany(cfg.db.GetTopmenuAbilityIDs))
	if err != nil {
		return Topmenu{}, err
	}

	response := Topmenu{
		ID:                topmenu.ID,
		Name:              topmenu.Name,
		Submenus:          submenus,
		OverdriveCommands: overdriveCommands,
		Overdrives:        overdrives,
		AeonCommands:      aeonCommands,
		Abilities:         abilities,
	}

	return response, nil
}

func (cfg *Config) retrieveTopmenus(r *http.Request, i handlerInput[seeding.Topmenu, Topmenu, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return i.resToListFunc(cfg, r, resources)
}
