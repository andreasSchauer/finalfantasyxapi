package assemble

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func GetElementRelationships(r *http.Request, element seedtypes.Element) (apitypes.Element, error) {
	var rel apitypes.Element
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.StatusProtection, err = apilookups.E.StatusConditions.DBGetRelationPtr(ctx, element, config.Cfg.DB.GetElementStatusConditionID)
		return err
	})

	g.Go(func() error {
		var err error
		rel.AutoAbilities, err = apilookups.E.AutoAbilities.DBGetRelations(ctx, element, config.Cfg.DB.GetElementAutoAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.PlayerAbilities, err = apilookups.E.PlayerAbilities.DBGetRelations(ctx, element, handlerinput.ToIntManyNull(config.Cfg.DB.GetElementPlayerAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.OverdriveAbilities, err = apilookups.E.OverdriveAbilities.DBGetRelations(ctx, element, handlerinput.ToIntManyNull(config.Cfg.DB.GetElementOverdriveAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.ItemAbilities, err = apilookups.E.ItemAbilities.DBGetRelations(ctx, element, handlerinput.ToIntManyNull(config.Cfg.DB.GetElementItemAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.EnemyAbilities, err = apilookups.E.EnemyAbilities.DBGetRelations(ctx, element, handlerinput.ToIntManyNull(config.Cfg.DB.GetElementEnemyAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.MonstersWeak, err = apilookups.E.Monsters.DBGetRelations(ctx, element, config.Cfg.DB.GetElementMonsterIDsWeak)
		return err
	})

	g.Go(func() error {
		var err error
		rel.MonstersHalved, err = apilookups.E.Monsters.DBGetRelations(ctx, element, config.Cfg.DB.GetElementMonsterIDsHalved)
		return err
	})

	g.Go(func() error {
		var err error
		rel.MonstersImmune, err = apilookups.E.Monsters.DBGetRelations(ctx, element, config.Cfg.DB.GetElementMonsterIDsImmune)
		return err
	})

	g.Go(func() error {
		var err error
		rel.MonstersAbsorb, err = apilookups.E.Monsters.DBGetRelations(ctx, element, config.Cfg.DB.GetElementMonsterIDsAbsorb)
		return err
	})

	err := g.Wait()
	if err != nil {
		return apitypes.Element{}, err
	}

	return rel, nil
}
