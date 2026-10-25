package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getArea(r *http.Request, i handlerInput[seeding.Area, Area, AreaAPIResource, AreaApiResourceList], id int32) (Area, error) {
	area, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return Area{}, err
	}

	connections, err := getAreaConnectedAreas(cfg, area)
	if err != nil {
		return Area{}, err
	}

	rel, err := getAreaRelationships(cfg, r, area)
	if err != nil {
		return Area{}, err
	}

	response := Area{
		ID:                area.ID,
		Name:              area.Name,
		Version:           area.Version,
		Specification:     area.Specification,
		DisplayName:       getAreaDisplayName(area),
		ParentLocation:    nameToNamedAPIResource(cfg, cfg.E.Locations, area.Location, nil),
		ParentSublocation: nameToNamedAPIResource(cfg, cfg.E.Sublocations, area.Sublocation, nil),
		Availability:      area.Availability,
		HasSaveSphere:     area.HasSaveSphere,
		AirshipDropOff:    area.AirshipDropOff,
		HasCompSphere:     area.HasCompilationSphere,
		CanRideChocobo:    area.CanRideChocobo,
		ConnectedAreas:    connections,
		LocRel:            rel,
	}

	return response, nil
}

func (cfg *Config) retrieveAreas(r *http.Request, i handlerInput[seeding.Area, Area, AreaAPIResource, AreaApiResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []IdFilter{
		idQuery(r, i, ids, QpnLocation, cfg.L.Locations, cfg.db.GetLocationAreaIDs),
		idQuery(r, i, ids, QpnSublocation, cfg.L.Sublocations, cfg.db.GetSublocationAreaIDs),
		idQuery(r, i, ids, QpnMonster, cfg.L.Monsters, cfg.db.GetAreaIDsWithMonster),
		idQueryWrapper(cfg, r, i, ids, QpnItem, cfg.L.Items, getAreasByItem),
		idQuery(r, i, ids, QpnKeyItem, cfg.L.KeyItems, cfg.db.GetAreaIDsWithKeyItem),
		idQuery(r, i, ids, QpnMonster, cfg.L.Monsters, cfg.db.GetAreasByMonster),
		idQuery(r, i, ids, QpnAutoAbility, cfg.L.AutoAbilities, cfg.db.GetAreaIDsWithAutoAbility),
		boolQuery(r, i, ids, QpnSaveSphere, cfg.db.GetAreaIDsWithSaveSphere),
		boolQuery(r, i, ids, QpnCompSphere, cfg.db.GetAreaIDsWithCompSphere),
		boolQuery(r, i, ids, QpnAirship, cfg.db.GetAreaIDsWithDropOff),
		boolQuery(r, i, ids, QpnChocobo, cfg.db.GetAreaIDsChocobo),
		boolQuery2(r, i, ids, QpnCharacters, cfg.db.GetAreaIDsWithCharacters),
		boolQuery2(r, i, ids, QpnAeons, cfg.db.GetAreaIDsWithAeons),
		boolQuery2(r, i, ids, QpnMonsters, cfg.db.GetAreaIDsWithMonsters),
		boolQuery2(r, i, ids, QpnBossFights, cfg.db.GetAreaIDsWithBosses),
		boolQuery2(r, i, ids, QpnShops, cfg.db.GetAreaIDsWithShops),
		boolQuery2(r, i, ids, QpnTreasures, cfg.db.GetAreaIDsWithTreasures),
		boolQuery2(r, i, ids, QpnSidequests, cfg.db.GetAreaIDsWithSidequests),
		boolQuery2(r, i, ids, QpnFMVs, cfg.db.GetAreaIDsWithFMVs),
	})
}
