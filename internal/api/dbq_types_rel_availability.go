package api

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type RelAvailabilityDbQuery func(ctx context.Context, p RelAvlParams) ([]int32, error)

func convGetItemMonsterIDs(cfg *Config) RelAvailabilityDbQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return cfg.db.GetItemMonsterIDs(ctx, database.GetItemMonsterIDsParams{
			ItemID:       p.ParentID,
			Availability: p.Availability,
			Repeatable:   p.Repeatable,
		})
	}
}

func convGetItemTreasureIDs(cfg *Config) RelAvailabilityDbQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return cfg.db.GetItemTreasureIDs(ctx, database.GetItemTreasureIDsParams{
			ItemID:       p.ParentID,
			Availability: p.Availability,
		})
	}
}

func convGetItemShopIDs(cfg *Config) RelAvailabilityDbQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return cfg.db.GetItemShopIDs(ctx, database.GetItemShopIDsParams{
			ItemID:       p.ParentID,
			Availability: p.Availability,
		})
	}
}

func convGetItemQuestIDs(cfg *Config) RelAvailabilityDbQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return cfg.db.GetItemQuestIDs(ctx, database.GetItemQuestIDsParams{
			ItemID:       p.ParentID,
			Availability: p.Availability,
			Repeatable:   p.Repeatable,
		})
	}
}

func convGetAutoAbilityItemMonsterIDs(cfg *Config) RelAvailabilityDbQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return cfg.db.GetAutoAbilityItemMonsterIDs(ctx, database.GetAutoAbilityItemMonsterIDsParams{
			AutoAbilityID: p.ParentID,
			Availability:  p.Availability,
			Repeatable:    p.Repeatable,
		})
	}
}

func convGetAutoAbilityMonsterIDs(cfg *Config) RelAvailabilityDbQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return cfg.db.GetAutoAbilityMonsterIDs(ctx, database.GetAutoAbilityMonsterIDsParams{
			AutoAbilityID: p.ParentID,
			Availability:  p.Availability,
			Repeatable:    p.Repeatable,
		})
	}
}

func convGetAutoAbilityTreasuresIDs(cfg *Config) RelAvailabilityDbQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return cfg.db.GetAutoAbilityTreasureIDs(ctx, database.GetAutoAbilityTreasureIDsParams{
			AutoAbilityID: p.ParentID,
			Availability:  p.Availability,
		})
	}
}

func convGetAutoAbilityShopIDsPre(cfg *Config) RelAvailabilityDbQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return cfg.db.GetAutoAbilityShopIDsPre(ctx, database.GetAutoAbilityShopIDsPreParams{
			AutoAbilityID: p.ParentID,
			Availability:  p.Availability,
		})
	}
}

func convGetAutoAbilityShopIDsPost(cfg *Config) RelAvailabilityDbQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return cfg.db.GetAutoAbilityShopIDsPost(ctx, database.GetAutoAbilityShopIDsPostParams{
			AutoAbilityID: p.ParentID,
			Availability:  p.Availability,
		})
	}
}

func convGetEquipmentTreasureIDs(cfg *Config) RelAvailabilityDbQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return cfg.db.GetEquipmentTreasureIDs(ctx, database.GetEquipmentTreasureIDsParams{
			EquipmentID:  p.ParentID,
			Availability: p.Availability,
		})
	}
}

func convGetEquipmentShopIDs(cfg *Config) RelAvailabilityDbQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return cfg.db.GetEquipmentShopIDs(ctx, database.GetEquipmentShopIDsParams{
			EquipmentID:  p.ParentID,
			Availability: p.Availability,
		})
	}
}
