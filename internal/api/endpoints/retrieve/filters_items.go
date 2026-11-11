package retrieve

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
)

func getMasterItemsByArea(r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := apilookups.E.AllItems
	queryParamMethods := i.QueryLookup[ts.QpnMethods]

	methods, err := queryParamMethods.ParseValueList(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := config.Cfg.DB.GetMasterItemIDsByArea(ctx, database.GetMasterItemIDsByAreaParams{
		AreaID:  id,
		Methods: h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, i.ErrorDBFilter(queryParamMethods, err)
	}

	return dbIDs, nil
}

func getMasterItemsBySublocation(r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := apilookups.E.AllItems
	queryParamMethods := i.QueryLookup[ts.QpnMethods]

	methods, err := queryParamMethods.ParseValueList(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := config.Cfg.DB.GetMasterItemIDsBySublocation(ctx, database.GetMasterItemIDsBySublocationParams{
		SublocationID: id,
		Methods:       h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, i.ErrorDBFilter(queryParamMethods, err)
	}

	return dbIDs, nil
}

func getMasterItemsByLocation(r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := apilookups.E.AllItems
	queryParamMethods := i.QueryLookup[ts.QpnMethods]

	methods, err := queryParamMethods.ParseValueList(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := config.Cfg.DB.GetMasterItemIDsByLocation(ctx, database.GetMasterItemIDsByLocationParams{
		LocationID: id,
		Methods:    h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, i.ErrorDBFilter(queryParamMethods, err)
	}

	return dbIDs, nil
}

func getItemsByArea(r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := apilookups.E.Items
	queryParamMethods := i.QueryLookup[ts.QpnMethods]

	methods, err := queryParamMethods.ParseValueList(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := config.Cfg.DB.GetItemIDsByArea(ctx, database.GetItemIDsByAreaParams{
		AreaID:  id,
		Methods: h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, i.ErrorDBFilter(queryParamMethods, err)
	}

	return dbIDs, nil
}

func getItemsBySublocation(r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := apilookups.E.Items
	queryParamMethods := i.QueryLookup[ts.QpnMethods]

	methods, err := queryParamMethods.ParseValueList(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := config.Cfg.DB.GetItemIDsBySublocation(ctx, database.GetItemIDsBySublocationParams{
		SublocationID: id,
		Methods:       h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, i.ErrorDBFilter(queryParamMethods, err)
	}

	return dbIDs, nil
}

func getItemsByLocation(r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := apilookups.E.Items
	queryParamMethods := i.QueryLookup[ts.QpnMethods]

	methods, err := queryParamMethods.ParseValueList(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := config.Cfg.DB.GetItemIDsByLocation(ctx, database.GetItemIDsByLocationParams{
		LocationID: id,
		Methods:    h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, i.ErrorDBFilter(queryParamMethods, err)
	}

	return dbIDs, nil
}
