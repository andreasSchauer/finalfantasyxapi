package assemble

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
	"golang.org/x/sync/errgroup"
)

type StatQueries struct {
	AutoAbilities      handlerinput.DBQueryIntMany
	PlayerAbilities    handlerinput.DBQueryIntMany
	OverdriveAbilities handlerinput.DBQueryIntMany
	ItemAbilities      handlerinput.DBQueryIntMany
	TriggerCommands    handlerinput.DBQueryIntMany
	StatusConditions   handlerinput.DBQueryIntMany
	Properties         handlerinput.DBQueryIntMany
}

func GetStatRelationships(r *http.Request, stat seedtypes.Stat) (apitypes.Stat, error) {
	var rel apitypes.Stat
	g, ctx := errgroup.WithContext(r.Context())

	queries, err := getStatQueries(r)
	if err != nil {
		return apitypes.Stat{}, err
	}

	g.Go(func() error {
		var err error
		rel.Spheres, err = apilookups.E.Spheres.DBGetRelations(ctx, stat, config.Cfg.DB.GetStatSphereIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.AutoAbilities, err = apilookups.E.AutoAbilities.DBGetRelations(ctx, stat, queries.AutoAbilities)
		return err
	})

	g.Go(func() error {
		var err error
		rel.PlayerAbilities, err = apilookups.E.PlayerAbilities.DBGetRelations(ctx, stat, queries.PlayerAbilities)
		return err
	})

	g.Go(func() error {
		var err error
		rel.OverdriveAbilities, err = apilookups.E.OverdriveAbilities.DBGetRelations(ctx, stat, queries.OverdriveAbilities)
		return err
	})

	g.Go(func() error {
		var err error
		rel.ItemAbilities, err = apilookups.E.ItemAbilities.DBGetRelations(ctx, stat, queries.ItemAbilities)
		return err
	})

	g.Go(func() error {
		var err error
		rel.TriggerCommands, err = apilookups.E.TriggerCommands.DBGetRelations(ctx, stat, queries.TriggerCommands)
		return err
	})

	g.Go(func() error {
		var err error
		rel.StatusConditions, err = apilookups.E.StatusConditions.DBGetRelations(ctx, stat, queries.StatusConditions)
		return err
	})

	g.Go(func() error {
		var err error
		if queries.Properties != nil {
			rel.Properties, err = apilookups.E.Properties.DBGetRelations(ctx, stat, queries.Properties)
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
	changesOnly, err := apilookups.Q.Stats[ts.QpnChangesOnly].ParseBool(r)
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
