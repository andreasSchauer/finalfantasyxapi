package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
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
		i.queryID(r, ids, typestrings.QpnLocation, lookup.L.Locations, cfg.db.GetLocationAreaIDs),
		i.queryID(r, ids, typestrings.QpnSublocation, lookup.L.Sublocations, cfg.db.GetSublocationAreaIDs),
		i.queryID(r, ids, typestrings.QpnMonster, lookup.L.Monsters, cfg.db.GetAreaIDsWithMonster),
		i.queryIDWrapper(cfg, r, ids, typestrings.QpnItem, lookup.L.Items, getAreasByItem),
		i.queryID(r, ids, typestrings.QpnKeyItem, lookup.L.KeyItems, cfg.db.GetAreaIDsWithKeyItem),
		i.queryID(r, ids, typestrings.QpnMonster, lookup.L.Monsters, cfg.db.GetAreasByMonster),
		i.queryID(r, ids, typestrings.QpnAutoAbility, lookup.L.AutoAbilities, cfg.db.GetAreaIDsWithAutoAbility),
		i.queryBool(r, ids, typestrings.QpnSaveSphere, cfg.db.GetAreaIDsWithSaveSphere),
		i.queryBool(r, ids, typestrings.QpnCompSphere, cfg.db.GetAreaIDsWithCompSphere),
		i.queryBool(r, ids, typestrings.QpnAirship, cfg.db.GetAreaIDsWithDropOff),
		i.queryBool(r, ids, typestrings.QpnChocobo, cfg.db.GetAreaIDsChocobo),
		i.queryBool2(r, ids, typestrings.QpnCharacters, cfg.db.GetAreaIDsWithCharacters),
		i.queryBool2(r, ids, typestrings.QpnAeons, cfg.db.GetAreaIDsWithAeons),
		i.queryBool2(r, ids, typestrings.QpnMonsters, cfg.db.GetAreaIDsWithMonsters),
		i.queryBool2(r, ids, typestrings.QpnBossFights, cfg.db.GetAreaIDsWithBosses),
		i.queryBool2(r, ids, typestrings.QpnShops, cfg.db.GetAreaIDsWithShops),
		i.queryBool2(r, ids, typestrings.QpnTreasures, cfg.db.GetAreaIDsWithTreasures),
		i.queryBool2(r, ids, typestrings.QpnSidequests, cfg.db.GetAreaIDsWithSidequests),
		i.queryBool2(r, ids, typestrings.QpnFMVs, cfg.db.GetAreaIDsWithFMVs),
	})
}
