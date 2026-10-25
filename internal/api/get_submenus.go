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
		Topmenu:     namePtrToNamedAPIResPtr(cfg, cfg.E.Topmenus, submenu.Topmenu, nil),
		Users:       namesToNamedAPIResources(cfg, cfg.E.CharacterClasses, submenu.Users),
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
		nameIdQuery(r, i, ids, QpnTopmenu, cfg.E.Topmenus.resTypeSingle, cfg.L.Topmenus, ToIntManyNull(cfg.db.GetTopmenuSubmenuIDs)),
	})
}
