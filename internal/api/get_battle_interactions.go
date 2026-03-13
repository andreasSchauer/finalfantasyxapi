package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

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