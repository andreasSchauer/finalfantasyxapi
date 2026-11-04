package api

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getMasterItemsByArea(server *Server, r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := E.AllItems
	queryParamMethods := i.queryLookup[ts.QpnMethods]

	methods, err := queryParamMethods.parseValueList(server, r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := server.Cfg.DB.GetMasterItemIDsByArea(ctx, database.GetMasterItemIDsByAreaParams{
		AreaID:  id,
		Methods: h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, newHTTPErrorDBFilter(i.resTypePlural, queryParamMethods, err)
	}

	return dbIDs, nil
}

func getMasterItemsBySublocation(server *Server, r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := E.AllItems
	queryParamMethods := i.queryLookup[ts.QpnMethods]

	methods, err := queryParamMethods.parseValueList(server, r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := server.Cfg.DB.GetMasterItemIDsBySublocation(ctx, database.GetMasterItemIDsBySublocationParams{
		SublocationID: id,
		Methods:       h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, newHTTPErrorDBFilter(i.resTypePlural, queryParamMethods, err)
	}

	return dbIDs, nil
}

func getMasterItemsByLocation(server *Server, r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := E.AllItems
	queryParamMethods := i.queryLookup[ts.QpnMethods]

	methods, err := queryParamMethods.parseValueList(server, r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := server.Cfg.DB.GetMasterItemIDsByLocation(ctx, database.GetMasterItemIDsByLocationParams{
		LocationID: id,
		Methods:    h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, newHTTPErrorDBFilter(i.resTypePlural, queryParamMethods, err)
	}

	return dbIDs, nil
}

func getItemsByArea(server *Server, r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := E.Items
	queryParamMethods := i.queryLookup[ts.QpnMethods]

	methods, err := queryParamMethods.parseValueList(server, r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := server.Cfg.DB.GetItemIDsByArea(ctx, database.GetItemIDsByAreaParams{
		AreaID:  id,
		Methods: h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, newHTTPErrorDBFilter(i.resTypePlural, queryParamMethods, err)
	}

	return dbIDs, nil
}

func getItemsBySublocation(server *Server, r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := E.Items
	queryParamMethods := i.queryLookup[ts.QpnMethods]

	methods, err := queryParamMethods.parseValueList(server, r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := server.Cfg.DB.GetItemIDsBySublocation(ctx, database.GetItemIDsBySublocationParams{
		SublocationID: id,
		Methods:       h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, newHTTPErrorDBFilter(i.resTypePlural, queryParamMethods, err)
	}

	return dbIDs, nil
}

func getItemsByLocation(server *Server, r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := E.Items
	queryParamMethods := i.queryLookup[ts.QpnMethods]

	methods, err := queryParamMethods.parseValueList(server, r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := server.Cfg.DB.GetItemIDsByLocation(ctx, database.GetItemIDsByLocationParams{
		LocationID: id,
		Methods:    h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, newHTTPErrorDBFilter(i.resTypePlural, queryParamMethods, err)
	}

	return dbIDs, nil
}
