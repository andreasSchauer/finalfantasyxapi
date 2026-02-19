package api

import (
	"context"
	"fmt"
	"net/http"

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
	Items         []ItemAmountSimple `json:"items"`
	Equipment     []EquipmentSimple  `json:"equipment"`
}

func getTreasuresLocSimple(cfg *Config, r *http.Request, resourceType string, id int32, dbQuery func(context.Context, int32) ([]int32, error)) (*TreasuresLocSimple, error) {
	treasureIDs, err := dbQuery(r.Context(), id)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't retrieve treasures of %s with id '%d'", resourceType, id), err)
	}

	if len(treasureIDs) == 0 {
		return nil, nil
	}

	treasures := populateTreasuresLocSimple(cfg, treasureIDs)
	return &treasures, nil
}

func populateTreasuresLocSimple(cfg *Config, treasureIDs []int32) TreasuresLocSimple {
	treasures := TreasuresLocSimple{
		TreasureCount: len(treasureIDs),
		Items:         []ItemAmountSimple{},
	}

	for _, treasureID := range treasureIDs {
		treasure, _ := seeding.GetResourceByID(treasureID, cfg.l.TreasuresID)

		switch treasure.LootType {
		case string(database.LootTypeGil):
			treasures.TotalGil += *treasure.GilAmount

		case string(database.LootTypeItem):
			for _, itemAmount := range treasure.Items {
				ia := convertItemAmountSimple(cfg, itemAmount)
				treasures.Items = append(treasures.Items, ia)
			}

		case string(database.LootTypeEquipment):
			equipment := treasure.Equipment
			es := convertEquipmentSimple(cfg, *equipment)
			treasures.Equipment = append(treasures.Equipment, es)
		}
	}

	treasures.Items = sortSimpleItemAmountsByID(cfg, treasures.Items)
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
