package api

import (
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func createSubmenuOpenedBy(cfg *Config, r *http.Request, submenu seeding.Submenu) (*MenuOpen, error) {
	ability, err := getResPtrDB(cfg, r, cfg.e.abilities, submenu, ToIntOneNull(cfg.db.GetSubmenuOpenedByAbilityID))
	if err != nil {
		return nil, err
	}

	aeonCommand, err := getResPtrDB(cfg, r, cfg.e.aeonCommands, submenu, ToIntOneNull(cfg.db.GetSubmenuOpenedByAeonCommandID))
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
		OverdriveCommands: h.SliceOrNil(overdriveCommands),
	}

	if menuOpen.IsZero() {
		return nil, nil
	}

	return &menuOpen, nil
}
