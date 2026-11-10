package sections

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/sections/sectiontypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
)

func CreateTreasureSimple(_ *http.Request, id int32, _ handlerinput.Subsection) (apiresource.SimpleResource, error) {
	i := apilookups.E.Treasures
	treasure, _ := lookup.GetResourceByID(id, i.ObjLookupID)

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
