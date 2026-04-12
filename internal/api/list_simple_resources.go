package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type SimpleResourceList struct {
	ListParams
	ParentResource APIResource      `json:"parent_resource,omitempty"`
	Results        []SimpleResource `json:"results"`
}

func (l SimpleResourceList) getListParams() ListParams {
	return l.ListParams
}

func (l SimpleResourceList) getResults() []SimpleResource {
	return l.Results
}

// I have the entire SubSectionFns here, so before calling createSimpleResources, I can create a relationship map
// maybe rename SubSectionFns into SubSection
// giveSubsection a map[id]map[resource_Type][]int32 or something like that?
// populating the map works like this:
//
//	ofo
//
// pass the entire struct into createSimpleResources (and the constructor), not just the function
func newSimpleResourceList[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], id int32, sectionName string, subsection Subsection) (SimpleResourceList, error) {
	dbIDs, err := subsection.dbQuery(r.Context(), id)
	if err != nil {
		return SimpleResourceList{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't retrieve %s of %s with id '%d'", sectionName, i.resourceType, id), err)
	}

	results, err := createSimpleResources(cfg, r, dbIDs, subsection)
	if err != nil {
		return SimpleResourceList{}, err
	}

	listParams, shownResults, err := createPaginatedList(cfg, r, results)
	if err != nil {
		return SimpleResourceList{}, err
	}

	subResList := SimpleResourceList{
		ListParams:     listParams,
		ParentResource: i.idToResFunc(cfg, i, id),
		Results:        shownResults,
	}

	return subResList, nil
}

func createSimpleResources(cfg *Config, r *http.Request, dbIDs []int32, subsection Subsection) ([]SimpleResource, error) {
	subs := []SimpleResource{}

	// create the relationMap here, if the function for it is not nil
	// createSubFn should take the relationMap with it

	for _, id := range dbIDs {
		subRes, err := subsection.createSubFn(cfg, r, id)
		if err != nil {
			return nil, err
		}

		subs = append(subs, subRes)
	}

	return subs, nil
}


func getAreaSectionRelations(cfg *Config, r *http.Request, areaIDs []int32) (map[int32]map[Relation][]int32, error) {
	i := cfg.e.areas
	relations := make(map[int32]map[Relation][]int32)

	treasureJunctions, err := getJunctions(r, areaIDs, i.resourceType, cfg.e.treasures.resourceType, cfg.db.GetAreaTreasureIdPairs, juncAreaTreasure)
	if err != nil {
		return nil, err
	}
	
	shopJunctions, err := getJunctions(r, areaIDs, i.resourceType, cfg.e.shops.resourceType, cfg.db.GetAreaShopIdPairs, juncAreaShop)
	if err != nil {
		return nil, err
	}
	
	monsterJunctions, err := getJunctions(r, areaIDs, i.resourceType, cfg.e.monsters.resourceType, cfg.db.GetAreaMonsterIdPairs, juncAreaMonster)
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

type Junction struct {
	ParentID int32
	ChildID int32
}

func juncAreaTreasure(junction database.GetAreaTreasureIdPairsRow) Junction {
	return Junction{
		ParentID: 	junction.AreaID,
		ChildID: 	junction.TreasureID,
	}
}

func juncAreaShop(junction database.GetAreaShopIdPairsRow) Junction {
	return Junction{
		ParentID: 	junction.AreaID,
		ChildID: 	junction.ShopID,
	}
}

func juncAreaMonster(junction database.GetAreaMonsterIdPairsRow) Junction {
	return Junction{
		ParentID: 	junction.AreaID,
		ChildID: 	junction.MonsterID,
	}
}

func getJunctions[R any](r *http.Request, ids []int32, parentResType, childResType string, dbQuery func (context.Context, []int32) ([]R, error), converter func(R) Junction) ([]Junction, error) {
	dbJunctions, err := dbQuery(r.Context(), ids)
	if err != nil {
		return nil, newHTTPErrorDbPairs(parentResType, childResType, err)
	}

	junctions := []Junction{}

	for _, dbJunction := range dbJunctions {
		junctions = append(junctions, converter(dbJunction))
	}

	return junctions, nil
}

func getJunctionIDs(parentID int32, junctions []Junction) ([]int32, []Junction) {
	ids := []int32{}

	for i, junction := range junctions {
		if junction.ParentID != parentID {
			return ids, junctions[i:]
		}
		ids = append(ids, junction.ChildID)
	}

	return ids, nil
}