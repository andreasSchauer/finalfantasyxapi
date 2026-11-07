package sectiontypes

import "github.com/andreasSchauer/finalfantasyxapi/internal/database"

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

type TreasuresLocSimple struct {
	TreasureCount int               `json:"treasure_count"`
	TotalGil      int32             `json:"total_gil"`
	Items         []string          `json:"items"`
	Equipment     []EquipmentSimple `json:"equipment"`
}

type EquipmentSimple struct {
	Name          string  `json:"name"`
	AutoAbilities *string `json:"auto_abilities"`
}
