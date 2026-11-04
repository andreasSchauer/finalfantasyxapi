package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (cfg *Config) getTreasure(r *http.Request, i handlerInput[seedtypes.Treasure, Treasure, UnnamedAPIResource], id int32) (Treasure, error) {
	treasure, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Treasure{}, err
	}

	response := Treasure{
		ID:              treasure.ID,
		Area:            idToAreaAPIResource(cfg, cfg.E.Areas, treasure.AreaID),
		Availability:    treasure.Availability,
		IsAnimaTreasure: treasure.IsAnimaTreasure,
		Notes:           treasure.Notes,
		TreasureType:    treasure.TreasureType,
		LootType:        treasure.LootType,
		GilAmount:       treasure.GilAmount,
		Items:           cfg.E.AllItems.toResAmtsFromNames(cfg, treasure.Items),
		Equipment:       convertObjPtr(cfg, treasure.Equipment, convertFoundEquipment),
	}

	return response, nil
}

func (cfg *Config) retrieveTreasures(r *http.Request, i handlerInput[seedtypes.Treasure, Treasure, UnnamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(cfg, r, ids, []IDFilter{
		i.queryID(r, ids, ts.QpnLocation, lookup.L.Locations, cfg.db.GetLocationTreasureIDs),
		i.queryID(r, ids, ts.QpnSublocation, lookup.L.Sublocations, cfg.db.GetSublocationTreasureIDs),
		i.queryID(r, ids, ts.QpnArea, lookup.L.Areas, cfg.db.GetAreaTreasureIDs),
		i.queryID(r, ids, ts.QpnItem, lookup.L.Items, cfg.db.GetTreasureIDsByItem),
		i.queryJoined(cfg, r, ids, []ts.QueryParamName{ts.QpnAutoAbility, ts.QpnEmptySlots, ts.QpnCharacter}, filterTreasuresEquipment),
		i.queryBool(r, ids, ts.QpnAnima, cfg.db.GetTreasureIDsByIsAnimaTreasure),
		i.queryEnum(r, enums.T.LootType, ids, ts.QpnLootType, cfg.db.GetTreasureIDsByLootType),
		i.queryEnum(r, enums.T.TreasureType, ids, ts.QpnTreasureType, cfg.db.GetTreasureIDsByTreasureType),
	})
}
