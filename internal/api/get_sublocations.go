package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (cfg *Config) getSublocation(r *http.Request, i handlerInput[seedtypes.Sublocation, Sublocation, NamedAPIResource], id int32) (Sublocation, error) {
	sublocation, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Sublocation{}, err
	}

	rel, err := getSublocationRelationships(cfg, r, sublocation)
	if err != nil {
		return Sublocation{}, err
	}

	response := Sublocation{
		ID:                    sublocation.ID,
		Name:                  sublocation.Name,
		ParentLocation:        nameToNamedAPIResource(cfg, cfg.E.Locations, sublocation.Location, nil),
		Availability:          sublocation.Availability,
		ConnectedSublocations: rel.ConnectedSublocations,
		Areas:                 rel.Areas,
		LocRel:                rel.LocRel,
	}

	return response, nil
}

func (cfg *Config) retrieveSublocations(r *http.Request, i handlerInput[seedtypes.Sublocation, Sublocation, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(cfg, r, ids, []IDFilter{
		i.queryID(r, ids, typestrings.QpnLocation, lookup.L.Locations, cfg.db.GetLocationSublocationIDs),
		i.queryID(r, ids, typestrings.QpnMonster, lookup.L.Monsters, cfg.db.GetSublocationIDsWithMonster),
		i.queryIDWrapper(cfg, r, ids, typestrings.QpnItem, lookup.L.Items, getSublocationsByItem),
		i.queryID(r, ids, typestrings.QpnKeyItem, lookup.L.KeyItems, cfg.db.GetSublocationIDsWithKeyItem),
		i.queryID(r, ids, typestrings.QpnAutoAbility, lookup.L.AutoAbilities, cfg.db.GetSublocationIDsWithAutoAbility),
		i.queryBool2(r, ids, typestrings.QpnCharacters, cfg.db.GetSublocationIDsWithCharacters),
		i.queryBool2(r, ids, typestrings.QpnAeons, cfg.db.GetSublocationIDsWithAeons),
		i.queryBool2(r, ids, typestrings.QpnMonsters, cfg.db.GetSublocationIDsWithMonsters),
		i.queryBool2(r, ids, typestrings.QpnBossFights, cfg.db.GetSublocationIDsWithBosses),
		i.queryBool2(r, ids, typestrings.QpnShops, cfg.db.GetSublocationIDsWithShops),
		i.queryBool2(r, ids, typestrings.QpnTreasures, cfg.db.GetSublocationIDsWithTreasures),
		i.queryBool2(r, ids, typestrings.QpnSidequests, cfg.db.GetSublocationIDsWithSidequests),
		i.queryBool2(r, ids, typestrings.QpnFMVs, cfg.db.GetSublocationIDsWithFMVs),
	})
}
