package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func filterAvlMasterItems(r *http.Request, inputIDs []int32) ([]int32, error) {
	i := E.AllItems

	avlParams, err := i.QueryAvlRepGetParams(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}
	if errs.QueryIsEmpty(err) {
		return inputIDs, nil
	}

	locContext, err := getLocContextParams(r, i)
	if err != nil {
		return nil, err
	}

	methods, err := i.queryLookup[ts.QpnMethods].ParseValueList(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := config.Cfg.DB.FilterMasterItemIDsByAvailability(r.Context(), database.FilterMasterItemIDsByAvailabilityParams{
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
		return nil, i.ErrorAvl(err)
	}

	return dbIDs, nil
}

func filterAvlItems(r *http.Request, inputIDs []int32) ([]int32, error) {
	i := E.Items

	avlParams, err := i.QueryAvlRepGetParams(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}
	if errs.QueryIsEmpty(err) {
		return inputIDs, nil
	}

	locContext, err := getLocContextParams(r, i)
	if err != nil {
		return nil, err
	}

	methods, err := i.queryLookup[ts.QpnMethods].ParseValueList(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := config.Cfg.DB.FilterItemIDsByAvailability(r.Context(), database.FilterItemIDsByAvailabilityParams{
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
		return nil, i.ErrorAvl(err)
	}

	return dbIDs, nil
}

func filterAvlKeyItems(r *http.Request, inputIDs []int32) ([]int32, error) {
	i := E.KeyItems

	avlParams, err := i.QueryAvlGetParams(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}
	if errs.QueryIsEmpty(err) {
		return inputIDs, nil
	}

	locContext, err := getLocContextParams(r, i)
	if err != nil {
		return nil, err
	}

	methods, err := i.queryLookup[ts.QpnMethods].ParseValueList(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := config.Cfg.DB.FilterKeyItemIDsByAvailability(r.Context(), database.FilterKeyItemIDsByAvailabilityParams{
		Ids:            inputIDs,
		Availability:   avlParams.availabilities,
		PreAirship:     avlParams.preAirship,
		LocContextID:   locContext.ID,
		LocContextType: locContext.Type,
		Methods:        h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, i.ErrorAvl(err)
	}

	return dbIDs, nil
}

func filterAvlSpheres(r *http.Request, inputIDs []int32) ([]int32, error) {
	i := E.Spheres

	avlParams, err := i.QueryAvlRepGetParams(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}
	if errs.QueryIsEmpty(err) {
		return inputIDs, nil
	}

	locContext, err := getLocContextParams(r, i)
	if err != nil {
		return nil, err
	}

	methods, err := i.queryLookup[ts.QpnMethods].ParseValueList(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := config.Cfg.DB.FilterSphereIDsByAvailability(r.Context(), database.FilterSphereIDsByAvailabilityParams{
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
		return nil, i.ErrorAvl(err)
	}

	return dbIDs, nil
}

func filterAvlPrimers(r *http.Request, inputIDs []int32) ([]int32, error) {
	i := E.Primers

	avlParams, err := i.QueryAvlGetParams(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}
	if errs.QueryIsEmpty(err) {
		return inputIDs, nil
	}

	dbIDs, err := config.Cfg.DB.FilterPrimerIDsByAvailability(r.Context(), database.FilterPrimerIDsByAvailabilityParams{
		Ids:          inputIDs,
		Availability: avlParams.availabilities,
		PreAirship:   avlParams.preAirship,
	})
	if err != nil {
		return nil, i.ErrorAvl(err)
	}

	return dbIDs, nil
}
