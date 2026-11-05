package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getTreasure(r *http.Request, i handlerInput[seedtypes.Treasure, Treasure, UnnamedAPIResource], id int32) (Treasure, error) {
	treasure, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Treasure{}, err
	}

	response := Treasure{
		ID:              treasure.ID,
		Area:            idToAreaAPIResource(E.Areas, treasure.AreaID),
		Availability:    treasure.Availability,
		IsAnimaTreasure: treasure.IsAnimaTreasure,
		Notes:           treasure.Notes,
		TreasureType:    treasure.TreasureType,
		LootType:        treasure.LootType,
		GilAmount:       treasure.GilAmount,
		Items:           E.AllItems.toResAmtsFromNames(treasure.Items),
		Equipment:       convertObjPtr(treasure.Equipment, convertFoundEquipment),
	}

	return response, nil
}

func retrieveTreasures(r *http.Request, i handlerInput[seedtypes.Treasure, Treasure, UnnamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(r, ids, []IDFilter{
		i.queryID(r, ids, ts.QpnLocation, lookup.L.Locations, config.Cfg.DB.GetLocationTreasureIDs),
		i.queryID(r, ids, ts.QpnSublocation, lookup.L.Sublocations, config.Cfg.DB.GetSublocationTreasureIDs),
		i.queryID(r, ids, ts.QpnArea, lookup.L.Areas, config.Cfg.DB.GetAreaTreasureIDs),
		i.queryID(r, ids, ts.QpnItem, lookup.L.Items, config.Cfg.DB.GetTreasureIDsByItem),
		i.queryJoined(r, ids, []ts.QueryParamName{ts.QpnAutoAbility, ts.QpnEmptySlots, ts.QpnCharacter}, filterTreasuresEquipment),
		i.queryBool(r, ids, ts.QpnAnima, config.Cfg.DB.GetTreasureIDsByIsAnimaTreasure),
		i.queryEnum(r, enums.T.LootType, ids, ts.QpnLootType, config.Cfg.DB.GetTreasureIDsByLootType),
		i.queryEnum(r, enums.T.TreasureType, ids, ts.QpnTreasureType, config.Cfg.DB.GetTreasureIDsByTreasureType),
	})
}
