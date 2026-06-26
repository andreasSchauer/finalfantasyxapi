package api

import (
	"context"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
	"golang.org/x/sync/errgroup"
)


func getSubmenuRelationships(cfg *Config, r *http.Request, submenu seeding.Submenu) (Submenu, error) {
	var rel Submenu
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.Abilities, err = getResourcesDbItem(cfg, ctx, cfg.e.abilities, submenu, ToIntManyNull(cfg.db.GetSubmenuAbilityIDs))
		return err
	})
	
	g.Go(func() error {
		var err error
		rel.OpenedBy, err = createSubmenuOpenedBy(cfg, ctx, submenu)
		return err
	})

	err := g.Wait()
	if err != nil {
		return Submenu{}, err
	}

	return rel, nil
}


func createSubmenuOpenedBy(cfg *Config, ctx context.Context, submenu seeding.Submenu) (*MenuOpen, error) {
	var menuOpen MenuOpen
	g, ctx := errgroup.WithContext(ctx)
	
	g.Go(func() error {
		var err error
		menuOpen.Ability, err = getResPtrDB(cfg, ctx, cfg.e.abilities, submenu, ToIntOneNull(cfg.db.GetSubmenuOpenedByAbilityID))
		return err
	})
	
	g.Go(func() error {
		var err error
		menuOpen.AeonCommand, err = getResPtrDB(cfg, ctx, cfg.e.aeonCommands, submenu, ToIntOneNull(cfg.db.GetSubmenuOpenedByAeonCommandID))
		return err
	})
	
	g.Go(func() error {
		overdriveCommands, err := getResourcesDbItem(cfg, ctx, cfg.e.overdriveCommands, submenu, ToIntManyNull(cfg.db.GetSubmenuOpenedByOverdriveCommandIDs))
		if err != nil {
			return err
		}
		menuOpen.OverdriveCommands = h.SliceOrNil(overdriveCommands)
		return nil
	})

	err := g.Wait()
	if err != nil {
		return nil, err
	}

	if menuOpen.IsZero() {
		return nil, nil
	}

	return &menuOpen, nil
}
