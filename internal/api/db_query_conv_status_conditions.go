package api

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
)

func convGetStatusConditionResistingMonsterIDs(resistance int32) handlerinput.DBQueryIntMany {
	return func(ctx context.Context, id int32) ([]int32, error) {
		return config.Cfg.DB.GetStatusConditionResistingMonsterIDs(ctx, database.GetStatusConditionResistingMonsterIDsParams{
			StatusConditionID: id,
			MinResistance:     resistance,
		})
	}
}

func convGetStatusConditionAbilityIDsInflicted(minRate, maxRate int32) handlerinput.DBQueryIntMany {
	return func(ctx context.Context, id int32) ([]int32, error) {
		return config.Cfg.DB.GetStatusConditionAbilityIDsInflicted(ctx, database.GetStatusConditionAbilityIDsInflictedParams{
			StatusConditionID: id,
			MinRate:           minRate,
			MaxRate:           maxRate,
		})
	}
}
