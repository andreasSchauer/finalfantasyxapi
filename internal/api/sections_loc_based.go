package api

import (
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/sectiontypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func idToShopLocSimple(shopID int32) sectiontypes.ShopLocSimple {
	shopLookup, _ := lookup.GetResourceByID(shopID, lookup.L.ShopsID)
	return sectiontypes.ShopLocSimple{
		ID:          shopID,
		Category:    database.ShopCategory(shopLookup.Category),
		PreAirship:  ah.ConvertObjPtr(shopLookup.PreAirship, convertShopSummarySimple),
		PostAirship: ah.ConvertObjPtr(shopLookup.PostAirship, convertShopSummarySimple),
	}
}

func convertShopSummarySimple(shop seedtypes.SubShop) sectiontypes.ShopSummarySimple {
	shopLoc := sectiontypes.ShopSummarySimple{}

	if len(shop.Items) != 0 {
		shopLoc.HasItems = true
	}

	if len(shop.Equipment) != 0 {
		shopLoc.HasEquipment = true
	}

	return shopLoc
}

func getTreasuresLocSimple(id int32, subsection handlerinput.Subsection) *sectiontypes.TreasuresLocSimple {
	treasureIDs := subsection.Relations[id][handlerinput.RelationTreasures]

	if len(treasureIDs) == 0 {
		return nil
	}

	treasures := populateTreasuresLocSimple(treasureIDs)
	return &treasures
}

func populateTreasuresLocSimple(treasureIDs []int32) sectiontypes.TreasuresLocSimple {
	treasures := sectiontypes.TreasuresLocSimple{
		TreasureCount: len(treasureIDs),
	}

	itemAmounts := []seedtypes.ItemAmount{}

	for _, treasureID := range treasureIDs {
		treasure, _ := lookup.GetResourceByID(treasureID, lookup.L.TreasuresID)

		switch treasure.LootType {
		case string(database.LootTypeGil):
			treasures.TotalGil += *treasure.GilAmount

		case string(database.LootTypeItem):
			itemAmounts = append(itemAmounts, treasure.Items...)

		case string(database.LootTypeEquipment):
			equipment := treasure.Equipment
			es := convertEquipmentSimple(*equipment)
			treasures.Equipment = append(treasures.Equipment, es)
		}
	}

	itemAmounts = sortItemAmountsByID(itemAmounts)
	treasures.Items = ah.ConvertObjSlice(itemAmounts, convertItemAmountSimple)
	return treasures
}

func convertEquipmentSimple(equipment seedtypes.TreasureEquipment) sectiontypes.EquipmentSimple {
	return sectiontypes.EquipmentSimple{
		Name:          equipment.Name,
		AutoAbilities: foundEquipmentAbilitiesStringPtr(equipment),
	}
}
