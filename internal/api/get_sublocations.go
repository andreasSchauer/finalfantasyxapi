package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getSublocation(r *http.Request, i handlerInput[seeding.Sublocation, Sublocation, NamedAPIResource], id int32) (Sublocation, error) {
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

func (cfg *Config) retrieveSublocations(r *http.Request, i handlerInput[seeding.Sublocation, Sublocation, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(cfg, r, ids, []IDFilter{
		i.queryID(r, ids, QpnLocation, cfg.L.Locations, cfg.db.GetLocationSublocationIDs),
		i.queryID(r, ids, QpnMonster, cfg.L.Monsters, cfg.db.GetSublocationIDsWithMonster),
		i.queryIDWrapper(cfg, r, ids, QpnItem, cfg.L.Items, getSublocationsByItem),
		i.queryID(r, ids, QpnKeyItem, cfg.L.KeyItems, cfg.db.GetSublocationIDsWithKeyItem),
		i.queryID(r, ids, QpnAutoAbility, cfg.L.AutoAbilities, cfg.db.GetSublocationIDsWithAutoAbility),
		i.queryBool2(r, ids, QpnCharacters, cfg.db.GetSublocationIDsWithCharacters),
		i.queryBool2(r, ids, QpnAeons, cfg.db.GetSublocationIDsWithAeons),
		i.queryBool2(r, ids, QpnMonsters, cfg.db.GetSublocationIDsWithMonsters),
		i.queryBool2(r, ids, QpnBossFights, cfg.db.GetSublocationIDsWithBosses),
		i.queryBool2(r, ids, QpnShops, cfg.db.GetSublocationIDsWithShops),
		i.queryBool2(r, ids, QpnTreasures, cfg.db.GetSublocationIDsWithTreasures),
		i.queryBool2(r, ids, QpnSidequests, cfg.db.GetSublocationIDsWithSidequests),
		i.queryBool2(r, ids, QpnFMVs, cfg.db.GetSublocationIDsWithFMVs),
	})
}
