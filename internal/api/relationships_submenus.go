package api

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getSubmenuRelationships(r *http.Request, submenu seedtypes.Submenu) (Submenu, error) {
	var rel Submenu
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.Abilities, err = E.Abilities.dbGetRelations(ctx, submenu, ToIntManyNull(config.Cfg.DB.GetSubmenuAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.OpenedBy, err = createSubmenuOpenedBy(ctx, submenu)
		return err
	})

	err := g.Wait()
	if err != nil {
		return Submenu{}, err
	}

	return rel, nil
}

func createSubmenuOpenedBy(ctx context.Context, submenu seedtypes.Submenu) (*MenuOpen, error) {
	var menuOpen MenuOpen
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		var err error
		menuOpen.Ability, err = E.Abilities.dbGetRelationPtr(ctx, submenu, ToIntOneNull(config.Cfg.DB.GetSubmenuOpenedByAbilityID))
		return err
	})

	g.Go(func() error {
		var err error
		menuOpen.AeonCommand, err = E.AeonCommands.dbGetRelationPtr(ctx, submenu, ToIntOneNull(config.Cfg.DB.GetSubmenuOpenedByAeonCommandID))
		return err
	})

	g.Go(func() error {
		overdriveCommands, err := E.OverdriveCommands.dbGetRelations(ctx, submenu, ToIntManyNull(config.Cfg.DB.GetSubmenuOpenedByOverdriveCommandIDs))
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
