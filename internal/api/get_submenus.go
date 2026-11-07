package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getSubmenu(r *http.Request, i HandlerInput[seedtypes.Submenu, apitypes.Submenu, apiresource.NamedAPIResource], id int32) (apitypes.Submenu, error) {
	submenu, err := i.QueryParamsVerifyGet(r, id)
	if err != nil {
		return apitypes.Submenu{}, err
	}

	rel, err := getSubmenuRelationships(r, submenu)
	if err != nil {
		return apitypes.Submenu{}, err
	}

	response := apitypes.Submenu{
		ID:          submenu.ID,
		Name:        submenu.Name,
		Description: submenu.Description,
		Effect:      submenu.Effect,
		Topmenu:     namePtrToNamedAPIResPtr(E.Topmenus, submenu.Topmenu, nil),
		Users:       namesToNamedAPIResources(E.CharacterClasses, submenu.Users),
		OpenedBy:    rel.OpenedBy,
		Abilities:   rel.Abilities,
	}

	return response, nil
}

func retrieveSubmenus(r *http.Request, i HandlerInput[seedtypes.Submenu, apitypes.Submenu, apiresource.NamedAPIResource]) ([]int32, error) {
	ids, err := i.QueryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []IDFilter{
		i.QueryNameID(r, ids, ts.QpnTopmenu, E.Topmenus.resTypeSingle, lookup.L.Topmenus, ToIntManyNull(config.Cfg.DB.GetTopmenuSubmenuIDs)),
	})
}
