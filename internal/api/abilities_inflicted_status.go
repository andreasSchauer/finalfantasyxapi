package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getAbilitiesInflictedStatus(cfg *Config, r *http.Request, id int32) ([]AbilityAPIResource, error) {
	i := cfg.e.abilities
	status, _ := seeding.GetResourceByID(id, cfg.l.StatusConditionsID)

	if status.Name == "delay" {
		return getResourcesDbID(cfg, r, i, 0, cfg.e.statusConditions.resourceType, queryNoInput(cfg.db.GetAbilityIDsDealsDelay))
	}

	return getResourcesDbItem(cfg, r, i, status, cfg.db.GetAbilityIDsByInflictedStatus)
}

func getPlayerAbilitiesInflictedStatus(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	i := cfg.e.playerAbilities
	status, _ := seeding.GetResourceByID(id, cfg.l.StatusConditionsID)

	if status.Name == "delay" {
		return getResourcesDbID(cfg, r, i, 0, cfg.e.statusConditions.resourceType, queryNoInput(cfg.db.GetPlayerAbilityIDsDealsDelay))
	}

	return getResourcesDbItem(cfg, r, i, status, cfg.db.GetPlayerAbilityIDsByInflictedStatus)
}

func getEnemyAbilitiesInflictedStatus(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	i := cfg.e.enemyAbilities
	status, _ := seeding.GetResourceByID(id, cfg.l.StatusConditionsID)

	if status.Name == "delay" {
		return getResourcesDbID(cfg, r, i, 0, cfg.e.statusConditions.resourceType, queryNoInput(cfg.db.GetEnemyAbilityIDsDealsDelay))
	}

	return getResourcesDbItem(cfg, r, i, status, cfg.db.GetEnemyAbilityIDsByInflictedStatus)
}

func getItemAbilitiesInflictedStatus(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	i := cfg.e.enemyAbilities
	status, _ := seeding.GetResourceByID(id, cfg.l.StatusConditionsID)

	if status.Name == "delay" {
		return getResourcesDbID(cfg, r, i, 0, cfg.e.statusConditions.resourceType, queryNoInput(cfg.db.GetItemAbilityIDsDealsDelay))
	}

	return getResourcesDbItem(cfg, r, i, status, cfg.db.GetItemAbilityIDsByInflictedStatus)
}

func getOverdriveAbilitiesInflictedStatus(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	i := cfg.e.overdriveAbilities
	status, _ := seeding.GetResourceByID(id, cfg.l.StatusConditionsID)

	if status.Name == "delay" {
		return getResourcesDbID(cfg, r, i, 0, cfg.e.statusConditions.resourceType, queryNoInput(cfg.db.GetOverdriveAbilityIDsDealsDelay))
	}

	return getResourcesDbItem(cfg, r, i, status, cfg.db.GetOverdriveAbilityIDsByInflictedStatus)
}
