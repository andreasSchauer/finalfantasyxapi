package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getTreasure(r *http.Request, i handlerInput[seeding.Treasure, Treasure, UnnamedAPIResource, UnnamedApiResourceList], id int32) (Treasure, error) {
	treasure, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return Treasure{}, err
	}

	response := Treasure{
		ID:              treasure.ID,
		Area:            idToAreaAPIResource(cfg, cfg.e.areas, treasure.AreaID),
		Availability:    enumToNamedAPIResource(cfg, cfg.e.availabilityType.endpoint, treasure.Availability, cfg.t.AvailabilityType),
		IsAnimaTreasure: treasure.IsAnimaTreasure,
		Notes:           treasure.Notes,
		TreasureType:    treasure.TreasureType,
		LootType:        enumToNamedAPIResource(cfg, cfg.e.lootType.endpoint, treasure.LootType, cfg.t.LootType),
		GilAmount:       treasure.GilAmount,
		Items:           nameAmtsToResAmts(cfg, cfg.e.allItems, treasure.Items),
		Equipment:       convertObjPtr(cfg, treasure.Equipment, convertFoundEquipment),
	}

	return response, nil
}

func (cfg *Config) retrieveTreasures(r *http.Request, i handlerInput[seeding.Treasure, Treasure, UnnamedAPIResource, UnnamedApiResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(cfg, r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []filteredIdList{
		fidl(idQuery(r, i, ids, qpnLocation, cfg.l.Locations, cfg.db.GetLocationTreasureIDs)),
		fidl(idQuery(r, i, ids, qpnSublocation, cfg.l.Sublocations, cfg.db.GetSublocationTreasureIDs)),
		fidl(idQuery(r, i, ids, qpnArea, cfg.l.Areas, cfg.db.GetAreaTreasureIDs)),
		fidl(idQuery(r, i, ids, qpnItem, cfg.l.Items, cfg.db.GetTreasureIDsByItem)),
		fidl(joinedQuery(cfg, r, i, ids, []QueryParamName{qpnAutoAbility, qpnEmptySlots, qpnCharacter}, filterTreasuresEquipment)),
		fidl(boolQuery(r, i, ids, qpnAnima, cfg.db.GetTreasureIDsByIsAnimaTreasure)),
		fidl(enumQuery(r, i, cfg.t.LootType, ids, qpnLootType, cfg.db.GetTreasureIDsByLootType)),
		fidl(enumQuery(r, i, cfg.t.TreasureType, ids, qpnTreasureType, cfg.db.GetTreasureIDsByTreasureType)),
	})
}
