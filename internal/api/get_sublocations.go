package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getSublocation(r *http.Request, i handlerInput[seeding.Sublocation, Sublocation, NamedAPIResource, NamedApiResourceList], id int32) (Sublocation, error) {
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
		ParentLocation:        nameToNamedAPIResource(cfg, cfg.e.locations, sublocation.Location.Name, nil),
		Availability:          sublocation.Availability,
		ConnectedSublocations: rel.ConnectedSublocations,
		Areas:                 rel.Areas,
		LocRel:                rel.LocRel,
	}

	return response, nil
}

func (cfg *Config) retrieveSublocations(r *http.Request, i handlerInput[seeding.Sublocation, Sublocation, NamedAPIResource, NamedApiResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []IdFilter{
		idQuery(r, i, ids, qpnLocation, cfg.l.Locations, cfg.db.GetLocationSublocationIDs),
		idQuery(r, i, ids, qpnMonster, cfg.l.Monsters, cfg.db.GetSublocationIDsWithMonster),
		idQueryWrapper(cfg, r, i, ids, qpnItem, cfg.l.Items, getSublocationsByItem),
		idQuery(r, i, ids, qpnKeyItem, cfg.l.KeyItems, cfg.db.GetSublocationIDsWithKeyItem),
		idQuery(r, i, ids, qpnAutoAbility, cfg.l.AutoAbilities, cfg.db.GetSublocationIDsWithAutoAbility),
		boolQuery2(r, i, ids, qpnCharacters, cfg.db.GetSublocationIDsWithCharacters),
		boolQuery2(r, i, ids, qpnAeons, cfg.db.GetSublocationIDsWithAeons),
		boolQuery2(r, i, ids, qpnMonsters, cfg.db.GetSublocationIDsWithMonsters),
		boolQuery2(r, i, ids, qpnBossFights, cfg.db.GetSublocationIDsWithBosses),
		boolQuery2(r, i, ids, qpnShops, cfg.db.GetSublocationIDsWithShops),
		boolQuery2(r, i, ids, qpnTreasures, cfg.db.GetSublocationIDsWithTreasures),
		boolQuery2(r, i, ids, qpnSidequests, cfg.db.GetSublocationIDsWithSidequests),
		boolQuery2(r, i, ids, qpnFMVs, cfg.db.GetSublocationIDsWithFMVs),
	})
}
