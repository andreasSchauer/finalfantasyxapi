package api

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/sectiontypes"
	"golang.org/x/sync/errgroup"
)


func createAreaSimple(r *http.Request, id int32, subsection Subsection) (apiresource.SimpleResource, error) {
	i := E.Areas
	area, _ := lookup.GetResourceByID(id, i.objLookupID)

	monsterIDs := subsection.relations[id][RelationMonsters]
	shopIDs := subsection.relations[id][RelationShops]

	areaSimple := sectiontypes.AreaSimple{
		ID:                area.ID,
		URL:               ah.CreateResourceURL(i.Endpoint, id),
		ParentLocation:    createSimpleRef(area.LocationID, area.Location, nil, nil),
		ParentSublocation: createSimpleRef(area.SublocationID, area.Sublocation, nil, nil),
		Name:              area.Name,
		Version:           area.Version,
		Specification:     area.Specification,
		HasSaveSphere:     area.HasSaveSphere,
		Availability:      area.Availability,
		Shops:             ah.ConvertObjSlice(shopIDs, idToShopLocSimple),
		Treasures:         getTreasuresLocSimple(id, subsection),
		Monsters:          ah.ConvertObjSlice(monsterIDs, idToMonsterSimpleString),
	}

	return areaSimple, nil
}

func getAreaSectionRelations(r *http.Request, areaIDs []int32) (map[int32]map[Relation][]int32, error) {
	i := E.Areas
	relations := make(map[int32]map[Relation][]int32)
	g, ctx := errgroup.WithContext(r.Context())

	var treasureJunctions []sectiontypes.Junction
	g.Go(func() error {
		var err error
		treasureJunctions, err = i.GetDBJunctions(ctx, areaIDs, E.Treasures.resTypeSingle, config.Cfg.DB.GetAreaTreasureIdPairs, juncAreaTreasure)
		return err
	})

	var shopJunctions []sectiontypes.Junction
	g.Go(func() error {
		var err error
		shopJunctions, err = i.GetDBJunctions(ctx, areaIDs, E.Shops.resTypeSingle, config.Cfg.DB.GetAreaShopIdPairs, juncAreaShop)
		return err
	})

	var monsterJunctions []sectiontypes.Junction
	g.Go(func() error {
		var err error
		monsterJunctions, err = i.GetDBJunctions(ctx, areaIDs, E.Monsters.resTypeSingle, config.Cfg.DB.GetAreaMonsterIdPairs, juncAreaMonster)
		return err
	})

	err := g.Wait()
	if err != nil {
		return nil, err
	}

	for _, areaID := range areaIDs {
		relationMap := make(map[Relation][]int32)

		relationMap[RelationTreasures], treasureJunctions = GetJunctionIDs(areaID, treasureJunctions)
		relationMap[RelationShops], shopJunctions = GetJunctionIDs(areaID, shopJunctions)
		relationMap[RelationMonsters], monsterJunctions = GetJunctionIDs(areaID, monsterJunctions)

		relations[areaID] = relationMap
	}

	return relations, nil
}
