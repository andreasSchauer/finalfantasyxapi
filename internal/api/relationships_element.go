package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
	"golang.org/x/sync/errgroup"
)

func getElementRelationships(cfg *Config, r *http.Request, element seeding.Element) (Element, error) {
	var rel Element
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error{
		var err error
		rel.StatusProtection, err = getResPtrDB(cfg, ctx, cfg.e.statusConditions, element, cfg.db.GetElementStatusConditionID)
		return err
	})
	
	g.Go(func() error{
		var err error
		rel.AutoAbilities, err = getResourcesDbItem(cfg, ctx, cfg.e.autoAbilities, element, cfg.db.GetElementAutoAbilityIDs)
		return err
	})

	g.Go(func() error{
		var err error
		rel.PlayerAbilities, err = getResourcesDbItem(cfg, ctx, cfg.e.playerAbilities, element, ToIntManyNull(cfg.db.GetElementPlayerAbilityIDs))
		return err
	})

	g.Go(func() error{
		var err error
		rel.OverdriveAbilities, err = getResourcesDbItem(cfg, ctx, cfg.e.overdriveAbilities, element, ToIntManyNull(cfg.db.GetElementOverdriveAbilityIDs))
		return err
	})

	g.Go(func() error{
		var err error
		rel.ItemAbilities, err = getResourcesDbItem(cfg, ctx, cfg.e.itemAbilities, element, ToIntManyNull(cfg.db.GetElementItemAbilityIDs))
		return err
	})

	g.Go(func() error{
		var err error
		rel.EnemyAbilities, err = getResourcesDbItem(cfg, ctx, cfg.e.enemyAbilities, element, ToIntManyNull(cfg.db.GetElementEnemyAbilityIDs))
		return err
	})

	g.Go(func() error{
		var err error
		rel.MonstersWeak, err = getResourcesDbItem(cfg, ctx, cfg.e.monsters, element, cfg.db.GetElementMonsterIDsWeak)
		return err
	})

	g.Go(func() error{
		var err error
		rel.MonstersHalved, err = getResourcesDbItem(cfg, ctx, cfg.e.monsters, element, cfg.db.GetElementMonsterIDsHalved)
		return err
	})

	g.Go(func() error{
		var err error
		rel.MonstersImmune, err = getResourcesDbItem(cfg, ctx, cfg.e.monsters, element, cfg.db.GetElementMonsterIDsImmune)
		return err
	})

	g.Go(func() error{
		var err error
		rel.MonstersAbsorb, err = getResourcesDbItem(cfg, ctx, cfg.e.monsters, element, cfg.db.GetElementMonsterIDsAbsorb)
		return err
	})

	err := g.Wait()
	if err != nil {
		return Element{}, err
	}

	return rel, nil
}
