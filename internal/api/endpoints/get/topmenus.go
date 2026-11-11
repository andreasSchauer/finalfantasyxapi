package get

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
	"golang.org/x/sync/errgroup"
)

func Topmenu(r *http.Request, i handlerinput.Main[seedtypes.Topmenu, apitypes.Topmenu, apiresource.Named], id int32) (apitypes.Topmenu, error) {
	topmenu, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Topmenu{}, err
	}

	rel, err := getTopmenuRelationships(r, topmenu)
	if err != nil {
		return apitypes.Topmenu{}, err
	}

	response := apitypes.Topmenu{
		ID:                topmenu.ID,
		Name:              topmenu.Name,
		Submenus:          rel.Submenus,
		OverdriveCommands: rel.OverdriveCommands,
		Overdrives:        rel.Overdrives,
		AeonCommands:      rel.AeonCommands,
		Abilities:         rel.Abilities,
	}

	return response, nil
}

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
