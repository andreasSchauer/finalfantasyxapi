package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/database"


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