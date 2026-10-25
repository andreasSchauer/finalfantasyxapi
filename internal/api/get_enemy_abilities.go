package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getEnemyAbility(r *http.Request, i handlerInput[seeding.EnemyAbility, EnemyAbility, NamedAPIResource, NamedApiResourceList], id int32) (EnemyAbility, error) {
	ability, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return EnemyAbility{}, err
	}

	monsters, err := getResourcesDbItem(cfg, r.Context(), cfg.E.Monsters, ability, cfg.db.GetEnemyAbilityMonsterIDs)
	if err != nil {
		return EnemyAbility{}, err
	}

	response := EnemyAbility{
		ID:                 ability.ID,
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		UntypedAbility:     idToTypedAPIResource(cfg, cfg.E.Abilities, ability.Ability.ID),
		Effect:             ability.Effect,
		Rank:               ability.Rank,
		AppearsInHelpBar:   ability.AppearsInHelpBar,
		CanCopycat:         ability.CanCopycat,
		Monsters:           monsters,
		BattleInteractions: convertObjSlice(cfg, ability.BattleInteractions, convertBattleInteraction),
	}

	return response, nil
}

func (cfg *Config) retrieveEnemyAbilities(r *http.Request, i handlerInput[seeding.EnemyAbility, EnemyAbility, NamedAPIResource, NamedApiResourceList]) ([]int32, error) {
	abilityType := database.AbilityTypeEnemyAbility
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []IdFilter{
		idQuery(r, i, ids, QpnMonster, cfg.L.Monsters, cfg.db.GetEnemyAbilityIDsByMonster),
		enumListQuery(cfg, r, i, cfg.T.DamageType, ids, QpnDamageType, getTypedAbilityIDsByDamageType(cfg, abilityType)),
		enumListQuery(cfg, r, i, cfg.T.AttackType, ids, QpnAttackType, getTypedAbilityIDsByAttackType(cfg, abilityType)),
		enumListQuery(cfg, r, i, cfg.T.TargetType, ids, QpnTargetType, getTypedAbilityIDsByTargetType(cfg, abilityType)),
		enumQuery(r, i, cfg.T.DamageFormula, ids, QpnDamageFormula, getTypedAbilityIDsByDamageFormula(cfg, abilityType)),
		intListQuery(cfg, r, i, ids, QpnRank, getTypedAbilityIDsByRank(cfg, abilityType)),
		nameIdListQueryNul(cfg, r, i, ids, QpnElement, cfg.E.Elements.resTypeSingle, cfg.L.Elements, getTypedAbilityIDsByElement(cfg, abilityType)),
		idQueryNul(r, i, ids, QpnStatusInflict, cfg.L.StatusConditions, getTypedAbilityIDsByInflictedStatus(cfg, abilityType)),
		idQueryNul(r, i, ids, QpnStatusRemove, cfg.L.StatusConditions, getTypedAbilityIDsByRemovedStatus(cfg, abilityType)),
		boolQuery(r, i, ids, QpnHelpBar, getTypedAbilityIDsByAppearsInHelpBar(cfg, abilityType)),
		boolQuery2(r, i, ids, QpnCanCrit, getTypedAbilityIDsCanCrit(cfg, abilityType)),
		boolQuery2(r, i, ids, QpnBDL, getTypedAbilityIDsBreakDmgLimit(cfg, abilityType)),
		boolQuery2(r, i, ids, QpnDarkable, getTypedAbilityIDsDarkable(cfg, abilityType)),
		boolQuery2(r, i, ids, QpnSilenceable, getTypedAbilityIDsSilenceable(cfg, abilityType)),
		boolQuery2(r, i, ids, QpnReflectable, getTypedAbilityIDsReflectable(cfg, abilityType)),
		boolQuery2(r, i, ids, QpnDelay, getTypedAbilityIDsDealsDelay(cfg, abilityType)),
	})
}
