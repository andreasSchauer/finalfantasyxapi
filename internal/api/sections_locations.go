package api

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api_helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
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

func createLocationSimple(r *http.Request, id int32, subsection Subsection) (SimpleResource, error) {
	i := E.Locations
	location, _ := lookup.GetResourceByID(id, i.objLookupID)

	monsterIDs := subsection.relations[id][RelationMonsters]
	shopIDs := subsection.relations[id][RelationShops]

	locationSimple := LocationSimple{
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

	var treasureJunctions []Junction
	g.Go(func() error {
		var err error
		treasureJunctions, err = getDBJunctions(ctx, locIDs, i.resTypeSingle, E.Treasures.resTypeSingle, config.Cfg.DB.GetLocationTreasureIdPairs, juncLocationTreasure)
		return err
	})

	var shopJunctions []Junction
	g.Go(func() error {
		var err error
		shopJunctions, err = getDBJunctions(ctx, locIDs, i.resTypeSingle, E.Shops.resTypeSingle, config.Cfg.DB.GetLocationShopIdPairs, juncLocationShop)
		return err
	})

	var monsterJunctions []Junction
	g.Go(func() error {
		var err error
		monsterJunctions, err = getDBJunctions(ctx, locIDs, i.resTypeSingle, E.Monsters.resTypeSingle, config.Cfg.DB.GetLocationMonsterIdPairs, juncLocationMonster)
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
