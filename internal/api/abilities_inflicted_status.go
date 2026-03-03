package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getAbilitiesInflictedStatus(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	i := cfg.e.abilities
	status, _ := seeding.GetResourceByID(id, cfg.l.StatusConditionsID)

	if status.Name == "delay" {
		return getResourcesDbNoInput(cfg, r, i, cfg.e.statusConditions.resourceType, cfg.db.GetAbilityIDsDealsDelay)
	}

	return getResourcesDB(cfg, r, i, status, cfg.db.GetAbilityIDsByInflictedStatus)
}

func getPlayerAbilitiesInflictedStatus(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	i := cfg.e.playerAbilities
	status, _ := seeding.GetResourceByID(id, cfg.l.StatusConditionsID)

	if status.Name == "delay" {
		return getResourcesDbNoInput(cfg, r, i, cfg.e.statusConditions.resourceType, cfg.db.GetPlayerAbilityIDsDealsDelay)
	}

	return getResourcesDB(cfg, r, i, status, cfg.db.GetPlayerAbilityIDsByInflictedStatus)
}

func getEnemyAbilitiesInflictedStatus(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	i := cfg.e.enemyAbilities
	status, _ := seeding.GetResourceByID(id, cfg.l.StatusConditionsID)

	if status.Name == "delay" {
		return getResourcesDbNoInput(cfg, r, i, cfg.e.statusConditions.resourceType, cfg.db.GetEnemyAbilityIDsDealsDelay)
	}

	return getResourcesDB(cfg, r, i, status, cfg.db.GetEnemyAbilityIDsByInflictedStatus)
}

func getItemAbilitiesInflictedStatus(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	i := cfg.e.enemyAbilities
	status, _ := seeding.GetResourceByID(id, cfg.l.StatusConditionsID)

	if status.Name == "delay" {
		return getResourcesDbNoInput(cfg, r, i, cfg.e.statusConditions.resourceType, cfg.db.GetItemAbilityIDsDealsDelay)
	}

	return getResourcesDB(cfg, r, i, status, cfg.db.GetItemAbilityIDsByInflictedStatus)
}

func getUnspecifiedAbilitiesInflictedStatus(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	i := cfg.e.unspecifiedAbilities
	status, _ := seeding.GetResourceByID(id, cfg.l.StatusConditionsID)

	if status.Name == "delay" {
		return getResourcesDbNoInput(cfg, r, i, cfg.e.statusConditions.resourceType, cfg.db.GetUnspecifiedAbilityIDsDealsDelay)
	}

	return getResourcesDB(cfg, r, i, status, cfg.db.GetUnspecifiedAbilityIDsByInflictedStatus)
}

func getOverdriveAbilitiesInflictedStatus(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	i := cfg.e.overdriveAbilities
	status, _ := seeding.GetResourceByID(id, cfg.l.StatusConditionsID)

	if status.Name == "delay" {
		return getResourcesDbNoInput(cfg, r, i, cfg.e.statusConditions.resourceType, cfg.db.GetOverdriveAbilityIDsDealsDelay)
	}

	return getResourcesDB(cfg, r, i, status, cfg.db.GetOverdriveAbilityIDsByInflictedStatus)
}
