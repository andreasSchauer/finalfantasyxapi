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
	
	response := apitypes.Topmenu{
		ID:                topmenu.ID,
		Name:              topmenu.Name,
	}

	err = getTopmenuRelationships(r, topmenu, &response)
	if err != nil {
		return apitypes.Topmenu{}, err
	}

	return response, nil
}

func getTopmenuRelationships(r *http.Request, topmenu seedtypes.Topmenu, response *apitypes.Topmenu) error {
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		response.Submenus, err = apilookups.E.Submenus.DBGetRelations(ctx, topmenu, handlerinput.ToIntManyNull(config.Cfg.DB.GetTopmenuSubmenuIDs))
		return err
	})

	g.Go(func() error {
		var err error
		response.OverdriveCommands, err = apilookups.E.OverdriveCommands.DBGetRelations(ctx, topmenu, handlerinput.ToIntManyNull(config.Cfg.DB.GetTopmenuOverdriveCommandIDs))
		return err
	})

	g.Go(func() error {
		var err error
		response.Overdrives, err = apilookups.E.Overdrives.DBGetRelations(ctx, topmenu, handlerinput.ToIntManyNull(config.Cfg.DB.GetTopmenuOverdriveIDs))
		return err
	})

	g.Go(func() error {
		var err error
		response.AeonCommands, err = apilookups.E.AeonCommands.DBGetRelations(ctx, topmenu, handlerinput.ToIntManyNull(config.Cfg.DB.GetTopmenuAeonCommandIDs))
		return err
	})

	g.Go(func() error {
		var err error
		response.Abilities, err = apilookups.E.Abilities.DBGetRelations(ctx, topmenu, handlerinput.ToIntManyNull(config.Cfg.DB.GetTopmenuAbilityIDs))
		return err
	})

	return g.Wait()
}
