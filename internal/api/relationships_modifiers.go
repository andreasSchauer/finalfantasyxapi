package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getModifierRelationships(cfg *Config, r *http.Request, modifier seedtypes.Modifier) (Modifier, error) {
	var rel Modifier
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.AutoAbilities, err = E.AutoAbilities.dbGetRelations(cfg, ctx, modifier, cfg.db.GetModifierAutoAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.PlayerAbilities, err = E.PlayerAbilities.dbGetRelations(cfg, ctx, modifier, cfg.db.GetModifierPlayerAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.OverdriveAbilities, err = E.OverdriveAbilities.dbGetRelations(cfg, ctx, modifier, cfg.db.GetModifierOverdriveAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.ItemAbilities, err = E.ItemAbilities.dbGetRelations(cfg, ctx, modifier, cfg.db.GetModifierItemAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.TriggerCommands, err = E.TriggerCommands.dbGetRelations(cfg, ctx, modifier, cfg.db.GetModifierTriggerCommandIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.EnemyAbilities, err = E.EnemyAbilities.dbGetRelations(cfg, ctx, modifier, cfg.db.GetModifierEnemyAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.StatusConditions, err = E.StatusConditions.dbGetRelations(cfg, ctx, modifier, cfg.db.GetModifierStatusConditionIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.Properties, err = E.Properties.dbGetRelations(cfg, ctx, modifier, cfg.db.GetModifierPropertyIDs)
		return err
	})

	err := g.Wait()
	if err != nil {
		return Modifier{}, err
	}

	return rel, nil
}
