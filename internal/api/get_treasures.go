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
		Area:            idToAreaAPIResource(cfg, cfg.e.areas, treasure.AreaID),
		Availability:    treasure.Availability,
		IsAnimaTreasure: treasure.IsAnimaTreasure,
		Notes:           treasure.Notes,
		TreasureType:    treasure.TreasureType,
		LootType:        treasure.LootType,
		GilAmount:       treasure.GilAmount,
		Items:           nameAmtsToResAmts(cfg, cfg.e.allItems, treasure.Items),
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
		idQuery(r, i, ids, qpnLocation, cfg.l.Locations, cfg.db.GetLocationTreasureIDs),
		idQuery(r, i, ids, qpnSublocation, cfg.l.Sublocations, cfg.db.GetSublocationTreasureIDs),
		idQuery(r, i, ids, qpnArea, cfg.l.Areas, cfg.db.GetAreaTreasureIDs),
		idQuery(r, i, ids, qpnItem, cfg.l.Items, cfg.db.GetTreasureIDsByItem),
		joinedQuery(cfg, r, i, ids, []QueryParamName{qpnAutoAbility, qpnEmptySlots, qpnCharacter}, filterTreasuresEquipment),
		boolQuery(r, i, ids, qpnAnima, cfg.db.GetTreasureIDsByIsAnimaTreasure),
		enumQuery(r, i, cfg.t.LootType, ids, qpnLootType, cfg.db.GetTreasureIDsByLootType),
		enumQuery(r, i, cfg.t.TreasureType, ids, qpnTreasureType, cfg.db.GetTreasureIDsByTreasureType),
	})
}
