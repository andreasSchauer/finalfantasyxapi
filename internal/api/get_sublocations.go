package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getSublocation(r *http.Request, i handlerInput[seeding.Sublocation, Sublocation, NamedAPIResource, NamedAPIResourceList], id int32) (Sublocation, error) {
	sublocation, err := verifyParamsAndGet(r, i, id)
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

func (cfg *Config) retrieveSublocations(r *http.Request, i handlerInput[seeding.Sublocation, Sublocation, NamedAPIResource, NamedAPIResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []IDFilter{
		idQuery(r, i, ids, QpnLocation, cfg.L.Locations, cfg.db.GetLocationSublocationIDs),
		idQuery(r, i, ids, QpnMonster, cfg.L.Monsters, cfg.db.GetSublocationIDsWithMonster),
		idQueryWrapper(cfg, r, i, ids, QpnItem, cfg.L.Items, getSublocationsByItem),
		idQuery(r, i, ids, QpnKeyItem, cfg.L.KeyItems, cfg.db.GetSublocationIDsWithKeyItem),
		idQuery(r, i, ids, QpnAutoAbility, cfg.L.AutoAbilities, cfg.db.GetSublocationIDsWithAutoAbility),
		boolQuery2(r, i, ids, QpnCharacters, cfg.db.GetSublocationIDsWithCharacters),
		boolQuery2(r, i, ids, QpnAeons, cfg.db.GetSublocationIDsWithAeons),
		boolQuery2(r, i, ids, QpnMonsters, cfg.db.GetSublocationIDsWithMonsters),
		boolQuery2(r, i, ids, QpnBossFights, cfg.db.GetSublocationIDsWithBosses),
		boolQuery2(r, i, ids, QpnShops, cfg.db.GetSublocationIDsWithShops),
		boolQuery2(r, i, ids, QpnTreasures, cfg.db.GetSublocationIDsWithTreasures),
		boolQuery2(r, i, ids, QpnSidequests, cfg.db.GetSublocationIDsWithSidequests),
		boolQuery2(r, i, ids, QpnFMVs, cfg.db.GetSublocationIDsWithFMVs),
	})
}
