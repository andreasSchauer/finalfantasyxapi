package get

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
	"golang.org/x/sync/errgroup"
)

func Element(r *http.Request, i handlerinput.Main[seedtypes.Element, apitypes.Element, apiresource.Named], id int32) (apitypes.Element, error) {
	element, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Element{}, err
	}

	response := apitypes.Element{
		ID:                 element.ID,
		Name:               element.Name,
		OppositeElement:    handlerinput.NamePtrToNamedAPIResPtr(i, element.OppositeElement, nil),
	}
	
	err = getElementRelationships(r, element, &response)
	if err != nil {
		return apitypes.Element{}, err
	}
	return response, nil
}

func getElementRelationships(r *http.Request, element seedtypes.Element, response *apitypes.Element) error {
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		response.StatusProtection, err = apilookups.E.StatusConditions.DBGetRelationPtr(ctx, element, config.Cfg.DB.GetElementStatusConditionID)
		return err
	})

	g.Go(func() error {
		var err error
		response.AutoAbilities, err = apilookups.E.AutoAbilities.DBGetRelations(ctx, element, config.Cfg.DB.GetElementAutoAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		response.PlayerAbilities, err = apilookups.E.PlayerAbilities.DBGetRelations(ctx, element, handlerinput.ToIntManyNull(config.Cfg.DB.GetElementPlayerAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		response.OverdriveAbilities, err = apilookups.E.OverdriveAbilities.DBGetRelations(ctx, element, handlerinput.ToIntManyNull(config.Cfg.DB.GetElementOverdriveAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		response.ItemAbilities, err = apilookups.E.ItemAbilities.DBGetRelations(ctx, element, handlerinput.ToIntManyNull(config.Cfg.DB.GetElementItemAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		response.EnemyAbilities, err = apilookups.E.EnemyAbilities.DBGetRelations(ctx, element, handlerinput.ToIntManyNull(config.Cfg.DB.GetElementEnemyAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		response.MonstersWeak, err = apilookups.E.Monsters.DBGetRelations(ctx, element, config.Cfg.DB.GetElementMonsterIDsWeak)
		return err
	})

	g.Go(func() error {
		var err error
		response.MonstersHalved, err = apilookups.E.Monsters.DBGetRelations(ctx, element, config.Cfg.DB.GetElementMonsterIDsHalved)
		return err
	})

	g.Go(func() error {
		var err error
		response.MonstersImmune, err = apilookups.E.Monsters.DBGetRelations(ctx, element, config.Cfg.DB.GetElementMonsterIDsImmune)
		return err
	})

	g.Go(func() error {
		var err error
		response.MonstersAbsorb, err = apilookups.E.Monsters.DBGetRelations(ctx, element, config.Cfg.DB.GetElementMonsterIDsAbsorb)
		return err
	})

	return g.Wait()
}
