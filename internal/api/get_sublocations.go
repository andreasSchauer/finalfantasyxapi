package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (s *Server) getSublocation(r *http.Request, i handlerInput[seedtypes.Sublocation, Sublocation, NamedAPIResource], id int32) (Sublocation, error) {
	sublocation, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Sublocation{}, err
	}

	rel, err := getSublocationRelationships(s, r, sublocation)
	if err != nil {
		return Sublocation{}, err
	}

	response := Sublocation{
		ID:                    sublocation.ID,
		Name:                  sublocation.Name,
		ParentLocation:        nameToNamedAPIResource(s, E.Locations, sublocation.Location, nil),
		Availability:          sublocation.Availability,
		ConnectedSublocations: rel.ConnectedSublocations,
		Areas:                 rel.Areas,
		LocRel:                rel.LocRel,
	}

	return response, nil
}

func (s *Server) retrieveSublocations(r *http.Request, i handlerInput[seedtypes.Sublocation, Sublocation, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(s, r, ids, []IDFilter{
		i.queryID(r, ids, ts.QpnLocation, lookup.L.Locations, s.Cfg.DB.GetLocationSublocationIDs),
		i.queryID(r, ids, ts.QpnMonster, lookup.L.Monsters, s.Cfg.DB.GetSublocationIDsWithMonster),
		i.queryIDWrapper(s, r, ids, ts.QpnItem, lookup.L.Items, getSublocationsByItem),
		i.queryID(r, ids, ts.QpnKeyItem, lookup.L.KeyItems, s.Cfg.DB.GetSublocationIDsWithKeyItem),
		i.queryID(r, ids, ts.QpnAutoAbility, lookup.L.AutoAbilities, s.Cfg.DB.GetSublocationIDsWithAutoAbility),
		i.queryBool2(r, ids, ts.QpnCharacters, s.Cfg.DB.GetSublocationIDsWithCharacters),
		i.queryBool2(r, ids, ts.QpnAeons, s.Cfg.DB.GetSublocationIDsWithAeons),
		i.queryBool2(r, ids, ts.QpnMonsters, s.Cfg.DB.GetSublocationIDsWithMonsters),
		i.queryBool2(r, ids, ts.QpnBossFights, s.Cfg.DB.GetSublocationIDsWithBosses),
		i.queryBool2(r, ids, ts.QpnShops, s.Cfg.DB.GetSublocationIDsWithShops),
		i.queryBool2(r, ids, ts.QpnTreasures, s.Cfg.DB.GetSublocationIDsWithTreasures),
		i.queryBool2(r, ids, ts.QpnSidequests, s.Cfg.DB.GetSublocationIDsWithSidequests),
		i.queryBool2(r, ids, ts.QpnFMVs, s.Cfg.DB.GetSublocationIDsWithFMVs),
	})
}
