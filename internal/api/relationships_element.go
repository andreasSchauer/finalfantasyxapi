package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getElementRelationships(r *http.Request, element seedtypes.Element) (apitypes.Element, error) {
	var rel apitypes.Element
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.StatusProtection, err = E.StatusConditions.DBGetRelationPtr(ctx, element, config.Cfg.DB.GetElementStatusConditionID)
		return err
	})

	g.Go(func() error {
		var err error
		rel.AutoAbilities, err = E.AutoAbilities.DBGetRelations(ctx, element, config.Cfg.DB.GetElementAutoAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.PlayerAbilities, err = E.PlayerAbilities.DBGetRelations(ctx, element, ToIntManyNull(config.Cfg.DB.GetElementPlayerAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.OverdriveAbilities, err = E.OverdriveAbilities.DBGetRelations(ctx, element, ToIntManyNull(config.Cfg.DB.GetElementOverdriveAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.ItemAbilities, err = E.ItemAbilities.DBGetRelations(ctx, element, ToIntManyNull(config.Cfg.DB.GetElementItemAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.EnemyAbilities, err = E.EnemyAbilities.DBGetRelations(ctx, element, ToIntManyNull(config.Cfg.DB.GetElementEnemyAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.MonstersWeak, err = E.Monsters.DBGetRelations(ctx, element, config.Cfg.DB.GetElementMonsterIDsWeak)
		return err
	})

	g.Go(func() error {
		var err error
		rel.MonstersHalved, err = E.Monsters.DBGetRelations(ctx, element, config.Cfg.DB.GetElementMonsterIDsHalved)
		return err
	})

	g.Go(func() error {
		var err error
		rel.MonstersImmune, err = E.Monsters.DBGetRelations(ctx, element, config.Cfg.DB.GetElementMonsterIDsImmune)
		return err
	})

	g.Go(func() error {
		var err error
		rel.MonstersAbsorb, err = E.Monsters.DBGetRelations(ctx, element, config.Cfg.DB.GetElementMonsterIDsAbsorb)
		return err
	})

	err := g.Wait()
	if err != nil {
		return apitypes.Element{}, err
	}

	return rel, nil
}
