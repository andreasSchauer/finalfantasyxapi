package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
	"golang.org/x/sync/errgroup"
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
	i := cfg.E.Sublocations
	sublocation, _ := seeding.GetResourceByID(id, i.objLookupID)

	monsterIDs := subsection.relations[id][RelationMonsters]
	shopIDs := subsection.relations[id][RelationShops]

	sublocationSimple := SublocationSimple{
		ID:             sublocation.ID,
		URL:            CreateResourceURL(cfg, i.Endpoint, id),
		ParentLocation: createSimpleRef(sublocation.LocationID, sublocation.Location, nil, nil),
		Name:           sublocation.Name,
		Shops:          convertObjSlice(cfg, shopIDs, idToShopLocSimple),
		Treasures:      getTreasuresLocSimple(cfg, id, subsection),
		Monsters:       convertObjSlice(cfg, monsterIDs, idToMonsterSimpleString),
	}

	return sublocationSimple, nil
}

func getSublocationSectionRelations(cfg *Config, r *http.Request, subLocIDs []int32) (map[int32]map[Relation][]int32, error) {
	i := cfg.E.Sublocations
	relations := make(map[int32]map[Relation][]int32)
	g, ctx := errgroup.WithContext(r.Context())

	var treasureJunctions []Junction
	g.Go(func() error {
		var err error
		treasureJunctions, err = getDBJunctions(ctx, subLocIDs, i.resTypeSingle, cfg.E.Treasures.resTypeSingle, cfg.db.GetSublocationTreasureIdPairs, juncSublocationTreasure)
		return err
	})

	var shopJunctions []Junction
	g.Go(func() error {
		var err error
		shopJunctions, err = getDBJunctions(ctx, subLocIDs, i.resTypeSingle, cfg.E.Shops.resTypeSingle, cfg.db.GetSublocationShopIdPairs, juncSublocationShop)
		return err
	})

	var monsterJunctions []Junction
	g.Go(func() error {
		var err error
		monsterJunctions, err = getDBJunctions(ctx, subLocIDs, i.resTypeSingle, cfg.E.Monsters.resTypeSingle, cfg.db.GetSublocationMonsterIdPairs, juncSublocationMonster)
		return err
	})

	err := g.Wait()
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
