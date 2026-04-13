package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type LocationSimple struct {
	ID        int32               `json:"id"`
	URL       string              `json:"url"`
	Name      string              `json:"name"`
	Shops     []ShopLocSimple     `json:"shops"`
	Treasures *TreasuresLocSimple `json:"treasures"`
	Monsters  []string            `json:"monsters"`
}

func (l LocationSimple) GetURL() string {
	return l.URL
}

func createLocationSimple(cfg *Config, r *http.Request, id int32, subsection Subsection) (SimpleResource, error) {
	i := cfg.e.locations
	location, _ := seeding.GetResourceByID(id, i.objLookupID)

	monsterIDs := subsection.relations[id][RelationMonsters]
	shopIDs := subsection.relations[id][RelationShops]

	locationSimple := LocationSimple{
		ID:        location.ID,
		URL:       createResourceURL(cfg, i.endpoint, id),
		Name:      location.Name,
		Shops:     convertObjSlice(cfg, shopIDs, idToShopLocSimple),
		Treasures: getTreasuresLocSimple(cfg, id, subsection),
		Monsters:  convertObjSlice(cfg, monsterIDs, idToMonsterSimpleString),
	}

	return locationSimple, nil
}

func getLocationSectionRelations(cfg *Config, r *http.Request, locIDs []int32) (map[int32]map[Relation][]int32, error) {
	i := cfg.e.locations
	relations := make(map[int32]map[Relation][]int32)

	treasureJunctions, err := getDbJunctions(r, locIDs, i.resourceType, cfg.e.treasures.resourceType, cfg.db.GetLocationTreasureIdPairs, juncLocationTreasure)
	if err != nil {
		return nil, err
	}

	shopJunctions, err := getDbJunctions(r, locIDs, i.resourceType, cfg.e.shops.resourceType, cfg.db.GetLocationShopIdPairs, juncLocationShop)
	if err != nil {
		return nil, err
	}

	monsterJunctions, err := getDbJunctions(r, locIDs, i.resourceType, cfg.e.monsters.resourceType, cfg.db.GetLocationMonsterIdPairs, juncLocationMonster)
	if err != nil {
		return nil, err
	}

	for _, locID := range locIDs {
		relationMap := make(map[Relation][]int32)

		relationMap[RelationTreasures], treasureJunctions = getJunctionIDs(locID, treasureJunctions)
		relationMap[RelationShops], shopJunctions = getJunctionIDs(locID, shopJunctions)
		relationMap[RelationMonsters], monsterJunctions = getJunctionIDs(locID, monsterJunctions)

		relations[locID] = relationMap
	}

	return relations, nil
}
