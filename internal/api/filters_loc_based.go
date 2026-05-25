package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


func getAreasByItem(cfg *Config, r *http.Request, id int32) ([]AreaAPIResource, error) {
	i := cfg.e.areas
	queryParamMethod := i.queryLookup["method"]

	methods, err := parseValueListQuery(cfg, r, queryParamMethod)
	if errIsNotEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetAreaIDsWithItemFromMethod(r.Context(), database.GetAreaIDsWithItemFromMethodParams{
		ID: 	id,
		Method: methods,
	})
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, queryParamMethod, err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)

	return resources, nil
}

func getSublocationsByItem(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	i := cfg.e.sublocations
	queryParamMethod := i.queryLookup["method"]

	methods, err := parseValueListQuery(cfg, r, queryParamMethod)
	if errIsNotEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetSublocationIDsWithItemFromMethod(r.Context(), database.GetSublocationIDsWithItemFromMethodParams{
		ID: 	id,
		Method: methods,
	})
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, queryParamMethod, err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)

	return resources, nil
}

func getLocationsByItem(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	i := cfg.e.locations
	queryParamMethod := i.queryLookup["method"]

	methods, err := parseValueListQuery(cfg, r, queryParamMethod)
	if errIsNotEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetLocationIDsWithItemFromMethod(r.Context(), database.GetLocationIDsWithItemFromMethodParams{
		ID: 	id,
		Method: methods,
	})
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resourceType, queryParamMethod, err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)

	return resources, nil
}