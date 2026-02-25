package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)



func getPlayerAbilitiesInflictedStatus(cfg *Config, r *http.Request, id int32) ([]NamedAPIResource, error) {
	i := cfg.e.playerAbilities
	status, _ := seeding.GetResourceByID(id, cfg.l.StatusConditionsID)

	if status.Name == "delay" {
		return getResourcesDbNoInput(cfg, r, i, cfg.e.statusConditions.resourceType, cfg.db.GetPlayerAbilityIDsDealsDelay)
	}

	return getResourcesDB(cfg, r, i, status, cfg.db.GetPlayerAbilityIDsByInflictedStatus)
}