package api

import (
	"net/http"

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
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[NamedAPIResource]{
		frl(idQuery(cfg, r, i, resources, "monster", len(cfg.l.Monsters), cfg.db.GetEnemyAbilityIDsByMonster)),
		frl(enumListQuery(cfg, r, i, cfg.t.DamageType, resources, "damage_type", cfg.db.GetEnemyAbilityIDsByDamageType)),
		frl(enumListQuery(cfg, r, i, cfg.t.AttackType, resources, "attack_type", cfg.db.GetEnemyAbilityIDsByAttackType)),
		frl(enumListQuery(cfg, r, i, cfg.t.TargetType, resources, "target_type", cfg.db.GetEnemyAbilityIDsByTargetType)),
		frl(enumQuery(cfg, r, i, cfg.t.DamageFormula, resources, "damage_formula", cfg.db.GetEnemyAbilityIDsByDamageFormula)),
		frl(intListQuery(cfg, r, i, resources, "rank", cfg.db.GetEnemyAbilityIDsByRank)),
		frl(nameIdListQueryNul(cfg, r, i, resources, "element", cfg.e.elements.resourceType, cfg.l.Elements, cfg.db.GetEnemyAbilityIDsByElement)),
		frl(idQueryNul(cfg, r, i, resources, "status_inflict", len(cfg.l.StatusConditions), cfg.db.GetEnemyAbilityIDsByInflictedStatus)),
		frl(idQueryNul(cfg, r, i, resources, "status_remove", len(cfg.l.StatusConditions), cfg.db.GetEnemyAbilityIDsByRemovedStatus)),
		frl(boolQuery(cfg, r, i, resources, "help_bar", cfg.db.GetEnemyAbilityIDsByAppearsInHelpBar)),
		frl(boolQuery2(cfg, r, i, resources, "can_crit", cfg.db.GetEnemyAbilityIDsCanCrit)),
		frl(boolQuery2(cfg, r, i, resources, "bdl", cfg.db.GetEnemyAbilityIDsBreakDmgLimit)),
		frl(boolQuery2(cfg, r, i, resources, "darkable", cfg.db.GetEnemyAbilityIDsDarkable)),
		frl(boolQuery2(cfg, r, i, resources, "silenceable", cfg.db.GetEnemyAbilityIDsSilenceable)),
		frl(boolQuery2(cfg, r, i, resources, "reflectable", cfg.db.GetEnemyAbilityIDsReflectable)),
		frl(boolQuery2(cfg, r, i, resources, "delay", cfg.db.GetEnemyAbilityIDsDealsDelay)),
	})
}
