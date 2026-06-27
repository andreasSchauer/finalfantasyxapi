package api

import (
	"context"
	"database/sql"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func getTypedAbilityIDsByName(cfg *Config, abilityType database.AbilityType) DbQueryStringMany {
	return func(ctx context.Context, name string) ([]int32, error) {
		return cfg.db.GetTypedAbilityIDsByName(ctx, database.GetTypedAbilityIDsByNameParams{
			Name: name,
			Type: abilityType,
		})
	}
}

func getTypedAbilityIDsByRank(cfg *Config, abilityType database.AbilityType) DbQueryIntList {
	return func(ctx context.Context, ranks []int32) ([]int32, error) {
		return cfg.db.GetTypedAbilityIDsByRank(ctx, database.GetTypedAbilityIDsByRankParams{
			Rank: ranks,
			Type: abilityType,
		})
	}
}

func getTypedAbilityIDsByCanCopycat(cfg *Config, abilityType database.AbilityType) DbQueryBool {
	return func(ctx context.Context, canCopycat bool) ([]int32, error) {
		return cfg.db.GetTypedAbilityIDsByCanCopycat(ctx, database.GetTypedAbilityIDsByCanCopycatParams{
			CanCopycat: canCopycat,
			Type: abilityType,
		})
	}
}

func getTypedAbilityIDsByAppearsInHelpBar(cfg *Config, abilityType database.AbilityType) DbQueryBool {
	return func(ctx context.Context, inHelpBar bool) ([]int32, error) {
		return cfg.db.GetTypedAbilityIDsByAppearsInHelpBar(ctx, database.GetTypedAbilityIDsByAppearsInHelpBarParams{
			AppearsInHelpBar: inHelpBar,
			Type: abilityType,
		})
	}
}

func getTypedAbilityIDsBasedOnUserAttack(cfg *Config, abilityType database.AbilityType) DbQueryNoInput {
	return func(ctx context.Context) ([]int32, error) {
		return cfg.db.GetTypedAbilityIDsBasedOnUserAttack(ctx, abilityType)
	}
}

func getTypedAbilityIDsByTargetType(cfg *Config, abilityType database.AbilityType) DbQueryEnumList[database.TargetType] {
	return func(ctx context.Context, targetTypes []database.TargetType) ([]int32, error) {
		return cfg.db.GetTypedAbilityIDsByTargetType(ctx, database.GetTypedAbilityIDsByTargetTypeParams{
			TargetType: targetTypes,
			Type: abilityType,
		})
	}
}

func getTypedAbilityIDsDarkable(cfg *Config, abilityType database.AbilityType) DbQueryNoInput {
	return func(ctx context.Context) ([]int32, error) {
		return cfg.db.GetTypedAbilityIDsDarkable(ctx, abilityType)
	}
}

func getTypedAbilityIDsSilenceable(cfg *Config, abilityType database.AbilityType) DbQueryNoInput {
	return func(ctx context.Context) ([]int32, error) {
		return cfg.db.GetTypedAbilityIDsSilenceable(ctx, abilityType)
	}
}

func getTypedAbilityIDsReflectable(cfg *Config, abilityType database.AbilityType) DbQueryNoInput {
	return func(ctx context.Context) ([]int32, error) {
		return cfg.db.GetTypedAbilityIDsReflectable(ctx, abilityType)
	}
}

func getTypedAbilityIDsDealsDelay(cfg *Config, abilityType database.AbilityType) DbQueryNoInput {
	return func(ctx context.Context) ([]int32, error) {
		return cfg.db.GetTypedAbilityIDsDealsDelay(ctx, abilityType)
	}
}

func getTypedAbilityIDsWithStatChanges(cfg *Config, abilityType database.AbilityType) DbQueryNoInput {
	return func(ctx context.Context) ([]int32, error) {
		return cfg.db.GetTypedAbilityIDsWithStatChanges(ctx, abilityType)
	}
}

func getTypedAbilityIDsWithModifierChanges(cfg *Config, abilityType database.AbilityType) DbQueryNoInput {
	return func(ctx context.Context) ([]int32, error) {
		return cfg.db.GetTypedAbilityIDsWithModifierChanges(ctx, abilityType)
	}
}

func getTypedAbilityIDsCanCrit(cfg *Config, abilityType database.AbilityType) DbQueryNoInput {
	return func(ctx context.Context) ([]int32, error) {
		return cfg.db.GetTypedAbilityIDsCanCrit(ctx, abilityType)
	}
}

func getTypedAbilityIDsBreakDmgLimit(cfg *Config, abilityType database.AbilityType) DbQueryNoInput {
	return func(ctx context.Context) ([]int32, error) {
		return cfg.db.GetTypedAbilityIDsBreakDmgLimit(ctx, abilityType)
	}
}

func getTypedAbilityIDsByDamageType(cfg *Config, abilityType database.AbilityType) DbQueryEnumList[database.DamageType] {
	return func(ctx context.Context, damageTypes []database.DamageType) ([]int32, error) {
		return cfg.db.GetTypedAbilityIDsByDamageType(ctx, database.GetTypedAbilityIDsByDamageTypeParams{
			DamageType: damageTypes,
			Type: abilityType,
		})
	}
}

func getTypedAbilityIDsByAttackType(cfg *Config, abilityType database.AbilityType) DbQueryEnumList[database.AttackType] {
	return func(ctx context.Context, attackTypes []database.AttackType) ([]int32, error) {
		return cfg.db.GetTypedAbilityIDsByAttackType(ctx, database.GetTypedAbilityIDsByAttackTypeParams{
			AttackType: attackTypes,
			Type: abilityType,
		})
	}
}

func getTypedAbilityIDsByDamageFormula(cfg *Config, abilityType database.AbilityType) DbQueryEnum[database.DamageFormula] {
	return func(ctx context.Context, formula database.DamageFormula) ([]int32, error) {
		return cfg.db.GetTypedAbilityIDsByDamageFormula(ctx, database.GetTypedAbilityIDsByDamageFormulaParams{
			DamageFormula: formula,
			Type: abilityType,
		})
	}
}

func getTypedAbilityIDsByInflictedStatus(cfg *Config, abilityType database.AbilityType) DbQueryNullIntMany {
	return func(ctx context.Context, status sql.NullInt32) ([]int32, error) {
		return cfg.db.GetTypedAbilityIDsByInflictedStatus(ctx, database.GetTypedAbilityIDsByInflictedStatusParams{
			StatusID: status,
			Type: abilityType,
		})
	}
}

func getTypedAbilityIDsByRemovedStatus(cfg *Config, abilityType database.AbilityType) DbQueryNullIntMany {
	return func(ctx context.Context, status sql.NullInt32) ([]int32, error) {
		return cfg.db.GetTypedAbilityIDsByRemovedStatus(ctx, database.GetTypedAbilityIDsByRemovedStatusParams{
			StatusID: status,
			Type: abilityType,
		})
	}
}

func getTypedAbilityIDsByElement(cfg *Config, abilityType database.AbilityType) DbQueryIntList {
	return func(ctx context.Context, elements []int32) ([]int32, error) {
		return cfg.db.GetTypedAbilityIDsByElement(ctx, database.GetTypedAbilityIDsByElementParams{
			ElementID: elements,
			Type: abilityType,
		})
	}
}