package retrieve

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func FilterAvlMasterItems(r *http.Request, inputIDs []int32) ([]int32, error) {
	i := apilookups.E.AllItems

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

	methods, err := i.QueryLookup[ts.QpnMethods].ParseValueList(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := config.Cfg.DB.FilterMasterItemIDsByAvailability(r.Context(), database.FilterMasterItemIDsByAvailabilityParams{
		Ids:            inputIDs,
		Availability:   avlParams.Availabilities,
		IsRepeatable:   avlParams.IsRepeatable,
		PreAirship:     avlParams.PreAirship,
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

func FilterAvlItems(r *http.Request, inputIDs []int32) ([]int32, error) {
	i := apilookups.E.Items

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

	methods, err := i.QueryLookup[ts.QpnMethods].ParseValueList(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := config.Cfg.DB.FilterItemIDsByAvailability(r.Context(), database.FilterItemIDsByAvailabilityParams{
		Ids:            inputIDs,
		Availability:   avlParams.Availabilities,
		IsRepeatable:   avlParams.IsRepeatable,
		PreAirship:     avlParams.PreAirship,
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

func FilterAvlKeyItems(r *http.Request, inputIDs []int32) ([]int32, error) {
	i := apilookups.E.KeyItems

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

	methods, err := i.QueryLookup[ts.QpnMethods].ParseValueList(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := config.Cfg.DB.FilterKeyItemIDsByAvailability(r.Context(), database.FilterKeyItemIDsByAvailabilityParams{
		Ids:            inputIDs,
		Availability:   avlParams.Availabilities,
		PreAirship:     avlParams.PreAirship,
		LocContextID:   locContext.ID,
		LocContextType: locContext.Type,
		Methods:        h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, i.ErrorAvl(err)
	}

	return dbIDs, nil
}

func FilterAvlSpheres(r *http.Request, inputIDs []int32) ([]int32, error) {
	i := apilookups.E.Spheres

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

	methods, err := i.QueryLookup[ts.QpnMethods].ParseValueList(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := config.Cfg.DB.FilterSphereIDsByAvailability(r.Context(), database.FilterSphereIDsByAvailabilityParams{
		Ids:            inputIDs,
		Availability:   avlParams.Availabilities,
		IsRepeatable:   avlParams.IsRepeatable,
		PreAirship:     avlParams.PreAirship,
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

func FilterAvlPrimers(r *http.Request, inputIDs []int32) ([]int32, error) {
	i := apilookups.E.Primers

	avlParams, err := i.QueryAvlGetParams(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}
	if errs.QueryIsEmpty(err) {
		return inputIDs, nil
	}

	dbIDs, err := config.Cfg.DB.FilterPrimerIDsByAvailability(r.Context(), database.FilterPrimerIDsByAvailabilityParams{
		Ids:          inputIDs,
		Availability: avlParams.Availabilities,
		PreAirship:   avlParams.PreAirship,
	})
	if err != nil {
		return nil, i.ErrorAvl(err)
	}

	return dbIDs, nil
}
