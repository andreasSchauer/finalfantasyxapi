package main

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type TreasureSub struct {
	ID            int32           `json:"id"`
	URL           string          `json:"url"`
	Area          string          `json:"area"`
	TreasureType  string          `json:"treasure_type"`
	LootType      string          `json:"loot_type"`
	IsPostAirship bool            `json:"is_post_airship"`
	Notes         *string         `json:"notes,omitempty"`
	GilAmount     *int32          `json:"gil_amount,omitempty"`
	Items         []ItemAmountSub `json:"items,omitempty"`
	Equipment     *EquipmentSub   `json:"equipment,omitempty"`
}

func (t TreasureSub) GetSectionName() string {
	return "treasures"
}

func (t TreasureSub) GetURL() string {
	return t.URL
}

func handleTreasuresSection(cfg *Config, _ *http.Request, dbIDs []int32) ([]SubResource, error) {
	i := cfg.e.treasures
	treasures := []TreasureSub{}

	for _, treasureID := range dbIDs {
		treasure, _ := seeding.GetResourceByID(treasureID, i.objLookupID)

		treasureSub := TreasureSub{
			ID:            treasure.ID,
			URL:           createResourceURL(cfg, i.endpoint, treasureID),
			Area:          idToLocAreaString(cfg, treasure.AreaID),
			TreasureType:  treasure.TreasureType,
			LootType:      treasure.LootType,
			IsPostAirship: treasure.IsPostAirship,
			Notes:         treasure.Notes,
			GilAmount:     treasure.GilAmount,
			Items:         convertObjSliceNullable(cfg, treasure.Items, convertSubItemAmount),
			Equipment:     convertObjPtr(cfg, treasure.Equipment, convertEquipmentSub),
		}

		treasures = append(treasures, treasureSub)
	}

	return toSubResourceSlice(treasures), nil
}
