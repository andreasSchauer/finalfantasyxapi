package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

type ShopLocSimple struct {
	ID          int32                 `json:"id"`
	Category    database.ShopCategory `json:"category"`
	PreAirship  *ShopSummarySimple    `json:"pre_airship"`
	PostAirship *ShopSummarySimple    `json:"post_airship"`
}

type ShopSummarySimple struct {
	HasItems     bool `json:"has_items"`
	HasEquipment bool `json:"has_equipment"`
}

func idToShopLocSimple(shopID int32) ShopLocSimple {
	shopLookup, _ := lookup.GetResourceByID(shopID, lookup.L.ShopsID)
	return ShopLocSimple{
		ID:          shopID,
		Category:    database.ShopCategory(shopLookup.Category),
		PreAirship:  convertObjPtr(shopLookup.PreAirship, convertShopSummarySimple),
		PostAirship: convertObjPtr(shopLookup.PostAirship, convertShopSummarySimple),
	}
}

func convertShopSummarySimple(shop seedtypes.SubShop) ShopSummarySimple {
	shopLoc := ShopSummarySimple{}

	if len(shop.Items) != 0 {
		shopLoc.HasItems = true
	}

	if len(shop.Equipment) != 0 {
		shopLoc.HasEquipment = true
	}

	return shopLoc
}

type TreasuresLocSimple struct {
	TreasureCount int               `json:"treasure_count"`
	TotalGil      int32             `json:"total_gil"`
	Items         []string          `json:"items"`
	Equipment     []EquipmentSimple `json:"equipment"`
}

func getTreasuresLocSimple(id int32, subsection Subsection) *TreasuresLocSimple {
	treasureIDs := subsection.relations[id][RelationTreasures]

	if len(treasureIDs) == 0 {
		return nil
	}

	treasures := populateTreasuresLocSimple(treasureIDs)
	return &treasures
}

func populateTreasuresLocSimple(treasureIDs []int32) TreasuresLocSimple {
	treasures := TreasuresLocSimple{
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
	treasures.Items = convertObjSlice(itemAmounts, convertItemAmountSimple)
	return treasures
}

type EquipmentSimple struct {
	Name          string  `json:"name"`
	AutoAbilities *string `json:"auto_abilities"`
}

func convertEquipmentSimple(equipment seedtypes.TreasureEquipment) EquipmentSimple {
	return EquipmentSimple{
		Name:          equipment.Name,
		AutoAbilities: foundEquipmentAbilitiesStringPtr(equipment),
	}
}
