package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getEnemyAbility(r *http.Request, i handlerInput[seeding.EnemyAbility, EnemyAbility, NamedAPIResource, NamedApiResourceList], id int32) (EnemyAbility, error) {
	ability, err := verifyParamsAndGet(cfg, r, i, id)
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

func (cfg *Config) retrieveEnemyAbilities(r *http.Request, i handlerInput[seeding.EnemyAbility, EnemyAbility, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	abilityType := database.AbilityTypeEnemyAbility
	ids, err := verifyParamsAndRetrieve(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterIDs(cfg, r, i, ids, []filteredIdList{
		fidl(idQuery(r, i, ids, "monster", cfg.l.Monsters, cfg.db.GetEnemyAbilityIDsByMonster)),
		fidl(enumListQuery(cfg, r, i, cfg.t.DamageType, ids, "damage_type", getTypedAbilityIDsByDamageType(cfg, abilityType))),
		fidl(enumListQuery(cfg, r, i, cfg.t.AttackType, ids, "attack_type", getTypedAbilityIDsByAttackType(cfg, abilityType))),
		fidl(enumListQuery(cfg, r, i, cfg.t.TargetType, ids, "target_type", getTypedAbilityIDsByTargetType(cfg, abilityType))),
		fidl(enumQuery(r, i, cfg.t.DamageFormula, ids, "damage_formula", getTypedAbilityIDsByDamageFormula(cfg, abilityType))),
		fidl(intListQuery(cfg, r, i, ids, "rank", getTypedAbilityIDsByRank(cfg, abilityType))),
		fidl(nameIdListQueryNul(cfg, r, i, ids, "element", cfg.e.elements.resourceType, cfg.l.Elements, getTypedAbilityIDsByElement(cfg, abilityType))),
		fidl(idQueryNul(r, i, ids, "status_inflict", cfg.l.StatusConditions, getTypedAbilityIDsByInflictedStatus(cfg, abilityType))),
		fidl(idQueryNul(r, i, ids, "status_remove", cfg.l.StatusConditions, getTypedAbilityIDsByRemovedStatus(cfg, abilityType))),
		fidl(boolQuery(r, i, ids, "help_bar", getTypedAbilityIDsByAppearsInHelpBar(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, "can_crit", getTypedAbilityIDsCanCrit(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, "bdl", getTypedAbilityIDsBreakDmgLimit(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, "darkable", getTypedAbilityIDsDarkable(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, "silenceable", getTypedAbilityIDsSilenceable(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, "reflectable", getTypedAbilityIDsReflectable(cfg, abilityType))),
		fidl(boolQuery2(r, i, ids, "delay", getTypedAbilityIDsDealsDelay(cfg, abilityType))),
	})
}
