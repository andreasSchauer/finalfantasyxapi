package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type Sublocation struct {
	ID                    int32              `json:"id"`
	Name                  string             `json:"name"`
	ParentLocation        NamedAPIResource   `json:"parent_location"`
	ConnectedSublocations []NamedAPIResource `json:"connected_sublocations"`
	Areas                 []AreaAPIResource  `json:"areas"`
	LocRel
}

func (cfg *Config) getSublocation(r *http.Request, i handlerInput[seeding.Sublocation, Sublocation, NamedAPIResource, NamedApiResourceList], id int32) (Sublocation, error) {
	sublocation, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return Sublocation{}, err
	}

	connectedSublocations, err := getResourcesDB(cfg, r, cfg.e.sublocations, sublocation, cfg.db.GetConnectedSublocationIDs)
	if err != nil {
		return Sublocation{}, err
	}

	areas, err := getResourcesDB(cfg, r, cfg.e.areas, sublocation, cfg.db.GetSublocationAreaIDs)
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
		ConnectedSublocations: connectedSublocations,
		Areas:                 areas,
		LocRel:                rel,
	}

	return response, nil
}

func (cfg *Config) retrieveSublocations(r *http.Request, i handlerInput[seeding.Sublocation, Sublocation, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[NamedAPIResource]{
		frl(idQuery(cfg, r, i, resources, "location", len(cfg.l.Locations), cfg.db.GetLocationSublocationIDs)),
		frl(idQueryWrapper(cfg, r, i, resources, "item", len(cfg.l.Items), getSublocationsByItem)),
		frl(idQueryWrapper(cfg, r, i, resources, "key_item", len(cfg.l.KeyItems), getSublocationsByKeyItem)),
		frl(boolQuery2(cfg, r, i, resources, "characters", cfg.db.GetSublocationIDsWithCharacters)),
		frl(boolQuery2(cfg, r, i, resources, "aeons", cfg.db.GetSublocationIDsWithAeons)),
		frl(boolQuery2(cfg, r, i, resources, "monsters", cfg.db.GetSublocationIDsWithMonsters)),
		frl(boolQuery2(cfg, r, i, resources, "boss_fights", cfg.db.GetSublocationIDsWithBosses)),
		frl(boolQuery2(cfg, r, i, resources, "shops", cfg.db.GetSublocationIDsWithShops)),
		frl(boolQuery2(cfg, r, i, resources, "treasures", cfg.db.GetSublocationIDsWithTreasures)),
		frl(boolQuery2(cfg, r, i, resources, "sidequests", cfg.db.GetSublocationIDsWithSidequests)),
		frl(boolQuery2(cfg, r, i, resources, "fmvs", cfg.db.GetSublocationIDsWithFMVs)),
	})
}
