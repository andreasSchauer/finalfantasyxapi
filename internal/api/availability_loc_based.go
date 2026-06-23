package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func filterAvlAreas(cfg *Config, r *http.Request, inputIDs []int32) ([]int32, error) {
	i := cfg.e.areas

	avlParams, err := checkAvlAndRep(cfg, r, i)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return inputIDs, nil
	}

	sources, err := getLocBasedSources(cfg, r, i, ViewSourceTypeArea)
	if err != nil {
		return nil, err
	}

	if sources.IsZero() && h.NullBoolIsZero(avlParams.isRepeatable) {
		dbIDs, err := cfg.db.FilterAreaIDsByAvailabilitySoft(r.Context(), avlParams.availabilities)
		if err != nil {
			return nil, newHTTPErrorAvailability(i.resTypePlural, err)
		}

		return dbIDs, nil
	}

	dbIDs, err := cfg.db.FilterAreaIDsByAvailability(r.Context(), database.FilterAreaIDsByAvailabilityParams{
		Ids:             inputIDs,
		Availability:    avlParams.availabilities,
		IsRepeatable:    avlParams.isRepeatable,
		PreAirship:      avlParams.preAirship,
		RequiredSources: sources.RequiredSources,
		ExcludedSources: sources.ExcludedSources,
		MonsterID:       sources.MonsterID,
		ItemID:          sources.ItemID,
		KeyItemID:       sources.KeyItemID,
		AutoAbilityID:   sources.AutoAbilityID,
		Methods:         sources.Methods,
	})
	if err != nil {
		return nil, newHTTPErrorAvailability(i.resTypePlural, err)
	}

	return dbIDs, nil
}

func filterAvlSublocations(cfg *Config, r *http.Request, inputIDs []int32) ([]int32, error) {
	i := cfg.e.sublocations

	avlParams, err := checkAvlAndRep(cfg, r, i)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return inputIDs, nil
	}

	sources, err := getLocBasedSources(cfg, r, i, ViewSourceTypeSublocation)
	if err != nil {
		return nil, err
	}

	if sources.IsZero() && h.NullBoolIsZero(avlParams.isRepeatable) {
		dbIDs, err := cfg.db.FilterSublocationIDsByAvailabilitySoft(r.Context(), avlParams.availabilities)
		if err != nil {
			return nil, newHTTPErrorAvailability(i.resTypePlural, err)
		}

		return dbIDs, nil
	}

	dbIDs, err := cfg.db.FilterSublocationIDsByAvailability(r.Context(), database.FilterSublocationIDsByAvailabilityParams{
		Ids:             inputIDs,
		Availability:    avlParams.availabilities,
		IsRepeatable:    avlParams.isRepeatable,
		PreAirship:      avlParams.preAirship,
		RequiredSources: sources.RequiredSources,
		ExcludedSources: sources.ExcludedSources,
		MonsterID:       sources.MonsterID,
		ItemID:          sources.ItemID,
		KeyItemID:       sources.KeyItemID,
		AutoAbilityID:   sources.AutoAbilityID,
		Methods:         sources.Methods,
	})
	if err != nil {
		return nil, newHTTPErrorAvailability(i.resTypePlural, err)
	}

	return dbIDs, nil
}

func filterAvlLocations(cfg *Config, r *http.Request, inputIDs []int32) ([]int32, error) {
	i := cfg.e.locations

	avlParams, err := checkAvlAndRep(cfg, r, i)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return inputIDs, nil
	}

	sources, err := getLocBasedSources(cfg, r, i, ViewSourceTypeLocation)
	if err != nil {
		return nil, err
	}

	if sources.IsZero() && h.NullBoolIsZero(avlParams.isRepeatable) {
		dbIDs, err := cfg.db.FilterLocationIDsByAvailabilitySoft(r.Context(), avlParams.availabilities)
		if err != nil {
			return nil, newHTTPErrorAvailability(i.resTypePlural, err)
		}

		return dbIDs, nil
	}

	dbIDs, err := cfg.db.FilterLocationIDsByAvailability(r.Context(), database.FilterLocationIDsByAvailabilityParams{
		Ids:             inputIDs,
		Availability:    avlParams.availabilities,
		IsRepeatable:    avlParams.isRepeatable,
		PreAirship:      avlParams.preAirship,
		RequiredSources: sources.RequiredSources,
		ExcludedSources: sources.ExcludedSources,
		MonsterID:       sources.MonsterID,
		ItemID:          sources.ItemID,
		KeyItemID:       sources.KeyItemID,
		AutoAbilityID:   sources.AutoAbilityID,
		Methods:         sources.Methods,
	})
	if err != nil {
		return nil, newHTTPErrorAvailability(i.resTypePlural, err)
	}

	return dbIDs, nil
}
