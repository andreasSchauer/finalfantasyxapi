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

	lootType, _ := newNamedAPIResourceFromType(cfg, cfg.e.lootType.endpoint, treasure.LootType, cfg.t.LootType)

	response := Treasure{
		ID:              treasure.ID,
		Area:            idToAreaAPIResource(cfg, cfg.e.areas, treasure.AreaID),
		IsPostAirship:   treasure.IsPostAirship,
		IsAnimaTreasure: treasure.IsAnimaTreasure,
		Notes:           treasure.Notes,
		TreasureType:    treasure.TreasureType,
		LootType:        lootType,
		GilAmount:       treasure.GilAmount,
		Items:           convertObjSlice(cfg, treasure.Items, convertItemAmount),
		Equipment:       convertObjPtr(cfg, treasure.Equipment, convertFoundEquipment),
	}

	return response, nil
}

func (cfg *Config) retrieveTreasures(r *http.Request, i handlerInput[seeding.Treasure, Treasure, UnnamedAPIResource, UnnamedApiResourceList]) (UnnamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return UnnamedApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[UnnamedAPIResource]{
		frl(idQuery(cfg, r, i, resources, "location", len(cfg.l.Locations), cfg.db.GetLocationTreasureIDs)),
		frl(idQuery(cfg, r, i, resources, "sublocation", len(cfg.l.Sublocations), cfg.db.GetSublocationTreasureIDs)),
		frl(idQuery(cfg, r, i, resources, "area", len(cfg.l.Areas), cfg.db.GetAreaTreasureIDs)),
		frl(boolQuery(cfg, r, i, resources, "anima", cfg.db.GetTreasureIDsByIsAnimaTreasure)),
		frl(boolQuery(cfg, r, i, resources, "airship", cfg.db.GetTreasureIDsByIsPostAirship)),
		frl(typeQuery(cfg, r, i, cfg.t.LootType, resources, "loot_type", cfg.db.GetTreasureIDsByLootType)),
		frl(typeQuery(cfg, r, i, cfg.t.TreasureType, resources, "treasure_type", cfg.db.GetTreasureIDsByTreasureType)),
	})
}
