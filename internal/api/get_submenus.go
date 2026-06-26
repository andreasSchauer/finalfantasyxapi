package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getSubmenu(r *http.Request, i handlerInput[seeding.Submenu, Submenu, NamedAPIResource, NamedApiResourceList], id int32) (Submenu, error) {
	submenu, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return Submenu{}, err
	}

	rel, err := getSubmenuRelationships(cfg, r, submenu)
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
		OpenedBy:    rel.OpenedBy,
		Abilities:   rel.Abilities,
	}

	return response, nil
}

func (cfg *Config) retrieveSubmenus(r *http.Request, i handlerInput[seeding.Submenu, Submenu, NamedAPIResource, NamedApiResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []IdFilter{
		nameIdQuery(r, i, ids, qpnTopmenu, cfg.e.topmenus.resTypeSing, cfg.l.Topmenus, ToIntManyNull(cfg.db.GetTopmenuSubmenuIDs)),
	})
}
