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
	
	g.Go(func() error{
		var err error
		rel.Members, err = getResourcesDbItem(cfg, ctx, cfg.e.playerUnits, class, cfg.db.GetCharacterClassUnitIDs)
		return err
	})

	g.Go(func() error{
		var err error
		rel.DefaultAbilities, err = getResourcesDbItem(cfg, ctx, cfg.e.abilities, class, cfg.db.GetCharacterClassDefaultAbilityIDs)
		if err != nil {
			return err
		}

		rel.LearnableAbilities, err = getClassLearnableAbilities(cfg, ctx, class, rel.DefaultAbilities)
		return err
	})

	g.Go(func() error{
		var err error
		rel.DefaultOverdrives, err = getResourcesDbItem(cfg, ctx, cfg.e.overdrives, class, cfg.db.GetCharacterClassDefaultOverdriveIDs)
		return err
	})

	g.Go(func() error{
		var err error
		rel.LearnableOverdrives, err = getResourcesDbItem(cfg, ctx, cfg.e.overdrives, class, cfg.db.GetCharacterClassLearnableOverdriveIDs)
		return err
	})
	
	g.Go(func() error{
		var err error
		rel.Submenus, err = getResourcesDbItem(cfg, ctx, cfg.e.submenus, class, cfg.db.GetCharacterClassSubmenuIDs)
		return err
	})
	
	err := g.Wait()
	if err != nil {
		return CharacterClass{}, err
	}

	return rel, nil
}

func getClassLearnableAbilities(cfg *Config, ctx context.Context, class seeding.CharacterClass, defaultAbilities []TypedAPIResource) ([]TypedAPIResource, error) {
	allAbilities, err := getResourcesDbItem(cfg, ctx, cfg.e.abilities, class, cfg.db.GetCharacterClassLearnableAbilityIDs)
	if err != nil {
		return nil, err
	}

	learnableAbilities := removeResourcesURL(allAbilities, defaultAbilities)
	return learnableAbilities, nil
}
