package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
	"golang.org/x/sync/errgroup"
)

func getModifierRelationships(cfg *Config, r *http.Request, modifier seeding.Modifier) (Modifier, error) {
	var rel Modifier
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.AutoAbilities, err = getResourcesDBItem(cfg, ctx, cfg.E.AutoAbilities, modifier, cfg.db.GetModifierAutoAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.PlayerAbilities, err = getResourcesDBItem(cfg, ctx, cfg.E.PlayerAbilities, modifier, cfg.db.GetModifierPlayerAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.OverdriveAbilities, err = getResourcesDBItem(cfg, ctx, cfg.E.OverdriveAbilities, modifier, cfg.db.GetModifierOverdriveAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.ItemAbilities, err = getResourcesDBItem(cfg, ctx, cfg.E.ItemAbilities, modifier, cfg.db.GetModifierItemAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.TriggerCommands, err = getResourcesDBItem(cfg, ctx, cfg.E.TriggerCommands, modifier, cfg.db.GetModifierTriggerCommandIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.EnemyAbilities, err = getResourcesDBItem(cfg, ctx, cfg.E.EnemyAbilities, modifier, cfg.db.GetModifierEnemyAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.StatusConditions, err = getResourcesDBItem(cfg, ctx, cfg.E.StatusConditions, modifier, cfg.db.GetModifierStatusConditionIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.Properties, err = getResourcesDBItem(cfg, ctx, cfg.E.Properties, modifier, cfg.db.GetModifierPropertyIDs)
		return err
	})

	err := g.Wait()
	if err != nil {
		return Modifier{}, err
	}

	return rel, nil
}
