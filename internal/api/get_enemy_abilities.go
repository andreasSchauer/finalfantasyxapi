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
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[NamedAPIResource]{
		frl(idQuery(cfg, r, i, resources, "monster", len(cfg.l.Monsters), cfg.db.GetEnemyAbilityIDsByMonster)),
		frl(enumListQuery(cfg, r, i, cfg.t.DamageType, resources, "damage_type", getTypedAbilityIDsByDamageType(cfg, abilityType))),
		frl(enumListQuery(cfg, r, i, cfg.t.AttackType, resources, "attack_type", getTypedAbilityIDsByAttackType(cfg, abilityType))),
		frl(enumListQuery(cfg, r, i, cfg.t.TargetType, resources, "target_type", getTypedAbilityIDsByTargetType(cfg, abilityType))),
		frl(enumQuery(cfg, r, i, cfg.t.DamageFormula, resources, "damage_formula", getTypedAbilityIDsByDamageFormula(cfg, abilityType))),
		frl(intListQuery(cfg, r, i, resources, "rank", getTypedAbilityIDsByRank(cfg, abilityType))),
		frl(nameIdListQueryNul(cfg, r, i, resources, "element", cfg.e.elements.resourceType, cfg.l.Elements, getTypedAbilityIDsByElement(cfg, abilityType))),
		frl(idQueryNul(cfg, r, i, resources, "status_inflict", len(cfg.l.StatusConditions), getTypedAbilityIDsByInflictedStatus(cfg, abilityType))),
		frl(idQueryNul(cfg, r, i, resources, "status_remove", len(cfg.l.StatusConditions), getTypedAbilityIDsByRemovedStatus(cfg, abilityType))),
		frl(boolQuery(cfg, r, i, resources, "help_bar", getTypedAbilityIDsByAppearsInHelpBar(cfg, abilityType))),
		frl(boolQuery2(cfg, r, i, resources, "can_crit", getTypedAbilityIDsCanCrit(cfg, abilityType))),
		frl(boolQuery2(cfg, r, i, resources, "bdl", getTypedAbilityIDsBreakDmgLimit(cfg, abilityType))),
		frl(boolQuery2(cfg, r, i, resources, "darkable", getTypedAbilityIDsDarkable(cfg, abilityType))),
		frl(boolQuery2(cfg, r, i, resources, "silenceable", getTypedAbilityIDsSilenceable(cfg, abilityType))),
		frl(boolQuery2(cfg, r, i, resources, "reflectable", getTypedAbilityIDsReflectable(cfg, abilityType))),
		frl(boolQuery2(cfg, r, i, resources, "delay", getTypedAbilityIDsDealsDelay(cfg, abilityType))),
	})
}
