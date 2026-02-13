package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type Location struct {
	ID                 int32              `json:"id"`
	Name               string             `json:"name"`
	ConnectedLocations []NamedAPIResource `json:"connected_locations"`
	Sublocations       []NamedAPIResource `json:"sublocations"`
	LocRel
}

func (cfg *Config) getLocation(r *http.Request, i handlerInput[seeding.Location, Location, NamedAPIResource, NamedApiResourceList], id int32) (Location, error) {
	location, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return Location{}, err
	}

	connectedLocations, err := getResourcesDB(cfg, r, cfg.e.locations, location, cfg.db.GetConnectedLocationIDs)
	if err != nil {
		return Location{}, err
	}

	sublocations, err := getResourcesDB(cfg, r, cfg.e.sublocations, location, cfg.db.GetLocationSublocationIDs)
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
		ConnectedLocations: connectedLocations,
		Sublocations:       sublocations,
		LocRel:             rel,
	}

	return response, nil
}

func (cfg *Config) retrieveLocations(r *http.Request, i handlerInput[seeding.Location, Location, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[NamedAPIResource]{
		frl(idQueryWrapper(cfg, r, i, resources, "item", len(cfg.l.Items), getLocationsByItem)),
		frl(idQueryWrapper(cfg, r, i, resources, "key_item", len(cfg.l.KeyItems), getLocationsByKeyItem)),
		frl(boolQuery2(cfg, r, i, resources, "characters", cfg.db.GetLocationIDsWithCharacters)),
		frl(boolQuery2(cfg, r, i, resources, "aeons", cfg.db.GetLocationIDsWithAeons)),
		frl(boolQuery2(cfg, r, i, resources, "monsters", cfg.db.GetLocationIDsWithMonsters)),
		frl(boolQuery2(cfg, r, i, resources, "boss_fights", cfg.db.GetLocationIDsWithBosses)),
		frl(boolQuery2(cfg, r, i, resources, "shops", cfg.db.GetLocationIDsWithShops)),
		frl(boolQuery2(cfg, r, i, resources, "treasures", cfg.db.GetLocationIDsWithTreasures)),
		frl(boolQuery2(cfg, r, i, resources, "sidequests", cfg.db.GetLocationIDsWithSidequests)),
		frl(boolQuery2(cfg, r, i, resources, "fmvs", cfg.db.GetLocationIDsWithFMVs)),
	})
}
