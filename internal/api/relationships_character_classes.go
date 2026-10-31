package api

import (
	"context"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
	"golang.org/x/sync/errgroup"
)

func getCharClassRelationships(cfg *Config, r *http.Request, class seeding.CharacterClass) (CharacterClass, error) {
	var rel CharacterClass
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.Members, err = getResourcesDBItem(cfg, ctx, cfg.E.PlayerUnits, class, cfg.db.GetCharacterClassUnitIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.DefaultAbilities, err = getResourcesDBItem(cfg, ctx, cfg.E.Abilities, class, cfg.db.GetCharacterClassDefaultAbilityIDs)
		if err != nil {
			return err
		}

		rel.LearnableAbilities, err = getClassLearnableAbilities(cfg, ctx, class, rel.DefaultAbilities)
		return err
	})

	g.Go(func() error {
		var err error
		rel.DefaultOverdrives, err = getResourcesDBItem(cfg, ctx, cfg.E.Overdrives, class, cfg.db.GetCharacterClassDefaultOverdriveIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.LearnableOverdrives, err = getResourcesDBItem(cfg, ctx, cfg.E.Overdrives, class, cfg.db.GetCharacterClassLearnableOverdriveIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.Submenus, err = getResourcesDBItem(cfg, ctx, cfg.E.Submenus, class, cfg.db.GetCharacterClassSubmenuIDs)
		return err
	})

	err := g.Wait()
	if err != nil {
		return CharacterClass{}, err
	}

	return rel, nil
}

func getClassLearnableAbilities(cfg *Config, ctx context.Context, class seeding.CharacterClass, defaultAbilities []TypedAPIResource) ([]TypedAPIResource, error) {
	allAbilities, err := getResourcesDBItem(cfg, ctx, cfg.E.Abilities, class, cfg.db.GetCharacterClassLearnableAbilityIDs)
	if err != nil {
		return nil, err
	}

	learnableAbilities := removeResourcesURL(allAbilities, defaultAbilities)
	return learnableAbilities, nil
}
