package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getTreasure(r *http.Request, i handlerInput[seeding.Treasure, Treasure, UnnamedAPIResource], id int32) (Treasure, error) {
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

func (cfg *Config) retrieveTreasures(r *http.Request, i handlerInput[seeding.Treasure, Treasure, UnnamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(cfg, r, ids, []IDFilter{
		i.queryID(r, ids, QpnLocation, cfg.L.Locations, cfg.db.GetLocationTreasureIDs),
		i.queryID(r, ids, QpnSublocation, cfg.L.Sublocations, cfg.db.GetSublocationTreasureIDs),
		i.queryID(r, ids, QpnArea, cfg.L.Areas, cfg.db.GetAreaTreasureIDs),
		i.queryID(r, ids, QpnItem, cfg.L.Items, cfg.db.GetTreasureIDsByItem),
		i.queryJoined(cfg, r, ids, []QueryParamName{QpnAutoAbility, QpnEmptySlots, QpnCharacter}, filterTreasuresEquipment),
		i.queryBool(r, ids, QpnAnima, cfg.db.GetTreasureIDsByIsAnimaTreasure),
		i.queryEnum(r, cfg.T.LootType, ids, QpnLootType, cfg.db.GetTreasureIDsByLootType),
		i.queryEnum(r, cfg.T.TreasureType, ids, QpnTreasureType, cfg.db.GetTreasureIDsByTreasureType),
	})
}
