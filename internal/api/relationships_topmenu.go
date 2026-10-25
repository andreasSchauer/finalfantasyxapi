package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
	"golang.org/x/sync/errgroup"
)

func getTopmenuRelationships(cfg *Config, r *http.Request, topmenu seeding.Topmenu) (Topmenu, error) {
	var rel Topmenu
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.Submenus, err = getResourcesDbItem(cfg, ctx, cfg.E.Submenus, topmenu, ToIntManyNull(cfg.db.GetTopmenuSubmenuIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.OverdriveCommands, err = getResourcesDbItem(cfg, ctx, cfg.E.OverdriveCommands, topmenu, ToIntManyNull(cfg.db.GetTopmenuOverdriveCommandIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Overdrives, err = getResourcesDbItem(cfg, ctx, cfg.E.Overdrives, topmenu, ToIntManyNull(cfg.db.GetTopmenuOverdriveIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.AeonCommands, err = getResourcesDbItem(cfg, ctx, cfg.E.AeonCommands, topmenu, ToIntManyNull(cfg.db.GetTopmenuAeonCommandIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Abilities, err = getResourcesDbItem(cfg, ctx, cfg.E.Abilities, topmenu, ToIntManyNull(cfg.db.GetTopmenuAbilityIDs))
		return err
	})

	err := g.Wait()
	if err != nil {
		return Topmenu{}, err
	}

	return rel, nil
}
