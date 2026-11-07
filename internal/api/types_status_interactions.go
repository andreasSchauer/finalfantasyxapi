package api

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

type StatusInteractionQueries struct {
	Abilities        handlerinput.DBQueryIntMany
	StatusConditions handlerinput.DBQueryIntMany
}

func getStatusInfliction(ctx context.Context, status seedtypes.StatusCondition, queries StatusInteractionQueries) (*apitypes.StatusInfliction, error) {
	var statusInfliction apitypes.StatusInfliction

	abilities, err := E.Abilities.DBGetRelations(ctx, status, queries.Abilities)
	if err != nil {
		return nil, err
	}
	statusInfliction = populateStatusInfliction(abilities)

	if statusInfliction.IsZero() {
		return nil, nil
	}

	return &statusInfliction, nil
}

func getStatusRemoval(ctx context.Context, status seedtypes.StatusCondition, queries StatusInteractionQueries) (*apitypes.StatusRemoval, error) {
	abilities, err := E.Abilities.DBGetRelations(ctx, status, queries.Abilities)
	if err != nil {
		return nil, err
	}
	statusRemoval := populateStatusRemoval(abilities)

	statusRemoval.StatusConditions, err = E.StatusConditions.DBGetRelations(ctx, status, queries.StatusConditions)
	if err != nil {
		return nil, err
	}

	if statusRemoval.IsZero() {
		return nil, nil
	}

	return &statusRemoval, nil
}

func populateStatusInfliction(abilities []apiresource.TypedAPIResource) apitypes.StatusInfliction {
	infliction := apitypes.StatusInfliction{
		PlayerAbilities:    []apiresource.NamedAPIResource{},
		OverdriveAbilities: []apiresource.NamedAPIResource{},
		ItemAbilities:      []apiresource.NamedAPIResource{},
		MiscAbilities:      []apiresource.NamedAPIResource{},
		EnemyAbilities:     []apiresource.NamedAPIResource{},
	}

	for _, ability := range abilities {
		obj := seedtypes.LookupObject{
			Name:    ability.Name,
			Version: ability.Version,
		}
		switch ability.Type {
		case string(database.AbilityTypePlayerAbility):
			playerAbility, _ := lookup.GetResource(obj, lookup.L.PlayerAbilities)
			res := handlerinput.NameToNamedAPIResource(E.PlayerAbilities, playerAbility.Name, playerAbility.Version)
			infliction.PlayerAbilities = append(infliction.PlayerAbilities, res)

		case string(database.AbilityTypeOverdriveAbility):
			overdriveAbility, _ := lookup.GetResource(obj, lookup.L.OverdriveAbilities)
			res := handlerinput.NameToNamedAPIResource(E.OverdriveAbilities, overdriveAbility.Name, overdriveAbility.Version)
			infliction.OverdriveAbilities = append(infliction.OverdriveAbilities, res)

		case string(database.AbilityTypeItemAbility):
			itemAbility, _ := lookup.GetResource(ability.Name, lookup.L.ItemAbilities)
			res := handlerinput.NameToNamedAPIResource(E.ItemAbilities, itemAbility.Name, itemAbility.Version)
			infliction.ItemAbilities = append(infliction.ItemAbilities, res)

		case string(database.AbilityTypeMiscAbility):
			miscAbility, _ := lookup.GetResource(obj, lookup.L.MiscAbilities)
			res := handlerinput.NameToNamedAPIResource(E.MiscAbilities, miscAbility.Name, miscAbility.Version)
			infliction.MiscAbilities = append(infliction.MiscAbilities, res)

		case string(database.AbilityTypeEnemyAbility):
			enemyAbility, _ := lookup.GetResource(obj, lookup.L.EnemyAbilities)
			res := handlerinput.NameToNamedAPIResource(E.EnemyAbilities, enemyAbility.Name, enemyAbility.Version)
			infliction.EnemyAbilities = append(infliction.EnemyAbilities, res)
		}
	}

	return infliction
}

func populateStatusRemoval(abilities []apiresource.TypedAPIResource) apitypes.StatusRemoval {
	removal := apitypes.StatusRemoval{
		PlayerAbilities:    []apiresource.NamedAPIResource{},
		OverdriveAbilities: []apiresource.NamedAPIResource{},
		ItemAbilities:      []apiresource.NamedAPIResource{},
		EnemyAbilities:     []apiresource.NamedAPIResource{},
	}

	for _, ability := range abilities {
		obj := seedtypes.LookupObject{
			Name:    ability.Name,
			Version: ability.Version,
		}
		switch ability.Type {
		case string(database.AbilityTypePlayerAbility):
			playerAbility, _ := lookup.GetResource(obj, lookup.L.PlayerAbilities)
			res := handlerinput.NameToNamedAPIResource(E.PlayerAbilities, playerAbility.Name, playerAbility.Version)
			removal.PlayerAbilities = append(removal.PlayerAbilities, res)

		case string(database.AbilityTypeOverdriveAbility):
			overdriveAbility, _ := lookup.GetResource(obj, lookup.L.OverdriveAbilities)
			res := handlerinput.NameToNamedAPIResource(E.OverdriveAbilities, overdriveAbility.Name, overdriveAbility.Version)
			removal.OverdriveAbilities = append(removal.OverdriveAbilities, res)

		case string(database.AbilityTypeItemAbility):
			itemAbility, _ := lookup.GetResource(ability.Name, lookup.L.ItemAbilities)
			res := handlerinput.NameToNamedAPIResource(E.ItemAbilities, itemAbility.Name, itemAbility.Version)
			removal.ItemAbilities = append(removal.ItemAbilities, res)

		case string(database.AbilityTypeEnemyAbility):
			enemyAbility, _ := lookup.GetResource(obj, lookup.L.EnemyAbilities)
			res := handlerinput.NameToNamedAPIResource(E.EnemyAbilities, enemyAbility.Name, enemyAbility.Version)
			removal.EnemyAbilities = append(removal.EnemyAbilities, res)
		}
	}

	return removal
}
