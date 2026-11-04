package api

import (
	"context"
	"database/sql"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func getTypedAbilityIDsByName(srv *Server, abilityType database.AbilityType) DBQueryStringMany {
	return func(ctx context.Context, name string) ([]int32, error) {
		return srv.Cfg.DB.GetTypedAbilityIDsByName(ctx, database.GetTypedAbilityIDsByNameParams{
			Name: name,
			Type: abilityType,
		})
	}
}

func getTypedAbilityIDsByRank(srv *Server, abilityType database.AbilityType) DBQueryIntList {
	return func(ctx context.Context, ranks []int32) ([]int32, error) {
		return srv.Cfg.DB.GetTypedAbilityIDsByRank(ctx, database.GetTypedAbilityIDsByRankParams{
			Rank: ranks,
			Type: abilityType,
		})
	}
}

func getTypedAbilityIDsByCanCopycat(srv *Server, abilityType database.AbilityType) DBQueryBool {
	return func(ctx context.Context, canCopycat bool) ([]int32, error) {
		return srv.Cfg.DB.GetTypedAbilityIDsByCanCopycat(ctx, database.GetTypedAbilityIDsByCanCopycatParams{
			CanCopycat: canCopycat,
			Type:       abilityType,
		})
	}
}

func getTypedAbilityIDsByAppearsInHelpBar(srv *Server, abilityType database.AbilityType) DBQueryBool {
	return func(ctx context.Context, inHelpBar bool) ([]int32, error) {
		return srv.Cfg.DB.GetTypedAbilityIDsByAppearsInHelpBar(ctx, database.GetTypedAbilityIDsByAppearsInHelpBarParams{
			AppearsInHelpBar: inHelpBar,
			Type:             abilityType,
		})
	}
}

func getTypedAbilityIDsBasedOnUserAttack(srv *Server, abilityType database.AbilityType) DBQueryNoInput {
	return func(ctx context.Context) ([]int32, error) {
		return srv.Cfg.DB.GetTypedAbilityIDsBasedOnUserAttack(ctx, abilityType)
	}
}

func getTypedAbilityIDsByTargetType(srv *Server, abilityType database.AbilityType) DBQueryEnumList[database.TargetType] {
	return func(ctx context.Context, targetTypes []database.TargetType) ([]int32, error) {
		return srv.Cfg.DB.GetTypedAbilityIDsByTargetType(ctx, database.GetTypedAbilityIDsByTargetTypeParams{
			TargetType: targetTypes,
			Type:       abilityType,
		})
	}
}

func getTypedAbilityIDsDarkable(srv *Server, abilityType database.AbilityType) DBQueryNoInput {
	return func(ctx context.Context) ([]int32, error) {
		return srv.Cfg.DB.GetTypedAbilityIDsDarkable(ctx, abilityType)
	}
}

func getTypedAbilityIDsSilenceable(srv *Server, abilityType database.AbilityType) DBQueryNoInput {
	return func(ctx context.Context) ([]int32, error) {
		return srv.Cfg.DB.GetTypedAbilityIDsSilenceable(ctx, abilityType)
	}
}

func getTypedAbilityIDsReflectable(srv *Server, abilityType database.AbilityType) DBQueryNoInput {
	return func(ctx context.Context) ([]int32, error) {
		return srv.Cfg.DB.GetTypedAbilityIDsReflectable(ctx, abilityType)
	}
}

func getTypedAbilityIDsDealsDelay(srv *Server, abilityType database.AbilityType) DBQueryNoInput {
	return func(ctx context.Context) ([]int32, error) {
		return srv.Cfg.DB.GetTypedAbilityIDsDealsDelay(ctx, abilityType)
	}
}

func getTypedAbilityIDsWithStatChanges(srv *Server, abilityType database.AbilityType) DBQueryNoInput {
	return func(ctx context.Context) ([]int32, error) {
		return srv.Cfg.DB.GetTypedAbilityIDsWithStatChanges(ctx, abilityType)
	}
}

func getTypedAbilityIDsWithModifierChanges(srv *Server, abilityType database.AbilityType) DBQueryNoInput {
	return func(ctx context.Context) ([]int32, error) {
		return srv.Cfg.DB.GetTypedAbilityIDsWithModifierChanges(ctx, abilityType)
	}
}

func getTypedAbilityIDsCanCrit(srv *Server, abilityType database.AbilityType) DBQueryNoInput {
	return func(ctx context.Context) ([]int32, error) {
		return srv.Cfg.DB.GetTypedAbilityIDsCanCrit(ctx, abilityType)
	}
}

func getTypedAbilityIDsBreakDmgLimit(srv *Server, abilityType database.AbilityType) DBQueryNoInput {
	return func(ctx context.Context) ([]int32, error) {
		return srv.Cfg.DB.GetTypedAbilityIDsBreakDmgLimit(ctx, abilityType)
	}
}

func getTypedAbilityIDsByDamageType(srv *Server, abilityType database.AbilityType) DBQueryEnumList[database.DamageType] {
	return func(ctx context.Context, damageTypes []database.DamageType) ([]int32, error) {
		return srv.Cfg.DB.GetTypedAbilityIDsByDamageType(ctx, database.GetTypedAbilityIDsByDamageTypeParams{
			DamageType: damageTypes,
			Type:       abilityType,
		})
	}
}

func getTypedAbilityIDsByAttackType(srv *Server, abilityType database.AbilityType) DBQueryEnumList[database.AttackType] {
	return func(ctx context.Context, attackTypes []database.AttackType) ([]int32, error) {
		return srv.Cfg.DB.GetTypedAbilityIDsByAttackType(ctx, database.GetTypedAbilityIDsByAttackTypeParams{
			AttackType: attackTypes,
			Type:       abilityType,
		})
	}
}

func getTypedAbilityIDsByDamageFormula(srv *Server, abilityType database.AbilityType) DBQueryEnum[database.DamageFormula] {
	return func(ctx context.Context, formula database.DamageFormula) ([]int32, error) {
		return srv.Cfg.DB.GetTypedAbilityIDsByDamageFormula(ctx, database.GetTypedAbilityIDsByDamageFormulaParams{
			DamageFormula: formula,
			Type:          abilityType,
		})
	}
}

func getTypedAbilityIDsByInflictedStatus(srv *Server, abilityType database.AbilityType) DBQueryNullIntMany {
	return func(ctx context.Context, status sql.NullInt32) ([]int32, error) {
		return srv.Cfg.DB.GetTypedAbilityIDsByInflictedStatus(ctx, database.GetTypedAbilityIDsByInflictedStatusParams{
			StatusID: status,
			Type:     abilityType,
		})
	}
}

func getTypedAbilityIDsByRemovedStatus(srv *Server, abilityType database.AbilityType) DBQueryNullIntMany {
	return func(ctx context.Context, status sql.NullInt32) ([]int32, error) {
		return srv.Cfg.DB.GetTypedAbilityIDsByRemovedStatus(ctx, database.GetTypedAbilityIDsByRemovedStatusParams{
			StatusID: status,
			Type:     abilityType,
		})
	}
}

func getTypedAbilityIDsByElement(srv *Server, abilityType database.AbilityType) DBQueryIntList {
	return func(ctx context.Context, elements []int32) ([]int32, error) {
		return srv.Cfg.DB.GetTypedAbilityIDsByElement(ctx, database.GetTypedAbilityIDsByElementParams{
			ElementID: elements,
			Type:      abilityType,
		})
	}
}
