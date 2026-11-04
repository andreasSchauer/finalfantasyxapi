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

func getStatRelationships(server *Server, r *http.Request, stat seedtypes.Stat) (Stat, error) {
	var rel Stat
	g, ctx := errgroup.WithContext(r.Context())

	queries, err := getStatQueries(server, r)
	if err != nil {
		return Stat{}, err
	}

	g.Go(func() error {
		var err error
		rel.Spheres, err = E.Spheres.dbGetRelations(server, ctx, stat, server.Cfg.DB.GetStatSphereIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.AutoAbilities, err = E.AutoAbilities.dbGetRelations(server, ctx, stat, queries.AutoAbilities)
		return err
	})

	g.Go(func() error {
		var err error
		rel.PlayerAbilities, err = E.PlayerAbilities.dbGetRelations(server, ctx, stat, queries.PlayerAbilities)
		return err
	})

	g.Go(func() error {
		var err error
		rel.OverdriveAbilities, err = E.OverdriveAbilities.dbGetRelations(server, ctx, stat, queries.OverdriveAbilities)
		return err
	})

	g.Go(func() error {
		var err error
		rel.ItemAbilities, err = E.ItemAbilities.dbGetRelations(server, ctx, stat, queries.ItemAbilities)
		return err
	})

	g.Go(func() error {
		var err error
		rel.TriggerCommands, err = E.TriggerCommands.dbGetRelations(server, ctx, stat, queries.TriggerCommands)
		return err
	})

	g.Go(func() error {
		var err error
		rel.StatusConditions, err = E.StatusConditions.dbGetRelations(server, ctx, stat, queries.StatusConditions)
		return err
	})

	g.Go(func() error {
		var err error
		if queries.Properties != nil {
			rel.Properties, err = E.Properties.dbGetRelations(server, ctx, stat, queries.Properties)
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

func getStatQueries(server *Server, r *http.Request) (StatQueries, error) {
	changesOnly, err := Q.stats[ts.QpnChangesOnly].parseBool(r)
	if errExceptEmptyQuery(err) {
		return StatQueries{}, err
	}

	var queries StatQueries

	if changesOnly {
		queries = StatQueries{
			AutoAbilities:      server.Cfg.DB.GetStatAutoAbilityIDsStatChange,
			PlayerAbilities:    server.Cfg.DB.GetStatPlayerAbilityIDsStatChange,
			OverdriveAbilities: server.Cfg.DB.GetStatOverdriveAbilityIDsStatChange,
			ItemAbilities:      server.Cfg.DB.GetStatItemAbilityIDsStatChange,
			TriggerCommands:    server.Cfg.DB.GetStatTriggerCommandIDsStatChange,
			StatusConditions:   server.Cfg.DB.GetStatStatusConditionIDsStatChange,
			Properties:         nil,
		}
	} else {
		queries = StatQueries{
			AutoAbilities:      server.Cfg.DB.GetStatAutoAbilityIDs,
			PlayerAbilities:    server.Cfg.DB.GetStatPlayerAbilityIDs,
			OverdriveAbilities: server.Cfg.DB.GetStatOverdriveAbilityIDs,
			ItemAbilities:      server.Cfg.DB.GetStatItemAbilityIDs,
			TriggerCommands:    server.Cfg.DB.GetStatTriggerCommandIDs,
			StatusConditions:   server.Cfg.DB.GetStatStatusConditionIDs,
			Properties:         server.Cfg.DB.GetStatPropertyIDs,
		}
	}

	return queries, nil
}
