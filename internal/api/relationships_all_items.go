package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getMasterItemObtainable(cfg *Config, r *http.Request, masterItem seeding.MasterItem) (ObtainableFrom, error) {
	bools, err := cfg.db.GetMasterItemObtainableBools(r.Context(), masterItem.ID)
	if err != nil {
		return ObtainableFrom{}, err
	}

	obtainable := ObtainableFrom{
		Monsters:  bools.Monsters,
		Treasures: bools.Treasures,
		Shops:     bools.Shops,
		Quests:    bools.Quests,
	}

	return obtainable, nil
}
