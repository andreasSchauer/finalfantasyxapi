package get

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
	"golang.org/x/sync/errgroup"
)

func GetStat(r *http.Request, i handlerinput.Main[seedtypes.Stat, apitypes.Stat, apiresource.Named], id int32) (apitypes.Stat, error) {
	stat, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Stat{}, err
	}

	rel, err := getStatRelationships(r, stat)
	if err != nil {
		return apitypes.Stat{}, err
	}

	response := apitypes.Stat{
		ID:                 stat.ID,
		Name:               stat.Name,
		Effect:             stat.Effect,
		MinVal:             stat.MinVal,
		MaxVal:             stat.MaxVal,
		MaxVal2:            stat.MaxVal2,
		ActivationSphere:   handlerinput.NameToNamedAPIResource(apilookups.E.Spheres, stat.ActivationSphere, nil),
		Spheres:            rel.Spheres,
		AutoAbilities:      rel.AutoAbilities,
		PlayerAbilities:    rel.PlayerAbilities,
		OverdriveAbilities: rel.OverdriveAbilities,
		ItemAbilities:      rel.ItemAbilities,
		TriggerCommands:    rel.TriggerCommands,
		StatusConditions:   rel.StatusConditions,
		Properties:         rel.Properties,
	}

	return response, nil
}

type StatQueries struct {
	AutoAbilities      handlerinput.DBQueryIntMany
	PlayerAbilities    handlerinput.DBQueryIntMany
	OverdriveAbilities handlerinput.DBQueryIntMany
	ItemAbilities      handlerinput.DBQueryIntMany
	TriggerCommands    handlerinput.DBQueryIntMany
	StatusConditions   handlerinput.DBQueryIntMany
	Properties         handlerinput.DBQueryIntMany
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
