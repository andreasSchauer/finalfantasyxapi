package api

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type RelAvailabilityDbQuery func(ctx context.Context, p RelAvlParams) ([]int32, error)

func getItemSourceIDs(cfg *Config, sourceType ViewSourceType) RelAvailabilityDbQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return cfg.db.GetItemSourceIDs(ctx, database.GetItemSourceIDsParams{
			ItemID:       p.ParentID,
			SourceType:   string(sourceType),
			Availability: p.Availability,
			Repeatable:   p.Repeatable,
		})
	}
}

func getKeyItemSourceIDs(cfg *Config, sourceType ViewSourceType) RelAvailabilityDbQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return cfg.db.GetKeyItemSourceIDs(ctx, database.GetKeyItemSourceIDsParams{
			KeyItemID:    p.ParentID,
			SourceType:   string(sourceType),
			Availability: p.Availability,
		})
	}
}

func getKeyItemAreaIDs(cfg *Config) RelAvailabilityDbQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return cfg.db.GetKeyItemAreaIDs(ctx, database.GetKeyItemAreaIDsParams{
			KeyItemID:    p.ParentID,
			Availability: p.Availability,
		})
	}
}

func getAutoAbilitySourceIDs(cfg *Config, sourceType ViewSourceType, shopType *database.ShopType) RelAvailabilityDbQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return cfg.db.GetAutoAbilitySourceIDs(ctx, database.GetAutoAbilitySourceIDsParams{
			AutoAbilityID: p.ParentID,
			SourceType:    string(sourceType),
			ShopType:      database.GetNullShopType(shopType),
			Availability:  p.Availability,
			Repeatable:    p.Repeatable,
		})
	}
}

func getAutoAbilityItemMonsterIDs(cfg *Config) RelAvailabilityDbQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return cfg.db.GetAutoAbilityItemMonsterIDs(ctx, database.GetAutoAbilityItemMonsterIDsParams{
			AutoAbilityID: p.ParentID,
			Availability:  p.Availability,
			Repeatable:    p.Repeatable,
		})
	}
}

func getEquipmentSourceIDs(cfg *Config, sourceType ViewSourceType) RelAvailabilityDbQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return cfg.db.GetEquipmentSourceIDs(ctx, database.GetEquipmentSourceIDsParams{
			EquipmentNameID: p.ParentID,
			SourceType:      string(sourceType),
			Availability:    p.Availability,
		})
	}
}

func getAreaRelSourceIDs(cfg *Config, sourceType ViewSourceType) RelAvailabilityDbQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return cfg.db.GetAreaRelSourceIDs(ctx, database.GetAreaRelSourceIDsParams{
			AreaID:       p.ParentID,
			SourceType:   string(sourceType),
			Availability: p.Availability,
			Repeatable:   p.Repeatable,
		})
	}
}

func getSublocationRelSourceIDs(cfg *Config, sourceType ViewSourceType) RelAvailabilityDbQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return cfg.db.GetSublocationRelSourceIDs(ctx, database.GetSublocationRelSourceIDsParams{
			SublocationID: p.ParentID,
			SourceType:    string(sourceType),
			Availability:  p.Availability,
			Repeatable:   p.Repeatable,
		})
	}
}

func getLocationRelSourceIDs(cfg *Config, sourceType ViewSourceType) RelAvailabilityDbQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return cfg.db.GetLocationRelSourceIDs(ctx, database.GetLocationRelSourceIDsParams{
			LocationID:   p.ParentID,
			SourceType:   string(sourceType),
			Availability: p.Availability,
			Repeatable:   p.Repeatable,
		})
	}
}

func getMonsterAreaIDs(cfg *Config) RelAvailabilityDbQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return cfg.db.GetMonsterAreaIDsRel(ctx, database.GetMonsterAreaIDsRelParams{
			MonsterID:    p.ParentID,
			Availability: p.Availability,
			Repeatable:   p.Repeatable,
		})
	}
}

func getMonsterMonsterFormationIDs(cfg *Config) RelAvailabilityDbQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return cfg.db.GetMonsterMonsterFormationIDsRel(ctx, database.GetMonsterMonsterFormationIDsRelParams{
			MonsterID:    p.ParentID,
			Availability: p.Availability,
			Repeatable:   p.Repeatable,
		})
	}
}
