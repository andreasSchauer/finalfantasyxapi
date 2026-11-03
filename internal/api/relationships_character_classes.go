package api

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getCharClassRelationships(cfg *Config, r *http.Request, class seedtypes.CharacterClass) (CharacterClass, error) {
	var rel CharacterClass
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.Members, err = cfg.E.PlayerUnits.dbGetRelations(cfg, ctx, class, cfg.db.GetCharacterClassUnitIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.DefaultAbilities, err = cfg.E.Abilities.dbGetRelations(cfg, ctx, class, cfg.db.GetCharacterClassDefaultAbilityIDs)
		if err != nil {
			return err
		}

		rel.LearnableAbilities, err = getClassLearnableAbilities(cfg, ctx, class, rel.DefaultAbilities)
		return err
	})

	g.Go(func() error {
		var err error
		rel.DefaultOverdrives, err = cfg.E.Overdrives.dbGetRelations(cfg, ctx, class, cfg.db.GetCharacterClassDefaultOverdriveIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.LearnableOverdrives, err = cfg.E.Overdrives.dbGetRelations(cfg, ctx, class, cfg.db.GetCharacterClassLearnableOverdriveIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.Submenus, err = cfg.E.Submenus.dbGetRelations(cfg, ctx, class, cfg.db.GetCharacterClassSubmenuIDs)
		return err
	})

	err := g.Wait()
	if err != nil {
		return CharacterClass{}, err
	}

	return rel, nil
}

func getClassLearnableAbilities(cfg *Config, ctx context.Context, class seedtypes.CharacterClass, defaultAbilities []TypedAPIResource) ([]TypedAPIResource, error) {
	allAbilities, err := cfg.E.Abilities.dbGetRelations(cfg, ctx, class, cfg.db.GetCharacterClassLearnableAbilityIDs)
	if err != nil {
		return nil, err
	}

	learnableAbilities := removeResourcesURL(allAbilities, defaultAbilities)
	return learnableAbilities, nil
}
