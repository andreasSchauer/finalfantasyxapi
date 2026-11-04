package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getTopmenuRelationships(server *Server, r *http.Request, topmenu seedtypes.Topmenu) (Topmenu, error) {
	var rel Topmenu
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.Submenus, err = E.Submenus.dbGetRelations(server, ctx, topmenu, ToIntManyNull(server.Cfg.DB.GetTopmenuSubmenuIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.OverdriveCommands, err = E.OverdriveCommands.dbGetRelations(server, ctx, topmenu, ToIntManyNull(server.Cfg.DB.GetTopmenuOverdriveCommandIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Overdrives, err = E.Overdrives.dbGetRelations(server, ctx, topmenu, ToIntManyNull(server.Cfg.DB.GetTopmenuOverdriveIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.AeonCommands, err = E.AeonCommands.dbGetRelations(server, ctx, topmenu, ToIntManyNull(server.Cfg.DB.GetTopmenuAeonCommandIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Abilities, err = E.Abilities.dbGetRelations(server, ctx, topmenu, ToIntManyNull(server.Cfg.DB.GetTopmenuAbilityIDs))
		return err
	})

	err := g.Wait()
	if err != nil {
		return Topmenu{}, err
	}

	return rel, nil
}
