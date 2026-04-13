package api

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type Junction struct {
	ParentID int32
	ChildID int32
}


// queries db for junctionPairs and converts them into a []Junction
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


// extracts the child IDs from a pre-sorted []Junction into a []int32 and removes the extracted pairs from the input slice
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

func juncMonsterArea(junction database.GetMonsterAreaIdPairsRow) Junction {
	return Junction{
		ParentID: 	junction.MonsterID,
		ChildID: 	junction.AreaID,
	}
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

func juncSublocationTreasure(junction database.GetSublocationTreasureIdPairsRow) Junction {
	return Junction{
		ParentID: 	junction.SublocationID,
		ChildID: 	junction.TreasureID,
	}
}

func juncSublocationShop(junction database.GetSublocationShopIdPairsRow) Junction {
	return Junction{
		ParentID: 	junction.SublocationID,
		ChildID: 	junction.ShopID,
	}
}

func juncSublocationMonster(junction database.GetSublocationMonsterIdPairsRow) Junction {
	return Junction{
		ParentID: 	junction.SublocationID,
		ChildID: 	junction.MonsterID,
	}
}

func juncLocationTreasure(junction database.GetLocationTreasureIdPairsRow) Junction {
	return Junction{
		ParentID: 	junction.LocationID,
		ChildID: 	junction.TreasureID,
	}
}

func juncLocationShop(junction database.GetLocationShopIdPairsRow) Junction {
	return Junction{
		ParentID: 	junction.LocationID,
		ChildID: 	junction.ShopID,
	}
}

func juncLocationMonster(junction database.GetLocationMonsterIdPairsRow) Junction {
	return Junction{
		ParentID: 	junction.LocationID,
		ChildID: 	junction.MonsterID,
	}
}

func juncAbilityRank(junction database.GetAbilityIdRankPairsRow) Junction {
	return Junction{
		ParentID: 	junction.AbilityID,
		ChildID: 	junction.Rank.Int32,
	}
}

func juncOverdriveAbilityRank(junction database.GetOverdriveAbilityIdRankPairsRow) Junction {
	return Junction{
		ParentID: 	junction.OverdriveAbilityID,
		ChildID: 	junction.Rank.Int32,
	}
}