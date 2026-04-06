package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getMasterItemRelationships(cfg *Config, r *http.Request, masterItem seeding.MasterItem) (MasterItem, error) {
	monsters, err := getResourcesDbItem(cfg, r, cfg.e.monsters, masterItem, cfg.db.GetMasterItemMonsterIDs)
	if err != nil {
		return MasterItem{}, err
	}
	
	treasures, err := getResourcesDbItem(cfg, r, cfg.e.treasures, masterItem, cfg.db.GetMasterItemTreasureIDs)
	if err != nil {
		return MasterItem{}, err
	}

	shops, err := getResourcesDbItem(cfg, r, cfg.e.shops, masterItem, cfg.db.GetMasterItemShopIDs)
	if err != nil {
		return MasterItem{}, err
	}

	quests, err := getResourcesDbItem(cfg, r, cfg.e.quests, masterItem, cfg.db.GetMasterItemQuestIDs)
	if err != nil {
		return MasterItem{}, err
	}

	rel := MasterItem{
		Monsters: 		monsters,
		Treasures:  	treasures,
		Shops: 			shops,
		Quests:  		quests,
	}

	return rel, nil
}
