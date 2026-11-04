package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func filterAvlMasterItems(server *Server, r *http.Request, inputIDs []int32) ([]int32, error) {
	i := E.AllItems

	avlParams, err := i.queryAvlRepGetParams(server, r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return inputIDs, nil
	}

	locContext, err := getLocContextParams(r, i)
	if err != nil {
		return nil, err
	}

	methods, err := i.queryLookup[ts.QpnMethods].parseValueList(server, r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := server.Cfg.DB.FilterMasterItemIDsByAvailability(r.Context(), database.FilterMasterItemIDsByAvailabilityParams{
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

func filterAvlItems(server *Server, r *http.Request, inputIDs []int32) ([]int32, error) {
	i := E.Items

	avlParams, err := i.queryAvlRepGetParams(server, r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return inputIDs, nil
	}

	locContext, err := getLocContextParams(r, i)
	if err != nil {
		return nil, err
	}

	methods, err := i.queryLookup[ts.QpnMethods].parseValueList(server, r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := server.Cfg.DB.FilterItemIDsByAvailability(r.Context(), database.FilterItemIDsByAvailabilityParams{
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

func filterAvlKeyItems(server *Server, r *http.Request, inputIDs []int32) ([]int32, error) {
	i := E.KeyItems

	avlParams, err := i.queryAvlGetParams(server, r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return inputIDs, nil
	}

	locContext, err := getLocContextParams(r, i)
	if err != nil {
		return nil, err
	}

	methods, err := i.queryLookup[ts.QpnMethods].parseValueList(server, r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := server.Cfg.DB.FilterKeyItemIDsByAvailability(r.Context(), database.FilterKeyItemIDsByAvailabilityParams{
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

func filterAvlSpheres(server *Server, r *http.Request, inputIDs []int32) ([]int32, error) {
	i := E.Spheres

	avlParams, err := i.queryAvlRepGetParams(server, r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return inputIDs, nil
	}

	locContext, err := getLocContextParams(r, i)
	if err != nil {
		return nil, err
	}

	methods, err := i.queryLookup[ts.QpnMethods].parseValueList(server, r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := server.Cfg.DB.FilterSphereIDsByAvailability(r.Context(), database.FilterSphereIDsByAvailabilityParams{
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

func filterAvlPrimers(server *Server, r *http.Request, inputIDs []int32) ([]int32, error) {
	i := E.Primers

	avlParams, err := i.queryAvlGetParams(server, r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return inputIDs, nil
	}

	dbIDs, err := server.Cfg.DB.FilterPrimerIDsByAvailability(r.Context(), database.FilterPrimerIDsByAvailabilityParams{
		Ids:          inputIDs,
		Availability: avlParams.availabilities,
		PreAirship:   avlParams.preAirship,
	})
	if err != nil {
		return nil, newHTTPErrorAvailability(i.resTypePlural, err)
	}

	return dbIDs, nil
}
