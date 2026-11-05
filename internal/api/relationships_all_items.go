package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func getMasterItemObtainable(r *http.Request, masterItem seedtypes.MasterItem) (ObtainableFrom, error) {
	availabilityParams, err := E.AllItems.queryAvlRelGetParams(r, masterItem.ID)
	if err != nil {
		return ObtainableFrom{}, err
	}

	bools, err := config.Cfg.DB.GetMasterItemObtainableBools(r.Context(), database.GetMasterItemObtainableBoolsParams{
		MasterItemID: availabilityParams.ParentID,
		Availability: availabilityParams.Availability,
		Repeatable:   availabilityParams.Repeatable,
	})
	if err != nil {
		return ObtainableFrom{}, err
	}

	obtainable := ObtainableFrom{
		Monsters:        bools.Monsters,
		Treasures:       bools.Treasures,
		Shops:           bools.Shops,
		Quests:          bools.Quests,
		BlitzballPrizes: bools.Blitzball,
	}

	return obtainable, nil
}
