package get

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiconv"
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func GetTreasure(r *http.Request, i handlerinput.Main[seedtypes.Treasure, apitypes.Treasure, apiresource.Unnamed], id int32) (apitypes.Treasure, error) {
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
		Equipment:       ah.ConvertObjPtr(treasure.Equipment, apiconv.ConvertFoundEquipment),
	}

	return response, nil
}
