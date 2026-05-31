package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func getMasterItemsByArea(cfg *Config, r *http.Request, id int32) ([]TypedAPIResource, error) {
	i := cfg.e.allItems

	method, err := getQueryValuePtr(r, "method", i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetMasterItemIDsByArea(r.Context(), database.GetMasterItemIDsByAreaParams{
		AreaID: id,
		Method: h.GetNullString(method),
	})
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, i.queryLookup["method"], err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)

	return resources, nil
}

func getMasterItemsBySublocation(cfg *Config, r *http.Request, id int32) ([]TypedAPIResource, error) {
	i := cfg.e.allItems

	method, err := getQueryValuePtr(r, "method", i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetMasterItemIDsBySublocation(r.Context(), database.GetMasterItemIDsBySublocationParams{
		SublocationID: id,
		Method:        h.GetNullString(method),
	})
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, i.queryLookup["method"], err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)

	return resources, nil
}

func getMasterItemsByLocation(cfg *Config, r *http.Request, id int32) ([]TypedAPIResource, error) {
	i := cfg.e.allItems

	method, err := getQueryValuePtr(r, "method", i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetMasterItemIDsByLocation(r.Context(), database.GetMasterItemIDsByLocationParams{
		LocationID: id,
		Method:     h.GetNullString(method),
	})
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, i.queryLookup["method"], err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)

	return resources, nil
}

func getItemsByArea(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	i := cfg.e.items

	method, err := getQueryValuePtr(r, "method", i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetItemIDsByArea(r.Context(), database.GetItemIDsByAreaParams{
		AreaID: id,
		Method: h.GetNullString(method),
	})
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, i.queryLookup["method"], err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)

	return resources, nil
}

func getItemsBySublocation(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	i := cfg.e.items

	method, err := getQueryValuePtr(r, "method", i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetItemIDsBySublocation(r.Context(), database.GetItemIDsBySublocationParams{
		SublocationID: id,
		Method:        h.GetNullString(method),
	})
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, i.queryLookup["method"], err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)

	return resources, nil
}

func getItemsByLocation(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	i := cfg.e.items

	method, err := getQueryValuePtr(r, "method", i.queryLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetItemIDsByLocation(r.Context(), database.GetItemIDsByLocationParams{
		LocationID: id,
		Method:     h.GetNullString(method),
	})
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, i.queryLookup["method"], err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)

	return resources, nil
}
