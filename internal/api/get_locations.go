package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getLocation(r *http.Request, i handlerinput.Main[seedtypes.Location, apitypes.Location, apiresource.Named], id int32) (apitypes.Location, error) {
	location, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Location{}, err
	}

	rel, err := getLocationRelationships(r, location)
	if err != nil {
		return apitypes.Location{}, err
	}

	response := apitypes.Location{
		ID:                 location.ID,
		Name:               location.Name,
		ConnectedLocations: rel.ConnectedLocations,
		Sublocations:       rel.Sublocations,
		Availability:       location.Availability,
		LocRel:             rel.LocRel,
	}

	return response, nil
}

func retrieveLocations(r *http.Request, i handlerinput.Main[seedtypes.Location, apitypes.Location, apiresource.Named]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryID(r, ids, ts.QpnMonster, lookup.L.Monsters, config.Cfg.DB.GetLocationIDsWithMonster),
		i.QueryIDWrapper(r, ids, ts.QpnItem, lookup.L.Items, getLocationsByItem),
		i.QueryID(r, ids, ts.QpnKeyItem, lookup.L.KeyItems, config.Cfg.DB.GetLocationIDsWithKeyItem),
		i.QueryID(r, ids, ts.QpnAutoAbility, lookup.L.AutoAbilities, config.Cfg.DB.GetLocationIDsWithAutoAbility),
		i.QueryBool2(r, ids, ts.QpnCharacters, config.Cfg.DB.GetLocationIDsWithCharacters),
		i.QueryBool2(r, ids, ts.QpnAeons, config.Cfg.DB.GetLocationIDsWithAeons),
		i.QueryBool2(r, ids, ts.QpnMonsters, config.Cfg.DB.GetLocationIDsWithMonsters),
		i.QueryBool2(r, ids, ts.QpnBossFights, config.Cfg.DB.GetLocationIDsWithBosses),
		i.QueryBool2(r, ids, ts.QpnShops, config.Cfg.DB.GetLocationIDsWithShops),
		i.QueryBool2(r, ids, ts.QpnTreasures, config.Cfg.DB.GetLocationIDsWithTreasures),
		i.QueryBool2(r, ids, ts.QpnSidequests, config.Cfg.DB.GetLocationIDsWithSidequests),
		i.QueryBool2(r, ids, ts.QpnFMVs, config.Cfg.DB.GetLocationIDsWithFMVs),
	})
}
