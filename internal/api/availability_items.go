package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func filterAvlMasterItems(cfg *Config, r *http.Request, inputIDs []int32) ([]int32, error) {
	i := cfg.e.allItems

	avlParams, err := checkAvlAndRep(cfg, r, i)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return inputIDs, nil
	}

	locContext, err := getLocContextParams(cfg, r, i)
	if err != nil {
		return nil, err
	}

	methods, err := parseValueListQuery(cfg, r, i.queryLookup[qpnMethods])
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.FilterMasterItemIDsByAvailability(r.Context(), database.FilterMasterItemIDsByAvailabilityParams{
		Ids:            inputIDs,
		Availability:   avlParams.availabilities,
		IsRepeatable:   avlParams.isRepeatable,
		PreAirship:     avlParams.preAirship,
		AvlType:        locContext.AvlType,
		LocContextID:   locContext.ID,
		LocContextType: locContext.Type,
		Methods:        h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, newHTTPErrorAvailability(i.resTypePlural, err)
	}

	return dbIDs, nil
}

func filterAvlItems(cfg *Config, r *http.Request, inputIDs []int32) ([]int32, error) {
	i := cfg.e.items

	avlParams, err := checkAvlAndRep(cfg, r, i)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return inputIDs, nil
	}

	locContext, err := getLocContextParams(cfg, r, i)
	if err != nil {
		return nil, err
	}

	methods, err := parseValueListQuery(cfg, r, i.queryLookup[qpnMethods])
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.FilterItemIDsByAvailability(r.Context(), database.FilterItemIDsByAvailabilityParams{
		Ids:            inputIDs,
		Availability:   avlParams.availabilities,
		IsRepeatable:   avlParams.isRepeatable,
		PreAirship:     avlParams.preAirship,
		AvlType:        locContext.AvlType,
		LocContextID:   locContext.ID,
		LocContextType: locContext.Type,
		Methods:        h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, newHTTPErrorAvailability(i.resTypePlural, err)
	}

	return dbIDs, nil
}

func filterAvlKeyItems(cfg *Config, r *http.Request, inputIDs []int32) ([]int32, error) {
	i := cfg.e.keyItems

	avlParams, err := checkAvl(cfg, r, i)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return inputIDs, nil
	}

	locContext, err := getLocContextParams(cfg, r, i)
	if err != nil {
		return nil, err
	}

	methods, err := parseValueListQuery(cfg, r, i.queryLookup[qpnMethods])
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.FilterKeyItemIDsByAvailability(r.Context(), database.FilterKeyItemIDsByAvailabilityParams{
		Ids:            inputIDs,
		Availability:   avlParams.availabilities,
		PreAirship:     avlParams.preAirship,
		LocContextID:   locContext.ID,
		LocContextType: locContext.Type,
		Methods:        h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, newHTTPErrorAvailability(i.resTypePlural, err)
	}

	return dbIDs, nil
}

func filterAvlSpheres(cfg *Config, r *http.Request, inputIDs []int32) ([]int32, error) {
	i := cfg.e.spheres

	avlParams, err := checkAvlAndRep(cfg, r, i)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return inputIDs, nil
	}

	locContext, err := getLocContextParams(cfg, r, i)
	if err != nil {
		return nil, err
	}

	methods, err := parseValueListQuery(cfg, r, i.queryLookup[qpnMethods])
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.FilterSphereIDsByAvailability(r.Context(), database.FilterSphereIDsByAvailabilityParams{
		Ids:            inputIDs,
		Availability:   avlParams.availabilities,
		IsRepeatable:   avlParams.isRepeatable,
		PreAirship:     avlParams.preAirship,
		AvlType:        locContext.AvlType,
		LocContextID:   locContext.ID,
		LocContextType: locContext.Type,
		Methods:        h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, newHTTPErrorAvailability(i.resTypePlural, err)
	}

	return dbIDs, nil
}

func filterAvlPrimers(cfg *Config, r *http.Request, inputIDs []int32) ([]int32, error) {
	i := cfg.e.primers

	avlParams, err := checkAvl(cfg, r, i)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return inputIDs, nil
	}

	dbIDs, err := cfg.db.FilterPrimerIDsByAvailability(r.Context(), database.FilterPrimerIDsByAvailabilityParams{
		Ids:          inputIDs,
		Availability: avlParams.availabilities,
		PreAirship:   avlParams.preAirship,
	})
	if err != nil {
		return nil, newHTTPErrorAvailability(i.resTypePlural, err)
	}

	return dbIDs, nil
}
