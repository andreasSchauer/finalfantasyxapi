package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getSubmenu(r *http.Request, i handlerinput.Main[seedtypes.Submenu, apitypes.Submenu, apiresource.Named], id int32) (apitypes.Submenu, error) {
	submenu, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Submenu{}, err
	}

	rel, err := GetSubmenuRelationships(r, submenu)
	if err != nil {
		return apitypes.Submenu{}, err
	}

	response := apitypes.Submenu{
		ID:          submenu.ID,
		Name:        submenu.Name,
		Description: submenu.Description,
		Effect:      submenu.Effect,
		Topmenu:     handlerinput.NamePtrToNamedAPIResPtr(apilookups.E.Topmenus, submenu.Topmenu, nil),
		Users:       handlerinput.NamesToNamedAPIResources(apilookups.E.CharacterClasses, submenu.Users),
		OpenedBy:    rel.OpenedBy,
		Abilities:   rel.Abilities,
	}

	return response, nil
}

func retrieveSubmenus(r *http.Request, i handlerinput.Main[seedtypes.Submenu, apitypes.Submenu, apiresource.Named]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryNameID(r, ids, ts.QpnTopmenu, apilookups.E.Topmenus.ResTypeSingle, lookup.L.Topmenus, handlerinput.ToIntManyNull(config.Cfg.DB.GetTopmenuSubmenuIDs)),
	})
}
