package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


func filterAvlMasterItems(cfg *Config, r *http.Request, resources []TypedAPIResource) ([]TypedAPIResource, error) {
	i := cfg.e.allItems

	avlParams, err := checkAvlAndRep(cfg, r, i)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return resources, nil
	}

	locContext, err := getLocContextParams(cfg, r, i)
	if err != nil {
		return nil, err
	}

	method, err := getQueryValuePtr(r, "method", i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.FilterMasterItemIDsByAvailability(r.Context(), database.FilterMasterItemIDsByAvailabilityParams{
		Ids:          	resToIDs(resources),
		Availability:   avlParams.availabilities,
		IsRepeatable:   avlParams.isRepeatable,
		PreAirship: 	avlParams.preAirship,
		AvlType:      	locContext.AvlType,
		LocContextID: 	locContext.ID,
		LocContextType: locContext.Type,
		Method: 		h.GetNullString(method),
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}


func filterAvlItems(cfg *Config, r *http.Request, resources []NamedAPIResource) ([]NamedAPIResource, error) {
	i := cfg.e.items

	avlParams, err := checkAvlAndRep(cfg, r, i)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return resources, nil
	}

	locContext, err := getLocContextParams(cfg, r, i)
	if err != nil {
		return nil, err
	}

	method, err := getQueryValuePtr(r, "method", i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.FilterItemIDsByAvailability(r.Context(), database.FilterItemIDsByAvailabilityParams{
		Ids:          	resToIDs(resources),
		Availability:   avlParams.availabilities,
		IsRepeatable:   avlParams.isRepeatable,
		PreAirship: 	avlParams.preAirship,
		AvlType:      	locContext.AvlType,
		LocContextID: 	locContext.ID,
		LocContextType: locContext.Type,
		Method: 		h.GetNullString(method),
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}


func filterAvlKeyItems(cfg *Config, r *http.Request, resources []NamedAPIResource) ([]NamedAPIResource, error) {
	i := cfg.e.keyItems

	avlParams, err := checkAvl(cfg, r, i)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return resources, nil
	}

	locContext, err := getLocContextParams(cfg, r, i)
	if err != nil {
		return nil, err
	}

	method, err := getQueryValuePtr(r, "method", i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.FilterKeyItemIDsByAvailability(r.Context(), database.FilterKeyItemIDsByAvailabilityParams{
		Ids:          	resToIDs(resources),
		Availability: 	avlParams.availabilities,
		PreAirship:		avlParams.preAirship,
		LocContextID: 	locContext.ID,
		LocContextType: locContext.Type,
		Method: 		h.GetNullString(method),
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}


func filterAvlSpheres(cfg *Config, r *http.Request, resources []NamedAPIResource) ([]NamedAPIResource, error) {
	i := cfg.e.spheres

	avlParams, err := checkAvlAndRep(cfg, r, i)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return resources, nil
	}

	locContext, err := getLocContextParams(cfg, r, i)
	if err != nil {
		return nil, err
	}

	method, err := getQueryValuePtr(r, "method", i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.FilterSphereIDsByAvailability(r.Context(), database.FilterSphereIDsByAvailabilityParams{
		Ids:          	resToIDs(resources),
		Availability:   avlParams.availabilities,
		IsRepeatable:   avlParams.isRepeatable,
		PreAirship: 	avlParams.preAirship,
		AvlType:      	locContext.AvlType,
		LocContextID: 	locContext.ID,
		LocContextType: locContext.Type,
		Method: 		h.GetNullString(method),
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}


func filterAvlPrimers(cfg *Config, r *http.Request, resources []NamedAPIResource) ([]NamedAPIResource, error) {
	i := cfg.e.primers

	avlParams, err := checkAvl(cfg, r, i)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if queryIsEmpty(err) {
		return resources, nil
	}

	dbIDs, err := cfg.db.FilterPrimerIDsByAvailability(r.Context(), database.FilterPrimerIDsByAvailabilityParams{
		Ids:          resToIDs(resources),
		Availability: 	avlParams.availabilities,
		PreAirship:		avlParams.preAirship,
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by availability", i.resourceType), err)
	}

	resNew := idsToAPIResources(cfg, i, dbIDs)
	return resNew, nil
}