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

	rel, err := getElementRelationships(r, element)
	if err != nil {
		return apitypes.Element{}, err
	}

	response := apitypes.Element{
		ID:                 element.ID,
		Name:               element.Name,
		OppositeElement:    handlerinput.NamePtrToNamedAPIResPtr(i, element.OppositeElement, nil),
		StatusProtection:   rel.StatusProtection,
		AutoAbilities:      rel.AutoAbilities,
		PlayerAbilities:    rel.PlayerAbilities,
		OverdriveAbilities: rel.OverdriveAbilities,
		ItemAbilities:      rel.ItemAbilities,
		EnemyAbilities:     rel.EnemyAbilities,
		MonstersWeak:       rel.MonstersWeak,
		MonstersHalved:     rel.MonstersHalved,
		MonstersImmune:     rel.MonstersImmune,
		MonstersAbsorb:     rel.MonstersAbsorb,
	}

	return response, nil
}

func getElementRelationships(r *http.Request, element seedtypes.Element) (apitypes.Element, error) {
	var rel apitypes.Element
	g, ctx := errgroup.WithContext(r.Context())

	g.Go(func() error {
		var err error
		rel.StatusProtection, err = apilookups.E.StatusConditions.DBGetRelationPtr(ctx, element, config.Cfg.DB.GetElementStatusConditionID)
		return err
	})

	g.Go(func() error {
		var err error
		rel.AutoAbilities, err = apilookups.E.AutoAbilities.DBGetRelations(ctx, element, config.Cfg.DB.GetElementAutoAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.PlayerAbilities, err = apilookups.E.PlayerAbilities.DBGetRelations(ctx, element, handlerinput.ToIntManyNull(config.Cfg.DB.GetElementPlayerAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.OverdriveAbilities, err = apilookups.E.OverdriveAbilities.DBGetRelations(ctx, element, handlerinput.ToIntManyNull(config.Cfg.DB.GetElementOverdriveAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.ItemAbilities, err = apilookups.E.ItemAbilities.DBGetRelations(ctx, element, handlerinput.ToIntManyNull(config.Cfg.DB.GetElementItemAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.EnemyAbilities, err = apilookups.E.EnemyAbilities.DBGetRelations(ctx, element, handlerinput.ToIntManyNull(config.Cfg.DB.GetElementEnemyAbilityIDs))
		return err
	})

	g.Go(func() error {
		var err error
		rel.MonstersWeak, err = apilookups.E.Monsters.DBGetRelations(ctx, element, config.Cfg.DB.GetElementMonsterIDsWeak)
		return err
	})

	g.Go(func() error {
		var err error
		rel.MonstersHalved, err = apilookups.E.Monsters.DBGetRelations(ctx, element, config.Cfg.DB.GetElementMonsterIDsHalved)
		return err
	})

	g.Go(func() error {
		var err error
		rel.MonstersImmune, err = apilookups.E.Monsters.DBGetRelations(ctx, element, config.Cfg.DB.GetElementMonsterIDsImmune)
		return err
	})

	g.Go(func() error {
		var err error
		rel.MonstersAbsorb, err = apilookups.E.Monsters.DBGetRelations(ctx, element, config.Cfg.DB.GetElementMonsterIDsAbsorb)
		return err
	})

	err := g.Wait()
	if err != nil {
		return apitypes.Element{}, err
	}

	return rel, nil
}
