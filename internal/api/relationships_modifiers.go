package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getModifierRelationships(server *Server, r *http.Request, modifier seedtypes.Modifier) (Modifier, error) {
	var rel Modifier
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.AutoAbilities, err = E.AutoAbilities.dbGetRelations(server, ctx, modifier, server.Cfg.DB.GetModifierAutoAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.PlayerAbilities, err = E.PlayerAbilities.dbGetRelations(server, ctx, modifier, server.Cfg.DB.GetModifierPlayerAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.OverdriveAbilities, err = E.OverdriveAbilities.dbGetRelations(server, ctx, modifier, server.Cfg.DB.GetModifierOverdriveAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.ItemAbilities, err = E.ItemAbilities.dbGetRelations(server, ctx, modifier, server.Cfg.DB.GetModifierItemAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.TriggerCommands, err = E.TriggerCommands.dbGetRelations(server, ctx, modifier, server.Cfg.DB.GetModifierTriggerCommandIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.EnemyAbilities, err = E.EnemyAbilities.dbGetRelations(server, ctx, modifier, server.Cfg.DB.GetModifierEnemyAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.StatusConditions, err = E.StatusConditions.dbGetRelations(server, ctx, modifier, server.Cfg.DB.GetModifierStatusConditionIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.Properties, err = E.Properties.dbGetRelations(server, ctx, modifier, server.Cfg.DB.GetModifierPropertyIDs)
		return err
	})

	err := g.Wait()
	if err != nil {
		return Modifier{}, err
	}

	return rel, nil
}
