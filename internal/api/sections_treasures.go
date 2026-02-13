package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type TreasureSub struct {
	ID            int32           `json:"id"`
	URL           string          `json:"url"`
	Area          string          `json:"area"`
	IsPostAirship bool            `json:"is_post_airship"`
	Notes         *string         `json:"notes,omitempty"`
	TreasureType  string          `json:"treasure_type"`
	LootType      string          `json:"loot_type"`
	GilAmount     *int32          `json:"gil_amount,omitempty"`
	Items         []ItemAmountSub `json:"items,omitempty"`
	Equipment     *EquipmentSub   `json:"equipment,omitempty"`
}

func (t TreasureSub) GetURL() string {
	return t.URL
}

func createTreasureSub(cfg *Config, _ *http.Request, id int32) (SubResource, error) {
	i := cfg.e.treasures
	treasure, _ := seeding.GetResourceByID(id, i.objLookupID)

	treasureSub := TreasureSub{
		ID:            treasure.ID,
		URL:           createResourceURL(cfg, i.endpoint, id),
		Area:          idToLocAreaString(cfg, treasure.AreaID),
		IsPostAirship: treasure.IsPostAirship,
		Notes:         treasure.Notes,
		TreasureType:  treasure.TreasureType,
		LootType:      treasure.LootType,
		GilAmount:     treasure.GilAmount,
		Items:         convertObjSliceNullable(cfg, treasure.Items, convertSubItemAmount),
		Equipment:     convertObjPtr(cfg, treasure.Equipment, convertEquipmentSub),
	}

	return treasureSub, nil
}
