package sections

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/sectiontypes"
	"golang.org/x/sync/errgroup"
)

func CreateLocationSimple(r *http.Request, id int32, subsection handlerinput.Subsection) (apiresource.SimpleResource, error) {
	i := apilookups.E.Locations
	location, _ := lookup.GetResourceByID(id, i.ObjLookupID)

	monsterIDs := subsection.Relations[id][handlerinput.RelationMonsters]
	shopIDs := subsection.Relations[id][handlerinput.RelationShops]

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

func GetLocationSectionRelations(r *http.Request, locIDs []int32) (map[int32]map[handlerinput.Relation][]int32, error) {
	i := apilookups.E.Locations
	relations := make(map[int32]map[handlerinput.Relation][]int32)
	g, ctx := errgroup.WithContext(r.Context())

	var treasureJunctions []sectiontypes.Junction
	g.Go(func() error {
		var err error
		treasureJunctions, err = i.GetDBJunctions(ctx, locIDs, apilookups.E.Treasures.ResTypeSingle, config.Cfg.DB.GetLocationTreasureIdPairs, juncLocationTreasure)
		return err
	})

	var shopJunctions []sectiontypes.Junction
	g.Go(func() error {
		var err error
		shopJunctions, err = i.GetDBJunctions(ctx, locIDs, apilookups.E.Shops.ResTypeSingle, config.Cfg.DB.GetLocationShopIdPairs, juncLocationShop)
		return err
	})

	var monsterJunctions []sectiontypes.Junction
	g.Go(func() error {
		var err error
		monsterJunctions, err = i.GetDBJunctions(ctx, locIDs, apilookups.E.Monsters.ResTypeSingle, config.Cfg.DB.GetLocationMonsterIdPairs, juncLocationMonster)
		return err
	})

	err := g.Wait()
	if err != nil {
		return nil, err
	}

	for _, locID := range locIDs {
		relationMap := make(map[handlerinput.Relation][]int32)

		relationMap[handlerinput.RelationTreasures], treasureJunctions = GetJunctionIDs(locID, treasureJunctions)
		relationMap[handlerinput.RelationShops], shopJunctions = GetJunctionIDs(locID, shopJunctions)
		relationMap[handlerinput.RelationMonsters], monsterJunctions = GetJunctionIDs(locID, monsterJunctions)

		relations[locID] = relationMap
	}

	return relations, nil
}
