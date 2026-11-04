package api

import (
	"context"
	"database/sql"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func getTypedAbilityIDsByName(server *Server, abilityType database.AbilityType) DBQueryStringMany {
	return func(ctx context.Context, name string) ([]int32, error) {
		return server.Cfg.DB.GetTypedAbilityIDsByName(ctx, database.GetTypedAbilityIDsByNameParams{
			Name: name,
			Type: abilityType,
		})
	}
}

func getTypedAbilityIDsByRank(server *Server, abilityType database.AbilityType) DBQueryIntList {
	return func(ctx context.Context, ranks []int32) ([]int32, error) {
		return server.Cfg.DB.GetTypedAbilityIDsByRank(ctx, database.GetTypedAbilityIDsByRankParams{
			Rank: ranks,
			Type: abilityType,
		})
	}
}

func getTypedAbilityIDsByCanCopycat(server *Server, abilityType database.AbilityType) DBQueryBool {
	return func(ctx context.Context, canCopycat bool) ([]int32, error) {
		return server.Cfg.DB.GetTypedAbilityIDsByCanCopycat(ctx, database.GetTypedAbilityIDsByCanCopycatParams{
			CanCopycat: canCopycat,
			Type:       abilityType,
		})
	}
}

func getTypedAbilityIDsByAppearsInHelpBar(server *Server, abilityType database.AbilityType) DBQueryBool {
	return func(ctx context.Context, inHelpBar bool) ([]int32, error) {
		return server.Cfg.DB.GetTypedAbilityIDsByAppearsInHelpBar(ctx, database.GetTypedAbilityIDsByAppearsInHelpBarParams{
			AppearsInHelpBar: inHelpBar,
			Type:             abilityType,
		})
	}
}

func getTypedAbilityIDsBasedOnUserAttack(server *Server, abilityType database.AbilityType) DBQueryNoInput {
	return func(ctx context.Context) ([]int32, error) {
		return server.Cfg.DB.GetTypedAbilityIDsBasedOnUserAttack(ctx, abilityType)
	}
}

func getTypedAbilityIDsByTargetType(server *Server, abilityType database.AbilityType) DBQueryEnumList[database.TargetType] {
	return func(ctx context.Context, targetTypes []database.TargetType) ([]int32, error) {
		return server.Cfg.DB.GetTypedAbilityIDsByTargetType(ctx, database.GetTypedAbilityIDsByTargetTypeParams{
			TargetType: targetTypes,
			Type:       abilityType,
		})
	}
}

func getTypedAbilityIDsDarkable(server *Server, abilityType database.AbilityType) DBQueryNoInput {
	return func(ctx context.Context) ([]int32, error) {
		return server.Cfg.DB.GetTypedAbilityIDsDarkable(ctx, abilityType)
	}
}

func getTypedAbilityIDsSilenceable(server *Server, abilityType database.AbilityType) DBQueryNoInput {
	return func(ctx context.Context) ([]int32, error) {
		return server.Cfg.DB.GetTypedAbilityIDsSilenceable(ctx, abilityType)
	}
}

func getTypedAbilityIDsReflectable(server *Server, abilityType database.AbilityType) DBQueryNoInput {
	return func(ctx context.Context) ([]int32, error) {
		return server.Cfg.DB.GetTypedAbilityIDsReflectable(ctx, abilityType)
	}
}

func getTypedAbilityIDsDealsDelay(server *Server, abilityType database.AbilityType) DBQueryNoInput {
	return func(ctx context.Context) ([]int32, error) {
		return server.Cfg.DB.GetTypedAbilityIDsDealsDelay(ctx, abilityType)
	}
}

func getTypedAbilityIDsWithStatChanges(server *Server, abilityType database.AbilityType) DBQueryNoInput {
	return func(ctx context.Context) ([]int32, error) {
		return server.Cfg.DB.GetTypedAbilityIDsWithStatChanges(ctx, abilityType)
	}
}

func getTypedAbilityIDsWithModifierChanges(server *Server, abilityType database.AbilityType) DBQueryNoInput {
	return func(ctx context.Context) ([]int32, error) {
		return server.Cfg.DB.GetTypedAbilityIDsWithModifierChanges(ctx, abilityType)
	}
}

func getTypedAbilityIDsCanCrit(server *Server, abilityType database.AbilityType) DBQueryNoInput {
	return func(ctx context.Context) ([]int32, error) {
		return server.Cfg.DB.GetTypedAbilityIDsCanCrit(ctx, abilityType)
	}
}

func getTypedAbilityIDsBreakDmgLimit(server *Server, abilityType database.AbilityType) DBQueryNoInput {
	return func(ctx context.Context) ([]int32, error) {
		return server.Cfg.DB.GetTypedAbilityIDsBreakDmgLimit(ctx, abilityType)
	}
}

func getTypedAbilityIDsByDamageType(server *Server, abilityType database.AbilityType) DBQueryEnumList[database.DamageType] {
	return func(ctx context.Context, damageTypes []database.DamageType) ([]int32, error) {
		return server.Cfg.DB.GetTypedAbilityIDsByDamageType(ctx, database.GetTypedAbilityIDsByDamageTypeParams{
			DamageType: damageTypes,
			Type:       abilityType,
		})
	}
}

func getTypedAbilityIDsByAttackType(server *Server, abilityType database.AbilityType) DBQueryEnumList[database.AttackType] {
	return func(ctx context.Context, attackTypes []database.AttackType) ([]int32, error) {
		return server.Cfg.DB.GetTypedAbilityIDsByAttackType(ctx, database.GetTypedAbilityIDsByAttackTypeParams{
			AttackType: attackTypes,
			Type:       abilityType,
		})
	}
}

func getTypedAbilityIDsByDamageFormula(server *Server, abilityType database.AbilityType) DBQueryEnum[database.DamageFormula] {
	return func(ctx context.Context, formula database.DamageFormula) ([]int32, error) {
		return server.Cfg.DB.GetTypedAbilityIDsByDamageFormula(ctx, database.GetTypedAbilityIDsByDamageFormulaParams{
			DamageFormula: formula,
			Type:          abilityType,
		})
	}
}

func getTypedAbilityIDsByInflictedStatus(server *Server, abilityType database.AbilityType) DBQueryNullIntMany {
	return func(ctx context.Context, status sql.NullInt32) ([]int32, error) {
		return server.Cfg.DB.GetTypedAbilityIDsByInflictedStatus(ctx, database.GetTypedAbilityIDsByInflictedStatusParams{
			StatusID: status,
			Type:     abilityType,
		})
	}
}

func getTypedAbilityIDsByRemovedStatus(server *Server, abilityType database.AbilityType) DBQueryNullIntMany {
	return func(ctx context.Context, status sql.NullInt32) ([]int32, error) {
		return server.Cfg.DB.GetTypedAbilityIDsByRemovedStatus(ctx, database.GetTypedAbilityIDsByRemovedStatusParams{
			StatusID: status,
			Type:     abilityType,
		})
	}
}

func getTypedAbilityIDsByElement(server *Server, abilityType database.AbilityType) DBQueryIntList {
	return func(ctx context.Context, elements []int32) ([]int32, error) {
		return server.Cfg.DB.GetTypedAbilityIDsByElement(ctx, database.GetTypedAbilityIDsByElementParams{
			ElementID: elements,
			Type:      abilityType,
		})
	}
}
