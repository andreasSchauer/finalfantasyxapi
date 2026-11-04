package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func filterAvlAreas(server *Server, r *http.Request, inputIDs []int32) ([]int32, error) {
	i := E.Areas

	avlParams, err := i.queryAvlRepGetParams(server, r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return inputIDs, nil
	}

	sources, err := getLocBasedSources(server, r, i, ViewSourceTypeArea)
	if err != nil {
		return nil, err
	}

	if sources.IsZero() && h.NullBoolIsZero(avlParams.isRepeatable) {
		dbIDs, err := server.Cfg.DB.FilterAreaIDsByAvailabilitySoft(r.Context(), avlParams.availabilities)
		if err != nil {
			return nil, newHTTPErrorAvailability(i.resTypePlural, err)
		}

		return dbIDs, nil
	}

	dbIDs, err := server.Cfg.DB.FilterAreaIDsByAvailability(r.Context(), database.FilterAreaIDsByAvailabilityParams{
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

func filterAvlSublocations(server *Server, r *http.Request, inputIDs []int32) ([]int32, error) {
	i := E.Sublocations

	avlParams, err := i.queryAvlRepGetParams(server, r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return inputIDs, nil
	}

	sources, err := getLocBasedSources(server, r, i, ViewSourceTypeSublocation)
	if err != nil {
		return nil, err
	}

	if sources.IsZero() && h.NullBoolIsZero(avlParams.isRepeatable) {
		dbIDs, err := server.Cfg.DB.FilterSublocationIDsByAvailabilitySoft(r.Context(), avlParams.availabilities)
		if err != nil {
			return nil, newHTTPErrorAvailability(i.resTypePlural, err)
		}

		return dbIDs, nil
	}

	dbIDs, err := server.Cfg.DB.FilterSublocationIDsByAvailability(r.Context(), database.FilterSublocationIDsByAvailabilityParams{
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

func filterAvlLocations(server *Server, r *http.Request, inputIDs []int32) ([]int32, error) {
	i := E.Locations

	avlParams, err := i.queryAvlRepGetParams(server, r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return inputIDs, nil
	}

	sources, err := getLocBasedSources(server, r, i, ViewSourceTypeLocation)
	if err != nil {
		return nil, err
	}

	if sources.IsZero() && h.NullBoolIsZero(avlParams.isRepeatable) {
		dbIDs, err := server.Cfg.DB.FilterLocationIDsByAvailabilitySoft(r.Context(), avlParams.availabilities)
		if err != nil {
			return nil, newHTTPErrorAvailability(i.resTypePlural, err)
		}

		return dbIDs, nil
	}

	dbIDs, err := server.Cfg.DB.FilterLocationIDsByAvailability(r.Context(), database.FilterLocationIDsByAvailabilityParams{
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
