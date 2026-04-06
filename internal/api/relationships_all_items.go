package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getMasterItemObtainable(cfg *Config, r *http.Request, masterItem seeding.MasterItem) (ObtainableFrom, error) {
	monsters, err := cfg.db.GetMasterItemMonstersBool(r.Context(), masterItem.ID)
	if err != nil {
		return ObtainableFrom{}, err
	}

	treasures, err := cfg.db.GetMasterItemTreasuresBool(r.Context(), masterItem.ID)
	if err != nil {
		return ObtainableFrom{}, err
	}

	shops, err := cfg.db.GetMasterItemShopsBool(r.Context(), masterItem.ID)
	if err != nil {
		return ObtainableFrom{}, err
	}

	quests, err := cfg.db.GetMasterItemQuestsBool(r.Context(), masterItem.ID)
	if err != nil {
		return ObtainableFrom{}, err
	}

	obtainable := ObtainableFrom{
		Monsters:  monsters,
		Treasures: treasures,
		Shops:     shops,
		Quests:    quests,
	}

	return obtainable, nil
}
