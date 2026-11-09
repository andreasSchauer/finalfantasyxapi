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

func CreateAreaSimple(r *http.Request, id int32, subsection handlerinput.Subsection) (apiresource.SimpleResource, error) {
	i := apilookups.E.Areas
	area, _ := lookup.GetResourceByID(id, i.ObjLookupID)

	monsterIDs := subsection.Relations[id][handlerinput.RelationMonsters]
	shopIDs := subsection.Relations[id][handlerinput.RelationShops]

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

func GetAreaSectionRelations(r *http.Request, areaIDs []int32) (map[int32]map[handlerinput.Relation][]int32, error) {
	i := apilookups.E.Areas
	relations := make(map[int32]map[handlerinput.Relation][]int32)
	g, ctx := errgroup.WithContext(r.Context())

	var treasureJunctions []sectiontypes.Junction
	g.Go(func() error {
		var err error
		treasureJunctions, err = i.GetDBJunctions(ctx, areaIDs, apilookups.E.Treasures.ResTypeSingle, config.Cfg.DB.GetAreaTreasureIdPairs, juncAreaTreasure)
		return err
	})

	var shopJunctions []sectiontypes.Junction
	g.Go(func() error {
		var err error
		shopJunctions, err = i.GetDBJunctions(ctx, areaIDs, apilookups.E.Shops.ResTypeSingle, config.Cfg.DB.GetAreaShopIdPairs, juncAreaShop)
		return err
	})

	var monsterJunctions []sectiontypes.Junction
	g.Go(func() error {
		var err error
		monsterJunctions, err = i.GetDBJunctions(ctx, areaIDs, apilookups.E.Monsters.ResTypeSingle, config.Cfg.DB.GetAreaMonsterIdPairs, juncAreaMonster)
		return err
	})

	err := g.Wait()
	if err != nil {
		return nil, err
	}

	for _, areaID := range areaIDs {
		relationMap := make(map[handlerinput.Relation][]int32)

		relationMap[handlerinput.RelationTreasures], treasureJunctions = GetJunctionIDs(areaID, treasureJunctions)
		relationMap[handlerinput.RelationShops], shopJunctions = GetJunctionIDs(areaID, shopJunctions)
		relationMap[handlerinput.RelationMonsters], monsterJunctions = GetJunctionIDs(areaID, monsterJunctions)

		relations[areaID] = relationMap
	}

	return relations, nil
}
