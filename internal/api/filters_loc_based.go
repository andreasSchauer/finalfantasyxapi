package api

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getAreasByItem(srv *Server, r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := E.Areas
	queryParamMethods := i.queryLookup[ts.QpnMethods]

	methods, err := queryParamMethods.parseValueList(srv, r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := srv.Cfg.DB.GetAreaIDsWithItemFromMethods(ctx, database.GetAreaIDsWithItemFromMethodsParams{
		ItemID:  id,
		Methods: h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, newHTTPErrorDBFilter(i.resTypePlural, queryParamMethods, err)
	}

	return dbIDs, nil
}

func getSublocationsByItem(srv *Server, r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := E.Sublocations
	queryParamMethods := i.queryLookup[ts.QpnMethods]

	methods, err := queryParamMethods.parseValueList(srv, r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := srv.Cfg.DB.GetSublocationIDsWithItemFromMethods(ctx, database.GetSublocationIDsWithItemFromMethodsParams{
		ItemID:  id,
		Methods: h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, newHTTPErrorDBFilter(i.resTypePlural, queryParamMethods, err)
	}

	return dbIDs, nil
}

func getLocationsByItem(srv *Server, r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := E.Locations
	queryParamMethods := i.queryLookup[ts.QpnMethods]

	methods, err := queryParamMethods.parseValueList(srv, r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := srv.Cfg.DB.GetLocationIDsWithItemFromMethods(ctx, database.GetLocationIDsWithItemFromMethodsParams{
		ItemID:  id,
		Methods: h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, newHTTPErrorDBFilter(i.resTypePlural, queryParamMethods, err)
	}

	return dbIDs, nil
}
