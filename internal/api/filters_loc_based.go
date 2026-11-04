package api

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getAreasByItem(cfg *Config, r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := E.Areas
	queryParamMethods := i.queryLookup[ts.QpnMethods]

	methods, err := queryParamMethods.parseValueList(cfg, r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetAreaIDsWithItemFromMethods(ctx, database.GetAreaIDsWithItemFromMethodsParams{
		ItemID:  id,
		Methods: h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, newHTTPErrorDBFilter(i.resTypePlural, queryParamMethods, err)
	}

	return dbIDs, nil
}

func getSublocationsByItem(cfg *Config, r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := E.Sublocations
	queryParamMethods := i.queryLookup[ts.QpnMethods]

	methods, err := queryParamMethods.parseValueList(cfg, r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetSublocationIDsWithItemFromMethods(ctx, database.GetSublocationIDsWithItemFromMethodsParams{
		ItemID:  id,
		Methods: h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, newHTTPErrorDBFilter(i.resTypePlural, queryParamMethods, err)
	}

	return dbIDs, nil
}

func getLocationsByItem(cfg *Config, r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := E.Locations
	queryParamMethods := i.queryLookup[ts.QpnMethods]

	methods, err := queryParamMethods.parseValueList(cfg, r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetLocationIDsWithItemFromMethods(ctx, database.GetLocationIDsWithItemFromMethodsParams{
		ItemID:  id,
		Methods: h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, newHTTPErrorDBFilter(i.resTypePlural, queryParamMethods, err)
	}

	return dbIDs, nil
}
