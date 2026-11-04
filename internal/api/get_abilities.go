package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (cfg *Config) getAbility(r *http.Request, i handlerInput[seedtypes.Ability, Ability, TypedAPIResource], id int32) (Ability, error) {
	ability, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Ability{}, err
	}

	monsters, err := cfg.E.Monsters.dbGetRelations(cfg, r.Context(), ability, cfg.db.GetAbilityMonsterIDs)
	if err != nil {
		return Ability{}, err
	}

	response := Ability{
		ID:                 ability.ID,
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		Type:               ability.Type,
		Rank:               ability.Rank,
		CanCopycat:         ability.CanCopycat,
		AppearsInHelpBar:   ability.AppearsInHelpBar,
		TypedAbility:       refToNamedAPIResource(cfg, ability.GetAbilityRef()),
		Monsters:           monsters,
		BattleInteractions: convertObjSlice(cfg, ability.BattleInteractions, convertBattleInteraction),
	}

	return response, nil
}

func (cfg *Config) retrieveAbilities(r *http.Request, i handlerInput[seedtypes.Ability, Ability, TypedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(cfg, r, ids, []IDFilter{
		i.queryEnumList(cfg, r, enums.T.AbilityType, ids, typestrings.QpnType, cfg.db.GetAbilityIDsByType),
		i.queryEnumList(cfg, r, enums.T.DamageType, ids, typestrings.QpnDamageType, cfg.db.GetAbilityIDsByDamageType),
		i.queryEnumList(cfg, r, enums.T.AttackType, ids, typestrings.QpnAttackType, cfg.db.GetAbilityIDsByAttackType),
		i.queryEnumList(cfg, r, enums.T.TargetType, ids, typestrings.QpnTargetType, cfg.db.GetAbilityIDsByTargetType),
		i.queryEnum(r, enums.T.DamageFormula, ids, typestrings.QpnDamageFormula, cfg.db.GetAbilityIDsByDamageFormula),
		i.queryIntList(cfg, r, ids, typestrings.QpnRank, cfg.db.GetAbilityIDsByRank),
		i.queryNameIDListNul(cfg, r, ids, typestrings.QpnElement, cfg.E.Elements.resTypeSingle, lookup.L.Elements, cfg.db.GetAbilityIDsByElement),
		i.queryIDNul(r, ids, typestrings.QpnStatusInflict, lookup.L.StatusConditions, cfg.db.GetAbilityIDsByInflictedStatus),
		i.queryIDNul(r, ids, typestrings.QpnStatusRemove, lookup.L.StatusConditions, cfg.db.GetAbilityIDsByRemovedStatus),
		i.queryID(r, ids, typestrings.QpnMonster, lookup.L.MonsterFormations, cfg.db.GetAbilityIDsByMonster),
		i.queryBool(r, ids, typestrings.QpnCopycat, cfg.db.GetAbilityIDsByCanCopycat),
		i.queryBool(r, ids, typestrings.QpnHelpBar, cfg.db.GetAbilityIDsByAppearsInHelpBar),
		i.queryBool2(r, ids, typestrings.QpnCanCrit, cfg.db.GetAbilityIDsCanCrit),
		i.queryBool2(r, ids, typestrings.QpnBDL, cfg.db.GetAbilityIDsBreakDmgLimit),
		i.queryBool2(r, ids, typestrings.QpnUserAtk, cfg.db.GetAbilityIDsBasedOnUserAttack),
		i.queryBool2(r, ids, typestrings.QpnDarkable, cfg.db.GetAbilityIDsDarkable),
		i.queryBool2(r, ids, typestrings.QpnSilenceable, cfg.db.GetAbilityIDsSilenceable),
		i.queryBool2(r, ids, typestrings.QpnReflectable, cfg.db.GetAbilityIDsReflectable),
		i.queryBool2(r, ids, typestrings.QpnDelay, cfg.db.GetAbilityIDsDealsDelay),
		i.queryBool2(r, ids, typestrings.QpnStatChanges, cfg.db.GetAbilityIDsWithStatChanges),
		i.queryBool2(r, ids, typestrings.QpnModChanges, cfg.db.GetAbilityIDsWithModifierChanges),
	})
}
