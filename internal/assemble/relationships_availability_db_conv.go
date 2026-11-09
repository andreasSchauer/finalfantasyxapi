package assemble

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getItemSourceIDs(sourceType ts.ViewSourceType) handlerinput.DBQueryRelAvl {
	return func(ctx context.Context, p handlerinput.RelAvlParams) ([]int32, error) {
		return config.Cfg.DB.GetItemSourceIDs(ctx, database.GetItemSourceIDsParams{
			ItemID:       p.ParentID,
			SourceType:   string(sourceType),
			Availability: p.Availability,
			Repeatable:   p.Repeatable,
		})
	}
}

func getKeyItemSourceIDs(sourceType ts.ViewSourceType) handlerinput.DBQueryRelAvl {
	return func(ctx context.Context, p handlerinput.RelAvlParams) ([]int32, error) {
		return config.Cfg.DB.GetKeyItemSourceIDs(ctx, database.GetKeyItemSourceIDsParams{
			KeyItemID:    p.ParentID,
			SourceType:   string(sourceType),
			Availability: p.Availability,
		})
	}
}

func getKeyItemAreaIDs() handlerinput.DBQueryRelAvl {
	return func(ctx context.Context, p handlerinput.RelAvlParams) ([]int32, error) {
		return config.Cfg.DB.GetKeyItemAreaIDs(ctx, database.GetKeyItemAreaIDsParams{
			KeyItemID:    p.ParentID,
			Availability: p.Availability,
		})
	}
}

func getAutoAbilitySourceIDs(sourceType ts.ViewSourceType, shopType *database.ShopType) handlerinput.DBQueryRelAvl {
	return func(ctx context.Context, p handlerinput.RelAvlParams) ([]int32, error) {
		return config.Cfg.DB.GetAutoAbilitySourceIDs(ctx, database.GetAutoAbilitySourceIDsParams{
			AutoAbilityID: p.ParentID,
			SourceType:    string(sourceType),
			ShopType:      database.GetNullShopType(shopType),
			Availability:  p.Availability,
			Repeatable:    p.Repeatable,
		})
	}
}

func getAutoAbilityItemMonsterIDs() handlerinput.DBQueryRelAvl {
	return func(ctx context.Context, p handlerinput.RelAvlParams) ([]int32, error) {
		return config.Cfg.DB.GetAutoAbilityItemMonsterIDs(ctx, database.GetAutoAbilityItemMonsterIDsParams{
			AutoAbilityID: p.ParentID,
			Availability:  p.Availability,
			Repeatable:    p.Repeatable,
		})
	}
}

func getEquipmentSourceIDs(sourceType ts.ViewSourceType) handlerinput.DBQueryRelAvl {
	return func(ctx context.Context, p handlerinput.RelAvlParams) ([]int32, error) {
		return config.Cfg.DB.GetEquipmentSourceIDs(ctx, database.GetEquipmentSourceIDsParams{
			EquipmentNameID: p.ParentID,
			SourceType:      string(sourceType),
			Availability:    p.Availability,
		})
	}
}

func getAreaRelSourceIDs(sourceType ts.ViewSourceType) handlerinput.DBQueryRelAvl {
	return func(ctx context.Context, p handlerinput.RelAvlParams) ([]int32, error) {
		return config.Cfg.DB.GetAreaRelSourceIDs(ctx, database.GetAreaRelSourceIDsParams{
			AreaID:       p.ParentID,
			SourceType:   string(sourceType),
			Availability: p.Availability,
			Repeatable:   p.Repeatable,
		})
	}
}

func getSublocationRelSourceIDs(sourceType ts.ViewSourceType) handlerinput.DBQueryRelAvl {
	return func(ctx context.Context, p handlerinput.RelAvlParams) ([]int32, error) {
		return config.Cfg.DB.GetSublocationRelSourceIDs(ctx, database.GetSublocationRelSourceIDsParams{
			SublocationID: p.ParentID,
			SourceType:    string(sourceType),
			Availability:  p.Availability,
			Repeatable:    p.Repeatable,
		})
	}
}

func getLocationRelSourceIDs(sourceType ts.ViewSourceType) handlerinput.DBQueryRelAvl {
	return func(ctx context.Context, p handlerinput.RelAvlParams) ([]int32, error) {
		return config.Cfg.DB.GetLocationRelSourceIDs(ctx, database.GetLocationRelSourceIDsParams{
			LocationID:   p.ParentID,
			SourceType:   string(sourceType),
			Availability: p.Availability,
			Repeatable:   p.Repeatable,
		})
	}
}

func getMonsterAreaIDs() handlerinput.DBQueryRelAvl {
	return func(ctx context.Context, p handlerinput.RelAvlParams) ([]int32, error) {
		return config.Cfg.DB.GetMonsterAreaIDsRel(ctx, database.GetMonsterAreaIDsRelParams{
			MonsterID:    p.ParentID,
			Availability: p.Availability,
			Repeatable:   p.Repeatable,
		})
	}
}

func getMonsterMonsterFormationIDs() handlerinput.DBQueryRelAvl {
	return func(ctx context.Context, p handlerinput.RelAvlParams) ([]int32, error) {
		return config.Cfg.DB.GetMonsterMonsterFormationIDsRel(ctx, database.GetMonsterMonsterFormationIDsRelParams{
			MonsterID:    p.ParentID,
			Availability: p.Availability,
			Repeatable:   p.Repeatable,
		})
	}
}
