package api

import (
	"context"
	"database/sql"
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

func getAvailabilityParams[T seeding.LookupableID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], parentID int32) (AvailabilityParams, error) {
	availabilities, err := getAvailabilities(cfg, r, i)
	if err != nil {
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

func getAvailabilities[T seeding.LookupableID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L]) ([]database.AvailabilityType, error) {
	queryParam := i.queryLookup["rel_availability"]
	
	availabilitySlice, err := parseEnumListQuery(r, i.endpoint, queryParam, cfg.t.AvailabilityType)
	if errIsNotEmptyQuery(err) {
		return nil, err
	}

	valMap := make(map[database.AvailabilityType]bool)
	var availabilities []database.AvailabilityType

	for _, val := range availabilitySlice {
		if valMap[val] {
			continue
		}

		if val == database.AvailabilityTypePostGame {
			valMap[database.AvailabilityTypeAlways] = true
			valMap[database.AvailabilityTypePost] = true
			continue
		}

		if val == database.AvailabilityTypeStoryOnly {
			valMap[database.AvailabilityTypeStory] = true
			valMap[database.AvailabilityTypePostStory] = true
			continue
		}

		valMap[val] = true
	}

	for val := range valMap {
		availabilities = append(availabilities, val)
	}

	return availabilities, nil
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

func convGetAutoAbilityItemMonsterIDs(cfg *Config) AvailabilityDbQuery {
	return func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetAutoAbilityItemMonsterIDs(ctx, database.GetAutoAbilityItemMonsterIDsParams{
			AutoAbilityID:	p.ParentID,
			Availability: 	p.Availability,
			Repeatable: 	p.Repeatable,
		})
	}
}

func convGetAutoAbilityMonsterIDs(cfg *Config) AvailabilityDbQuery {
	return func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetAutoAbilityMonsterIDs(ctx, database.GetAutoAbilityMonsterIDsParams{
			AutoAbilityID:	p.ParentID,
			Availability: 	p.Availability,
			Repeatable: 	p.Repeatable,
		})
	}
}

func convGetAutoAbilityTreasuresIDs(cfg *Config) AvailabilityDbQuery {
	return func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetAutoAbilityTreasureIDs(ctx, database.GetAutoAbilityTreasureIDsParams{
			AutoAbilityID:	p.ParentID,
			Availability: 	p.Availability,
		})
	}
}

func convGetAutoAbilityShopIDsPre(cfg *Config) AvailabilityDbQuery {
	return func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetAutoAbilityShopIDsPre(ctx, database.GetAutoAbilityShopIDsPreParams{
			AutoAbilityID:	p.ParentID,
			Availability: 	p.Availability,
		})
	}
}

func convGetAutoAbilityShopIDsPost(cfg *Config) AvailabilityDbQuery {
	return func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetAutoAbilityShopIDsPost(ctx, database.GetAutoAbilityShopIDsPostParams{
			AutoAbilityID:	p.ParentID,
			Availability: 	p.Availability,
		})
	}
}


func convGetEquipmentTreasureIDs(cfg *Config) AvailabilityDbQuery {
	return func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetEquipmentTreasureIDs(ctx, database.GetEquipmentTreasureIDsParams{
			EquipmentID:	p.ParentID,
			Availability: 	p.Availability,
		})
	}
}

func convGetEquipmentShopIDs(cfg *Config) AvailabilityDbQuery {
	return func(ctx context.Context, p AvailabilityParams) ([]int32, error) {
		return cfg.db.GetEquipmentShopIDs(ctx, database.GetEquipmentShopIDsParams{
			EquipmentID:	p.ParentID,
			Availability: 	p.Availability,
		})
	}
}