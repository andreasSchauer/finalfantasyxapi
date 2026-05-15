package api

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type AvailabilityDbQuery func(ctx context.Context, p AvailabilityParams) ([]int32, error)



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
