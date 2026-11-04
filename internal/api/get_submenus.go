package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (s *Server) getSubmenu(r *http.Request, i handlerInput[seedtypes.Submenu, Submenu, NamedAPIResource], id int32) (Submenu, error) {
	submenu, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Submenu{}, err
	}

	rel, err := getSubmenuRelationships(s, r, submenu)
	if err != nil {
		return Submenu{}, err
	}

	response := Submenu{
		ID:          submenu.ID,
		Name:        submenu.Name,
		Description: submenu.Description,
		Effect:      submenu.Effect,
		Topmenu:     namePtrToNamedAPIResPtr(s, E.Topmenus, submenu.Topmenu, nil),
		Users:       namesToNamedAPIResources(s, E.CharacterClasses, submenu.Users),
		OpenedBy:    rel.OpenedBy,
		Abilities:   rel.Abilities,
	}

	return response, nil
}

func (s *Server) retrieveSubmenus(r *http.Request, i handlerInput[seedtypes.Submenu, Submenu, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(s, r, ids, []IDFilter{
		i.queryNameID(r, ids, ts.QpnTopmenu, E.Topmenus.resTypeSingle, lookup.L.Topmenus, ToIntManyNull(s.Cfg.DB.GetTopmenuSubmenuIDs)),
	})
}
