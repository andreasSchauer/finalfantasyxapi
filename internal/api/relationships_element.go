package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
	"golang.org/x/sync/errgroup"
)

func getElementRelationships(cfg *Config, r *http.Request, element seeding.Element) (Element, error) {
	var rel Element
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.StatusProtection, err = getResPtrDB(cfg, ctx, cfg.E.StatusConditions, element, cfg.db.GetElementStatusConditionID)
		return err
	})

	g.Go(func() error {
		var err error
		rel.AutoAbilities, err = getResourcesDBItem(cfg, ctx, cfg.E.AutoAbilities, element, cfg.db.GetElementAutoAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.PlayerAbilities, err = getResourcesDBItem(cfg, ctx, cfg.E.PlayerAbilities, element, ToIntManyNull(cfg.db.GetElementPlayerAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.OverdriveAbilities, err = getResourcesDBItem(cfg, ctx, cfg.E.OverdriveAbilities, element, ToIntManyNull(cfg.db.GetElementOverdriveAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.ItemAbilities, err = getResourcesDBItem(cfg, ctx, cfg.E.ItemAbilities, element, ToIntManyNull(cfg.db.GetElementItemAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.EnemyAbilities, err = getResourcesDBItem(cfg, ctx, cfg.E.EnemyAbilities, element, ToIntManyNull(cfg.db.GetElementEnemyAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.MonstersWeak, err = getResourcesDBItem(cfg, ctx, cfg.E.Monsters, element, cfg.db.GetElementMonsterIDsWeak)
		return err
	})

	g.Go(func() error {
		var err error
		rel.MonstersHalved, err = getResourcesDBItem(cfg, ctx, cfg.E.Monsters, element, cfg.db.GetElementMonsterIDsHalved)
		return err
	})

	g.Go(func() error {
		var err error
		rel.MonstersImmune, err = getResourcesDBItem(cfg, ctx, cfg.E.Monsters, element, cfg.db.GetElementMonsterIDsImmune)
		return err
	})

	g.Go(func() error {
		var err error
		rel.MonstersAbsorb, err = getResourcesDBItem(cfg, ctx, cfg.E.Monsters, element, cfg.db.GetElementMonsterIDsAbsorb)
		return err
	})

	err := g.Wait()
	if err != nil {
		return Element{}, err
	}

	return rel, nil
}
