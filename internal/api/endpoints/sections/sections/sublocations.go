package sections

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/sections/sectiontypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"golang.org/x/sync/errgroup"
)

func CreateSublocationSimple(r *http.Request, id int32, subsection handlerinput.Subsection) (apiresource.SimpleResource, error) {
	i := apilookups.E.Sublocations
	sublocation, _ := lookup.GetResourceByID(id, i.ObjLookupID)

	monsterIDs := subsection.Relations[id][handlerinput.RelationMonsters]
	shopIDs := subsection.Relations[id][handlerinput.RelationShops]

	sublocationSimple := sectiontypes.SublocationSimple{
		ID:             sublocation.ID,
		URL:            ah.CreateResourceURL(i.Endpoint, id),
		ParentLocation: createSimpleRef(sublocation.LocationID, sublocation.Location, nil, nil),
		Name:           sublocation.Name,
		Shops:          ah.ConvertObjSlice(shopIDs, idToShopLocSimple),
		Treasures:      getTreasuresLocSimple(id, subsection),
		Monsters:       ah.ConvertObjSlice(monsterIDs, idToMonsterSimpleString),
	}

	return sublocationSimple, nil
}

func GetSublocationSectionRelations(r *http.Request, subLocIDs []int32) (map[int32]map[handlerinput.Relation][]int32, error) {
	i := apilookups.E.Sublocations
	relations := make(map[int32]map[handlerinput.Relation][]int32)
	g, ctx := errgroup.WithContext(r.Context())

	var treasureJunctions []sectiontypes.Junction
	g.Go(func() error {
		var err error
		treasureJunctions, err = i.GetDBJunctions(ctx, subLocIDs, apilookups.E.Treasures.ResTypeSingle, config.Cfg.DB.GetSublocationTreasureIdPairs, juncSublocationTreasure)
		return err
	})

	var shopJunctions []sectiontypes.Junction
	g.Go(func() error {
		var err error
		shopJunctions, err = i.GetDBJunctions(ctx, subLocIDs, apilookups.E.Shops.ResTypeSingle, config.Cfg.DB.GetSublocationShopIdPairs, juncSublocationShop)
		return err
	})

	var monsterJunctions []sectiontypes.Junction
	g.Go(func() error {
		var err error
		monsterJunctions, err = i.GetDBJunctions(ctx, subLocIDs, apilookups.E.Monsters.ResTypeSingle, config.Cfg.DB.GetSublocationMonsterIdPairs, juncSublocationMonster)
		return err
	})

	err := g.Wait()
	if err != nil {
		return nil, err
	}

	for _, subLocID := range subLocIDs {
		relationMap := make(map[handlerinput.Relation][]int32)

		relationMap[handlerinput.RelationTreasures], treasureJunctions = GetJunctionIDs(subLocID, treasureJunctions)
		relationMap[handlerinput.RelationShops], shopJunctions = GetJunctionIDs(subLocID, shopJunctions)
		relationMap[handlerinput.RelationMonsters], monsterJunctions = GetJunctionIDs(subLocID, monsterJunctions)

		relations[subLocID] = relationMap
	}

	return relations, nil
}
