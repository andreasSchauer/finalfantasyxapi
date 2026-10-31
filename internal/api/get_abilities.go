package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getAbility(r *http.Request, i handlerInput[seeding.Ability, Ability, TypedAPIResource, TypedAPIResourceList], id int32) (Ability, error) {
	ability, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return Ability{}, err
	}

	monsters, err := getResourcesDBItem(cfg, r.Context(), cfg.E.Monsters, ability, cfg.db.GetAbilityMonsterIDs)
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

func (cfg *Config) retrieveAbilities(r *http.Request, i handlerInput[seeding.Ability, Ability, TypedAPIResource, TypedAPIResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []IDFilter{
		enumListQuery(cfg, r, i, cfg.T.AbilityType, ids, QpnType, cfg.db.GetAbilityIDsByType),
		enumListQuery(cfg, r, i, cfg.T.DamageType, ids, QpnDamageType, cfg.db.GetAbilityIDsByDamageType),
		enumListQuery(cfg, r, i, cfg.T.AttackType, ids, QpnAttackType, cfg.db.GetAbilityIDsByAttackType),
		enumListQuery(cfg, r, i, cfg.T.TargetType, ids, QpnTargetType, cfg.db.GetAbilityIDsByTargetType),
		enumQuery(r, i, cfg.T.DamageFormula, ids, QpnDamageFormula, cfg.db.GetAbilityIDsByDamageFormula),
		intListQuery(cfg, r, i, ids, QpnRank, cfg.db.GetAbilityIDsByRank),
		nameIDListQueryNul(cfg, r, i, ids, QpnElement, cfg.E.Elements.resTypeSingle, cfg.L.Elements, cfg.db.GetAbilityIDsByElement),
		idQueryNul(r, i, ids, QpnStatusInflict, cfg.L.StatusConditions, cfg.db.GetAbilityIDsByInflictedStatus),
		idQueryNul(r, i, ids, QpnStatusRemove, cfg.L.StatusConditions, cfg.db.GetAbilityIDsByRemovedStatus),
		idQuery(r, i, ids, QpnMonster, cfg.L.MonsterFormations, cfg.db.GetAbilityIDsByMonster),
		boolQuery(r, i, ids, QpnCopycat, cfg.db.GetAbilityIDsByCanCopycat),
		boolQuery(r, i, ids, QpnHelpBar, cfg.db.GetAbilityIDsByAppearsInHelpBar),
		boolQuery2(r, i, ids, QpnCanCrit, cfg.db.GetAbilityIDsCanCrit),
		boolQuery2(r, i, ids, QpnBDL, cfg.db.GetAbilityIDsBreakDmgLimit),
		boolQuery2(r, i, ids, QpnUserAtk, cfg.db.GetAbilityIDsBasedOnUserAttack),
		boolQuery2(r, i, ids, QpnDarkable, cfg.db.GetAbilityIDsDarkable),
		boolQuery2(r, i, ids, QpnSilenceable, cfg.db.GetAbilityIDsSilenceable),
		boolQuery2(r, i, ids, QpnReflectable, cfg.db.GetAbilityIDsReflectable),
		boolQuery2(r, i, ids, QpnDelay, cfg.db.GetAbilityIDsDealsDelay),
		boolQuery2(r, i, ids, QpnStatChanges, cfg.db.GetAbilityIDsWithStatChanges),
		boolQuery2(r, i, ids, QpnModChanges, cfg.db.GetAbilityIDsWithModifierChanges),
	})
}
