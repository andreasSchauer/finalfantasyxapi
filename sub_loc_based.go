package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type ShopLocSub struct {
	Category    database.ShopCategory `json:"category"`
	PreAirship  *ShopSummarySub       `json:"pre_airship"`
	PostAirship *ShopSummarySub       `json:"post_airship"`
}

type ShopSummarySub struct {
	HasItems     bool `json:"has_items"`
	HasEquipment bool `json:"has_equipment"`
}

func idToShopLocSub(cfg *Config, shopID int32) ShopLocSub {
	shopLookup, _ := seeding.GetResourceByID(shopID, cfg.l.ShopsID)
	return ShopLocSub{
		Category:    database.ShopCategory(shopLookup.Category),
		PreAirship:  convertObjPtr(cfg, shopLookup.PreAirship, convertShopSummarySub),
		PostAirship: convertObjPtr(cfg, shopLookup.PostAirship, convertShopSummarySub),
	}
}

func convertShopSummarySub(_ *Config, shop seeding.SubShop) ShopSummarySub {
	shopLoc := ShopSummarySub{}

	if len(shop.Items) != 0 {
		shopLoc.HasItems = true
	}

	if len(shop.Equipment) != 0 {
		shopLoc.HasEquipment = true
	}

	return shopLoc
}


type TreasuresLocSub struct {
	TreasureCount int             `json:"treasure_count"`
	TotalGil      int32           `json:"total_gil"`
	Items         []ItemAmountSub `json:"items"`
	Equipment     []EquipmentSub  `json:"equipment"`
}

func getTreasuresLocSub(cfg *Config, r *http.Request, resourceType string, id int32, dbQuery func(context.Context, int32) ([]int32, error)) (*TreasuresLocSub, error) {
	treasureIDs, err := dbQuery(r.Context(), id)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't retrieve treasures of %s with id '%d'", resourceType, id), err)
	}

	if len(treasureIDs) == 0 {
		return nil, nil
	}

	treasures := populateTreasuresLocSub(cfg, treasureIDs)
	return &treasures, nil
}

func populateTreasuresLocSub(cfg *Config, treasureIDs []int32) TreasuresLocSub {
	treasures := TreasuresLocSub{
		TreasureCount: len(treasureIDs),
		Items:         []ItemAmountSub{},
	}

	for _, treasureID := range treasureIDs {
		treasure, _ := seeding.GetResourceByID(treasureID, cfg.l.TreasuresID)

		switch treasure.LootType {
		case string(database.LootTypeGil):
			treasures.TotalGil += *treasure.GilAmount

		case string(database.LootTypeItem):
			for _, itemAmount := range treasure.Items {
				ia := convertSubItemAmount(cfg, itemAmount)
				treasures.Items = append(treasures.Items, ia)
			}

		case string(database.LootTypeEquipment):
			equipment := treasure.Equipment
			es := convertEquipmentSub(cfg, *equipment)
			treasures.Equipment = append(treasures.Equipment, es)
		}
	}

	treasures.Items = sortItemAmountSubsByID(cfg, treasures.Items)
	return treasures
}


type EquipmentSub struct {
	Name             string   `json:"name"`
	AutoAbilities    []string `json:"auto_abilities"`
	EmptySlotsAmount int32    `json:"empty_slots_amount"`
}

func convertEquipmentSub(cfg *Config, equipment seeding.FoundEquipment) EquipmentSub {
	return EquipmentSub{
		Name:             equipment.Name,
		AutoAbilities:    sortNamesByID(equipment.Abilities, cfg.l.AutoAbilities),
		EmptySlotsAmount: equipment.EmptySlotsAmount,
	}
}