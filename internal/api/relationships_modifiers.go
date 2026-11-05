package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getModifierRelationships(srv *Server, r *http.Request, modifier seedtypes.Modifier) (Modifier, error) {
	var rel Modifier
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.AutoAbilities, err = E.AutoAbilities.dbGetRelations(srv, ctx, modifier, config.Cfg.DB.GetModifierAutoAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.PlayerAbilities, err = E.PlayerAbilities.dbGetRelations(srv, ctx, modifier, config.Cfg.DB.GetModifierPlayerAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.OverdriveAbilities, err = E.OverdriveAbilities.dbGetRelations(srv, ctx, modifier, config.Cfg.DB.GetModifierOverdriveAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.ItemAbilities, err = E.ItemAbilities.dbGetRelations(srv, ctx, modifier, config.Cfg.DB.GetModifierItemAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.TriggerCommands, err = E.TriggerCommands.dbGetRelations(srv, ctx, modifier, config.Cfg.DB.GetModifierTriggerCommandIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.EnemyAbilities, err = E.EnemyAbilities.dbGetRelations(srv, ctx, modifier, config.Cfg.DB.GetModifierEnemyAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.StatusConditions, err = E.StatusConditions.dbGetRelations(srv, ctx, modifier, config.Cfg.DB.GetModifierStatusConditionIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.Properties, err = E.Properties.dbGetRelations(srv, ctx, modifier, config.Cfg.DB.GetModifierPropertyIDs)
		return err
	})

	err := g.Wait()
	if err != nil {
		return Modifier{}, err
	}

	return rel, nil
}
