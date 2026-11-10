package retrieve

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

func GetAreasByItem(r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := apilookups.E.Areas
	queryParamMethods := i.QueryLookup[ts.QpnMethods]

	methods, err := queryParamMethods.ParseValueList(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := config.Cfg.DB.GetAreaIDsWithItemFromMethods(ctx, database.GetAreaIDsWithItemFromMethodsParams{
		ItemID:  id,
		Methods: h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, i.ErrorDBFilter(queryParamMethods, err)
	}

	return dbIDs, nil
}

func GetSublocationsByItem(r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := apilookups.E.Sublocations
	queryParamMethods := i.QueryLookup[ts.QpnMethods]

	methods, err := queryParamMethods.ParseValueList(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := config.Cfg.DB.GetSublocationIDsWithItemFromMethods(ctx, database.GetSublocationIDsWithItemFromMethodsParams{
		ItemID:  id,
		Methods: h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, i.ErrorDBFilter(queryParamMethods, err)
	}

	return dbIDs, nil
}

func GetLocationsByItem(r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := apilookups.E.Locations
	queryParamMethods := i.QueryLookup[ts.QpnMethods]

	methods, err := queryParamMethods.ParseValueList(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := config.Cfg.DB.GetLocationIDsWithItemFromMethods(ctx, database.GetLocationIDsWithItemFromMethodsParams{
		ItemID:  id,
		Methods: h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, i.ErrorDBFilter(queryParamMethods, err)
	}

	return dbIDs, nil
}
