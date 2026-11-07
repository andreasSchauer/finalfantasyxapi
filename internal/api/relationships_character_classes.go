package api

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getCharClassRelationships(r *http.Request, class seedtypes.CharacterClass) (apitypes.CharacterClass, error) {
	var rel apitypes.CharacterClass
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.Members, err = E.PlayerUnits.DBGetRelations(ctx, class, config.Cfg.DB.GetCharacterClassUnitIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.DefaultAbilities, err = E.Abilities.DBGetRelations(ctx, class, config.Cfg.DB.GetCharacterClassDefaultAbilityIDs)
		if err != nil {
			return err
		}

		rel.LearnableAbilities, err = getClassLearnableAbilities(ctx, class, rel.DefaultAbilities)
		return err
	})

	g.Go(func() error {
		var err error
		rel.DefaultOverdrives, err = E.Overdrives.DBGetRelations(ctx, class, config.Cfg.DB.GetCharacterClassDefaultOverdriveIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.LearnableOverdrives, err = E.Overdrives.DBGetRelations(ctx, class, config.Cfg.DB.GetCharacterClassLearnableOverdriveIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.Submenus, err = E.Submenus.DBGetRelations(ctx, class, config.Cfg.DB.GetCharacterClassSubmenuIDs)
		return err
	})

	err := g.Wait()
	if err != nil {
		return apitypes.CharacterClass{}, err
	}

	return rel, nil
}

func getClassLearnableAbilities(ctx context.Context, class seedtypes.CharacterClass, defaultAbilities []apiresource.TypedAPIResource) ([]apiresource.TypedAPIResource, error) {
	allAbilities, err := E.Abilities.DBGetRelations(ctx, class, config.Cfg.DB.GetCharacterClassLearnableAbilityIDs)
	if err != nil {
		return nil, err
	}

	learnableAbilities := RemoveResourcesURL(allAbilities, defaultAbilities)
	return learnableAbilities, nil
}
