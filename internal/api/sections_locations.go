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


func createLocationSimple(r *http.Request, id int32, subsection Subsection) (apiresource.SimpleResource, error) {
	i := E.Locations
	location, _ := lookup.GetResourceByID(id, i.objLookupID)

	monsterIDs := subsection.relations[id][RelationMonsters]
	shopIDs := subsection.relations[id][RelationShops]

	locationSimple := sectiontypes.LocationSimple{
		ID:        location.ID,
		URL:       ah.CreateResourceURL(i.Endpoint, id),
		Name:      location.Name,
		Shops:     ah.ConvertObjSlice(shopIDs, idToShopLocSimple),
		Treasures: getTreasuresLocSimple(id, subsection),
		Monsters:  ah.ConvertObjSlice(monsterIDs, idToMonsterSimpleString),
	}

	return locationSimple, nil
}

func getLocationSectionRelations(r *http.Request, locIDs []int32) (map[int32]map[Relation][]int32, error) {
	i := E.Locations
	relations := make(map[int32]map[Relation][]int32)
	g, ctx := errgroup.WithContext(r.Context())

	var treasureJunctions []sectiontypes.Junction
	g.Go(func() error {
		var err error
		treasureJunctions, err = i.GetDBJunctions(ctx, locIDs, E.Treasures.resTypeSingle, config.Cfg.DB.GetLocationTreasureIdPairs, juncLocationTreasure)
		return err
	})

	var shopJunctions []sectiontypes.Junction
	g.Go(func() error {
		var err error
		shopJunctions, err = i.GetDBJunctions(ctx, locIDs, E.Shops.resTypeSingle, config.Cfg.DB.GetLocationShopIdPairs, juncLocationShop)
		return err
	})

	var monsterJunctions []sectiontypes.Junction
	g.Go(func() error {
		var err error
		monsterJunctions, err = i.GetDBJunctions(ctx, locIDs, E.Monsters.resTypeSingle, config.Cfg.DB.GetLocationMonsterIdPairs, juncLocationMonster)
		return err
	})

	err := g.Wait()
	if err != nil {
		return nil, err
	}

	for _, locID := range locIDs {
		relationMap := make(map[Relation][]int32)

		relationMap[RelationTreasures], treasureJunctions = GetJunctionIDs(locID, treasureJunctions)
		relationMap[RelationShops], shopJunctions = GetJunctionIDs(locID, shopJunctions)
		relationMap[RelationMonsters], monsterJunctions = GetJunctionIDs(locID, monsterJunctions)

		relations[locID] = relationMap
	}

	return relations, nil
}
