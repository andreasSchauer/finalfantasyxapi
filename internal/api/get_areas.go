package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
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
		i.queryID(r, ids, QpnLocation, cfg.L.Locations, cfg.db.GetLocationAreaIDs),
		i.queryID(r, ids, QpnSublocation, cfg.L.Sublocations, cfg.db.GetSublocationAreaIDs),
		i.queryID(r, ids, QpnMonster, cfg.L.Monsters, cfg.db.GetAreaIDsWithMonster),
		i.queryIDWrapper(cfg, r, ids, QpnItem, cfg.L.Items, getAreasByItem),
		i.queryID(r, ids, QpnKeyItem, cfg.L.KeyItems, cfg.db.GetAreaIDsWithKeyItem),
		i.queryID(r, ids, QpnMonster, cfg.L.Monsters, cfg.db.GetAreasByMonster),
		i.queryID(r, ids, QpnAutoAbility, cfg.L.AutoAbilities, cfg.db.GetAreaIDsWithAutoAbility),
		i.queryBool(r, ids, QpnSaveSphere, cfg.db.GetAreaIDsWithSaveSphere),
		i.queryBool(r, ids, QpnCompSphere, cfg.db.GetAreaIDsWithCompSphere),
		i.queryBool(r, ids, QpnAirship, cfg.db.GetAreaIDsWithDropOff),
		i.queryBool(r, ids, QpnChocobo, cfg.db.GetAreaIDsChocobo),
		i.queryBool2(r, ids, QpnCharacters, cfg.db.GetAreaIDsWithCharacters),
		i.queryBool2(r, ids, QpnAeons, cfg.db.GetAreaIDsWithAeons),
		i.queryBool2(r, ids, QpnMonsters, cfg.db.GetAreaIDsWithMonsters),
		i.queryBool2(r, ids, QpnBossFights, cfg.db.GetAreaIDsWithBosses),
		i.queryBool2(r, ids, QpnShops, cfg.db.GetAreaIDsWithShops),
		i.queryBool2(r, ids, QpnTreasures, cfg.db.GetAreaIDsWithTreasures),
		i.queryBool2(r, ids, QpnSidequests, cfg.db.GetAreaIDsWithSidequests),
		i.queryBool2(r, ids, QpnFMVs, cfg.db.GetAreaIDsWithFMVs),
	})
}
