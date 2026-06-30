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
		ParentLocation:    nameToNamedAPIResource(cfg, cfg.e.locations, area.Sublocation.Location.Name, nil),
		ParentSublocation: nameToNamedAPIResource(cfg, cfg.e.sublocations, area.Sublocation.Name, nil),
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
		idQuery(r, i, ids, qpnLocation, cfg.l.Locations, cfg.db.GetLocationAreaIDs),
		idQuery(r, i, ids, qpnSublocation, cfg.l.Sublocations, cfg.db.GetSublocationAreaIDs),
		idQuery(r, i, ids, qpnMonster, cfg.l.Monsters, cfg.db.GetAreaIDsWithMonster),
		idQueryWrapper(cfg, r, i, ids, qpnItem, cfg.l.Items, getAreasByItem),
		idQuery(r, i, ids, qpnKeyItem, cfg.l.KeyItems, cfg.db.GetAreaIDsWithKeyItem),
		idQuery(r, i, ids, qpnMonster, cfg.l.Monsters, cfg.db.GetAreasByMonster),
		idQuery(r, i, ids, qpnAutoAbility, cfg.l.AutoAbilities, cfg.db.GetAreaIDsWithAutoAbility),
		boolQuery(r, i, ids, qpnSaveSphere, cfg.db.GetAreaIDsWithSaveSphere),
		boolQuery(r, i, ids, qpnCompSphere, cfg.db.GetAreaIDsWithCompSphere),
		boolQuery(r, i, ids, qpnAirship, cfg.db.GetAreaIDsWithDropOff),
		boolQuery(r, i, ids, qpnChocobo, cfg.db.GetAreaIDsChocobo),
		boolQuery2(r, i, ids, qpnCharacters, cfg.db.GetAreaIDsWithCharacters),
		boolQuery2(r, i, ids, qpnAeons, cfg.db.GetAreaIDsWithAeons),
		boolQuery2(r, i, ids, qpnMonsters, cfg.db.GetAreaIDsWithMonsters),
		boolQuery2(r, i, ids, qpnBossFights, cfg.db.GetAreaIDsWithBosses),
		boolQuery2(r, i, ids, qpnShops, cfg.db.GetAreaIDsWithShops),
		boolQuery2(r, i, ids, qpnTreasures, cfg.db.GetAreaIDsWithTreasures),
		boolQuery2(r, i, ids, qpnSidequests, cfg.db.GetAreaIDsWithSidequests),
		boolQuery2(r, i, ids, qpnFMVs, cfg.db.GetAreaIDsWithFMVs),
	})
}
