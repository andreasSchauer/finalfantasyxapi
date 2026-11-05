package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getSublocation(r *http.Request, i handlerInput[seedtypes.Sublocation, Sublocation, apiresource.NamedAPIResource], id int32) (Sublocation, error) {
	sublocation, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Sublocation{}, err
	}

	rel, err := getSublocationRelationships(r, sublocation)
	if err != nil {
		return Sublocation{}, err
	}

	response := Sublocation{
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

func retrieveSublocations(r *http.Request, i handlerInput[seedtypes.Sublocation, Sublocation, apiresource.NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(r, ids, []IDFilter{
		i.queryID(r, ids, ts.QpnLocation, lookup.L.Locations, config.Cfg.DB.GetLocationSublocationIDs),
		i.queryID(r, ids, ts.QpnMonster, lookup.L.Monsters, config.Cfg.DB.GetSublocationIDsWithMonster),
		i.queryIDWrapper(r, ids, ts.QpnItem, lookup.L.Items, getSublocationsByItem),
		i.queryID(r, ids, ts.QpnKeyItem, lookup.L.KeyItems, config.Cfg.DB.GetSublocationIDsWithKeyItem),
		i.queryID(r, ids, ts.QpnAutoAbility, lookup.L.AutoAbilities, config.Cfg.DB.GetSublocationIDsWithAutoAbility),
		i.queryBool2(r, ids, ts.QpnCharacters, config.Cfg.DB.GetSublocationIDsWithCharacters),
		i.queryBool2(r, ids, ts.QpnAeons, config.Cfg.DB.GetSublocationIDsWithAeons),
		i.queryBool2(r, ids, ts.QpnMonsters, config.Cfg.DB.GetSublocationIDsWithMonsters),
		i.queryBool2(r, ids, ts.QpnBossFights, config.Cfg.DB.GetSublocationIDsWithBosses),
		i.queryBool2(r, ids, ts.QpnShops, config.Cfg.DB.GetSublocationIDsWithShops),
		i.queryBool2(r, ids, ts.QpnTreasures, config.Cfg.DB.GetSublocationIDsWithTreasures),
		i.queryBool2(r, ids, ts.QpnSidequests, config.Cfg.DB.GetSublocationIDsWithSidequests),
		i.queryBool2(r, ids, ts.QpnFMVs, config.Cfg.DB.GetSublocationIDsWithFMVs),
	})
}
