package api

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getCharClassRelationships(srv *Server, r *http.Request, class seedtypes.CharacterClass) (CharacterClass, error) {
	var rel CharacterClass
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.Members, err = E.PlayerUnits.dbGetRelations(srv, ctx, class, config.Cfg.DB.GetCharacterClassUnitIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.DefaultAbilities, err = E.Abilities.dbGetRelations(srv, ctx, class, config.Cfg.DB.GetCharacterClassDefaultAbilityIDs)
		if err != nil {
			return err
		}

		rel.LearnableAbilities, err = getClassLearnableAbilities(srv, ctx, class, rel.DefaultAbilities)
		return err
	})

	g.Go(func() error {
		var err error
		rel.DefaultOverdrives, err = E.Overdrives.dbGetRelations(srv, ctx, class, config.Cfg.DB.GetCharacterClassDefaultOverdriveIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.LearnableOverdrives, err = E.Overdrives.dbGetRelations(srv, ctx, class, config.Cfg.DB.GetCharacterClassLearnableOverdriveIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.Submenus, err = E.Submenus.dbGetRelations(srv, ctx, class, config.Cfg.DB.GetCharacterClassSubmenuIDs)
		return err
	})

	err := g.Wait()
	if err != nil {
		return CharacterClass{}, err
	}

	return rel, nil
}

func getClassLearnableAbilities(srv *Server, ctx context.Context, class seedtypes.CharacterClass, defaultAbilities []TypedAPIResource) ([]TypedAPIResource, error) {
	allAbilities, err := E.Abilities.dbGetRelations(srv, ctx, class, config.Cfg.DB.GetCharacterClassLearnableAbilityIDs)
	if err != nil {
		return nil, err
	}

	learnableAbilities := removeResourcesURL(allAbilities, defaultAbilities)
	return learnableAbilities, nil
}
