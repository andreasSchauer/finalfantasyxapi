package api

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type RelAvailabilityDBQuery func(ctx context.Context, p RelAvlParams) ([]int32, error)

func getItemSourceIDs(srv *Server, sourceType ViewSourceType) RelAvailabilityDBQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return srv.Cfg.DB.GetItemSourceIDs(ctx, database.GetItemSourceIDsParams{
			ItemID:       p.ParentID,
			SourceType:   string(sourceType),
			Availability: p.Availability,
			Repeatable:   p.Repeatable,
		})
	}
}

func getKeyItemSourceIDs(srv *Server, sourceType ViewSourceType) RelAvailabilityDBQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return srv.Cfg.DB.GetKeyItemSourceIDs(ctx, database.GetKeyItemSourceIDsParams{
			KeyItemID:    p.ParentID,
			SourceType:   string(sourceType),
			Availability: p.Availability,
		})
	}
}

func getKeyItemAreaIDs(srv *Server) RelAvailabilityDBQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return srv.Cfg.DB.GetKeyItemAreaIDs(ctx, database.GetKeyItemAreaIDsParams{
			KeyItemID:    p.ParentID,
			Availability: p.Availability,
		})
	}
}

func getAutoAbilitySourceIDs(srv *Server, sourceType ViewSourceType, shopType *database.ShopType) RelAvailabilityDBQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return srv.Cfg.DB.GetAutoAbilitySourceIDs(ctx, database.GetAutoAbilitySourceIDsParams{
			AutoAbilityID: p.ParentID,
			SourceType:    string(sourceType),
			ShopType:      database.GetNullShopType(shopType),
			Availability:  p.Availability,
			Repeatable:    p.Repeatable,
		})
	}
}

func getAutoAbilityItemMonsterIDs(srv *Server) RelAvailabilityDBQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return srv.Cfg.DB.GetAutoAbilityItemMonsterIDs(ctx, database.GetAutoAbilityItemMonsterIDsParams{
			AutoAbilityID: p.ParentID,
			Availability:  p.Availability,
			Repeatable:    p.Repeatable,
		})
	}
}

func getEquipmentSourceIDs(srv *Server, sourceType ViewSourceType) RelAvailabilityDBQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return srv.Cfg.DB.GetEquipmentSourceIDs(ctx, database.GetEquipmentSourceIDsParams{
			EquipmentNameID: p.ParentID,
			SourceType:      string(sourceType),
			Availability:    p.Availability,
		})
	}
}

func getAreaRelSourceIDs(srv *Server, sourceType ViewSourceType) RelAvailabilityDBQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return srv.Cfg.DB.GetAreaRelSourceIDs(ctx, database.GetAreaRelSourceIDsParams{
			AreaID:       p.ParentID,
			SourceType:   string(sourceType),
			Availability: p.Availability,
			Repeatable:   p.Repeatable,
		})
	}
}

func getSublocationRelSourceIDs(srv *Server, sourceType ViewSourceType) RelAvailabilityDBQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return srv.Cfg.DB.GetSublocationRelSourceIDs(ctx, database.GetSublocationRelSourceIDsParams{
			SublocationID: p.ParentID,
			SourceType:    string(sourceType),
			Availability:  p.Availability,
			Repeatable:    p.Repeatable,
		})
	}
}

func getLocationRelSourceIDs(srv *Server, sourceType ViewSourceType) RelAvailabilityDBQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return srv.Cfg.DB.GetLocationRelSourceIDs(ctx, database.GetLocationRelSourceIDsParams{
			LocationID:   p.ParentID,
			SourceType:   string(sourceType),
			Availability: p.Availability,
			Repeatable:   p.Repeatable,
		})
	}
}

func getMonsterAreaIDs(srv *Server) RelAvailabilityDBQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return srv.Cfg.DB.GetMonsterAreaIDsRel(ctx, database.GetMonsterAreaIDsRelParams{
			MonsterID:    p.ParentID,
			Availability: p.Availability,
			Repeatable:   p.Repeatable,
		})
	}
}

func getMonsterMonsterFormationIDs(srv *Server) RelAvailabilityDBQuery {
	return func(ctx context.Context, p RelAvlParams) ([]int32, error) {
		return srv.Cfg.DB.GetMonsterMonsterFormationIDsRel(ctx, database.GetMonsterMonsterFormationIDsRelParams{
			MonsterID:    p.ParentID,
			Availability: p.Availability,
			Repeatable:   p.Repeatable,
		})
	}
}
