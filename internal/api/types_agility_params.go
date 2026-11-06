package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func getAgilityTierDB(r *http.Request, stats []apitypes.BaseStat) (seedtypes.AgilityTier, error) {
	agilityStat := getBaseStat("agility", stats)
	agility := agilityStat.Value

	dbAgilityTier, err := config.Cfg.DB.GetAgilityTierByAgility(r.Context(), agility)
	if err != nil {
		return seedtypes.AgilityTier{}, errs.NewHTTPError(http.StatusInternalServerError, "couldn't extract agility parameters", err)
	}

	agilityTier, _ := lookup.GetResourceByID(dbAgilityTier.ID, lookup.L.AgilityTiersID)

	return agilityTier, nil
}
