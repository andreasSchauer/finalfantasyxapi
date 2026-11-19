package get

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
	"golang.org/x/sync/errgroup"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Modifier(r *http.Request, i handlerinput.Main[seedtypes.Modifier, apitypes.Modifier, apiresource.Named], id int32) (apitypes.Modifier, error) {
	modifier, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Modifier{}, err
	}
	
	response := apitypes.Modifier{
		ID:                 modifier.ID,
		Name:               modifier.Name,
		Effect:             modifier.Effect,
		Category:           modifier.Category,
		DefaultValue:       modifier.DefaultValue,
	}

	err = getModifierRelationships(r, modifier, &response)
	if err != nil {
		return apitypes.Modifier{}, err
	}
	
	return response, nil
}

func getModifierRelationships(r *http.Request, modifier seedtypes.Modifier, response *apitypes.Modifier) error {
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		response.AutoAbilities, err = apilookups.E.AutoAbilities.DBGetRelations(ctx, modifier, config.Cfg.DB.GetModifierAutoAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		response.PlayerAbilities, err = apilookups.E.PlayerAbilities.DBGetRelations(ctx, modifier, config.Cfg.DB.GetModifierPlayerAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		response.OverdriveAbilities, err = apilookups.E.OverdriveAbilities.DBGetRelations(ctx, modifier, config.Cfg.DB.GetModifierOverdriveAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		response.ItemAbilities, err = apilookups.E.ItemAbilities.DBGetRelations(ctx, modifier, config.Cfg.DB.GetModifierItemAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		response.TriggerCommands, err = apilookups.E.TriggerCommands.DBGetRelations(ctx, modifier, config.Cfg.DB.GetModifierTriggerCommandIDs)
		return err
	})

	g.Go(func() error {
		var err error
		response.EnemyAbilities, err = apilookups.E.EnemyAbilities.DBGetRelations(ctx, modifier, config.Cfg.DB.GetModifierEnemyAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		response.StatusConditions, err = apilookups.E.StatusConditions.DBGetRelations(ctx, modifier, config.Cfg.DB.GetModifierStatusConditionIDs)
		return err
	})

	g.Go(func() error {
		var err error
		response.Properties, err = apilookups.E.Properties.DBGetRelations(ctx, modifier, config.Cfg.DB.GetModifierPropertyIDs)
		return err
	})

	return g.Wait()
}