package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getLocation(r *http.Request, i handlerInput[seeding.Location, Location, NamedAPIResource, NamedApiResourceList], id int32) (Location, error) {
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
		Availability:       enumToNamedAPIResource(cfg, cfg.e.availabilityType.endpoint, location.Availability, cfg.t.AvailabilityType),
		LocRel:             rel.LocRel,
	}

	return response, nil
}

func (cfg *Config) retrieveLocations(r *http.Request, i handlerInput[seeding.Location, Location, NamedAPIResource, NamedApiResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []IdFilter{
		idQuery(r, i, ids, qpnMonster, cfg.l.Monsters, cfg.db.GetLocationIDsWithMonster),
		idQueryWrapper(cfg, r, i, ids, qpnItem, cfg.l.Items, getLocationsByItem),
		idQuery(r, i, ids, qpnKeyItem, cfg.l.KeyItems, cfg.db.GetLocationIDsWithKeyItem),
		idQuery(r, i, ids, qpnAutoAbility, cfg.l.AutoAbilities, cfg.db.GetLocationIDsWithAutoAbility),
		boolQuery2(r, i, ids, qpnCharacters, cfg.db.GetLocationIDsWithCharacters),
		boolQuery2(r, i, ids, qpnAeons, cfg.db.GetLocationIDsWithAeons),
		boolQuery2(r, i, ids, qpnMonsters, cfg.db.GetLocationIDsWithMonsters),
		boolQuery2(r, i, ids, qpnBossFights, cfg.db.GetLocationIDsWithBosses),
		boolQuery2(r, i, ids, qpnShops, cfg.db.GetLocationIDsWithShops),
		boolQuery2(r, i, ids, qpnTreasures, cfg.db.GetLocationIDsWithTreasures),
		boolQuery2(r, i, ids, qpnSidequests, cfg.db.GetLocationIDsWithSidequests),
		boolQuery2(r, i, ids, qpnFMVs, cfg.db.GetLocationIDsWithFMVs),
	})
}
