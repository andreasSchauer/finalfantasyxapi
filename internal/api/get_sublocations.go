package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getSublocation(r *http.Request, i handlerInput[seeding.Sublocation, Sublocation, NamedAPIResource, NamedApiResourceList], id int32) (Sublocation, error) {
	sublocation, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return Sublocation{}, err
	}

	connectedSublocations, err := getResourcesDbItem(cfg, r, cfg.e.sublocations, sublocation, cfg.db.GetConnectedSublocationIDs)
	if err != nil {
		return Sublocation{}, err
	}

	areas, err := getResourcesDbItem(cfg, r, cfg.e.areas, sublocation, cfg.db.GetSublocationAreaIDs)
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
		Availability:          enumToNamedAPIResource(cfg, cfg.e.availabilityType.endpoint, sublocation.Availability, cfg.t.AvailabilityType),
		ConnectedSublocations: connectedSublocations,
		Areas:                 areas,
		LocRel:                rel,
	}

	return response, nil
}

func (cfg *Config) retrieveSublocations(r *http.Request, i handlerInput[seeding.Sublocation, Sublocation, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	ids, err := verifyParamsAndRetrieve(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterIDs(cfg, r, i, ids, []filteredIdList{
		fidl(idQuery(r, i, ids, qpnLocation, cfg.l.Locations, cfg.db.GetLocationSublocationIDs)),
		fidl(idQuery(r, i, ids, qpnMonster, cfg.l.Monsters, cfg.db.GetSublocationIDsWithMonster)),
		fidl(idQueryWrapper(cfg, r, i, ids, qpnItem, cfg.l.Items, getSublocationsByItem)),
		fidl(idQuery(r, i, ids, qpnKeyItem, cfg.l.KeyItems, cfg.db.GetSublocationIDsWithKeyItem)),
		fidl(idQuery(r, i, ids, qpnAutoAbility, cfg.l.AutoAbilities, cfg.db.GetSublocationIDsWithAutoAbility)),
		fidl(boolQuery2(r, i, ids, qpnCharacters, cfg.db.GetSublocationIDsWithCharacters)),
		fidl(boolQuery2(r, i, ids, qpnAeons, cfg.db.GetSublocationIDsWithAeons)),
		fidl(boolQuery2(r, i, ids, qpnMonsters, cfg.db.GetSublocationIDsWithMonsters)),
		fidl(boolQuery2(r, i, ids, qpnBossFights, cfg.db.GetSublocationIDsWithBosses)),
		fidl(boolQuery2(r, i, ids, qpnShops, cfg.db.GetSublocationIDsWithShops)),
		fidl(boolQuery2(r, i, ids, qpnTreasures, cfg.db.GetSublocationIDsWithTreasures)),
		fidl(boolQuery2(r, i, ids, qpnSidequests, cfg.db.GetSublocationIDsWithSidequests)),
		fidl(boolQuery2(r, i, ids, qpnFMVs, cfg.db.GetSublocationIDsWithFMVs)),
	})
}
