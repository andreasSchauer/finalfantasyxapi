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

func convGetStatusConditionPlayerAbilityIDsInflicted(cfg *Config, minRate, maxRate int32) DbQueryIntMany {
	return func (ctx context.Context, id int32) ([]int32, error) {
		return cfg.db.GetStatusConditionPlayerAbilityIDsInflicted(ctx, database.GetStatusConditionPlayerAbilityIDsInflictedParams{
			StatusConditionID: 	id,
			MinRate: 			minRate,
			MaxRate: 			maxRate,
		})
	}
}

func convGetStatusConditionOverdriveAbilityIDsInflicted(cfg *Config, minRate, maxRate int32) DbQueryIntMany {
	return func (ctx context.Context, id int32) ([]int32, error) {
		return cfg.db.GetStatusConditionOverdriveAbilityIDsInflicted(ctx, database.GetStatusConditionOverdriveAbilityIDsInflictedParams{
			StatusConditionID: 	id,
			MinRate: 			minRate,
			MaxRate: 			maxRate,
		})
	}
}

func convGetStatusConditionItemAbilityIDsInflicted(cfg *Config, minRate, maxRate int32) DbQueryIntMany {
	return func (ctx context.Context, id int32) ([]int32, error) {
		return cfg.db.GetStatusConditionItemAbilityIDsInflicted(ctx, database.GetStatusConditionItemAbilityIDsInflictedParams{
			StatusConditionID: 	id,
			MinRate: 			minRate,
			MaxRate: 			maxRate,
		})
	}
}

func convGetStatusConditionUnspecifiedAbilityIDsInflicted(cfg *Config, minRate, maxRate int32) DbQueryIntMany {
	return func (ctx context.Context, id int32) ([]int32, error) {
		return cfg.db.GetStatusConditionUnspecifiedAbilityIDsInflicted(ctx, database.GetStatusConditionUnspecifiedAbilityIDsInflictedParams{
			StatusConditionID: 	id,
			MinRate: 			minRate,
			MaxRate: 			maxRate,
		})
	}
}

func convGetStatusConditionEnemyAbilityIDsInflicted(cfg *Config, minRate, maxRate int32) DbQueryIntMany {
	return func (ctx context.Context, id int32) ([]int32, error) {
		return cfg.db.GetStatusConditionEnemyAbilityIDsInflicted(ctx, database.GetStatusConditionEnemyAbilityIDsInflictedParams{
			StatusConditionID: 	id,
			MinRate: 			minRate,
			MaxRate: 			maxRate,
		})
	}
}