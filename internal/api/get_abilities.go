package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
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
		i.queryEnumList(cfg, r, cfg.T.AbilityType, ids, QpnType, cfg.db.GetAbilityIDsByType),
		i.queryEnumList(cfg, r, cfg.T.DamageType, ids, QpnDamageType, cfg.db.GetAbilityIDsByDamageType),
		i.queryEnumList(cfg, r, cfg.T.AttackType, ids, QpnAttackType, cfg.db.GetAbilityIDsByAttackType),
		i.queryEnumList(cfg, r, cfg.T.TargetType, ids, QpnTargetType, cfg.db.GetAbilityIDsByTargetType),
		i.queryEnum(r, cfg.T.DamageFormula, ids, QpnDamageFormula, cfg.db.GetAbilityIDsByDamageFormula),
		i.queryIntList(cfg, r, ids, QpnRank, cfg.db.GetAbilityIDsByRank),
		i.queryNameIDListNul(cfg, r, ids, QpnElement, cfg.E.Elements.resTypeSingle, cfg.L.Elements, cfg.db.GetAbilityIDsByElement),
		i.queryIDNul(r, ids, QpnStatusInflict, cfg.L.StatusConditions, cfg.db.GetAbilityIDsByInflictedStatus),
		i.queryIDNul(r, ids, QpnStatusRemove, cfg.L.StatusConditions, cfg.db.GetAbilityIDsByRemovedStatus),
		i.queryID(r, ids, QpnMonster, cfg.L.MonsterFormations, cfg.db.GetAbilityIDsByMonster),
		i.queryBool(r, ids, QpnCopycat, cfg.db.GetAbilityIDsByCanCopycat),
		i.queryBool(r, ids, QpnHelpBar, cfg.db.GetAbilityIDsByAppearsInHelpBar),
		i.queryBool2(r, ids, QpnCanCrit, cfg.db.GetAbilityIDsCanCrit),
		i.queryBool2(r, ids, QpnBDL, cfg.db.GetAbilityIDsBreakDmgLimit),
		i.queryBool2(r, ids, QpnUserAtk, cfg.db.GetAbilityIDsBasedOnUserAttack),
		i.queryBool2(r, ids, QpnDarkable, cfg.db.GetAbilityIDsDarkable),
		i.queryBool2(r, ids, QpnSilenceable, cfg.db.GetAbilityIDsSilenceable),
		i.queryBool2(r, ids, QpnReflectable, cfg.db.GetAbilityIDsReflectable),
		i.queryBool2(r, ids, QpnDelay, cfg.db.GetAbilityIDsDealsDelay),
		i.queryBool2(r, ids, QpnStatChanges, cfg.db.GetAbilityIDsWithStatChanges),
		i.queryBool2(r, ids, QpnModChanges, cfg.db.GetAbilityIDsWithModifierChanges),
	})
}
