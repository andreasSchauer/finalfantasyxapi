package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getLocation(r *http.Request, i handlerInput[seedtypes.Location, Location, apiresource.NamedAPIResource], id int32) (Location, error) {
	location, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Location{}, err
	}

	rel, err := getLocationRelationships(r, location)
	if err != nil {
		return Location{}, err
	}

	response := Location{
		ID:                 location.ID,
		Name:               location.Name,
		ConnectedLocations: rel.ConnectedLocations,
		Sublocations:       rel.Sublocations,
		Availability:       location.Availability,
		LocRel:             rel.LocRel,
	}

	return response, nil
}

func retrieveLocations(r *http.Request, i handlerInput[seedtypes.Location, Location, apiresource.NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(r, ids, []IDFilter{
		i.queryID(r, ids, ts.QpnMonster, lookup.L.Monsters, config.Cfg.DB.GetLocationIDsWithMonster),
		i.queryIDWrapper(r, ids, ts.QpnItem, lookup.L.Items, getLocationsByItem),
		i.queryID(r, ids, ts.QpnKeyItem, lookup.L.KeyItems, config.Cfg.DB.GetLocationIDsWithKeyItem),
		i.queryID(r, ids, ts.QpnAutoAbility, lookup.L.AutoAbilities, config.Cfg.DB.GetLocationIDsWithAutoAbility),
		i.queryBool2(r, ids, ts.QpnCharacters, config.Cfg.DB.GetLocationIDsWithCharacters),
		i.queryBool2(r, ids, ts.QpnAeons, config.Cfg.DB.GetLocationIDsWithAeons),
		i.queryBool2(r, ids, ts.QpnMonsters, config.Cfg.DB.GetLocationIDsWithMonsters),
		i.queryBool2(r, ids, ts.QpnBossFights, config.Cfg.DB.GetLocationIDsWithBosses),
		i.queryBool2(r, ids, ts.QpnShops, config.Cfg.DB.GetLocationIDsWithShops),
		i.queryBool2(r, ids, ts.QpnTreasures, config.Cfg.DB.GetLocationIDsWithTreasures),
		i.queryBool2(r, ids, ts.QpnSidequests, config.Cfg.DB.GetLocationIDsWithSidequests),
		i.queryBool2(r, ids, ts.QpnFMVs, config.Cfg.DB.GetLocationIDsWithFMVs),
	})
}
