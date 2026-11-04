package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (srv *Server) getSublocation(r *http.Request, i handlerInput[seedtypes.Sublocation, Sublocation, NamedAPIResource], id int32) (Sublocation, error) {
	sublocation, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Sublocation{}, err
	}

	rel, err := getSublocationRelationships(srv, r, sublocation)
	if err != nil {
		return Sublocation{}, err
	}

	response := Sublocation{
		ID:                    sublocation.ID,
		Name:                  sublocation.Name,
		ParentLocation:        nameToNamedAPIResource(srv, E.Locations, sublocation.Location, nil),
		Availability:          sublocation.Availability,
		ConnectedSublocations: rel.ConnectedSublocations,
		Areas:                 rel.Areas,
		LocRel:                rel.LocRel,
	}

	return response, nil
}

func (srv *Server) retrieveSublocations(r *http.Request, i handlerInput[seedtypes.Sublocation, Sublocation, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(srv, r, ids, []IDFilter{
		i.queryID(r, ids, ts.QpnLocation, lookup.L.Locations, srv.Cfg.DB.GetLocationSublocationIDs),
		i.queryID(r, ids, ts.QpnMonster, lookup.L.Monsters, srv.Cfg.DB.GetSublocationIDsWithMonster),
		i.queryIDWrapper(srv, r, ids, ts.QpnItem, lookup.L.Items, getSublocationsByItem),
		i.queryID(r, ids, ts.QpnKeyItem, lookup.L.KeyItems, srv.Cfg.DB.GetSublocationIDsWithKeyItem),
		i.queryID(r, ids, ts.QpnAutoAbility, lookup.L.AutoAbilities, srv.Cfg.DB.GetSublocationIDsWithAutoAbility),
		i.queryBool2(r, ids, ts.QpnCharacters, srv.Cfg.DB.GetSublocationIDsWithCharacters),
		i.queryBool2(r, ids, ts.QpnAeons, srv.Cfg.DB.GetSublocationIDsWithAeons),
		i.queryBool2(r, ids, ts.QpnMonsters, srv.Cfg.DB.GetSublocationIDsWithMonsters),
		i.queryBool2(r, ids, ts.QpnBossFights, srv.Cfg.DB.GetSublocationIDsWithBosses),
		i.queryBool2(r, ids, ts.QpnShops, srv.Cfg.DB.GetSublocationIDsWithShops),
		i.queryBool2(r, ids, ts.QpnTreasures, srv.Cfg.DB.GetSublocationIDsWithTreasures),
		i.queryBool2(r, ids, ts.QpnSidequests, srv.Cfg.DB.GetSublocationIDsWithSidequests),
		i.queryBool2(r, ids, ts.QpnFMVs, srv.Cfg.DB.GetSublocationIDsWithFMVs),
	})
}
