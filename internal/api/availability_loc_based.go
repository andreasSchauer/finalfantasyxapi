package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func filterAvlAreas(r *http.Request, inputIDs []int32) ([]int32, error) {
	i := apilookups.E.Areas

	avlParams, err := i.QueryAvlRepGetParams(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}
	if errs.QueryIsEmpty(err) {
		return inputIDs, nil
	}

	sources, err := getLocBasedSources(r, i, ViewSourceTypeArea)
	if err != nil {
		return nil, err
	}

	if sources.IsZero() && h.NullBoolIsZero(avlParams.IsRepeatable) {
		dbIDs, err := config.Cfg.DB.FilterAreaIDsByAvailabilitySoft(r.Context(), avlParams.Availabilities)
		if err != nil {
			return nil, i.ErrorAvl(err)
		}

		return dbIDs, nil
	}

	dbIDs, err := config.Cfg.DB.FilterAreaIDsByAvailability(r.Context(), database.FilterAreaIDsByAvailabilityParams{
		Ids:             inputIDs,
		Availability:    avlParams.Availabilities,
		IsRepeatable:    avlParams.IsRepeatable,
		PreAirship:      avlParams.PreAirship,
		RequiredSources: sources.RequiredSources,
		ExcludedSources: sources.ExcludedSources,
		MonsterID:       sources.MonsterID,
		ItemID:          sources.ItemID,
		KeyItemID:       sources.KeyItemID,
		AutoAbilityID:   sources.AutoAbilityID,
		Methods:         sources.Methods,
	})
	if err != nil {
		return nil, i.ErrorAvl(err)
	}

	return dbIDs, nil
}

func filterAvlSublocations(r *http.Request, inputIDs []int32) ([]int32, error) {
	i := apilookups.E.Sublocations

	avlParams, err := i.QueryAvlRepGetParams(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}
	if errs.QueryIsEmpty(err) {
		return inputIDs, nil
	}

	sources, err := getLocBasedSources(r, i, ViewSourceTypeSublocation)
	if err != nil {
		return nil, err
	}

	if sources.IsZero() && h.NullBoolIsZero(avlParams.IsRepeatable) {
		dbIDs, err := config.Cfg.DB.FilterSublocationIDsByAvailabilitySoft(r.Context(), avlParams.Availabilities)
		if err != nil {
			return nil, i.ErrorAvl(err)
		}

		return dbIDs, nil
	}

	dbIDs, err := config.Cfg.DB.FilterSublocationIDsByAvailability(r.Context(), database.FilterSublocationIDsByAvailabilityParams{
		Ids:             inputIDs,
		Availability:    avlParams.Availabilities,
		IsRepeatable:    avlParams.IsRepeatable,
		PreAirship:      avlParams.PreAirship,
		RequiredSources: sources.RequiredSources,
		ExcludedSources: sources.ExcludedSources,
		MonsterID:       sources.MonsterID,
		ItemID:          sources.ItemID,
		KeyItemID:       sources.KeyItemID,
		AutoAbilityID:   sources.AutoAbilityID,
		Methods:         sources.Methods,
	})
	if err != nil {
		return nil, i.ErrorAvl(err)
	}

	return dbIDs, nil
}

func filterAvlLocations(r *http.Request, inputIDs []int32) ([]int32, error) {
	i := apilookups.E.Locations

	avlParams, err := i.QueryAvlRepGetParams(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}
	if errs.QueryIsEmpty(err) {
		return inputIDs, nil
	}

	sources, err := getLocBasedSources(r, i, ViewSourceTypeLocation)
	if err != nil {
		return nil, err
	}

	if sources.IsZero() && h.NullBoolIsZero(avlParams.IsRepeatable) {
		dbIDs, err := config.Cfg.DB.FilterLocationIDsByAvailabilitySoft(r.Context(), avlParams.Availabilities)
		if err != nil {
			return nil, i.ErrorAvl(err)
		}

		return dbIDs, nil
	}

	dbIDs, err := config.Cfg.DB.FilterLocationIDsByAvailability(r.Context(), database.FilterLocationIDsByAvailabilityParams{
		Ids:             inputIDs,
		Availability:    avlParams.Availabilities,
		IsRepeatable:    avlParams.IsRepeatable,
		PreAirship:      avlParams.PreAirship,
		RequiredSources: sources.RequiredSources,
		ExcludedSources: sources.ExcludedSources,
		MonsterID:       sources.MonsterID,
		ItemID:          sources.ItemID,
		KeyItemID:       sources.KeyItemID,
		AutoAbilityID:   sources.AutoAbilityID,
		Methods:         sources.Methods,
	})
	if err != nil {
		return nil, i.ErrorAvl(err)
	}

	return dbIDs, nil
}
