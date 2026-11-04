package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
	"golang.org/x/sync/errgroup"
)

type StatQueries struct {
	AutoAbilities      DBQueryIntMany
	PlayerAbilities    DBQueryIntMany
	OverdriveAbilities DBQueryIntMany
	ItemAbilities      DBQueryIntMany
	TriggerCommands    DBQueryIntMany
	StatusConditions   DBQueryIntMany
	Properties         DBQueryIntMany
}

func getStatRelationships(srv *Server, r *http.Request, stat seedtypes.Stat) (Stat, error) {
	var rel Stat
	g, ctx := errgroup.WithContext(r.Context())

	queries, err := getStatQueries(srv, r)
	if err != nil {
		return Stat{}, err
	}

	g.Go(func() error {
		var err error
		rel.Spheres, err = E.Spheres.dbGetRelations(srv, ctx, stat, srv.Cfg.DB.GetStatSphereIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.AutoAbilities, err = E.AutoAbilities.dbGetRelations(srv, ctx, stat, queries.AutoAbilities)
		return err
	})

	g.Go(func() error {
		var err error
		rel.PlayerAbilities, err = E.PlayerAbilities.dbGetRelations(srv, ctx, stat, queries.PlayerAbilities)
		return err
	})

	g.Go(func() error {
		var err error
		rel.OverdriveAbilities, err = E.OverdriveAbilities.dbGetRelations(srv, ctx, stat, queries.OverdriveAbilities)
		return err
	})

	g.Go(func() error {
		var err error
		rel.ItemAbilities, err = E.ItemAbilities.dbGetRelations(srv, ctx, stat, queries.ItemAbilities)
		return err
	})

	g.Go(func() error {
		var err error
		rel.TriggerCommands, err = E.TriggerCommands.dbGetRelations(srv, ctx, stat, queries.TriggerCommands)
		return err
	})

	g.Go(func() error {
		var err error
		rel.StatusConditions, err = E.StatusConditions.dbGetRelations(srv, ctx, stat, queries.StatusConditions)
		return err
	})

	g.Go(func() error {
		var err error
		if queries.Properties != nil {
			rel.Properties, err = E.Properties.dbGetRelations(srv, ctx, stat, queries.Properties)
			return err
		}
		return nil
	})

	err = g.Wait()
	if err != nil {
		return Stat{}, err
	}

	return rel, nil
}

func getStatQueries(srv *Server, r *http.Request) (StatQueries, error) {
	changesOnly, err := Q.stats[ts.QpnChangesOnly].parseBool(r)
	if errExceptEmptyQuery(err) {
		return StatQueries{}, err
	}

	var queries StatQueries

	if changesOnly {
		queries = StatQueries{
			AutoAbilities:      srv.Cfg.DB.GetStatAutoAbilityIDsStatChange,
			PlayerAbilities:    srv.Cfg.DB.GetStatPlayerAbilityIDsStatChange,
			OverdriveAbilities: srv.Cfg.DB.GetStatOverdriveAbilityIDsStatChange,
			ItemAbilities:      srv.Cfg.DB.GetStatItemAbilityIDsStatChange,
			TriggerCommands:    srv.Cfg.DB.GetStatTriggerCommandIDsStatChange,
			StatusConditions:   srv.Cfg.DB.GetStatStatusConditionIDsStatChange,
			Properties:         nil,
		}
	} else {
		queries = StatQueries{
			AutoAbilities:      srv.Cfg.DB.GetStatAutoAbilityIDs,
			PlayerAbilities:    srv.Cfg.DB.GetStatPlayerAbilityIDs,
			OverdriveAbilities: srv.Cfg.DB.GetStatOverdriveAbilityIDs,
			ItemAbilities:      srv.Cfg.DB.GetStatItemAbilityIDs,
			TriggerCommands:    srv.Cfg.DB.GetStatTriggerCommandIDs,
			StatusConditions:   srv.Cfg.DB.GetStatStatusConditionIDs,
			Properties:         srv.Cfg.DB.GetStatPropertyIDs,
		}
	}

	return queries, nil
}
