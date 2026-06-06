package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getMasterItemObtainable(cfg *Config, r *http.Request, masterItem seeding.MasterItem) (ObtainableFrom, error) {
	availabilityParams, err := getRelAvailabilityParams(cfg, r, cfg.e.allItems, masterItem.ID)
	if err != nil {
		return ObtainableFrom{}, err
	}

	bools, err := cfg.db.GetMasterItemObtainableBools(r.Context(), database.GetMasterItemObtainableBoolsParams{
		MasterItemID: availabilityParams.ParentID,
		Availability: availabilityParams.Availability,
		Repeatable:   availabilityParams.Repeatable,
	})
	if err != nil {
		return ObtainableFrom{}, err
	}

	obtainable := ObtainableFrom{
		Monsters:  			bools.Monsters,
		Treasures: 			bools.Treasures,
		Shops:     			bools.Shops,
		Quests:    			bools.Quests,
		BlitzballPrizes: 	bools.Blitzball,
	}

	return obtainable, nil
}
