package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (s *Server) getTreasure(r *http.Request, i handlerInput[seedtypes.Treasure, Treasure, UnnamedAPIResource], id int32) (Treasure, error) {
	treasure, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Treasure{}, err
	}

	response := Treasure{
		ID:              treasure.ID,
		Area:            idToAreaAPIResource(s, E.Areas, treasure.AreaID),
		Availability:    treasure.Availability,
		IsAnimaTreasure: treasure.IsAnimaTreasure,
		Notes:           treasure.Notes,
		TreasureType:    treasure.TreasureType,
		LootType:        treasure.LootType,
		GilAmount:       treasure.GilAmount,
		Items:           E.AllItems.toResAmtsFromNames(s, treasure.Items),
		Equipment:       convertObjPtr(s, treasure.Equipment, convertFoundEquipment),
	}

	return response, nil
}

func (s *Server) retrieveTreasures(r *http.Request, i handlerInput[seedtypes.Treasure, Treasure, UnnamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(s, r, ids, []IDFilter{
		i.queryID(r, ids, ts.QpnLocation, lookup.L.Locations, s.Cfg.DB.GetLocationTreasureIDs),
		i.queryID(r, ids, ts.QpnSublocation, lookup.L.Sublocations, s.Cfg.DB.GetSublocationTreasureIDs),
		i.queryID(r, ids, ts.QpnArea, lookup.L.Areas, s.Cfg.DB.GetAreaTreasureIDs),
		i.queryID(r, ids, ts.QpnItem, lookup.L.Items, s.Cfg.DB.GetTreasureIDsByItem),
		i.queryJoined(s, r, ids, []ts.QueryParamName{ts.QpnAutoAbility, ts.QpnEmptySlots, ts.QpnCharacter}, filterTreasuresEquipment),
		i.queryBool(r, ids, ts.QpnAnima, s.Cfg.DB.GetTreasureIDsByIsAnimaTreasure),
		i.queryEnum(r, enums.T.LootType, ids, ts.QpnLootType, s.Cfg.DB.GetTreasureIDsByLootType),
		i.queryEnum(r, enums.T.TreasureType, ids, ts.QpnTreasureType, s.Cfg.DB.GetTreasureIDsByTreasureType),
	})
}
