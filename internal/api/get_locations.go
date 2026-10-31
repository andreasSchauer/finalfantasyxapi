package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getLocation(r *http.Request, i handlerInput[seeding.Location, Location, NamedAPIResource, NamedAPIResourceList], id int32) (Location, error) {
	location, err := verifyParamsAndGet(r, i, id)
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

func (cfg *Config) retrieveLocations(r *http.Request, i handlerInput[seeding.Location, Location, NamedAPIResource, NamedAPIResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []IDFilter{
		idQuery(r, i, ids, QpnMonster, cfg.L.Monsters, cfg.db.GetLocationIDsWithMonster),
		idQueryWrapper(cfg, r, i, ids, QpnItem, cfg.L.Items, getLocationsByItem),
		idQuery(r, i, ids, QpnKeyItem, cfg.L.KeyItems, cfg.db.GetLocationIDsWithKeyItem),
		idQuery(r, i, ids, QpnAutoAbility, cfg.L.AutoAbilities, cfg.db.GetLocationIDsWithAutoAbility),
		boolQuery2(r, i, ids, QpnCharacters, cfg.db.GetLocationIDsWithCharacters),
		boolQuery2(r, i, ids, QpnAeons, cfg.db.GetLocationIDsWithAeons),
		boolQuery2(r, i, ids, QpnMonsters, cfg.db.GetLocationIDsWithMonsters),
		boolQuery2(r, i, ids, QpnBossFights, cfg.db.GetLocationIDsWithBosses),
		boolQuery2(r, i, ids, QpnShops, cfg.db.GetLocationIDsWithShops),
		boolQuery2(r, i, ids, QpnTreasures, cfg.db.GetLocationIDsWithTreasures),
		boolQuery2(r, i, ids, QpnSidequests, cfg.db.GetLocationIDsWithSidequests),
		boolQuery2(r, i, ids, QpnFMVs, cfg.db.GetLocationIDsWithFMVs),
	})
}
