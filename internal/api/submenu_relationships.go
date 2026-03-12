package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func createSubmenuOpenedBy(cfg *Config, r *http.Request, submenu seeding.Submenu) (*MenuOpen, error) {
	ability, err := getAbilityResPtrDB(cfg, r, submenu, cfg.db.GetSubmenuOpenedByAbilityID)
	if err != nil {
		return nil, err
	}

	aeonCommand, err := getResPtrDbNullable(cfg, r, cfg.e.aeonCommands, submenu, cfg.db.GetSubmenuOpenedByAeonCommandID)
	if err != nil {
		return nil, err
	}

	overdriveCommands, err := getResourcesDB(cfg, r, cfg.e.overdriveCommands, submenu, cfg.db.GetSubmenuOpenedByOverdriveCommandIDs)
	if err != nil {
		return nil, err
	}

	menuOpen := MenuOpen{
			Ability: 			ability,
			AeonCommand: 		aeonCommand,
			OverdriveCommands: 	overdriveCommands,
	}

	if menuOpen.IsZero() {
		return nil, nil
	}

	return &menuOpen, nil
}