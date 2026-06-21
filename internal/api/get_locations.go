package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getLocation(r *http.Request, i handlerInput[seeding.Location, Location, NamedAPIResource, NamedApiResourceList], id int32) (Location, error) {
	location, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return Location{}, err
	}

	connectedLocations, err := getResourcesDbItem(cfg, r, cfg.e.locations, location, cfg.db.GetConnectedLocationIDs)
	if err != nil {
		return Location{}, err
	}

	sublocations, err := getResourcesDbItem(cfg, r, cfg.e.sublocations, location, cfg.db.GetLocationSublocationIDs)
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
		Availability:       enumToNamedAPIResource(cfg, cfg.e.availabilityType.endpoint, location.Availability, cfg.t.AvailabilityType),
		ConnectedLocations: connectedLocations,
		Sublocations:       sublocations,
		LocRel:             rel,
	}

	return response, nil
}

func (cfg *Config) retrieveLocations(r *http.Request, i handlerInput[seeding.Location, Location, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	ids, err := verifyParamsAndRetrieve(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterIDs(cfg, r, i, ids, []filteredIdList{
		fidl(idQuery(r, i, ids, "monster", cfg.l.Monsters, cfg.db.GetLocationIDsWithMonster)),
		fidl(idQueryWrapper(cfg, r, i, ids, "item", cfg.l.Items, getLocationsByItem)),
		fidl(idQuery(r, i, ids, "key_item", cfg.l.KeyItems, cfg.db.GetLocationIDsWithKeyItem)),
		fidl(idQuery(r, i, ids, "auto_ability", cfg.l.AutoAbilities, cfg.db.GetLocationIDsWithAutoAbility)),
		fidl(boolQuery2(r, i, ids, "characters", cfg.db.GetLocationIDsWithCharacters)),
		fidl(boolQuery2(r, i, ids, "aeons", cfg.db.GetLocationIDsWithAeons)),
		fidl(boolQuery2(r, i, ids, "monsters", cfg.db.GetLocationIDsWithMonsters)),
		fidl(boolQuery2(r, i, ids, "boss_fights", cfg.db.GetLocationIDsWithBosses)),
		fidl(boolQuery2(r, i, ids, "shops", cfg.db.GetLocationIDsWithShops)),
		fidl(boolQuery2(r, i, ids, "treasures", cfg.db.GetLocationIDsWithTreasures)),
		fidl(boolQuery2(r, i, ids, "sidequests", cfg.db.GetLocationIDsWithSidequests)),
		fidl(boolQuery2(r, i, ids, "fmvs", cfg.db.GetLocationIDsWithFMVs)),
	})
}
