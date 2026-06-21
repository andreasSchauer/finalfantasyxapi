package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func getMasterItemsByArea(cfg *Config, r *http.Request, id int32) ([]int32, error) {
	i := cfg.e.allItems
	queryParamMethods := i.queryLookup["methods"]

	methods, err := parseValueListQuery(cfg, r, queryParamMethods)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetMasterItemIDsByArea(r.Context(), database.GetMasterItemIDsByAreaParams{
		AreaID:  id,
		Methods: h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, queryParamMethods, err)
	}

	return dbIDs, nil
}

func getMasterItemsBySublocation(cfg *Config, r *http.Request, id int32) ([]int32, error) {
	i := cfg.e.allItems
	queryParamMethods := i.queryLookup["methods"]

	methods, err := parseValueListQuery(cfg, r, queryParamMethods)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetMasterItemIDsBySublocation(r.Context(), database.GetMasterItemIDsBySublocationParams{
		SublocationID:  id,
		Methods: 		h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, queryParamMethods, err)
	}

	return dbIDs, nil
}

func getMasterItemsByLocation(cfg *Config, r *http.Request, id int32) ([]int32, error) {
	i := cfg.e.allItems
	queryParamMethods := i.queryLookup["methods"]

	methods, err := parseValueListQuery(cfg, r, queryParamMethods)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetMasterItemIDsByLocation(r.Context(), database.GetMasterItemIDsByLocationParams{
		LocationID: id,
		Methods: 	h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, queryParamMethods, err)
	}

	return dbIDs, nil
}

func getItemsByArea(cfg *Config, r *http.Request, id int32) ([]int32, error) {
	i := cfg.e.items
	queryParamMethods := i.queryLookup["methods"]

	methods, err := parseValueListQuery(cfg, r, queryParamMethods)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetItemIDsByArea(r.Context(), database.GetItemIDsByAreaParams{
		AreaID:  id,
		Methods: h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, queryParamMethods, err)
	}

	return dbIDs, nil
}

func getItemsBySublocation(cfg *Config, r *http.Request, id int32) ([]int32, error) {
	i := cfg.e.items
	queryParamMethods := i.queryLookup["methods"]

	methods, err := parseValueListQuery(cfg, r, queryParamMethods)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetItemIDsBySublocation(r.Context(), database.GetItemIDsBySublocationParams{
		SublocationID: id,
		Methods: 	   h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, queryParamMethods, err)
	}

	return dbIDs, nil
}

func getItemsByLocation(cfg *Config, r *http.Request, id int32) ([]int32, error) {
	i := cfg.e.items
	queryParamMethods := i.queryLookup["methods"]

	methods, err := parseValueListQuery(cfg, r, queryParamMethods)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetItemIDsByLocation(r.Context(), database.GetItemIDsByLocationParams{
		LocationID: id,
		Methods: 	h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, queryParamMethods, err)
	}

	return dbIDs, nil
}
