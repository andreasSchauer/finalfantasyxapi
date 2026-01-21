package main

import (
	"fmt"
	"net/http"
	
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

// can be used for various other functions related to abilities
// put into own file
func createAbilityResource(cfg *Config, name string, version *int32, abilityType database.AbilityType) (NamedAPIResource, error) {
	var res NamedAPIResource
	var err error

	switch abilityType {
	case database.AbilityTypePlayerAbility:
		res = nameToNamedAPIResource(cfg, cfg.e.playerAbilities, name, version)

	case database.AbilityTypeEnemyAbility:
		res = nameToNamedAPIResource(cfg, cfg.e.enemyAbilities, name, version)

	case database.AbilityTypeOverdriveAbility:
		res = nameToNamedAPIResource(cfg, cfg.e.overdriveAbilities, name, version)

	case database.AbilityTypeItemAbility:
		res = nameToNamedAPIResource(cfg, cfg.e.itemAbilities, name, version)

	case database.AbilityTypeTriggerCommand:
		res = nameToNamedAPIResource(cfg, cfg.e.triggerCommands, name, version)

	default:
		err = newHTTPError(http.StatusInternalServerError, fmt.Sprintf("ability of type '%s' does not exist.", abilityType), nil)
	}

	if err != nil {
		return NamedAPIResource{}, err
	}

	return res, nil
}