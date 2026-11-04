package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (srv *Server) getLocation(r *http.Request, i handlerInput[seedtypes.Location, Location, NamedAPIResource], id int32) (Location, error) {
	location, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Location{}, err
	}

	rel, err := getLocationRelationships(srv, r, location)
	if err != nil {
		return Location{}, err
	}

	response := Location{
		ID:                 location.ID,
		Name:               location.Name,
		ConnectedLocations: rel.ConnectedLocations,
		Sublocations:       rel.Sublocations,
		Availability:       location.Availability,
		LocRel:             rel.LocRel,
	}

	return response, nil
}

func (srv *Server) retrieveLocations(r *http.Request, i handlerInput[seedtypes.Location, Location, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(srv, r, ids, []IDFilter{
		i.queryID(r, ids, ts.QpnMonster, lookup.L.Monsters, srv.Cfg.DB.GetLocationIDsWithMonster),
		i.queryIDWrapper(srv, r, ids, ts.QpnItem, lookup.L.Items, getLocationsByItem),
		i.queryID(r, ids, ts.QpnKeyItem, lookup.L.KeyItems, srv.Cfg.DB.GetLocationIDsWithKeyItem),
		i.queryID(r, ids, ts.QpnAutoAbility, lookup.L.AutoAbilities, srv.Cfg.DB.GetLocationIDsWithAutoAbility),
		i.queryBool2(r, ids, ts.QpnCharacters, srv.Cfg.DB.GetLocationIDsWithCharacters),
		i.queryBool2(r, ids, ts.QpnAeons, srv.Cfg.DB.GetLocationIDsWithAeons),
		i.queryBool2(r, ids, ts.QpnMonsters, srv.Cfg.DB.GetLocationIDsWithMonsters),
		i.queryBool2(r, ids, ts.QpnBossFights, srv.Cfg.DB.GetLocationIDsWithBosses),
		i.queryBool2(r, ids, ts.QpnShops, srv.Cfg.DB.GetLocationIDsWithShops),
		i.queryBool2(r, ids, ts.QpnTreasures, srv.Cfg.DB.GetLocationIDsWithTreasures),
		i.queryBool2(r, ids, ts.QpnSidequests, srv.Cfg.DB.GetLocationIDsWithSidequests),
		i.queryBool2(r, ids, ts.QpnFMVs, srv.Cfg.DB.GetLocationIDsWithFMVs),
	})
}
