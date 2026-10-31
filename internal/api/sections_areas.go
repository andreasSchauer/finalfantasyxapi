package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
	"golang.org/x/sync/errgroup"
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
	Availability      string              `json:"availability"`
	Shops             []ShopLocSimple     `json:"shops"`
	Treasures         *TreasuresLocSimple `json:"treasures"`
	Monsters          []string            `json:"monsters"`
}

func (a AreaSimple) GetURL() string {
	return a.URL
}

func createAreaSimple(cfg *Config, r *http.Request, id int32, subsection Subsection) (SimpleResource, error) {
	i := cfg.E.Areas
	area, _ := seeding.GetResourceByID(id, i.objLookupID)

	monsterIDs := subsection.relations[id][RelationMonsters]
	shopIDs := subsection.relations[id][RelationShops]

	areaSimple := AreaSimple{
		ID:                area.ID,
		URL:               CreateResourceURL(cfg, i.Endpoint, id),
		ParentLocation:    createSimpleRef(area.LocationID, area.Location, nil, nil),
		ParentSublocation: createSimpleRef(area.SublocationID, area.Sublocation, nil, nil),
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
	i := cfg.E.Areas
	relations := make(map[int32]map[Relation][]int32)
	g, ctx := errgroup.WithContext(r.Context())

	var treasureJunctions []Junction
	g.Go(func() error {
		var err error
		treasureJunctions, err = getDBJunctions(ctx, areaIDs, i.resTypeSingle, cfg.E.Treasures.resTypeSingle, cfg.db.GetAreaTreasureIdPairs, juncAreaTreasure)
		return err
	})

	var shopJunctions []Junction
	g.Go(func() error {
		var err error
		shopJunctions, err = getDBJunctions(ctx, areaIDs, i.resTypeSingle, cfg.E.Shops.resTypeSingle, cfg.db.GetAreaShopIdPairs, juncAreaShop)
		return err
	})

	var monsterJunctions []Junction
	g.Go(func() error {
		var err error
		monsterJunctions, err = getDBJunctions(ctx, areaIDs, i.resTypeSingle, cfg.E.Monsters.resTypeSingle, cfg.db.GetAreaMonsterIdPairs, juncAreaMonster)
		return err
	})

	err := g.Wait()
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
