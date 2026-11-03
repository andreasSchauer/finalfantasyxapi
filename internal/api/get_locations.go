package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
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
		i.queryID(r, ids, QpnMonster, cfg.L.Monsters, cfg.db.GetLocationIDsWithMonster),
		i.queryIDWrapper(cfg, r, ids, QpnItem, cfg.L.Items, getLocationsByItem),
		i.queryID(r, ids, QpnKeyItem, cfg.L.KeyItems, cfg.db.GetLocationIDsWithKeyItem),
		i.queryID(r, ids, QpnAutoAbility, cfg.L.AutoAbilities, cfg.db.GetLocationIDsWithAutoAbility),
		i.queryBool2(r, ids, QpnCharacters, cfg.db.GetLocationIDsWithCharacters),
		i.queryBool2(r, ids, QpnAeons, cfg.db.GetLocationIDsWithAeons),
		i.queryBool2(r, ids, QpnMonsters, cfg.db.GetLocationIDsWithMonsters),
		i.queryBool2(r, ids, QpnBossFights, cfg.db.GetLocationIDsWithBosses),
		i.queryBool2(r, ids, QpnShops, cfg.db.GetLocationIDsWithShops),
		i.queryBool2(r, ids, QpnTreasures, cfg.db.GetLocationIDsWithTreasures),
		i.queryBool2(r, ids, QpnSidequests, cfg.db.GetLocationIDsWithSidequests),
		i.queryBool2(r, ids, QpnFMVs, cfg.db.GetLocationIDsWithFMVs),
	})
}
