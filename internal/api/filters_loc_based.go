package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func getAreasByItem(cfg *Config, r *http.Request, id int32) ([]AreaAPIResource, error) {
	i := cfg.e.areas
	queryParamMethods := i.queryLookup["methods"]

	methods, err := parseValueListQuery(cfg, r, queryParamMethods)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetAreaIDsWithItemFromMethods(r.Context(), database.GetAreaIDsWithItemFromMethodsParams{
		ItemID:     id,
		Methods: 	h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, queryParamMethods, err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)

	return resources, nil
}

func getSublocationsByItem(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	i := cfg.e.sublocations
	queryParamMethods := i.queryLookup["methods"]

	methods, err := parseValueListQuery(cfg, r, queryParamMethods)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetSublocationIDsWithItemFromMethods(r.Context(), database.GetSublocationIDsWithItemFromMethodsParams{
		ItemID:     id,
		Methods: 	h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, queryParamMethods, err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)

	return resources, nil
}

func getLocationsByItem(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	i := cfg.e.locations
	queryParamMethods := i.queryLookup["methods"]

	methods, err := parseValueListQuery(cfg, r, queryParamMethods)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetLocationIDsWithItemFromMethods(r.Context(), database.GetLocationIDsWithItemFromMethodsParams{
		ItemID:     id,
		Methods: 	h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, queryParamMethods, err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)

	return resources, nil
}
