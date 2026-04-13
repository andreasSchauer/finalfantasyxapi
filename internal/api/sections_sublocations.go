package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type SublocationSimple struct {
	ID             int32               `json:"id"`
	URL            string              `json:"url"`
	Name           string              `json:"name"`
	ParentLocation SimpleRef           `json:"parent_location"`
	Shops          []ShopLocSimple     `json:"shops"`
	Treasures      *TreasuresLocSimple `json:"treasures"`
	Monsters       []string            `json:"monsters"`
}

func (s SublocationSimple) GetURL() string {
	return s.URL
}

func createSublocationSimple(cfg *Config, r *http.Request, id int32, subsection Subsection) (SimpleResource, error) {
	i := cfg.e.sublocations
	sublocation, _ := seeding.GetResourceByID(id, i.objLookupID)

	monsterIDs := subsection.relations[id][RelationMonsters]
	shopIDs := subsection.relations[id][RelationShops]

	sublocationSimple := SublocationSimple{
		ID:             sublocation.ID,
		URL:            createResourceURL(cfg, i.endpoint, id),
		ParentLocation: createSimpleRef(sublocation.Location.ID, sublocation.Location.Name, nil, nil),
		Name:           sublocation.Name,
		Shops:          convertObjSlice(cfg, shopIDs, idToShopLocSimple),
		Treasures:      getTreasuresLocSimple(cfg, id, subsection),
		Monsters:       convertObjSlice(cfg, monsterIDs, idToMonsterSimpleString),
	}

	return sublocationSimple, nil
}

func getSublocationSectionRelations(cfg *Config, r *http.Request, subLocIDs []int32) (map[int32]map[Relation][]int32, error) {
	i := cfg.e.sublocations
	relations := make(map[int32]map[Relation][]int32)

	treasureJunctions, err := getDbJunctions(r, subLocIDs, i.resourceType, cfg.e.treasures.resourceType, cfg.db.GetSublocationTreasureIdPairs, juncSublocationTreasure)
	if err != nil {
		return nil, err
	}

	shopJunctions, err := getDbJunctions(r, subLocIDs, i.resourceType, cfg.e.shops.resourceType, cfg.db.GetSublocationShopIdPairs, juncSublocationShop)
	if err != nil {
		return nil, err
	}

	monsterJunctions, err := getDbJunctions(r, subLocIDs, i.resourceType, cfg.e.monsters.resourceType, cfg.db.GetSublocationMonsterIdPairs, juncSublocationMonster)
	if err != nil {
		return nil, err
	}

	for _, subLocID := range subLocIDs {
		relationMap := make(map[Relation][]int32)

		relationMap[RelationTreasures], treasureJunctions = getJunctionIDs(subLocID, treasureJunctions)
		relationMap[RelationShops], shopJunctions = getJunctionIDs(subLocID, shopJunctions)
		relationMap[RelationMonsters], monsterJunctions = getJunctionIDs(subLocID, monsterJunctions)

		relations[subLocID] = relationMap
	}

	return relations, nil
}
