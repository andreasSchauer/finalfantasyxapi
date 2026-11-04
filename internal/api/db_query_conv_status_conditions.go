package api

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func convGetStatusConditionResistingMonsterIDs(srv *Server, resistance int32) DBQueryIntMany {
	return func(ctx context.Context, id int32) ([]int32, error) {
		return srv.Cfg.DB.GetStatusConditionResistingMonsterIDs(ctx, database.GetStatusConditionResistingMonsterIDsParams{
			StatusConditionID: id,
			MinResistance:     resistance,
		})
	}
}

func convGetStatusConditionAbilityIDsInflicted(srv *Server, minRate, maxRate int32) DBQueryIntMany {
	return func(ctx context.Context, id int32) ([]int32, error) {
		return srv.Cfg.DB.GetStatusConditionAbilityIDsInflicted(ctx, database.GetStatusConditionAbilityIDsInflictedParams{
			StatusConditionID: id,
			MinRate:           minRate,
			MaxRate:           maxRate,
		})
	}
}
