package api

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type AvailabilityDbQuery func(ctx context.Context, p AvailabilityParams) ([]int32, error)

type AvailabilityParams struct {
	ParentID     int32
	Availability []database.AvailabilityType
	Repeatable   sql.NullBool
	Method		 sql.NullString
}

func getAvailabilityParams[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], parentID int32) (AvailabilityParams, error) {
	queryParam := i.queryLookup["rel_availability"]

	availabilities, err := parseEnumListQuery(cfg, r, i.endpoint, queryParam, cfg.t.AvailabilityType)
	if errIsNotEmptyQuery(err) {
		return AvailabilityParams{}, err
	}

	repeatable, err := getQueryBoolPtr(r, "repeatable", cfg.e.monsters.queryLookup)
	if err != nil {
		return AvailabilityParams{}, err
	}

	availabilityParams := AvailabilityParams{
		ParentID:     parentID,
		Availability: availabilities,
		Repeatable:   h.GetNullBool(repeatable),
	}

	return availabilityParams, nil
}

func runAvailabilityQuery[T, K seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], item K, params AvailabilityParams, dbQuery AvailabilityDbQuery) ([]A, error) {
	dbIDs, err := dbQuery(r.Context(), params)
	if err != nil {
		return nil, newHTTPErrorDB(i.resourceType, item, err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)
	return resources, nil
}

func convGetItemMonsterIDs(cfg *Config) AvailabilityDbQuery {
	return func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetItemMonsterIDs(ctx, database.GetItemMonsterIDsParams{
			ItemID:       p.ParentID,
			Availability: p.Availability,
			Repeatable:   p.Repeatable,
		})
	}
}

func convGetItemTreasureIDs(cfg *Config) AvailabilityDbQuery {
	return func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetItemTreasureIDs(ctx, database.GetItemTreasureIDsParams{
			ItemID:       p.ParentID,
			Availability: p.Availability,
		})
	}
}

func convGetItemShopIDs(cfg *Config) AvailabilityDbQuery {
	return func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetItemShopIDs(ctx, database.GetItemShopIDsParams{
			ItemID:       p.ParentID,
			Availability: p.Availability,
		})
	}
}

func convGetItemQuestIDs(cfg *Config) AvailabilityDbQuery {
	return func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetItemQuestIDs(ctx, database.GetItemQuestIDsParams{
			ItemID:       p.ParentID,
			Availability: p.Availability,
			Repeatable:   p.Repeatable,
		})
	}
}

func convGetAutoAbilityItemMonsterIDs(cfg *Config) AvailabilityDbQuery {
	return func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetAutoAbilityItemMonsterIDs(ctx, database.GetAutoAbilityItemMonsterIDsParams{
			AutoAbilityID: p.ParentID,
			Availability:  p.Availability,
			Repeatable:    p.Repeatable,
		})
	}
}

func convGetAutoAbilityMonsterIDs(cfg *Config) AvailabilityDbQuery {
	return func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetAutoAbilityMonsterIDs(ctx, database.GetAutoAbilityMonsterIDsParams{
			AutoAbilityID: p.ParentID,
			Availability:  p.Availability,
			Repeatable:    p.Repeatable,
		})
	}
}

func convGetAutoAbilityTreasuresIDs(cfg *Config) AvailabilityDbQuery {
	return func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetAutoAbilityTreasureIDs(ctx, database.GetAutoAbilityTreasureIDsParams{
			AutoAbilityID: p.ParentID,
			Availability:  p.Availability,
		})
	}
}

func convGetAutoAbilityShopIDsPre(cfg *Config) AvailabilityDbQuery {
	return func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetAutoAbilityShopIDsPre(ctx, database.GetAutoAbilityShopIDsPreParams{
			AutoAbilityID: p.ParentID,
			Availability:  p.Availability,
		})
	}
}

func convGetAutoAbilityShopIDsPost(cfg *Config) AvailabilityDbQuery {
	return func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetAutoAbilityShopIDsPost(ctx, database.GetAutoAbilityShopIDsPostParams{
			AutoAbilityID: p.ParentID,
			Availability:  p.Availability,
		})
	}
}

func convGetEquipmentTreasureIDs(cfg *Config) AvailabilityDbQuery {
	return func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetEquipmentTreasureIDs(ctx, database.GetEquipmentTreasureIDsParams{
			EquipmentID:  p.ParentID,
			Availability: p.Availability,
		})
	}
}

func convGetEquipmentShopIDs(cfg *Config) AvailabilityDbQuery {
	return func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetEquipmentShopIDs(ctx, database.GetEquipmentShopIDsParams{
			EquipmentID:  p.ParentID,
			Availability: p.Availability,
		})
	}
}





func getMasterItemIDsByLocation(cfg *Config, r *http.Request, id int32) ([]TypedAPIResource, error) {
	dbQuery := func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetMasterItemIDsByLocation(ctx, database.GetMasterItemIDsByLocationParams{
			LocationID:  	p.ParentID,
			Availability: 	p.Availability,
			Method: 		p.Method,
		})
	}
	return execAvailabilityQuery(cfg, r, cfg.e.allItems, id, cfg.e.locations.resourceType, dbQuery)
}

func getMasterItemIDsBySublocation(cfg *Config, r *http.Request, id int32) ([]TypedAPIResource, error) {
	dbQuery := func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetMasterItemIDsBySublocation(ctx, database.GetMasterItemIDsBySublocationParams{
			SublocationID:  p.ParentID,
			Availability: 	p.Availability,
			Method: 		p.Method,
		})
	}
	return execAvailabilityQuery(cfg, r, cfg.e.allItems, id, cfg.e.sublocations.resourceType, dbQuery)
}

func getMasterItemIDsByArea(cfg *Config, r *http.Request, id int32) ([]TypedAPIResource, error) {
	dbQuery := func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetMasterItemIDsByArea(ctx, database.GetMasterItemIDsByAreaParams{
			AreaID:  		p.ParentID,
			Availability: 	p.Availability,
			Method: 		p.Method,
		})
	}
	return execAvailabilityQuery(cfg, r, cfg.e.allItems, id, cfg.e.areas.resourceType, dbQuery)
}

func getItemIDsByLocation(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	dbQuery := func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetItemIDsByLocation(ctx, database.GetItemIDsByLocationParams{
			LocationID:  	p.ParentID,
			Availability: 	p.Availability,
			Method: 		p.Method,
		})
	}
	return execAvailabilityQuery(cfg, r, cfg.e.items, id, cfg.e.locations.resourceType, dbQuery)
}

func getItemIDsBySublocation(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	dbQuery := func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetItemIDsBySublocation(ctx, database.GetItemIDsBySublocationParams{
			SublocationID:  p.ParentID,
			Availability: 	p.Availability,
			Method: 		p.Method,
		})
	}
	return execAvailabilityQuery(cfg, r, cfg.e.items, id, cfg.e.sublocations.resourceType, dbQuery)
}

func getItemIDsByArea(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	dbQuery := func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetItemIDsByArea(ctx, database.GetItemIDsByAreaParams{
			AreaID:  		p.ParentID,
			Availability: 	p.Availability,
			Method: 		p.Method,
		})
	}
	return execAvailabilityQuery(cfg, r, cfg.e.items, id, cfg.e.areas.resourceType, dbQuery)
}


func execAvailabilityQuery[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], id int32, pResType string, dbQuery AvailabilityDbQuery) ([]A, error) {
	paramAvailability := i.queryLookup["availability"]

	availabilities, err := parseEnumListQuery(cfg, r, i.endpoint, paramAvailability, cfg.t.AvailabilityType)
	if errIsNotEmptyQuery(err) {
		return nil, err
	}

	methodPtr, err := getQueryValuePtr(r, "method", i.queryLookup)
	if err != nil {
		return nil, err
	}

	dbIDs, err := dbQuery(r.Context(), AvailabilityParams{
		ParentID: 	  id,
		Availability: availabilities,
		Method: 	  h.GetNullString(methodPtr),
	})
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't filter %ss by %s id '%d'.", i.resourceType, pResType, id), err)
	}

	resources := idsToAPIResources(cfg, i, dbIDs)
	return resources, nil
}