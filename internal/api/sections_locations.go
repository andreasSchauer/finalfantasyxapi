package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
	"golang.org/x/sync/errgroup"
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
	g, ctx := errgroup.WithContext(r.Context())

	var treasureJunctions []Junction
	g.Go(func() error {
		var err error
		treasureJunctions, err = getDbJunctions(ctx, locIDs, i.resTypeSing, cfg.e.treasures.resTypeSing, cfg.db.GetLocationTreasureIdPairs, juncLocationTreasure)
		return err
	})
	
	var shopJunctions []Junction
	g.Go(func() error {
		var err error
		shopJunctions, err = getDbJunctions(ctx, locIDs, i.resTypeSing, cfg.e.shops.resTypeSing, cfg.db.GetLocationShopIdPairs, juncLocationShop)
		return err
	})
	
	var monsterJunctions []Junction
	g.Go(func() error {
		var err error
		monsterJunctions, err = getDbJunctions(ctx, locIDs, i.resTypeSing, cfg.e.monsters.resTypeSing, cfg.db.GetLocationMonsterIdPairs, juncLocationMonster)
		return err
	})
	

	err := g.Wait()
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
