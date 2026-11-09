package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getTopmenuRelationships(r *http.Request, topmenu seedtypes.Topmenu) (apitypes.Topmenu, error) {
	var rel apitypes.Topmenu
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.Submenus, err = apilookups.E.Submenus.DBGetRelations(ctx, topmenu, handlerinput.ToIntManyNull(config.Cfg.DB.GetTopmenuSubmenuIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.OverdriveCommands, err = apilookups.E.OverdriveCommands.DBGetRelations(ctx, topmenu, handlerinput.ToIntManyNull(config.Cfg.DB.GetTopmenuOverdriveCommandIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Overdrives, err = apilookups.E.Overdrives.DBGetRelations(ctx, topmenu, handlerinput.ToIntManyNull(config.Cfg.DB.GetTopmenuOverdriveIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.AeonCommands, err = apilookups.E.AeonCommands.DBGetRelations(ctx, topmenu, handlerinput.ToIntManyNull(config.Cfg.DB.GetTopmenuAeonCommandIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.Abilities, err = apilookups.E.Abilities.DBGetRelations(ctx, topmenu, handlerinput.ToIntManyNull(config.Cfg.DB.GetTopmenuAbilityIDs))
		return err
	})

	err := g.Wait()
	if err != nil {
		return apitypes.Topmenu{}, err
	}

	return rel, nil
}
