package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getTopmenuRelationships(srv *Server, r *http.Request, topmenu seedtypes.Topmenu) (Topmenu, error) {
	var rel Topmenu
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.Submenus, err = E.Submenus.dbGetRelations(srv, ctx, topmenu, ToIntManyNull(config.Cfg.DB.GetTopmenuSubmenuIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.OverdriveCommands, err = E.OverdriveCommands.dbGetRelations(srv, ctx, topmenu, ToIntManyNull(config.Cfg.DB.GetTopmenuOverdriveCommandIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Overdrives, err = E.Overdrives.dbGetRelations(srv, ctx, topmenu, ToIntManyNull(config.Cfg.DB.GetTopmenuOverdriveIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.AeonCommands, err = E.AeonCommands.dbGetRelations(srv, ctx, topmenu, ToIntManyNull(config.Cfg.DB.GetTopmenuAeonCommandIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Abilities, err = E.Abilities.dbGetRelations(srv, ctx, topmenu, ToIntManyNull(config.Cfg.DB.GetTopmenuAbilityIDs))
		return err
	})

	err := g.Wait()
	if err != nil {
		return Topmenu{}, err
	}

	return rel, nil
}
