package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (srv *Server) getArea(r *http.Request, i handlerInput[seedtypes.Area, Area, AreaAPIResource], id int32) (Area, error) {
	area, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Area{}, err
	}

	connections, err := getAreaConnectedAreas(srv, area)
	if err != nil {
		return Area{}, err
	}

	rel, err := getAreaRelationships(srv, r, area)
	if err != nil {
		return Area{}, err
	}

	response := Area{
		ID:                area.ID,
		Name:              area.Name,
		Version:           area.Version,
		Specification:     area.Specification,
		DisplayName:       getAreaDisplayName(area),
		ParentLocation:    nameToNamedAPIResource(srv, E.Locations, area.Location, nil),
		ParentSublocation: nameToNamedAPIResource(srv, E.Sublocations, area.Sublocation, nil),
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

func (srv *Server) retrieveAreas(r *http.Request, i handlerInput[seedtypes.Area, Area, AreaAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(srv, r, ids, []IDFilter{
		i.queryID(r, ids, ts.QpnLocation, lookup.L.Locations, config.Cfg.DB.GetLocationAreaIDs),
		i.queryID(r, ids, ts.QpnSublocation, lookup.L.Sublocations, config.Cfg.DB.GetSublocationAreaIDs),
		i.queryID(r, ids, ts.QpnMonster, lookup.L.Monsters, config.Cfg.DB.GetAreaIDsWithMonster),
		i.queryIDWrapper(srv, r, ids, ts.QpnItem, lookup.L.Items, getAreasByItem),
		i.queryID(r, ids, ts.QpnKeyItem, lookup.L.KeyItems, config.Cfg.DB.GetAreaIDsWithKeyItem),
		i.queryID(r, ids, ts.QpnMonster, lookup.L.Monsters, config.Cfg.DB.GetAreasByMonster),
		i.queryID(r, ids, ts.QpnAutoAbility, lookup.L.AutoAbilities, config.Cfg.DB.GetAreaIDsWithAutoAbility),
		i.queryBool(r, ids, ts.QpnSaveSphere, config.Cfg.DB.GetAreaIDsWithSaveSphere),
		i.queryBool(r, ids, ts.QpnCompSphere, config.Cfg.DB.GetAreaIDsWithCompSphere),
		i.queryBool(r, ids, ts.QpnAirship, config.Cfg.DB.GetAreaIDsWithDropOff),
		i.queryBool(r, ids, ts.QpnChocobo, config.Cfg.DB.GetAreaIDsChocobo),
		i.queryBool2(r, ids, ts.QpnCharacters, config.Cfg.DB.GetAreaIDsWithCharacters),
		i.queryBool2(r, ids, ts.QpnAeons, config.Cfg.DB.GetAreaIDsWithAeons),
		i.queryBool2(r, ids, ts.QpnMonsters, config.Cfg.DB.GetAreaIDsWithMonsters),
		i.queryBool2(r, ids, ts.QpnBossFights, config.Cfg.DB.GetAreaIDsWithBosses),
		i.queryBool2(r, ids, ts.QpnShops, config.Cfg.DB.GetAreaIDsWithShops),
		i.queryBool2(r, ids, ts.QpnTreasures, config.Cfg.DB.GetAreaIDsWithTreasures),
		i.queryBool2(r, ids, ts.QpnSidequests, config.Cfg.DB.GetAreaIDsWithSidequests),
		i.queryBool2(r, ids, ts.QpnFMVs, config.Cfg.DB.GetAreaIDsWithFMVs),
	})
}
