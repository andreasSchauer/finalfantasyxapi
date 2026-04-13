package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type AreaSimple struct {
	ID                int32               `json:"id"`
	URL               string              `json:"url"`
	Name              string              `json:"name"`
	ParentLocation    SimpleRef           `json:"parent_location"`
	ParentSublocation SimpleRef           `json:"parent_sublocation"`
	Version           *int32              `json:"version,omitempty"`
	Specification     *string             `json:"specification,omitempty"`
	HasSaveSphere     bool                `json:"has_save_sphere"`
	Availability	  string			  `json:"availability"`
	Shops             []ShopLocSimple     `json:"shops"`
	Treasures         *TreasuresLocSimple `json:"treasures"`
	Monsters          []string            `json:"monsters"`
}

func (a AreaSimple) GetURL() string {
	return a.URL
}

func createAreaSimple(cfg *Config, r *http.Request, id int32, subsection Subsection) (SimpleResource, error) {
	i := cfg.e.areas
	area, _ := seeding.GetResourceByID(id, i.objLookupID)

	monsterIDs := subsection.relations[id][RelationMonsters]
	shopIDs := subsection.relations[id][RelationShops]

	areaSimple := AreaSimple{
		ID:                area.ID,
		URL:               createResourceURL(cfg, i.endpoint, id),
		ParentLocation:    createSimpleRef(area.Sublocation.Location.ID, area.Sublocation.Location.Name, nil, nil),
		ParentSublocation: createSimpleRef(area.Sublocation.ID, area.Sublocation.Name, nil, nil),
		Name:              area.Name,
		Version:           area.Version,
		Specification:     area.Specification,
		HasSaveSphere:     area.HasSaveSphere,
		Availability:      area.Availability,
		Shops:             convertObjSlice(cfg, shopIDs, idToShopLocSimple),
		Treasures:         getTreasuresLocSimple(cfg, id, subsection),
		Monsters:          convertObjSlice(cfg, monsterIDs, idToMonsterSimpleString),
	}

	return areaSimple, nil
}

func getAreaSectionRelations(cfg *Config, r *http.Request, areaIDs []int32) (map[int32]map[Relation][]int32, error) {
	i := cfg.e.areas
	relations := make(map[int32]map[Relation][]int32)

	treasureJunctions, err := getJunctions(r, areaIDs, i.resourceType, cfg.e.treasures.resourceType, cfg.db.GetAreaTreasureIdPairs, juncAreaTreasure)
	if err != nil {
		return nil, err
	}
	
	shopJunctions, err := getJunctions(r, areaIDs, i.resourceType, cfg.e.shops.resourceType, cfg.db.GetAreaShopIdPairs, juncAreaShop)
	if err != nil {
		return nil, err
	}
	
	monsterJunctions, err := getJunctions(r, areaIDs, i.resourceType, cfg.e.monsters.resourceType, cfg.db.GetAreaMonsterIdPairs, juncAreaMonster)
	if err != nil {
		return nil, err
	}

	for _, areaID := range areaIDs {
		relationMap := make(map[Relation][]int32)

		relationMap[RelationTreasures], treasureJunctions = getJunctionIDs(areaID, treasureJunctions)
		relationMap[RelationShops], shopJunctions = getJunctionIDs(areaID, shopJunctions)
		relationMap[RelationMonsters], monsterJunctions = getJunctionIDs(areaID, monsterJunctions)

		relations[areaID] = relationMap
	}

	return relations, nil
}