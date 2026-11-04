package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (cfg *Config) getArea(r *http.Request, i handlerInput[seedtypes.Area, Area, AreaAPIResource], id int32) (Area, error) {
	area, err := i.queryParamsVerifyGet(r, id)
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

func (cfg *Config) retrieveAreas(r *http.Request, i handlerInput[seedtypes.Area, Area, AreaAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(cfg, r, ids, []IDFilter{
		i.queryID(r, ids, ts.QpnLocation, lookup.L.Locations, cfg.db.GetLocationAreaIDs),
		i.queryID(r, ids, ts.QpnSublocation, lookup.L.Sublocations, cfg.db.GetSublocationAreaIDs),
		i.queryID(r, ids, ts.QpnMonster, lookup.L.Monsters, cfg.db.GetAreaIDsWithMonster),
		i.queryIDWrapper(cfg, r, ids, ts.QpnItem, lookup.L.Items, getAreasByItem),
		i.queryID(r, ids, ts.QpnKeyItem, lookup.L.KeyItems, cfg.db.GetAreaIDsWithKeyItem),
		i.queryID(r, ids, ts.QpnMonster, lookup.L.Monsters, cfg.db.GetAreasByMonster),
		i.queryID(r, ids, ts.QpnAutoAbility, lookup.L.AutoAbilities, cfg.db.GetAreaIDsWithAutoAbility),
		i.queryBool(r, ids, ts.QpnSaveSphere, cfg.db.GetAreaIDsWithSaveSphere),
		i.queryBool(r, ids, ts.QpnCompSphere, cfg.db.GetAreaIDsWithCompSphere),
		i.queryBool(r, ids, ts.QpnAirship, cfg.db.GetAreaIDsWithDropOff),
		i.queryBool(r, ids, ts.QpnChocobo, cfg.db.GetAreaIDsChocobo),
		i.queryBool2(r, ids, ts.QpnCharacters, cfg.db.GetAreaIDsWithCharacters),
		i.queryBool2(r, ids, ts.QpnAeons, cfg.db.GetAreaIDsWithAeons),
		i.queryBool2(r, ids, ts.QpnMonsters, cfg.db.GetAreaIDsWithMonsters),
		i.queryBool2(r, ids, ts.QpnBossFights, cfg.db.GetAreaIDsWithBosses),
		i.queryBool2(r, ids, ts.QpnShops, cfg.db.GetAreaIDsWithShops),
		i.queryBool2(r, ids, ts.QpnTreasures, cfg.db.GetAreaIDsWithTreasures),
		i.queryBool2(r, ids, ts.QpnSidequests, cfg.db.GetAreaIDsWithSidequests),
		i.queryBool2(r, ids, ts.QpnFMVs, cfg.db.GetAreaIDsWithFMVs),
	})
}
