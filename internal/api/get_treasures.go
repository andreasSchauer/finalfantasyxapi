package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getTreasure(r *http.Request, i handlerInput[seeding.Treasure, Treasure, UnnamedAPIResource, UnnamedApiResourceList], id int32) (Treasure, error) {
	treasure, err := verifyParamsAndGet(r, i, id)
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
		Items:           nameAmtsToResAmts(cfg, cfg.E.AllItems, treasure.Items),
		Equipment:       convertObjPtr(cfg, treasure.Equipment, convertFoundEquipment),
	}

	return response, nil
}

func (cfg *Config) retrieveTreasures(r *http.Request, i handlerInput[seeding.Treasure, Treasure, UnnamedAPIResource, UnnamedApiResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []IdFilter{
		idQuery(r, i, ids, QpnLocation, cfg.L.Locations, cfg.db.GetLocationTreasureIDs),
		idQuery(r, i, ids, QpnSublocation, cfg.L.Sublocations, cfg.db.GetSublocationTreasureIDs),
		idQuery(r, i, ids, QpnArea, cfg.L.Areas, cfg.db.GetAreaTreasureIDs),
		idQuery(r, i, ids, QpnItem, cfg.L.Items, cfg.db.GetTreasureIDsByItem),
		joinedQuery(cfg, r, i, ids, []QueryParamName{QpnAutoAbility, QpnEmptySlots, QpnCharacter}, filterTreasuresEquipment),
		boolQuery(r, i, ids, QpnAnima, cfg.db.GetTreasureIDsByIsAnimaTreasure),
		enumQuery(r, i, cfg.T.LootType, ids, QpnLootType, cfg.db.GetTreasureIDsByLootType),
		enumQuery(r, i, cfg.T.TreasureType, ids, QpnTreasureType, cfg.db.GetTreasureIDsByTreasureType),
	})
}
