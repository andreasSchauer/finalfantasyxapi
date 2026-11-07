package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
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

func getStatRelationships(r *http.Request, stat seedtypes.Stat) (apitypes.Stat, error) {
	var rel apitypes.Stat
	g, ctx := errgroup.WithContext(r.Context())

	queries, err := getStatQueries(r)
	if err != nil {
		return apitypes.Stat{}, err
	}

	g.Go(func() error {
		var err error
		rel.Spheres, err = E.Spheres.DBGetRelations(ctx, stat, config.Cfg.DB.GetStatSphereIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.AutoAbilities, err = E.AutoAbilities.DBGetRelations(ctx, stat, queries.AutoAbilities)
		return err
	})

	g.Go(func() error {
		var err error
		rel.PlayerAbilities, err = E.PlayerAbilities.DBGetRelations(ctx, stat, queries.PlayerAbilities)
		return err
	})

	g.Go(func() error {
		var err error
		rel.OverdriveAbilities, err = E.OverdriveAbilities.DBGetRelations(ctx, stat, queries.OverdriveAbilities)
		return err
	})

	g.Go(func() error {
		var err error
		rel.ItemAbilities, err = E.ItemAbilities.DBGetRelations(ctx, stat, queries.ItemAbilities)
		return err
	})

	g.Go(func() error {
		var err error
		rel.TriggerCommands, err = E.TriggerCommands.DBGetRelations(ctx, stat, queries.TriggerCommands)
		return err
	})

	g.Go(func() error {
		var err error
		rel.StatusConditions, err = E.StatusConditions.DBGetRelations(ctx, stat, queries.StatusConditions)
		return err
	})

	g.Go(func() error {
		var err error
		if queries.Properties != nil {
			rel.Properties, err = E.Properties.DBGetRelations(ctx, stat, queries.Properties)
			return err
		}
		return nil
	})

	err = g.Wait()
	if err != nil {
		return apitypes.Stat{}, err
	}

	return rel, nil
}

func getStatQueries(r *http.Request) (StatQueries, error) {
	changesOnly, err := Q.Stats[ts.QpnChangesOnly].ParseBool(r)
	if errs.ErrExceptEmptyQuery(err) {
		return StatQueries{}, err
	}

	var queries StatQueries

	if changesOnly {
		queries = StatQueries{
			AutoAbilities:      config.Cfg.DB.GetStatAutoAbilityIDsStatChange,
			PlayerAbilities:    config.Cfg.DB.GetStatPlayerAbilityIDsStatChange,
			OverdriveAbilities: config.Cfg.DB.GetStatOverdriveAbilityIDsStatChange,
			ItemAbilities:      config.Cfg.DB.GetStatItemAbilityIDsStatChange,
			TriggerCommands:    config.Cfg.DB.GetStatTriggerCommandIDsStatChange,
			StatusConditions:   config.Cfg.DB.GetStatStatusConditionIDsStatChange,
			Properties:         nil,
		}
	} else {
		queries = StatQueries{
			AutoAbilities:      config.Cfg.DB.GetStatAutoAbilityIDs,
			PlayerAbilities:    config.Cfg.DB.GetStatPlayerAbilityIDs,
			OverdriveAbilities: config.Cfg.DB.GetStatOverdriveAbilityIDs,
			ItemAbilities:      config.Cfg.DB.GetStatItemAbilityIDs,
			TriggerCommands:    config.Cfg.DB.GetStatTriggerCommandIDs,
			StatusConditions:   config.Cfg.DB.GetStatStatusConditionIDs,
			Properties:         config.Cfg.DB.GetStatPropertyIDs,
		}
	}

	return queries, nil
}
