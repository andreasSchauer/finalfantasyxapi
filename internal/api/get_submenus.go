package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getSubmenu(r *http.Request, i handlerInput[seeding.Submenu, Submenu, NamedAPIResource, NamedApiResourceList], id int32) (Submenu, error) {
	submenu, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return Submenu{}, err
	}
	
	abilities, err := getAbilityResourcesDbNullable(cfg, r, submenu, cfg.db.GetSubmenuAbilityIDs)
	if err != nil {
		return Submenu{}, err
	}

	menuOpen, err := createSubmenuOpenedBy(cfg, r, submenu)
	if err != nil {
		return Submenu{}, err
	}

	response := Submenu{
		ID:             submenu.ID,
		Name:           submenu.Name,
		Description:    submenu.Description,
		Effect: 		submenu.Effect,
		Topmenu: 		namePtrToNamedAPIResPtr(cfg, cfg.e.topmenus, submenu.Topmenu, nil),
		Users: 			namesToNamedAPIResources(cfg, cfg.e.characterClasses, submenu.Users),
		OpenedBy: 		menuOpen,
		Abilities: 		abilities,
	}

	return response, nil
}

func createSubmenuOpenedBy(cfg *Config, r *http.Request, submenu seeding.Submenu) (*MenuOpen, error) {
	ability, err := getAbilityResPtrDB(cfg, r, submenu, cfg.db.GetSubmenuOpenedByAbilityID)
	if err != nil {
		return nil, err
	}

	aeonCommand, err := getResPtrDbNullable(cfg, r, cfg.e.aeonCommands, submenu, cfg.db.GetSubmenuOpenedByAeonCommandID)
	if err != nil {
		return nil, err
	}

	overdriveCommands, err := getResourcesDB(cfg, r, cfg.e.overdriveCommands, submenu, cfg.db.GetSubmenuOpenedByOverdriveCommandIDs)
	if err != nil {
		return nil, err
	}

	menuOpen := MenuOpen{
			Ability: 			ability,
			AeonCommand: 		aeonCommand,
			OverdriveCommands: 	overdriveCommands,
	}

	if menuOpen.IsZero() {
		return nil, nil
	}

	return &menuOpen, nil
}

func (cfg *Config) retrieveSubmenus(r *http.Request, i handlerInput[seeding.Submenu, Submenu, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[NamedAPIResource]{
		frl(nameOrIdQuery(cfg, r, i, resources, "topmenu", cfg.e.topmenus.resourceType, cfg.l.Topmenus, cfg.db.GetTopmenuSubmenuIDs)),
	})
}
