package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getElementRelationships(cfg *Config, r *http.Request, element seeding.Element) (Element, error) {
	statusProtection, err := getResPtrDB(cfg, r, cfg.e.statusConditions, element, cfg.db.GetElementStatusConditionID)
	if err != nil {
		return Element{}, err
	}

	autoAbilities, err := getResourcesDbItem(cfg, r, cfg.e.autoAbilities, element, cfg.db.GetElementAutoAbilityIDs)
	if err != nil {
		return Element{}, err
	}

	playerAbilities, err := getResourcesDbItem(cfg, r, cfg.e.playerAbilities, element, cfg.db.GetElementPlayerAbilityIDs)
	if err != nil {
		return Element{}, err
	}

	overdriveAbilities, err := getResourcesDbItem(cfg, r, cfg.e.overdriveAbilities, element, cfg.db.GetElementOverdriveAbilityIDs)
	if err != nil {
		return Element{}, err
	}

	itemAbilities, err := getResourcesDbItem(cfg, r, cfg.e.itemAbilities, element, cfg.db.GetElementItemAbilityIDs)
	if err != nil {
		return Element{}, err
	}

	enemyAbilities, err := getResourcesDbItem(cfg, r, cfg.e.enemyAbilities, element, cfg.db.GetElementEnemyAbilityIDs)
	if err != nil {
		return Element{}, err
	}

	monstersWeak, err := getResourcesDbItem(cfg, r, cfg.e.monsters, element, cfg.db.GetElementMonsterIDsWeak)
	if err != nil {
		return Element{}, err
	}

	monstersHalved, err := getResourcesDbItem(cfg, r, cfg.e.monsters, element, cfg.db.GetElementMonsterIDsHalved)
	if err != nil {
		return Element{}, err
	}

	monstersImmune, err := getResourcesDbItem(cfg, r, cfg.e.monsters, element, cfg.db.GetElementMonsterIDsImmune)
	if err != nil {
		return Element{}, err
	}

	monstersAbsorb, err := getResourcesDbItem(cfg, r, cfg.e.monsters, element, cfg.db.GetElementMonsterIDsAbsorb)
	if err != nil {
		return Element{}, err
	}

	rel := Element{
		StatusProtection: 	statusProtection,
		AutoAbilities: 		autoAbilities,
		PlayerAbilities: 	playerAbilities,
		OverdriveAbilities: overdriveAbilities,
		ItemAbilities: 		itemAbilities,
		EnemyAbilities: 	enemyAbilities,
		MonstersWeak: 		monstersWeak,
		MonstersHalved: 	monstersHalved,
		MonstersImmune: 	monstersImmune,
		MonstersAbsorb: 	monstersAbsorb,
	}

	return rel, nil
}