package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (srv *Server) getSubmenu(r *http.Request, i handlerInput[seedtypes.Submenu, Submenu, NamedAPIResource], id int32) (Submenu, error) {
	submenu, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Submenu{}, err
	}

	rel, err := getSubmenuRelationships(srv, r, submenu)
	if err != nil {
		return Submenu{}, err
	}

	response := Submenu{
		ID:          submenu.ID,
		Name:        submenu.Name,
		Description: submenu.Description,
		Effect:      submenu.Effect,
		Topmenu:     namePtrToNamedAPIResPtr(srv, E.Topmenus, submenu.Topmenu, nil),
		Users:       namesToNamedAPIResources(srv, E.CharacterClasses, submenu.Users),
		OpenedBy:    rel.OpenedBy,
		Abilities:   rel.Abilities,
	}

	return response, nil
}

func (srv *Server) retrieveSubmenus(r *http.Request, i handlerInput[seedtypes.Submenu, Submenu, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(srv, r, ids, []IDFilter{
		i.queryNameID(r, ids, ts.QpnTopmenu, E.Topmenus.resTypeSingle, lookup.L.Topmenus, ToIntManyNull(config.Cfg.DB.GetTopmenuSubmenuIDs)),
	})
}
