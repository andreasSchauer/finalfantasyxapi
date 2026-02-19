package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type TreasureSimple struct {
	ID            int32              `json:"id"`
	URL           string             `json:"url"`
	Area          string             `json:"area"`
	IsPostAirship bool               `json:"is_post_airship"`
	Notes         *string            `json:"notes,omitempty"`
	TreasureType  string             `json:"treasure_type"`
	LootType      string             `json:"loot_type"`
	GilAmount     *int32             `json:"gil_amount,omitempty"`
	Items         []ItemAmountSimple `json:"items,omitempty"`
	Equipment     *EquipmentSimple   `json:"equipment,omitempty"`
}

func (t TreasureSimple) GetURL() string {
	return t.URL
}

func createTreasureSimple(cfg *Config, _ *http.Request, id int32) (SimpleResource, error) {
	i := cfg.e.treasures
	treasure, _ := seeding.GetResourceByID(id, i.objLookupID)

	treasureSimple := TreasureSimple{
		ID:            treasure.ID,
		URL:           createResourceURL(cfg, i.endpoint, id),
		Area:          idToLocAreaString(cfg, treasure.AreaID),
		IsPostAirship: treasure.IsPostAirship,
		Notes:         treasure.Notes,
		TreasureType:  treasure.TreasureType,
		LootType:      treasure.LootType,
		GilAmount:     treasure.GilAmount,
		Items:         convertObjSliceNullable(cfg, treasure.Items, convertItemAmountSimple),
		Equipment:     convertObjPtr(cfg, treasure.Equipment, convertEquipmentSimple),
	}

	return treasureSimple, nil
}
