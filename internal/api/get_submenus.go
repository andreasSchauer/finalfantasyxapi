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

	abilities, err := getResourcesDbItem(cfg, r, cfg.e.abilities, submenu, NullToIntMany(cfg.db.GetSubmenuAbilityIDs))
	if err != nil {
		return Submenu{}, err
	}

	menuOpen, err := createSubmenuOpenedBy(cfg, r, submenu)
	if err != nil {
		return Submenu{}, err
	}

	response := Submenu{
		ID:          submenu.ID,
		Name:        submenu.Name,
		Description: submenu.Description,
		Effect:      submenu.Effect,
		Topmenu:     namePtrToNamedAPIResPtr(cfg, cfg.e.topmenus, submenu.Topmenu, nil),
		Users:       namesToNamedAPIResources(cfg, cfg.e.characterClasses, submenu.Users),
		OpenedBy:    menuOpen,
		Abilities:   abilities,
	}

	return response, nil
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
