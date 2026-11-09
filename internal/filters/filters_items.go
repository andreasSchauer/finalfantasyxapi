package filters

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func GetMasterItemsByArea(r *http.Request, ctx context.Context, id int32) ([]int32, error) {
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

func GetMasterItemsBySublocation(r *http.Request, ctx context.Context, id int32) ([]int32, error) {
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

func GetMasterItemsByLocation(r *http.Request, ctx context.Context, id int32) ([]int32, error) {
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

func GetItemsByArea(r *http.Request, ctx context.Context, id int32) ([]int32, error) {
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

func GetItemsBySublocation(r *http.Request, ctx context.Context, id int32) ([]int32, error) {
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

func GetItemsByLocation(r *http.Request, ctx context.Context, id int32) ([]int32, error) {
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
