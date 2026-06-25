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

	monsters, err := getResourcesDbItem(cfg, r, cfg.e.monsters, ability, cfg.db.GetEnemyAbilityMonsterIDs)
	if err != nil {
		return EnemyAbility{}, err
	}

	response := EnemyAbility{
		ID:                 ability.ID,
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		UntypedAbility:     idToTypedAPIResource(cfg, cfg.e.abilities, ability.Ability.ID),
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

	return filterIDs(cfg, r, i, ids, []filteredIdList{
		fidl(idQuery(r, i, ids, qpnMonster, cfg.l.Monsters, cfg.db.GetEnemyAbilityIDsByMonster)),
		fidl(enumListQuery(cfg, r, i, cfg.t.DamageType, ids, qpnDamageType, getTypedAbilityIDsByDamageType(cfg, abilityType))),
		fidl(enumListQuery(cfg, r, i, cfg.t.AttackType, ids, qpnAttackType, getTypedAbilityIDsByAttackType(cfg, abilityType))),
		fidl(enumListQuery(cfg, r, i, cfg.t.TargetType, ids, qpnTargetType, getTypedAbilityIDsByTargetType(cfg, abilityType))),
		fidl(enumQuery(r, i, cfg.t.DamageFormula, ids, qpnDamageFormula, getTypedAbilityIDsByDamageFormula(cfg, abilityType))),
		fidl(intListQuery(cfg, r, i, ids, qpnRank, getTypedAbilityIDsByRank(cfg, abilityType))),
		fidl(nameIdListQueryNul(cfg, r, i, ids, qpnElement, cfg.e.elements.resTypeSing, cfg.l.Elements, getTypedAbilityIDsByElement(cfg, abilityType))),
		fidl(idQueryNul(r, i, ids, qpnStatusInflict, cfg.l.StatusConditions, getTypedAbilityIDsByInflictedStatus(cfg, abilityType))),
		fidl(idQueryNul(r, i, ids, qpnStatusRemove, cfg.l.StatusConditions, getTypedAbilityIDsByRemovedStatus(cfg, abilityType))),
		fidl(boolQuery(r, i, ids, qpnHelpBar, getTypedAbilityIDsByAppearsInHelpBar(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, qpnCanCrit, getTypedAbilityIDsCanCrit(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, qpnBDL, getTypedAbilityIDsBreakDmgLimit(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, qpnDarkable, getTypedAbilityIDsDarkable(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, qpnSilenceable, getTypedAbilityIDsSilenceable(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, qpnReflectable, getTypedAbilityIDsReflectable(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, qpnDelay, getTypedAbilityIDsDealsDelay(cfg, abilityType))),
	})
}
