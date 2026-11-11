package get

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apiconvert"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Treasure(r *http.Request, i handlerinput.Main[seedtypes.Treasure, apitypes.Treasure, apiresource.Unnamed], id int32) (apitypes.Treasure, error) {
	treasure, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Treasure{}, err
	}

	response := apitypes.Treasure{
		ID:              treasure.ID,
		Area:            handlerinput.IDToAreaAPIResource(apilookups.E.Areas, treasure.AreaID),
		Availability:    treasure.Availability,
		IsAnimaTreasure: treasure.IsAnimaTreasure,
		Notes:           treasure.Notes,
		TreasureType:    treasure.TreasureType,
		LootType:        treasure.LootType,
		GilAmount:       treasure.GilAmount,
		Items:           apilookups.E.AllItems.ToResAmtsFromNames(treasure.Items),
		Equipment:       ah.ConvertObjPtr(treasure.Equipment, apiconvert.FoundEquipment),
	}

	return response, nil
}
