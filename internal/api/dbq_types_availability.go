package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type AvailabilityDbQuery func(ctx context.Context, p AvailabilityParams) ([]int32, error)
type AvailabilityDbQueryBool func(ctx context.Context, p AvailabilityBoolParams) ([]int32, error)

type AvailabilityBoolParams struct {
	Boolean      bool
	Availability []database.AvailabilityType
}


func getMasterItemIDsByLocation(cfg *Config, r *http.Request, id int32) ([]TypedAPIResource, error) {
	dbQuery := func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetMasterItemIDsByLocation(ctx, database.GetMasterItemIDsByLocationParams{
			LocationID:   p.ParentID,
			Availability: p.Availability,
			Method:       p.Method,
		})
	}
	return runAvailabilityQuery(cfg, r, cfg.e.allItems, id, cfg.e.locations.resourceType, dbQuery)
}

func getMasterItemIDsBySublocation(cfg *Config, r *http.Request, id int32) ([]TypedAPIResource, error) {
	dbQuery := func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetMasterItemIDsBySublocation(ctx, database.GetMasterItemIDsBySublocationParams{
			SublocationID: p.ParentID,
			Availability:  p.Availability,
			Method:        p.Method,
		})
	}
	return runAvailabilityQuery(cfg, r, cfg.e.allItems, id, cfg.e.sublocations.resourceType, dbQuery)
}

func getMasterItemIDsByArea(cfg *Config, r *http.Request, id int32) ([]TypedAPIResource, error) {
	dbQuery := func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetMasterItemIDsByArea(ctx, database.GetMasterItemIDsByAreaParams{
			AreaID:       p.ParentID,
			Availability: p.Availability,
			Method:       p.Method,
		})
	}
	return runAvailabilityQuery(cfg, r, cfg.e.allItems, id, cfg.e.areas.resourceType, dbQuery)
}



func getItemIDsByLocation(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	dbQuery := func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetItemIDsByLocation(ctx, database.GetItemIDsByLocationParams{
			LocationID:   p.ParentID,
			Availability: p.Availability,
			Method:       p.Method,
		})
	}
	return runAvailabilityQuery(cfg, r, cfg.e.items, id, cfg.e.locations.resourceType, dbQuery)
}

func getItemIDsBySublocation(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	dbQuery := func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetItemIDsBySublocation(ctx, database.GetItemIDsBySublocationParams{
			SublocationID: p.ParentID,
			Availability:  p.Availability,
			Method:        p.Method,
		})
	}
	return runAvailabilityQuery(cfg, r, cfg.e.items, id, cfg.e.sublocations.resourceType, dbQuery)
}

func getItemIDsByArea(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	dbQuery := func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetItemIDsByArea(ctx, database.GetItemIDsByAreaParams{
			AreaID:       p.ParentID,
			Availability: p.Availability,
			Method:       p.Method,
		})
	}
	return runAvailabilityQuery(cfg, r, cfg.e.items, id, cfg.e.areas.resourceType, dbQuery)
}

func runAvlBoolQuery[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput [T, R, A, L], boolean bool, dbQuery AvailabilityDbQueryBool) ([]A, error) {
	queryParam := i.queryLookup["availability"]

	availabilities, err := parseEnumListQuery(cfg, r, i.endpoint, queryParam, cfg.t.AvailabilityType)
	if errIsNotEmptyQuery(err) {
		return nil, err
	}

	params := AvailabilityBoolParams{
		Boolean: 		boolean,
		Availability: 	availabilities,
	}

	dbIDs, err := dbQuery(r.Context(), params)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss.", i.resourceType), err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)
	return resources, nil
}