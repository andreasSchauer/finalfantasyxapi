package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (s *Server) getLocation(r *http.Request, i handlerInput[seedtypes.Location, Location, NamedAPIResource], id int32) (Location, error) {
	location, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Location{}, err
	}

	rel, err := getLocationRelationships(s, r, location)
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

func (s *Server) retrieveLocations(r *http.Request, i handlerInput[seedtypes.Location, Location, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(s, r, ids, []IDFilter{
		i.queryID(r, ids, ts.QpnMonster, lookup.L.Monsters, s.Cfg.DB.GetLocationIDsWithMonster),
		i.queryIDWrapper(s, r, ids, ts.QpnItem, lookup.L.Items, getLocationsByItem),
		i.queryID(r, ids, ts.QpnKeyItem, lookup.L.KeyItems, s.Cfg.DB.GetLocationIDsWithKeyItem),
		i.queryID(r, ids, ts.QpnAutoAbility, lookup.L.AutoAbilities, s.Cfg.DB.GetLocationIDsWithAutoAbility),
		i.queryBool2(r, ids, ts.QpnCharacters, s.Cfg.DB.GetLocationIDsWithCharacters),
		i.queryBool2(r, ids, ts.QpnAeons, s.Cfg.DB.GetLocationIDsWithAeons),
		i.queryBool2(r, ids, ts.QpnMonsters, s.Cfg.DB.GetLocationIDsWithMonsters),
		i.queryBool2(r, ids, ts.QpnBossFights, s.Cfg.DB.GetLocationIDsWithBosses),
		i.queryBool2(r, ids, ts.QpnShops, s.Cfg.DB.GetLocationIDsWithShops),
		i.queryBool2(r, ids, ts.QpnTreasures, s.Cfg.DB.GetLocationIDsWithTreasures),
		i.queryBool2(r, ids, ts.QpnSidequests, s.Cfg.DB.GetLocationIDsWithSidequests),
		i.queryBool2(r, ids, ts.QpnFMVs, s.Cfg.DB.GetLocationIDsWithFMVs),
	})
}
