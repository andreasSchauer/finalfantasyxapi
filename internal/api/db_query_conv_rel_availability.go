package api

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type RelAvailabilityDBQuery func(ctx context.Context, p RelAvlParams) ([]int32, error)

func getItemSourceIDs(server *Server, sourceType ViewSourceType) RelAvailabilityDBQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return server.Cfg.DB.GetItemSourceIDs(ctx, database.GetItemSourceIDsParams{
			ItemID:       p.ParentID,
			SourceType:   string(sourceType),
			Availability: p.Availability,
			Repeatable:   p.Repeatable,
		})
	}
}

func getKeyItemSourceIDs(server *Server, sourceType ViewSourceType) RelAvailabilityDBQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return server.Cfg.DB.GetKeyItemSourceIDs(ctx, database.GetKeyItemSourceIDsParams{
			KeyItemID:    p.ParentID,
			SourceType:   string(sourceType),
			Availability: p.Availability,
		})
	}
}

func getKeyItemAreaIDs(server *Server) RelAvailabilityDBQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return server.Cfg.DB.GetKeyItemAreaIDs(ctx, database.GetKeyItemAreaIDsParams{
			KeyItemID:    p.ParentID,
			Availability: p.Availability,
		})
	}
}

func getAutoAbilitySourceIDs(server *Server, sourceType ViewSourceType, shopType *database.ShopType) RelAvailabilityDBQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return server.Cfg.DB.GetAutoAbilitySourceIDs(ctx, database.GetAutoAbilitySourceIDsParams{
			AutoAbilityID: p.ParentID,
			SourceType:    string(sourceType),
			ShopType:      database.GetNullShopType(shopType),
			Availability:  p.Availability,
			Repeatable:    p.Repeatable,
		})
	}
}

func getAutoAbilityItemMonsterIDs(server *Server) RelAvailabilityDBQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return server.Cfg.DB.GetAutoAbilityItemMonsterIDs(ctx, database.GetAutoAbilityItemMonsterIDsParams{
			AutoAbilityID: p.ParentID,
			Availability:  p.Availability,
			Repeatable:    p.Repeatable,
		})
	}
}

func getEquipmentSourceIDs(server *Server, sourceType ViewSourceType) RelAvailabilityDBQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return server.Cfg.DB.GetEquipmentSourceIDs(ctx, database.GetEquipmentSourceIDsParams{
			EquipmentNameID: p.ParentID,
			SourceType:      string(sourceType),
			Availability:    p.Availability,
		})
	}
}

func getAreaRelSourceIDs(server *Server, sourceType ViewSourceType) RelAvailabilityDBQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return server.Cfg.DB.GetAreaRelSourceIDs(ctx, database.GetAreaRelSourceIDsParams{
			AreaID:       p.ParentID,
			SourceType:   string(sourceType),
			Availability: p.Availability,
			Repeatable:   p.Repeatable,
		})
	}
}

func getSublocationRelSourceIDs(server *Server, sourceType ViewSourceType) RelAvailabilityDBQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return server.Cfg.DB.GetSublocationRelSourceIDs(ctx, database.GetSublocationRelSourceIDsParams{
			SublocationID: p.ParentID,
			SourceType:    string(sourceType),
			Availability:  p.Availability,
			Repeatable:    p.Repeatable,
		})
	}
}

func getLocationRelSourceIDs(server *Server, sourceType ViewSourceType) RelAvailabilityDBQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return server.Cfg.DB.GetLocationRelSourceIDs(ctx, database.GetLocationRelSourceIDsParams{
			LocationID:   p.ParentID,
			SourceType:   string(sourceType),
			Availability: p.Availability,
			Repeatable:   p.Repeatable,
		})
	}
}

func getMonsterAreaIDs(server *Server) RelAvailabilityDBQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return server.Cfg.DB.GetMonsterAreaIDsRel(ctx, database.GetMonsterAreaIDsRelParams{
			MonsterID:    p.ParentID,
			Availability: p.Availability,
			Repeatable:   p.Repeatable,
		})
	}
}

func getMonsterMonsterFormationIDs(server *Server) RelAvailabilityDBQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return server.Cfg.DB.GetMonsterMonsterFormationIDsRel(ctx, database.GetMonsterMonsterFormationIDsRelParams{
			MonsterID:    p.ParentID,
			Availability: p.Availability,
			Repeatable:   p.Repeatable,
		})
	}
}
