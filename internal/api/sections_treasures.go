package api

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/sectiontypes"
)

func createTreasureSimple(_ *http.Request, id int32, _ Subsection) (apiresource.SimpleResource, error) {
	i := E.Treasures
	treasure, _ := lookup.GetResourceByID(id, i.objLookupID)

	treasureSimple := sectiontypes.TreasureSimple{
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
