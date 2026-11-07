package api

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getTreasure(r *http.Request, i handlerinput.HandlerInput[seedtypes.Treasure, apitypes.Treasure, apiresource.UnnamedAPIResource], id int32) (apitypes.Treasure, error) {
	treasure, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Treasure{}, err
	}

	response := apitypes.Treasure{
		ID:              treasure.ID,
		Area:            handlerinput.IDToAreaAPIResource(E.Areas, treasure.AreaID),
		Availability:    treasure.Availability,
		IsAnimaTreasure: treasure.IsAnimaTreasure,
		Notes:           treasure.Notes,
		TreasureType:    treasure.TreasureType,
		LootType:        treasure.LootType,
		GilAmount:       treasure.GilAmount,
		Items:           E.AllItems.ToResAmtsFromNames(treasure.Items),
		Equipment:       ah.ConvertObjPtr(treasure.Equipment, convertFoundEquipment),
	}

	return response, nil
}

func retrieveTreasures(r *http.Request, i handlerinput.HandlerInput[seedtypes.Treasure, apitypes.Treasure, apiresource.UnnamedAPIResource]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryID(r, ids, ts.QpnLocation, lookup.L.Locations, config.Cfg.DB.GetLocationTreasureIDs),
		i.QueryID(r, ids, ts.QpnSublocation, lookup.L.Sublocations, config.Cfg.DB.GetSublocationTreasureIDs),
		i.QueryID(r, ids, ts.QpnArea, lookup.L.Areas, config.Cfg.DB.GetAreaTreasureIDs),
		i.QueryID(r, ids, ts.QpnItem, lookup.L.Items, config.Cfg.DB.GetTreasureIDsByItem),
		i.QueryJoined(r, ids, []ts.QueryParamName{ts.QpnAutoAbility, ts.QpnEmptySlots, ts.QpnCharacter}, filterTreasuresEquipment),
		i.QueryBool(r, ids, ts.QpnAnima, config.Cfg.DB.GetTreasureIDsByIsAnimaTreasure),
		i.QueryEnum(r, enums.T.LootType, ids, ts.QpnLootType, config.Cfg.DB.GetTreasureIDsByLootType),
		i.QueryEnum(r, enums.T.TreasureType, ids, ts.QpnTreasureType, config.Cfg.DB.GetTreasureIDsByTreasureType),
	})
}
