package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	th "github.com/andreasSchauer/finalfantasyxapi/internal/typehelpers"
)

func getUnitAgilityParams(r *http.Request, baseStats []apitypes.BaseStat) (apitypes.AgilityParams, error) {
	agilityTier, err := GetAgilityTierDB(r, baseStats)
	if err != nil {
		return apitypes.AgilityParams{}, err
	}
	agilityStat := th.GetBaseStat("agility", baseStats)
	agility := agilityStat.Value

	var minICV *int32

	for _, subTier := range agilityTier.CharacterMinICVs {
		if agility >= subTier.MinAgility && agility <= subTier.MaxAgility {
			minICV = subTier.CharacterMinICV
			break
		}
	}

	agilityParams := apitypes.AgilityParams{
		AgilityTier: handlerinput.IDToUnnamedAPIResource(apilookups.E.AgilityTiers, agilityTier.ID),
		TickSpeed:   agilityTier.TickSpeed,
		MinICV:      minICV,
		MaxICV:      agilityTier.CharacterMaxICV,
	}

	return agilityParams, nil
}

func GetAgilityTierDB(r *http.Request, stats []apitypes.BaseStat) (seedtypes.AgilityTier, error) {
	agilityStat := th.GetBaseStat("agility", stats)
	agility := agilityStat.Value

	dbAgilityTier, err := config.Cfg.DB.GetAgilityTierByAgility(r.Context(), agility)
	if err != nil {
		return seedtypes.AgilityTier{}, errs.NewHTTPError(http.StatusInternalServerError, "couldn't extract agility parameters", err)
	}

	agilityTier, _ := lookup.GetResourceByID(dbAgilityTier.ID, lookup.L.AgilityTiersID)

	return agilityTier, nil
}
