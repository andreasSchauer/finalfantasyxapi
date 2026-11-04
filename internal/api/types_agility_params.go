package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

type AgilityParams struct {
	AgilityTier UnnamedAPIResource `json:"agility_tier"`
	TickSpeed   int32              `json:"tick_speed"`
	MinICV      *int32             `json:"min_icv"`
	MaxICV      *int32             `json:"max_icv"`
}

func getAgilityTierDB(server *Server, r *http.Request, stats []BaseStat) (seedtypes.AgilityTier, error) {
	agilityStat := getBaseStat("agility", stats)
	agility := agilityStat.Value

	dbAgilityTier, err := server.Cfg.DB.GetAgilityTierByAgility(r.Context(), agility)
	if err != nil {
		return seedtypes.AgilityTier{}, newHTTPError(http.StatusInternalServerError, "couldn't extract agility parameters", err)
	}

	agilityTier, _ := lookup.GetResourceByID(dbAgilityTier.ID, lookup.L.AgilityTiersID)

	return agilityTier, nil
}
