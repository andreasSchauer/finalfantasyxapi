package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
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

func idToShopLocSimple(cfg *Config, shopID int32) ShopLocSimple {
	shopLookup, _ := seeding.GetResourceByID(shopID, cfg.l.ShopsID)
	return ShopLocSimple{
		ID:          shopID,
		Category:    database.ShopCategory(shopLookup.Category),
		PreAirship:  convertObjPtr(cfg, shopLookup.PreAirship, convertShopSummarySimple),
		PostAirship: convertObjPtr(cfg, shopLookup.PostAirship, convertShopSummarySimple),
	}
}

func convertShopSummarySimple(_ *Config, shop seeding.SubShop) ShopSummarySimple {
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
	TreasureCount int                `json:"treasure_count"`
	TotalGil      int32              `json:"total_gil"`
	Items         []string 			 `json:"items"`
	Equipment     []EquipmentSimple  `json:"equipment"`
}

func getTreasuresLocSimple(cfg *Config, id int32, subsection Subsection) *TreasuresLocSimple {
	treasureIDs := subsection.relations[id][RelationTreasures]

	if len(treasureIDs) == 0 {
		return nil
	}

	treasures := populateTreasuresLocSimple(cfg, treasureIDs)
	return &treasures
}

func populateTreasuresLocSimple(cfg *Config, treasureIDs []int32) TreasuresLocSimple {
	treasures := TreasuresLocSimple{
		TreasureCount: len(treasureIDs),
	}

	itemAmounts := []seeding.ItemAmount{}

	for _, treasureID := range treasureIDs {
		treasure, _ := seeding.GetResourceByID(treasureID, cfg.l.TreasuresID)

		switch treasure.LootType {
		case string(database.LootTypeGil):
			treasures.TotalGil += *treasure.GilAmount

		case string(database.LootTypeItem):
			for _, itemAmount := range treasure.Items {
				itemAmounts = append(itemAmounts, itemAmount)
			}

		case string(database.LootTypeEquipment):
			equipment := treasure.Equipment
			es := convertEquipmentSimple(cfg, *equipment)
			treasures.Equipment = append(treasures.Equipment, es)
		}
	}

	itemAmounts = sortItemAmountsByID(cfg, itemAmounts)
	treasures.Items = convertObjSlice(cfg, itemAmounts, convertItemAmountSimple)
	return treasures
}

type EquipmentSimple struct {
	Name          string  `json:"name"`
	AutoAbilities *string `json:"auto_abilities"`
}

func convertEquipmentSimple(_ *Config, equipment seeding.FoundEquipment) EquipmentSimple {
	return EquipmentSimple{
		Name:          equipment.Name,
		AutoAbilities: foundEquipmentAbilitiesStringPtr(equipment),
	}
}
