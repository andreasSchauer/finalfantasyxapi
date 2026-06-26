package api

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func getAreasByItem(cfg *Config, r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := cfg.e.areas
	queryParamMethods := i.queryLookup[qpnMethods]

	methods, err := parseValueListQuery(cfg, r, queryParamMethods)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetAreaIDsWithItemFromMethods(ctx, database.GetAreaIDsWithItemFromMethodsParams{
		ItemID:  id,
		Methods: h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resTypePlural, queryParamMethods, err)
	}

	return dbIDs, nil
}

func getSublocationsByItem(cfg *Config, r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := cfg.e.sublocations
	queryParamMethods := i.queryLookup[qpnMethods]

	methods, err := parseValueListQuery(cfg, r, queryParamMethods)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetSublocationIDsWithItemFromMethods(ctx, database.GetSublocationIDsWithItemFromMethodsParams{
		ItemID:  id,
		Methods: h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resTypePlural, queryParamMethods, err)
	}

	return dbIDs, nil
}

func getLocationsByItem(cfg *Config, r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := cfg.e.locations
	queryParamMethods := i.queryLookup[qpnMethods]

	methods, err := parseValueListQuery(cfg, r, queryParamMethods)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetLocationIDsWithItemFromMethods(ctx, database.GetLocationIDsWithItemFromMethodsParams{
		ItemID:  id,
		Methods: h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resTypePlural, queryParamMethods, err)
	}

	return dbIDs, nil
}
