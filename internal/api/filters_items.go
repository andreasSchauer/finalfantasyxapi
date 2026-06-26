package api

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func getMasterItemsByArea(cfg *Config, r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := cfg.e.allItems
	queryParamMethods := i.queryLookup[qpnMethods]

	methods, err := parseValueListQuery(cfg, r, queryParamMethods)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetMasterItemIDsByArea(ctx, database.GetMasterItemIDsByAreaParams{
		AreaID:  id,
		Methods: h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resTypePlural, queryParamMethods, err)
	}

	return dbIDs, nil
}

func getMasterItemsBySublocation(cfg *Config, r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := cfg.e.allItems
	queryParamMethods := i.queryLookup[qpnMethods]

	methods, err := parseValueListQuery(cfg, r, queryParamMethods)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetMasterItemIDsBySublocation(ctx, database.GetMasterItemIDsBySublocationParams{
		SublocationID: id,
		Methods:       h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resTypePlural, queryParamMethods, err)
	}

	return dbIDs, nil
}

func getMasterItemsByLocation(cfg *Config, r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := cfg.e.allItems
	queryParamMethods := i.queryLookup[qpnMethods]

	methods, err := parseValueListQuery(cfg, r, queryParamMethods)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetMasterItemIDsByLocation(ctx, database.GetMasterItemIDsByLocationParams{
		LocationID: id,
		Methods:    h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resTypePlural, queryParamMethods, err)
	}

	return dbIDs, nil
}

func getItemsByArea(cfg *Config, r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := cfg.e.items
	queryParamMethods := i.queryLookup[qpnMethods]

	methods, err := parseValueListQuery(cfg, r, queryParamMethods)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetItemIDsByArea(ctx, database.GetItemIDsByAreaParams{
		AreaID:  id,
		Methods: h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resTypePlural, queryParamMethods, err)
	}

	return dbIDs, nil
}

func getItemsBySublocation(cfg *Config, r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := cfg.e.items
	queryParamMethods := i.queryLookup[qpnMethods]

	methods, err := parseValueListQuery(cfg, r, queryParamMethods)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetItemIDsBySublocation(ctx, database.GetItemIDsBySublocationParams{
		SublocationID: id,
		Methods:       h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resTypePlural, queryParamMethods, err)
	}

	return dbIDs, nil
}

func getItemsByLocation(cfg *Config, r *http.Request, ctx context.Context, id int32) ([]int32, error) {
	i := cfg.e.items
	queryParamMethods := i.queryLookup[qpnMethods]

	methods, err := parseValueListQuery(cfg, r, queryParamMethods)
	if errExceptEmptyQuery(err) {
		return nil, err
	}

	dbIDs, err := cfg.db.GetItemIDsByLocation(ctx, database.GetItemIDsByLocationParams{
		LocationID: id,
		Methods:    h.SliceOrNil(methods),
	})
	if err != nil {
		return nil, newHTTPErrorDbFilter(i.resTypePlural, queryParamMethods, err)
	}

	return dbIDs, nil
}
