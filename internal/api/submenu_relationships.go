package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func createSubmenuOpenedBy(cfg *Config, r *http.Request, submenu seeding.Submenu) (*MenuOpen, error) {
	ability, err := getAbilityResPtrDB(cfg, r, submenu, queryOne(cfg.db.GetSubmenuOpenedByAbilityID))
	if err != nil {
		return nil, err
	}

	aeonCommand, err := getResPtrDB(cfg, r, cfg.e.aeonCommands, submenu, queryOne(cfg.db.GetSubmenuOpenedByAeonCommandID))
	if err != nil {
		return nil, err
	}

	overdriveCommands, err := getResourcesDbItem(cfg, r, cfg.e.overdriveCommands, submenu, cfg.db.GetSubmenuOpenedByOverdriveCommandIDs)
	if err != nil {
		return nil, err
	}

	menuOpen := MenuOpen{
		Ability:           ability,
		AeonCommand:       aeonCommand,
		OverdriveCommands: sliceOrNil(overdriveCommands),
	}

	if menuOpen.IsZero() {
		return nil, nil
	}

	return &menuOpen, nil
}
