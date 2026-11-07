package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/sectiontypes"
)

func juncMonsterArea(junction database.GetMonsterAreaIdPairsRow) sectiontypes.Junction {
	return sectiontypes.Junction{
		ParentID: junction.MonsterID,
		ChildID:  junction.AreaID,
	}
}

func juncAreaTreasure(junction database.GetAreaTreasureIdPairsRow) sectiontypes.Junction {
	return sectiontypes.Junction{
		ParentID: junction.AreaID,
		ChildID:  junction.TreasureID,
	}
}

func juncAreaShop(junction database.GetAreaShopIdPairsRow) sectiontypes.Junction {
	return sectiontypes.Junction{
		ParentID: junction.AreaID,
		ChildID:  junction.ShopID,
	}
}

func juncAreaMonster(junction database.GetAreaMonsterIdPairsRow) sectiontypes.Junction {
	return sectiontypes.Junction{
		ParentID: junction.AreaID,
		ChildID:  junction.MonsterID,
	}
}

func juncSublocationTreasure(junction database.GetSublocationTreasureIdPairsRow) sectiontypes.Junction {
	return sectiontypes.Junction{
		ParentID: junction.SublocationID,
		ChildID:  junction.TreasureID,
	}
}

func juncSublocationShop(junction database.GetSublocationShopIdPairsRow) sectiontypes.Junction {
	return sectiontypes.Junction{
		ParentID: junction.SublocationID,
		ChildID:  junction.ShopID,
	}
}

func juncSublocationMonster(junction database.GetSublocationMonsterIdPairsRow) sectiontypes.Junction {
	return sectiontypes.Junction{
		ParentID: junction.SublocationID,
		ChildID:  junction.MonsterID,
	}
}

func juncLocationTreasure(junction database.GetLocationTreasureIdPairsRow) sectiontypes.Junction {
	return sectiontypes.Junction{
		ParentID: junction.LocationID,
		ChildID:  junction.TreasureID,
	}
}

func juncLocationShop(junction database.GetLocationShopIdPairsRow) sectiontypes.Junction {
	return sectiontypes.Junction{
		ParentID: junction.LocationID,
		ChildID:  junction.ShopID,
	}
}

func juncLocationMonster(junction database.GetLocationMonsterIdPairsRow) sectiontypes.Junction {
	return sectiontypes.Junction{
		ParentID: junction.LocationID,
		ChildID:  junction.MonsterID,
	}
}