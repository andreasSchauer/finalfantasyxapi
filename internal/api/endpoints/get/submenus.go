package get

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"golang.org/x/sync/errgroup"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Submenu(r *http.Request, i handlerinput.Main[seedtypes.Submenu, apitypes.Submenu, apiresource.Named], id int32) (apitypes.Submenu, error) {
	submenu, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Submenu{}, err
	}

	rel, err := getSubmenuRelationships(r, submenu)
	if err != nil {
		return apitypes.Submenu{}, err
	}

	response := apitypes.Submenu{
		ID:          submenu.ID,
		Name:        submenu.Name,
		Description: submenu.Description,
		Effect:      submenu.Effect,
		Topmenu:     handlerinput.NamePtrToNamedAPIResPtr(apilookups.E.Topmenus, submenu.Topmenu, nil),
		Users:       handlerinput.NamesToNamedAPIResources(apilookups.E.CharacterClasses, submenu.Users),
		OpenedBy:    rel.OpenedBy,
		Abilities:   rel.Abilities,
	}

	return response, nil
}

func getSubmenuRelationships(r *http.Request, submenu seedtypes.Submenu) (apitypes.Submenu, error) {
	var rel apitypes.Submenu
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.Abilities, err = apilookups.E.Abilities.DBGetRelations(ctx, submenu, handlerinput.ToIntManyNull(config.Cfg.DB.GetSubmenuAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.OpenedBy, err = createSubmenuOpenedBy(ctx, submenu)
		return err
	})

	err := g.Wait()
	if err != nil {
		return apitypes.Submenu{}, err
	}

	return rel, nil
}

func createSubmenuOpenedBy(ctx context.Context, submenu seedtypes.Submenu) (*apitypes.MenuOpen, error) {
	var menuOpen apitypes.MenuOpen
	g, ctx := errgroup.WithContext(ctx)

	g.Go(func() error {
		var err error
		menuOpen.Ability, err = apilookups.E.Abilities.DBGetRelationPtr(ctx, submenu, handlerinput.ToIntOneNull(config.Cfg.DB.GetSubmenuOpenedByAbilityID))
		return err
	})

	g.Go(func() error {
		var err error
		menuOpen.AeonCommand, err = apilookups.E.AeonCommands.DBGetRelationPtr(ctx, submenu, handlerinput.ToIntOneNull(config.Cfg.DB.GetSubmenuOpenedByAeonCommandID))
		return err
	})

	g.Go(func() error {
		overdriveCommands, err := apilookups.E.OverdriveCommands.DBGetRelations(ctx, submenu, handlerinput.ToIntManyNull(config.Cfg.DB.GetSubmenuOpenedByOverdriveCommandIDs))
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
