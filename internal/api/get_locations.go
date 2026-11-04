package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (cfg *Config) getLocation(r *http.Request, i handlerInput[seedtypes.Location, Location, NamedAPIResource], id int32) (Location, error) {
	location, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Location{}, err
	}

	rel, err := getLocationRelationships(cfg, r, location)
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

func (cfg *Config) retrieveLocations(r *http.Request, i handlerInput[seedtypes.Location, Location, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(cfg, r, ids, []IDFilter{
		i.queryID(r, ids, typestrings.QpnMonster, lookup.L.Monsters, cfg.db.GetLocationIDsWithMonster),
		i.queryIDWrapper(cfg, r, ids, typestrings.QpnItem, lookup.L.Items, getLocationsByItem),
		i.queryID(r, ids, typestrings.QpnKeyItem, lookup.L.KeyItems, cfg.db.GetLocationIDsWithKeyItem),
		i.queryID(r, ids, typestrings.QpnAutoAbility, lookup.L.AutoAbilities, cfg.db.GetLocationIDsWithAutoAbility),
		i.queryBool2(r, ids, typestrings.QpnCharacters, cfg.db.GetLocationIDsWithCharacters),
		i.queryBool2(r, ids, typestrings.QpnAeons, cfg.db.GetLocationIDsWithAeons),
		i.queryBool2(r, ids, typestrings.QpnMonsters, cfg.db.GetLocationIDsWithMonsters),
		i.queryBool2(r, ids, typestrings.QpnBossFights, cfg.db.GetLocationIDsWithBosses),
		i.queryBool2(r, ids, typestrings.QpnShops, cfg.db.GetLocationIDsWithShops),
		i.queryBool2(r, ids, typestrings.QpnTreasures, cfg.db.GetLocationIDsWithTreasures),
		i.queryBool2(r, ids, typestrings.QpnSidequests, cfg.db.GetLocationIDsWithSidequests),
		i.queryBool2(r, ids, typestrings.QpnFMVs, cfg.db.GetLocationIDsWithFMVs),
	})
}
