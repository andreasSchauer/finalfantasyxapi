package api

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
)

type TreasureSimple struct {
	ID           int32            `json:"id"`
	URL          string           `json:"url"`
	Area         string           `json:"area"`
	Availability string           `json:"availability"`
	Notes        *string          `json:"notes,omitempty"`
	TreasureType string           `json:"treasure_type"`
	LootType     string           `json:"loot_type"`
	GilAmount    *int32           `json:"gil_amount,omitempty"`
	Items        []string         `json:"items,omitempty"`
	Equipment    *EquipmentSimple `json:"equipment,omitempty"`
}

func (t TreasureSimple) GetURL() string {
	return t.URL
}

func createTreasureSimple(_ *http.Request, id int32, _ Subsection) (apiresource.SimpleResource, error) {
	i := E.Treasures
	treasure, _ := lookup.GetResourceByID(id, i.objLookupID)

	treasureSimple := TreasureSimple{
		ID:           treasure.ID,
		URL:          ah.CreateResourceURL(i.Endpoint, id),
		Area:         idToLocAreaString(treasure.AreaID),
		Availability: treasure.Availability,
		Notes:        treasure.Notes,
		TreasureType: treasure.TreasureType,
		LootType:     treasure.LootType,
		GilAmount:    treasure.GilAmount,
		Items:        ah.ConvertObjSliceOrNil(treasure.Items, convertItemAmountSimple),
		Equipment:    ah.ConvertObjPtr(treasure.Equipment, convertEquipmentSimple),
	}

	return treasureSimple, nil
}
