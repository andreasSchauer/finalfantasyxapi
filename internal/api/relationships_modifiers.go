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
		rel.AutoAbilities, err = getResourcesDbItem(cfg, ctx, cfg.e.autoAbilities, modifier, cfg.db.GetModifierAutoAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.PlayerAbilities, err = getResourcesDbItem(cfg, ctx, cfg.e.playerAbilities, modifier, cfg.db.GetModifierPlayerAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.OverdriveAbilities, err = getResourcesDbItem(cfg, ctx, cfg.e.overdriveAbilities, modifier, cfg.db.GetModifierOverdriveAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.ItemAbilities, err = getResourcesDbItem(cfg, ctx, cfg.e.itemAbilities, modifier, cfg.db.GetModifierItemAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.TriggerCommands, err = getResourcesDbItem(cfg, ctx, cfg.e.triggerCommands, modifier, cfg.db.GetModifierTriggerCommandIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.EnemyAbilities, err = getResourcesDbItem(cfg, ctx, cfg.e.enemyAbilities, modifier, cfg.db.GetModifierEnemyAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.StatusConditions, err = getResourcesDbItem(cfg, ctx, cfg.e.statusConditions, modifier, cfg.db.GetModifierStatusConditionIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.Properties, err = getResourcesDbItem(cfg, ctx, cfg.e.properties, modifier, cfg.db.GetModifierPropertyIDs)
		return err
	})

	err := g.Wait()
	if err != nil {
		return Modifier{}, err
	}

	return rel, nil
}
