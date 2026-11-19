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

func Stat(r *http.Request, i handlerinput.Main[seedtypes.Stat, apitypes.Stat, apiresource.Named], id int32) (apitypes.Stat, error) {
	stat, err := i.VerifyQueriesGet(r, id)
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
	}

	err = getStatRelationships(r, stat, &response)
	if err != nil {
		return apitypes.Stat{}, err
	}

	return response, nil
}

type statQueries struct {
	autoAbilities      handlerinput.DBQueryIntMany
	playerAbilities    handlerinput.DBQueryIntMany
	overdriveAbilities handlerinput.DBQueryIntMany
	itemAbilities      handlerinput.DBQueryIntMany
	triggerCommands    handlerinput.DBQueryIntMany
	statusConditions   handlerinput.DBQueryIntMany
	properties         handlerinput.DBQueryIntMany
}

func getStatRelationships(r *http.Request, stat seedtypes.Stat, response *apitypes.Stat) error {
	g, ctx := errgroup.WithContext(r.Context())

	queries, err := getStatQueries(r)
	if err != nil {
		return err
	}

	g.Go(func() error {
		var err error
		response.Spheres, err = apilookups.E.Spheres.DBGetRelations(ctx, stat, config.Cfg.DB.GetStatSphereIDs)
		return err
	})

	g.Go(func() error {
		var err error
		response.AutoAbilities, err = apilookups.E.AutoAbilities.DBGetRelations(ctx, stat, queries.autoAbilities)
		return err
	})

	g.Go(func() error {
		var err error
		response.PlayerAbilities, err = apilookups.E.PlayerAbilities.DBGetRelations(ctx, stat, queries.playerAbilities)
		return err
	})

	g.Go(func() error {
		var err error
		response.OverdriveAbilities, err = apilookups.E.OverdriveAbilities.DBGetRelations(ctx, stat, queries.overdriveAbilities)
		return err
	})

	g.Go(func() error {
		var err error
		response.ItemAbilities, err = apilookups.E.ItemAbilities.DBGetRelations(ctx, stat, queries.itemAbilities)
		return err
	})

	g.Go(func() error {
		var err error
		response.TriggerCommands, err = apilookups.E.TriggerCommands.DBGetRelations(ctx, stat, queries.triggerCommands)
		return err
	})

	g.Go(func() error {
		var err error
		response.StatusConditions, err = apilookups.E.StatusConditions.DBGetRelations(ctx, stat, queries.statusConditions)
		return err
	})

	g.Go(func() error {
		var err error
		if queries.properties != nil {
			response.Properties, err = apilookups.E.Properties.DBGetRelations(ctx, stat, queries.properties)
			return err
		}
		return nil
	})

	return g.Wait()
}

func getStatQueries(r *http.Request) (statQueries, error) {
	changesOnly, err := apilookups.Q.Stats[ts.QpnChangesOnly].ParseBool(r)
	if errs.ErrExceptEmptyQuery(err) {
		return statQueries{}, err
	}

	var queries statQueries

	if changesOnly {
		queries = statQueries{
			autoAbilities:      config.Cfg.DB.GetStatAutoAbilityIDsStatChange,
			playerAbilities:    config.Cfg.DB.GetStatPlayerAbilityIDsStatChange,
			overdriveAbilities: config.Cfg.DB.GetStatOverdriveAbilityIDsStatChange,
			itemAbilities:      config.Cfg.DB.GetStatItemAbilityIDsStatChange,
			triggerCommands:    config.Cfg.DB.GetStatTriggerCommandIDsStatChange,
			statusConditions:   config.Cfg.DB.GetStatStatusConditionIDsStatChange,
			properties:         nil,
		}
	} else {
		queries = statQueries{
			autoAbilities:      config.Cfg.DB.GetStatAutoAbilityIDs,
			playerAbilities:    config.Cfg.DB.GetStatPlayerAbilityIDs,
			overdriveAbilities: config.Cfg.DB.GetStatOverdriveAbilityIDs,
			itemAbilities:      config.Cfg.DB.GetStatItemAbilityIDs,
			triggerCommands:    config.Cfg.DB.GetStatTriggerCommandIDs,
			statusConditions:   config.Cfg.DB.GetStatStatusConditionIDs,
			properties:         config.Cfg.DB.GetStatPropertyIDs,
		}
	}

	return queries, nil
}
