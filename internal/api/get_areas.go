package api

import (
	"net/http"

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
		i.queryID(r, ids, ts.QpnLocation, lookup.L.Locations, srv.Cfg.DB.GetLocationAreaIDs),
		i.queryID(r, ids, ts.QpnSublocation, lookup.L.Sublocations, srv.Cfg.DB.GetSublocationAreaIDs),
		i.queryID(r, ids, ts.QpnMonster, lookup.L.Monsters, srv.Cfg.DB.GetAreaIDsWithMonster),
		i.queryIDWrapper(srv, r, ids, ts.QpnItem, lookup.L.Items, getAreasByItem),
		i.queryID(r, ids, ts.QpnKeyItem, lookup.L.KeyItems, srv.Cfg.DB.GetAreaIDsWithKeyItem),
		i.queryID(r, ids, ts.QpnMonster, lookup.L.Monsters, srv.Cfg.DB.GetAreasByMonster),
		i.queryID(r, ids, ts.QpnAutoAbility, lookup.L.AutoAbilities, srv.Cfg.DB.GetAreaIDsWithAutoAbility),
		i.queryBool(r, ids, ts.QpnSaveSphere, srv.Cfg.DB.GetAreaIDsWithSaveSphere),
		i.queryBool(r, ids, ts.QpnCompSphere, srv.Cfg.DB.GetAreaIDsWithCompSphere),
		i.queryBool(r, ids, ts.QpnAirship, srv.Cfg.DB.GetAreaIDsWithDropOff),
		i.queryBool(r, ids, ts.QpnChocobo, srv.Cfg.DB.GetAreaIDsChocobo),
		i.queryBool2(r, ids, ts.QpnCharacters, srv.Cfg.DB.GetAreaIDsWithCharacters),
		i.queryBool2(r, ids, ts.QpnAeons, srv.Cfg.DB.GetAreaIDsWithAeons),
		i.queryBool2(r, ids, ts.QpnMonsters, srv.Cfg.DB.GetAreaIDsWithMonsters),
		i.queryBool2(r, ids, ts.QpnBossFights, srv.Cfg.DB.GetAreaIDsWithBosses),
		i.queryBool2(r, ids, ts.QpnShops, srv.Cfg.DB.GetAreaIDsWithShops),
		i.queryBool2(r, ids, ts.QpnTreasures, srv.Cfg.DB.GetAreaIDsWithTreasures),
		i.queryBool2(r, ids, ts.QpnSidequests, srv.Cfg.DB.GetAreaIDsWithSidequests),
		i.queryBool2(r, ids, ts.QpnFMVs, srv.Cfg.DB.GetAreaIDsWithFMVs),
	})
}
