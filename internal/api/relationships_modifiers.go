package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getModifierRelationships(cfg *Config, r *http.Request, modifier seeding.Modifier) (Modifier, error) {
	autoAbilities, err := getResourcesDbItem(cfg, r, cfg.e.autoAbilities, modifier, cfg.db.GetModifierAutoAbilityIDs)
	if err != nil {
		return Modifier{}, err
	}

	playerAbilities, err := getResourcesDbItem(cfg, r, cfg.e.playerAbilities, modifier, cfg.db.GetModifierPlayerAbilityIDs)
	if err != nil {
		return Modifier{}, err
	}

	overdriveAbilities, err := getResourcesDbItem(cfg, r, cfg.e.overdriveAbilities, modifier, cfg.db.GetModifierOverdriveAbilityIDs)
	if err != nil {
		return Modifier{}, err
	}

	itemAbilities, err := getResourcesDbItem(cfg, r, cfg.e.itemAbilities, modifier, cfg.db.GetModifierItemAbilityIDs)
	if err != nil {
		return Modifier{}, err
	}

	triggerCommands, err := getResourcesDbItem(cfg, r, cfg.e.triggerCommands, modifier, cfg.db.GetModifierTriggerCommandIDs)
	if err != nil {
		return Modifier{}, err
	}

	enemyAbilities, err := getResourcesDbItem(cfg, r, cfg.e.enemyAbilities, modifier, cfg.db.GetModifierEnemyAbilityIDs)
	if err != nil {
		return Modifier{}, err
	}

	statusConditions, err := getResourcesDbItem(cfg, r, cfg.e.statusConditions, modifier, cfg.db.GetModifierStatusConditionIDs)
	if err != nil {
		return Modifier{}, err
	}

	properties, err := getResourcesDbItem(cfg, r, cfg.e.properties, modifier, cfg.db.GetModifierPropertyIDs)
	if err != nil {
		return Modifier{}, err
	}

	rel := Modifier{
		AutoAbilities: 		autoAbilities,
		PlayerAbilities: 	playerAbilities,
		OverdriveAbilities: overdriveAbilities,
		ItemAbilities: 		itemAbilities,
		TriggerCommands: 	triggerCommands,
		EnemyAbilities: 	enemyAbilities,
		StatusConditions: 	statusConditions,
		Properties: 		properties,
	}

	return rel, nil
}