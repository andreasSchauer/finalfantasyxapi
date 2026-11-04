package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	"golang.org/x/sync/errgroup"
)

func getElementRelationships(srv *Server, r *http.Request, element seedtypes.Element) (Element, error) {
	var rel Element
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.StatusProtection, err = E.StatusConditions.dbGetRelationPtr(srv, ctx, element, srv.Cfg.DB.GetElementStatusConditionID)
		return err
	})

	g.Go(func() error {
		var err error
		rel.AutoAbilities, err = E.AutoAbilities.dbGetRelations(srv, ctx, element, srv.Cfg.DB.GetElementAutoAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.PlayerAbilities, err = E.PlayerAbilities.dbGetRelations(srv, ctx, element, ToIntManyNull(srv.Cfg.DB.GetElementPlayerAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.OverdriveAbilities, err = E.OverdriveAbilities.dbGetRelations(srv, ctx, element, ToIntManyNull(srv.Cfg.DB.GetElementOverdriveAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.ItemAbilities, err = E.ItemAbilities.dbGetRelations(srv, ctx, element, ToIntManyNull(srv.Cfg.DB.GetElementItemAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.EnemyAbilities, err = E.EnemyAbilities.dbGetRelations(srv, ctx, element, ToIntManyNull(srv.Cfg.DB.GetElementEnemyAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.MonstersWeak, err = E.Monsters.dbGetRelations(srv, ctx, element, srv.Cfg.DB.GetElementMonsterIDsWeak)
		return err
	})

	g.Go(func() error {
		var err error
		rel.MonstersHalved, err = E.Monsters.dbGetRelations(srv, ctx, element, srv.Cfg.DB.GetElementMonsterIDsHalved)
		return err
	})

	g.Go(func() error {
		var err error
		rel.MonstersImmune, err = E.Monsters.dbGetRelations(srv, ctx, element, srv.Cfg.DB.GetElementMonsterIDsImmune)
		return err
	})

	g.Go(func() error {
		var err error
		rel.MonstersAbsorb, err = E.Monsters.dbGetRelations(srv, ctx, element, srv.Cfg.DB.GetElementMonsterIDsAbsorb)
		return err
	})

	err := g.Wait()
	if err != nil {
		return Element{}, err
	}

	return rel, nil
}
