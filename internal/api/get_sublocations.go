package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getSublocation(r *http.Request, i HandlerInput[seedtypes.Sublocation, apitypes.Sublocation, apiresource.NamedAPIResource], id int32) (apitypes.Sublocation, error) {
	sublocation, err := i.QueryParamsVerifyGet(r, id)
	if err != nil {
		return apitypes.Sublocation{}, err
	}

	rel, err := getSublocationRelationships(r, sublocation)
	if err != nil {
		return apitypes.Sublocation{}, err
	}

	response := apitypes.Sublocation{
		ID:                    sublocation.ID,
		Name:                  sublocation.Name,
		ParentLocation:        nameToNamedAPIResource(E.Locations, sublocation.Location, nil),
		Availability:          sublocation.Availability,
		ConnectedSublocations: rel.ConnectedSublocations,
		Areas:                 rel.Areas,
		LocRel:                rel.LocRel,
	}

	return response, nil
}

func retrieveSublocations(r *http.Request, i HandlerInput[seedtypes.Sublocation, apitypes.Sublocation, apiresource.NamedAPIResource]) ([]int32, error) {
	ids, err := i.QueryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []IDFilter{
		i.QueryID(r, ids, ts.QpnLocation, lookup.L.Locations, config.Cfg.DB.GetLocationSublocationIDs),
		i.QueryID(r, ids, ts.QpnMonster, lookup.L.Monsters, config.Cfg.DB.GetSublocationIDsWithMonster),
		i.QueryIDWrapper(r, ids, ts.QpnItem, lookup.L.Items, getSublocationsByItem),
		i.QueryID(r, ids, ts.QpnKeyItem, lookup.L.KeyItems, config.Cfg.DB.GetSublocationIDsWithKeyItem),
		i.QueryID(r, ids, ts.QpnAutoAbility, lookup.L.AutoAbilities, config.Cfg.DB.GetSublocationIDsWithAutoAbility),
		i.QueryBool2(r, ids, ts.QpnCharacters, config.Cfg.DB.GetSublocationIDsWithCharacters),
		i.QueryBool2(r, ids, ts.QpnAeons, config.Cfg.DB.GetSublocationIDsWithAeons),
		i.QueryBool2(r, ids, ts.QpnMonsters, config.Cfg.DB.GetSublocationIDsWithMonsters),
		i.QueryBool2(r, ids, ts.QpnBossFights, config.Cfg.DB.GetSublocationIDsWithBosses),
		i.QueryBool2(r, ids, ts.QpnShops, config.Cfg.DB.GetSublocationIDsWithShops),
		i.QueryBool2(r, ids, ts.QpnTreasures, config.Cfg.DB.GetSublocationIDsWithTreasures),
		i.QueryBool2(r, ids, ts.QpnSidequests, config.Cfg.DB.GetSublocationIDsWithSidequests),
		i.QueryBool2(r, ids, ts.QpnFMVs, config.Cfg.DB.GetSublocationIDsWithFMVs),
	})
}
