package assemble

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func GetModifierRelationships(r *http.Request, modifier seedtypes.Modifier) (apitypes.Modifier, error) {
	var rel apitypes.Modifier
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.AutoAbilities, err = apilookups.E.AutoAbilities.DBGetRelations(ctx, modifier, config.Cfg.DB.GetModifierAutoAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.PlayerAbilities, err = apilookups.E.PlayerAbilities.DBGetRelations(ctx, modifier, config.Cfg.DB.GetModifierPlayerAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.OverdriveAbilities, err = apilookups.E.OverdriveAbilities.DBGetRelations(ctx, modifier, config.Cfg.DB.GetModifierOverdriveAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.ItemAbilities, err = apilookups.E.ItemAbilities.DBGetRelations(ctx, modifier, config.Cfg.DB.GetModifierItemAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.TriggerCommands, err = apilookups.E.TriggerCommands.DBGetRelations(ctx, modifier, config.Cfg.DB.GetModifierTriggerCommandIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.EnemyAbilities, err = apilookups.E.EnemyAbilities.DBGetRelations(ctx, modifier, config.Cfg.DB.GetModifierEnemyAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.StatusConditions, err = apilookups.E.StatusConditions.DBGetRelations(ctx, modifier, config.Cfg.DB.GetModifierStatusConditionIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.Properties, err = apilookups.E.Properties.DBGetRelations(ctx, modifier, config.Cfg.DB.GetModifierPropertyIDs)
		return err
	})

	err := g.Wait()
	if err != nil {
		return apitypes.Modifier{}, err
	}

	return rel, nil
}
