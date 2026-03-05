package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getAbility(r *http.Request, i handlerInput[seeding.Ability, Ability, NamedAPIResource, NamedApiResourceList], id int32) (Ability, error) {
	ability, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return Ability{}, err
	}

	abilityType, _ := newNamedAPIResourceFromType(cfg, cfg.e.abilityType.endpoint, string(ability.Type), cfg.t.AbilityType)

	monsters, err := getResourcesDB(cfg, r, cfg.e.monsters, ability, cfg.db.GetAbilityMonsterIDs)
	if err != nil {
		return Ability{}, err
	}

	response := Ability{
		ID:                 ability.ID,
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		Type:               abilityType,
		TypedAbility:       createAbilityResource(cfg, ability.Name, ability.Version, ability.Type),
		Monsters:           monsters,
		BattleInteractions: getAbilityBattleInteractions(cfg, ability),
	}

	switch ability.Type {
	case database.AbilityTypeOverdriveAbility:
		attributes, err := cfg.db.GetOverdriveAbilityAttributes(r.Context(), id)
		if err != nil {
			return Ability{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get attributes for %s", ability), err)
		}

		response.Rank = h.NullInt32ToPtr(attributes.Rank)
		response.CanCopycat = attributes.CanCopycat
		response.AppearsInHelpBar = attributes.AppearsInHelpBar

	default:
		response.Rank = ability.Rank
		response.CanCopycat = ability.CanCopycat
		response.AppearsInHelpBar = ability.AppearsInHelpBar
	}

	return response, nil
}

func getAbilityBattleInteractions(cfg *Config, ability seeding.Ability) []BattleInteraction {
	seedBattleInteractions := lookupAbilityBattleInteractions(cfg, ability)
	return convertObjSlice(cfg, seedBattleInteractions, convertBattleInteraction)
}

func getAbilityBattleInteractionsSimple(cfg *Config, ability seeding.Ability) []BattleInteractionSimple {
	seedBattleInteractions := lookupAbilityBattleInteractions(cfg, ability)
	return convertObjSlice(cfg, seedBattleInteractions, convertBattleInteractionSimple)
}

func lookupAbilityBattleInteractions(cfg *Config, ability seeding.Ability) []seeding.BattleInteraction {
	var seedBattleInteractions []seeding.BattleInteraction
	abilityRef := ability.GetAbilityRef().Untyped()

	switch ability.Type {
	case database.AbilityTypeEnemyAbility:
		lookup, _ := seeding.GetResource(abilityRef, cfg.l.EnemyAbilities)
		seedBattleInteractions = lookup.BattleInteractions

	case database.AbilityTypeItemAbility:
		lookup, _ := seeding.GetResource(ability.Name, cfg.l.Items)
		seedBattleInteractions = lookup.BattleInteractions

	case database.AbilityTypeUnspecifiedAbility:
		lookup, _ := seeding.GetResource(abilityRef, cfg.l.UnspecifiedAbilities)
		seedBattleInteractions = lookup.BattleInteractions

	case database.AbilityTypeOverdriveAbility:
		lookup, _ := seeding.GetResource(abilityRef, cfg.l.OverdriveAbilities)
		seedBattleInteractions = lookup.BattleInteractions

	case database.AbilityTypePlayerAbility:
		lookup, _ := seeding.GetResource(abilityRef, cfg.l.PlayerAbilities)
		seedBattleInteractions = lookup.BattleInteractions

	case database.AbilityTypeTriggerCommand:
		lookup, _ := seeding.GetResource(abilityRef, cfg.l.TriggerCommands)
		seedBattleInteractions = lookup.BattleInteractions
	}

	return seedBattleInteractions
}

func (cfg *Config) retrieveAbilities(r *http.Request, i handlerInput[seeding.Ability, Ability, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[NamedAPIResource]{
		frl(typeQuery(cfg, r, i, cfg.t.AbilityType, resources, "type", cfg.db.GetAbilityIDsByType)),
		frl(typeQuery(cfg, r, i, cfg.t.DamageType, resources, "damage_type", cfg.db.GetAbilityIDsByDamageType)),
		frl(typeQuery(cfg, r, i, cfg.t.AttackType, resources, "attack_type", cfg.db.GetAbilityIDsByAttackType)),
		frl(typeQuery(cfg, r, i, cfg.t.TargetType, resources, "target_type", cfg.db.GetAbilityIDsByTargetType)),
		frl(typeQuery(cfg, r, i, cfg.t.DamageFormula, resources, "damage_formula", cfg.db.GetAbilityIDsByDamageFormula)),
		frl(intQueryNullable(cfg, r, i, resources, "rank", cfg.db.GetAbilityIDsByRank)),
		frl(nameOrIdQuery(cfg, r, i, resources, "element", cfg.e.elements.resourceType, cfg.l.Elements, cfg.db.GetAbilityIDsByElement)),
		frl(idQueryWrapper(cfg, r, i, resources, "status_inflict", len(cfg.l.StatusConditions), getAbilitiesInflictedStatus)),
		frl(idQuery(cfg, r, i, resources, "status_remove", len(cfg.l.StatusConditions), cfg.db.GetAbilityIDsByRemovedStatus)),
		frl(boolQuery(cfg, r, i, resources, "copycat", cfg.db.GetAbilityIDsByCanCopycat)),
		frl(boolQuery(cfg, r, i, resources, "help_bar", cfg.db.GetAbilityIDsByAppearsInHelpBar)),
		frl(boolQuery2(cfg, r, i, resources, "phys_atk", cfg.db.GetAbilityIDsBasedOnPhysAttack)),
		frl(boolQuery2(cfg, r, i, resources, "darkable", cfg.db.GetAbilityIDsDarkable)),
		frl(boolQuery2(cfg, r, i, resources, "silenceable", cfg.db.GetAbilityIDsSilenceable)),
		frl(boolQuery2(cfg, r, i, resources, "reflectable", cfg.db.GetAbilityIDsReflectable)),
		frl(boolQuery2(cfg, r, i, resources, "delay", cfg.db.GetAbilityIDsDealsDelay)),
		frl(boolQuery2(cfg, r, i, resources, "stat_changes", cfg.db.GetAbilityIDsWithStatChanges)),
		frl(boolQuery2(cfg, r, i, resources, "mod_changes", cfg.db.GetAbilityIDsWithModifierChanges)),
	})
}
