package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getElementRelationships(server *Server, r *http.Request, element seedtypes.Element) (Element, error) {
	var rel Element
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.StatusProtection, err = E.StatusConditions.dbGetRelationPtr(server, ctx, element, server.Cfg.DB.GetElementStatusConditionID)
		return err
	})

	g.Go(func() error {
		var err error
		rel.AutoAbilities, err = E.AutoAbilities.dbGetRelations(server, ctx, element, server.Cfg.DB.GetElementAutoAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.PlayerAbilities, err = E.PlayerAbilities.dbGetRelations(server, ctx, element, ToIntManyNull(server.Cfg.DB.GetElementPlayerAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.OverdriveAbilities, err = E.OverdriveAbilities.dbGetRelations(server, ctx, element, ToIntManyNull(server.Cfg.DB.GetElementOverdriveAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.ItemAbilities, err = E.ItemAbilities.dbGetRelations(server, ctx, element, ToIntManyNull(server.Cfg.DB.GetElementItemAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.EnemyAbilities, err = E.EnemyAbilities.dbGetRelations(server, ctx, element, ToIntManyNull(server.Cfg.DB.GetElementEnemyAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.MonstersWeak, err = E.Monsters.dbGetRelations(server, ctx, element, server.Cfg.DB.GetElementMonsterIDsWeak)
		return err
	})

	g.Go(func() error {
		var err error
		rel.MonstersHalved, err = E.Monsters.dbGetRelations(server, ctx, element, server.Cfg.DB.GetElementMonsterIDsHalved)
		return err
	})

	g.Go(func() error {
		var err error
		rel.MonstersImmune, err = E.Monsters.dbGetRelations(server, ctx, element, server.Cfg.DB.GetElementMonsterIDsImmune)
		return err
	})

	g.Go(func() error {
		var err error
		rel.MonstersAbsorb, err = E.Monsters.dbGetRelations(server, ctx, element, server.Cfg.DB.GetElementMonsterIDsAbsorb)
		return err
	})

	err := g.Wait()
	if err != nil {
		return Element{}, err
	}

	return rel, nil
}
