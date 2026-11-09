package api

import (
	"context"
	"database/sql"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
)

func getTypedAbilityIDsByName(abilityType database.AbilityType) handlerinput.DBQueryStringMany {
	return func(ctx context.Context, name string) ([]int32, error) {
		return config.Cfg.DB.GetTypedAbilityIDsByName(ctx, database.GetTypedAbilityIDsByNameParams{
			Name: name,
			Type: abilityType,
		})
	}
}

func getTypedAbilityIDsByRank(abilityType database.AbilityType) handlerinput.DBQueryIntList {
	return func(ctx context.Context, ranks []int32) ([]int32, error) {
		return config.Cfg.DB.GetTypedAbilityIDsByRank(ctx, database.GetTypedAbilityIDsByRankParams{
			Rank: ranks,
			Type: abilityType,
		})
	}
}

func getTypedAbilityIDsByCanCopycat(abilityType database.AbilityType) handlerinput.DBQueryBool {
	return func(ctx context.Context, canCopycat bool) ([]int32, error) {
		return config.Cfg.DB.GetTypedAbilityIDsByCanCopycat(ctx, database.GetTypedAbilityIDsByCanCopycatParams{
			CanCopycat: canCopycat,
			Type:       abilityType,
		})
	}
}

func getTypedAbilityIDsByAppearsInHelpBar(abilityType database.AbilityType) handlerinput.DBQueryBool {
	return func(ctx context.Context, inHelpBar bool) ([]int32, error) {
		return config.Cfg.DB.GetTypedAbilityIDsByAppearsInHelpBar(ctx, database.GetTypedAbilityIDsByAppearsInHelpBarParams{
			AppearsInHelpBar: inHelpBar,
			Type:             abilityType,
		})
	}
}

func getTypedAbilityIDsBasedOnUserAttack(abilityType database.AbilityType) handlerinput.DBQueryNoInput {
	return func(ctx context.Context) ([]int32, error) {
		return config.Cfg.DB.GetTypedAbilityIDsBasedOnUserAttack(ctx, abilityType)
	}
}

func getTypedAbilityIDsByTargetType(abilityType database.AbilityType) handlerinput.DBQueryEnumList[database.TargetType] {
	return func(ctx context.Context, targetTypes []database.TargetType) ([]int32, error) {
		return config.Cfg.DB.GetTypedAbilityIDsByTargetType(ctx, database.GetTypedAbilityIDsByTargetTypeParams{
			TargetType: targetTypes,
			Type:       abilityType,
		})
	}
}

func getTypedAbilityIDsDarkable(abilityType database.AbilityType) handlerinput.DBQueryNoInput {
	return func(ctx context.Context) ([]int32, error) {
		return config.Cfg.DB.GetTypedAbilityIDsDarkable(ctx, abilityType)
	}
}

func getTypedAbilityIDsSilenceable(abilityType database.AbilityType) handlerinput.DBQueryNoInput {
	return func(ctx context.Context) ([]int32, error) {
		return config.Cfg.DB.GetTypedAbilityIDsSilenceable(ctx, abilityType)
	}
}

func getTypedAbilityIDsReflectable(abilityType database.AbilityType) handlerinput.DBQueryNoInput {
	return func(ctx context.Context) ([]int32, error) {
		return config.Cfg.DB.GetTypedAbilityIDsReflectable(ctx, abilityType)
	}
}

func getTypedAbilityIDsDealsDelay(abilityType database.AbilityType) handlerinput.DBQueryNoInput {
	return func(ctx context.Context) ([]int32, error) {
		return config.Cfg.DB.GetTypedAbilityIDsDealsDelay(ctx, abilityType)
	}
}

func getTypedAbilityIDsWithStatChanges(abilityType database.AbilityType) handlerinput.DBQueryNoInput {
	return func(ctx context.Context) ([]int32, error) {
		return config.Cfg.DB.GetTypedAbilityIDsWithStatChanges(ctx, abilityType)
	}
}

func getTypedAbilityIDsWithModifierChanges(abilityType database.AbilityType) handlerinput.DBQueryNoInput {
	return func(ctx context.Context) ([]int32, error) {
		return config.Cfg.DB.GetTypedAbilityIDsWithModifierChanges(ctx, abilityType)
	}
}

func getTypedAbilityIDsCanCrit(abilityType database.AbilityType) handlerinput.DBQueryNoInput {
	return func(ctx context.Context) ([]int32, error) {
		return config.Cfg.DB.GetTypedAbilityIDsCanCrit(ctx, abilityType)
	}
}

func getTypedAbilityIDsBreakDmgLimit(abilityType database.AbilityType) handlerinput.DBQueryNoInput {
	return func(ctx context.Context) ([]int32, error) {
		return config.Cfg.DB.GetTypedAbilityIDsBreakDmgLimit(ctx, abilityType)
	}
}

func getTypedAbilityIDsByDamageType(abilityType database.AbilityType) handlerinput.DBQueryEnumList[database.DamageType] {
	return func(ctx context.Context, damageTypes []database.DamageType) ([]int32, error) {
		return config.Cfg.DB.GetTypedAbilityIDsByDamageType(ctx, database.GetTypedAbilityIDsByDamageTypeParams{
			DamageType: damageTypes,
			Type:       abilityType,
		})
	}
}

func getTypedAbilityIDsByAttackType(abilityType database.AbilityType) handlerinput.DBQueryEnumList[database.AttackType] {
	return func(ctx context.Context, attackTypes []database.AttackType) ([]int32, error) {
		return config.Cfg.DB.GetTypedAbilityIDsByAttackType(ctx, database.GetTypedAbilityIDsByAttackTypeParams{
			AttackType: attackTypes,
			Type:       abilityType,
		})
	}
}

func getTypedAbilityIDsByDamageFormula(abilityType database.AbilityType) handlerinput.DBQueryEnum[database.DamageFormula] {
	return func(ctx context.Context, formula database.DamageFormula) ([]int32, error) {
		return config.Cfg.DB.GetTypedAbilityIDsByDamageFormula(ctx, database.GetTypedAbilityIDsByDamageFormulaParams{
			DamageFormula: formula,
			Type:          abilityType,
		})
	}
}

func getTypedAbilityIDsByInflictedStatus(abilityType database.AbilityType) handlerinput.DBQueryNullIntMany {
	return func(ctx context.Context, status sql.NullInt32) ([]int32, error) {
		return config.Cfg.DB.GetTypedAbilityIDsByInflictedStatus(ctx, database.GetTypedAbilityIDsByInflictedStatusParams{
			StatusID: status,
			Type:     abilityType,
		})
	}
}

func getTypedAbilityIDsByRemovedStatus(abilityType database.AbilityType) handlerinput.DBQueryNullIntMany {
	return func(ctx context.Context, status sql.NullInt32) ([]int32, error) {
		return config.Cfg.DB.GetTypedAbilityIDsByRemovedStatus(ctx, database.GetTypedAbilityIDsByRemovedStatusParams{
			StatusID: status,
			Type:     abilityType,
		})
	}
}

func getTypedAbilityIDsByElement(abilityType database.AbilityType) handlerinput.DBQueryIntList {
	return func(ctx context.Context, elements []int32) ([]int32, error) {
		return config.Cfg.DB.GetTypedAbilityIDsByElement(ctx, database.GetTypedAbilityIDsByElementParams{
			ElementID: elements,
			Type:      abilityType,
		})
	}
}
