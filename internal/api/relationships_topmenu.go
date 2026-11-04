package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getTopmenuRelationships(cfg *Config, r *http.Request, topmenu seedtypes.Topmenu) (Topmenu, error) {
	var rel Topmenu
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.Submenus, err = E.Submenus.dbGetRelations(cfg, ctx, topmenu, ToIntManyNull(cfg.db.GetTopmenuSubmenuIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.OverdriveCommands, err = E.OverdriveCommands.dbGetRelations(cfg, ctx, topmenu, ToIntManyNull(cfg.db.GetTopmenuOverdriveCommandIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Overdrives, err = E.Overdrives.dbGetRelations(cfg, ctx, topmenu, ToIntManyNull(cfg.db.GetTopmenuOverdriveIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.AeonCommands, err = E.AeonCommands.dbGetRelations(cfg, ctx, topmenu, ToIntManyNull(cfg.db.GetTopmenuAeonCommandIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Abilities, err = E.Abilities.dbGetRelations(cfg, ctx, topmenu, ToIntManyNull(cfg.db.GetTopmenuAbilityIDs))
		return err
	})

	err := g.Wait()
	if err != nil {
		return Topmenu{}, err
	}

	return rel, nil
}
