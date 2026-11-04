package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (s *Server) getArea(r *http.Request, i handlerInput[seedtypes.Area, Area, AreaAPIResource], id int32) (Area, error) {
	area, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Area{}, err
	}

	connections, err := getAreaConnectedAreas(s, area)
	if err != nil {
		return Area{}, err
	}

	rel, err := getAreaRelationships(s, r, area)
	if err != nil {
		return Area{}, err
	}

	response := Area{
		ID:                area.ID,
		Name:              area.Name,
		Version:           area.Version,
		Specification:     area.Specification,
		DisplayName:       getAreaDisplayName(area),
		ParentLocation:    nameToNamedAPIResource(s, E.Locations, area.Location, nil),
		ParentSublocation: nameToNamedAPIResource(s, E.Sublocations, area.Sublocation, nil),
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

func (s *Server) retrieveAreas(r *http.Request, i handlerInput[seedtypes.Area, Area, AreaAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(s, r, ids, []IDFilter{
		i.queryID(r, ids, ts.QpnLocation, lookup.L.Locations, s.Cfg.DB.GetLocationAreaIDs),
		i.queryID(r, ids, ts.QpnSublocation, lookup.L.Sublocations, s.Cfg.DB.GetSublocationAreaIDs),
		i.queryID(r, ids, ts.QpnMonster, lookup.L.Monsters, s.Cfg.DB.GetAreaIDsWithMonster),
		i.queryIDWrapper(s, r, ids, ts.QpnItem, lookup.L.Items, getAreasByItem),
		i.queryID(r, ids, ts.QpnKeyItem, lookup.L.KeyItems, s.Cfg.DB.GetAreaIDsWithKeyItem),
		i.queryID(r, ids, ts.QpnMonster, lookup.L.Monsters, s.Cfg.DB.GetAreasByMonster),
		i.queryID(r, ids, ts.QpnAutoAbility, lookup.L.AutoAbilities, s.Cfg.DB.GetAreaIDsWithAutoAbility),
		i.queryBool(r, ids, ts.QpnSaveSphere, s.Cfg.DB.GetAreaIDsWithSaveSphere),
		i.queryBool(r, ids, ts.QpnCompSphere, s.Cfg.DB.GetAreaIDsWithCompSphere),
		i.queryBool(r, ids, ts.QpnAirship, s.Cfg.DB.GetAreaIDsWithDropOff),
		i.queryBool(r, ids, ts.QpnChocobo, s.Cfg.DB.GetAreaIDsChocobo),
		i.queryBool2(r, ids, ts.QpnCharacters, s.Cfg.DB.GetAreaIDsWithCharacters),
		i.queryBool2(r, ids, ts.QpnAeons, s.Cfg.DB.GetAreaIDsWithAeons),
		i.queryBool2(r, ids, ts.QpnMonsters, s.Cfg.DB.GetAreaIDsWithMonsters),
		i.queryBool2(r, ids, ts.QpnBossFights, s.Cfg.DB.GetAreaIDsWithBosses),
		i.queryBool2(r, ids, ts.QpnShops, s.Cfg.DB.GetAreaIDsWithShops),
		i.queryBool2(r, ids, ts.QpnTreasures, s.Cfg.DB.GetAreaIDsWithTreasures),
		i.queryBool2(r, ids, ts.QpnSidequests, s.Cfg.DB.GetAreaIDsWithSidequests),
		i.queryBool2(r, ids, ts.QpnFMVs, s.Cfg.DB.GetAreaIDsWithFMVs),
	})
}
