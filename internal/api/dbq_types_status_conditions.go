package api

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func convGetStatusConditionResistingMonsterIDs(cfg *Config, resistance int32) DbQueryIntMany {
	return func (ctx context.Context, id int32) ([]int32, error) {
		return cfg.db.GetStatusConditionResistingMonsterIDs(ctx, database.GetStatusConditionResistingMonsterIDsParams{
			StatusConditionID: 	id,
			MinResistance: 		resistance,
		})
	}
}

func convGetStatusConditionAbilityIDsInflicted(cfg *Config, minRate, maxRate int32) DbQueryIntMany {
	return func (ctx context.Context, id int32) ([]int32, error) {
		return cfg.db.GetStatusConditionAbilityIDsInflicted(ctx, database.GetStatusConditionAbilityIDsInflictedParams{
			StatusConditionID: 	id,
			MinRate: 			minRate,
			MaxRate: 			maxRate,
		})
	}
}