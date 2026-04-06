package api

import (
	"context"
	"database/sql"
	"errors"
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
}

func getAvailabilityParams[T seeding.LookupableID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], item T) (AvailabilityParams, error) {
	queryParamAvailability := i.queryLookup["rel_availability"]
	availabilitySlice, err := parseEnumListQuery(r, i.endpoint, queryParamAvailability, cfg.t.AvailabilityType)
	if err != nil && !errors.Is(err, errEmptyQuery) {
		return AvailabilityParams{}, err
	}

	repeatable, err := getQueryBoolPtr(r, "repeatable", cfg.e.monsters.queryLookup)
	if err != nil {
		return AvailabilityParams{}, err
	}

	availabilityParams := AvailabilityParams{
		ParentID:     item.GetID(),
		Availability: availabilitySlice,
		Repeatable:   h.GetNullBool(repeatable),
	}

	return availabilityParams, nil
}

func runAvailabilityQuery[T, K seeding.LookupableID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], item K, params AvailabilityParams, dbQuery AvailabilityDbQuery) ([]A, error) {
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
