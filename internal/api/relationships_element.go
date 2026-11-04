package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getElementRelationships(cfg *Config, r *http.Request, element seedtypes.Element) (Element, error) {
	var rel Element
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.StatusProtection, err = E.StatusConditions.dbGetRelationPtr(cfg, ctx, element, cfg.db.GetElementStatusConditionID)
		return err
	})

	g.Go(func() error {
		var err error
		rel.AutoAbilities, err = E.AutoAbilities.dbGetRelations(cfg, ctx, element, cfg.db.GetElementAutoAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.PlayerAbilities, err = E.PlayerAbilities.dbGetRelations(cfg, ctx, element, ToIntManyNull(cfg.db.GetElementPlayerAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.OverdriveAbilities, err = E.OverdriveAbilities.dbGetRelations(cfg, ctx, element, ToIntManyNull(cfg.db.GetElementOverdriveAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.ItemAbilities, err = E.ItemAbilities.dbGetRelations(cfg, ctx, element, ToIntManyNull(cfg.db.GetElementItemAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.EnemyAbilities, err = E.EnemyAbilities.dbGetRelations(cfg, ctx, element, ToIntManyNull(cfg.db.GetElementEnemyAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.MonstersWeak, err = E.Monsters.dbGetRelations(cfg, ctx, element, cfg.db.GetElementMonsterIDsWeak)
		return err
	})

	g.Go(func() error {
		var err error
		rel.MonstersHalved, err = E.Monsters.dbGetRelations(cfg, ctx, element, cfg.db.GetElementMonsterIDsHalved)
		return err
	})

	g.Go(func() error {
		var err error
		rel.MonstersImmune, err = E.Monsters.dbGetRelations(cfg, ctx, element, cfg.db.GetElementMonsterIDsImmune)
		return err
	})

	g.Go(func() error {
		var err error
		rel.MonstersAbsorb, err = E.Monsters.dbGetRelations(cfg, ctx, element, cfg.db.GetElementMonsterIDsAbsorb)
		return err
	})

	err := g.Wait()
	if err != nil {
		return Element{}, err
	}

	return rel, nil
}
