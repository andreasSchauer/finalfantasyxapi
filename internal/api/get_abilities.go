package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getAbility(r *http.Request, i handlerInput[seeding.Ability, Ability, TypedAPIResource, TypedAPIResourceList], id int32) (Ability, error) {
	ability, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return Ability{}, err
	}

	monsters, err := getResourcesDbItem(cfg, r, cfg.e.monsters, ability, cfg.db.GetAbilityMonsterIDs)
	if err != nil {
		return Ability{}, err
	}

	response := Ability{
		ID:                 ability.ID,
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		Type:               enumToNamedAPIResource(cfg, cfg.e.abilityType.endpoint, string(ability.Type), cfg.t.AbilityType),
		Rank:               ability.Rank,
		CanCopycat:         ability.CanCopycat,
		AppearsInHelpBar:   ability.AppearsInHelpBar,
		TypedAbility:       refToNamedApiResource(cfg, ability.GetAbilityRef()),
		Monsters:           monsters,
		BattleInteractions: convertObjSlice(cfg, ability.BattleInteractions, convertBattleInteraction),
	}

	return response, nil
}

func (cfg *Config) retrieveAbilities(r *http.Request, i handlerInput[seeding.Ability, Ability, TypedAPIResource, TypedAPIResourceList]) (TypedAPIResourceList, error) {
	ids, err := verifyParamsAndRetrieve(cfg, r, i)
	if err != nil {
		return TypedAPIResourceList{}, err
	}

	return filterIDs(cfg, r, i, ids, []filteredIdList{
		fidl(enumListQuery(cfg, r, i, cfg.t.AbilityType, ids, qpnType, cfg.db.GetAbilityIDsByType)),
		fidl(enumListQuery(cfg, r, i, cfg.t.DamageType, ids, qpnDamageType, cfg.db.GetAbilityIDsByDamageType)),
		fidl(enumListQuery(cfg, r, i, cfg.t.AttackType, ids, qpnAttackType, cfg.db.GetAbilityIDsByAttackType)),
		fidl(enumListQuery(cfg, r, i, cfg.t.TargetType, ids, qpnTargetType, cfg.db.GetAbilityIDsByTargetType)),
		fidl(enumQuery(r, i, cfg.t.DamageFormula, ids, qpnDamageFormula, cfg.db.GetAbilityIDsByDamageFormula)),
		fidl(intListQuery(cfg, r, i, ids, qpnRank, cfg.db.GetAbilityIDsByRank)),
		fidl(nameIdListQueryNul(cfg, r, i, ids, qpnElement, cfg.e.elements.resTypeSing, cfg.l.Elements, cfg.db.GetAbilityIDsByElement)),
		fidl(idQueryNul(r, i, ids, qpnStatusInflict, cfg.l.StatusConditions, cfg.db.GetAbilityIDsByInflictedStatus)),
		fidl(idQueryNul(r, i, ids, qpnStatusRemove, cfg.l.StatusConditions, cfg.db.GetAbilityIDsByRemovedStatus)),
		fidl(idQuery(r, i, ids, qpnMonster, cfg.l.MonsterFormations, cfg.db.GetAbilityIDsByMonster)),
		fidl(boolQuery(r, i, ids, qpnCopycat, cfg.db.GetAbilityIDsByCanCopycat)),
		fidl(boolQuery(r, i, ids, qpnHelpBar, cfg.db.GetAbilityIDsByAppearsInHelpBar)),
		fidl(boolQuery2(r, i, ids, qpnCanCrit, cfg.db.GetAbilityIDsCanCrit)),
		fidl(boolQuery2(r, i, ids, qpnBDL, cfg.db.GetAbilityIDsBreakDmgLimit)),
		fidl(boolQuery2(r, i, ids, qpnUserAtk, cfg.db.GetAbilityIDsBasedOnUserAttack)),
		fidl(boolQuery2(r, i, ids, qpnDarkable, cfg.db.GetAbilityIDsDarkable)),
		fidl(boolQuery2(r, i, ids, qpnSilenceable, cfg.db.GetAbilityIDsSilenceable)),
		fidl(boolQuery2(r, i, ids, qpnReflectable, cfg.db.GetAbilityIDsReflectable)),
		fidl(boolQuery2(r, i, ids, qpnDelay, cfg.db.GetAbilityIDsDealsDelay)),
		fidl(boolQuery2(r, i, ids, qpnStatChanges, cfg.db.GetAbilityIDsWithStatChanges)),
		fidl(boolQuery2(r, i, ids, qpnModChanges, cfg.db.GetAbilityIDsWithModifierChanges)),
	})
}
