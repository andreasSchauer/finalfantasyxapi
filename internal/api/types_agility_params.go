package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)


type AgilityParams struct {
	TickSpeed int32  `json:"tick_speed"`
	MinICV    *int32 `json:"min_icv"`
	MaxICV    *int32 `json:"max_icv"`
}

func getAgilityTier(cfg *Config, r *http.Request, stats []BaseStat) (seeding.AgilityTier, error) {
	agilityStat := getBaseStat(cfg, "agility", stats)
	agility := agilityStat.Value

	dbAgilityTier, err := cfg.db.GetAgilityTierByAgility(r.Context(), agility)
	if err != nil {
		return seeding.AgilityTier{}, newHTTPError(http.StatusInternalServerError, "couldn't extract agility parameters.", err)
	}

	agilityTier, _ := seeding.GetResourceByID(dbAgilityTier.ID, cfg.l.AgilityTiersID)

	return agilityTier, nil
}