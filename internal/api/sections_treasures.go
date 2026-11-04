package api

import (
	"net/http"

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

func createTreasureSimple(srv *Server, _ *http.Request, id int32, _ Subsection) (SimpleResource, error) {
	i := E.Treasures
	treasure, _ := lookup.GetResourceByID(id, i.objLookupID)

	treasureSimple := TreasureSimple{
		ID:           treasure.ID,
		URL:          CreateResourceURL(srv, i.Endpoint, id),
		Area:         idToLocAreaString(srv, treasure.AreaID),
		Availability: treasure.Availability,
		Notes:        treasure.Notes,
		TreasureType: treasure.TreasureType,
		LootType:     treasure.LootType,
		GilAmount:    treasure.GilAmount,
		Items:        convertObjSliceOrNil(srv, treasure.Items, convertItemAmountSimple),
		Equipment:    convertObjPtr(srv, treasure.Equipment, convertEquipmentSimple),
	}

	return treasureSimple, nil
}
