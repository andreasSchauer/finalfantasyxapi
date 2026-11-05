package api

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

type StatusInfliction struct {
	PlayerAbilities    []apiresource.NamedAPIResource `json:"player_abilities"`
	OverdriveAbilities []apiresource.NamedAPIResource `json:"overdrive_abilities"`
	ItemAbilities      []apiresource.NamedAPIResource `json:"item_abilities"`
	MiscAbilities      []apiresource.NamedAPIResource `json:"misc_abilities"`
	EnemyAbilities     []apiresource.NamedAPIResource `json:"enemy_abilities"`
}

func (s StatusInfliction) IsZero() bool {
	return len(s.PlayerAbilities) == 0 &&
		len(s.OverdriveAbilities) == 0 &&
		len(s.ItemAbilities) == 0 &&
		len(s.MiscAbilities) == 0 &&
		len(s.EnemyAbilities) == 0
}

type StatusRemoval struct {
	PlayerAbilities    []apiresource.NamedAPIResource `json:"player_abilities"`
	OverdriveAbilities []apiresource.NamedAPIResource `json:"overdrive_abilities"`
	ItemAbilities      []apiresource.NamedAPIResource `json:"item_abilities"`
	EnemyAbilities     []apiresource.NamedAPIResource `json:"enemy_abilities"`
	StatusConditions   []apiresource.NamedAPIResource `json:"status_conditions"`
}

func (s StatusRemoval) IsZero() bool {
	return len(s.PlayerAbilities) == 0 &&
		len(s.OverdriveAbilities) == 0 &&
		len(s.ItemAbilities) == 0 &&
		len(s.EnemyAbilities) == 0 &&
		len(s.StatusConditions) == 0
}

type StatusInteractionQueries struct {
	Abilities        DBQueryIntMany
	StatusConditions DBQueryIntMany
}

func getStatusInfliction(ctx context.Context, status seedtypes.StatusCondition, queries StatusInteractionQueries) (*StatusInfliction, error) {
	var statusInfliction StatusInfliction

	abilities, err := E.Abilities.dbGetRelations(ctx, status, queries.Abilities)
	if err != nil {
		return nil, err
	}
	statusInfliction = populateStatusInfliction(abilities)

	if statusInfliction.IsZero() {
		return nil, nil
	}

	return &statusInfliction, nil
}

func getStatusRemoval(ctx context.Context, status seedtypes.StatusCondition, queries StatusInteractionQueries) (*StatusRemoval, error) {
	abilities, err := E.Abilities.dbGetRelations(ctx, status, queries.Abilities)
	if err != nil {
		return nil, err
	}
	statusRemoval := populateStatusRemoval(abilities)

	statusRemoval.StatusConditions, err = E.StatusConditions.dbGetRelations(ctx, status, queries.StatusConditions)
	if err != nil {
		return nil, err
	}

	if statusRemoval.IsZero() {
		return nil, nil
	}

	return &statusRemoval, nil
}

func populateStatusInfliction(abilities []apiresource.TypedAPIResource) StatusInfliction {
	infliction := StatusInfliction{
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
			res := nameToNamedAPIResource(E.PlayerAbilities, playerAbility.Name, playerAbility.Version)
			infliction.PlayerAbilities = append(infliction.PlayerAbilities, res)

		case string(database.AbilityTypeOverdriveAbility):
			overdriveAbility, _ := lookup.GetResource(obj, lookup.L.OverdriveAbilities)
			res := nameToNamedAPIResource(E.OverdriveAbilities, overdriveAbility.Name, overdriveAbility.Version)
			infliction.OverdriveAbilities = append(infliction.OverdriveAbilities, res)

		case string(database.AbilityTypeItemAbility):
			itemAbility, _ := lookup.GetResource(ability.Name, lookup.L.ItemAbilities)
			res := nameToNamedAPIResource(E.ItemAbilities, itemAbility.Name, itemAbility.Version)
			infliction.ItemAbilities = append(infliction.ItemAbilities, res)

		case string(database.AbilityTypeMiscAbility):
			miscAbility, _ := lookup.GetResource(obj, lookup.L.MiscAbilities)
			res := nameToNamedAPIResource(E.MiscAbilities, miscAbility.Name, miscAbility.Version)
			infliction.MiscAbilities = append(infliction.MiscAbilities, res)

		case string(database.AbilityTypeEnemyAbility):
			enemyAbility, _ := lookup.GetResource(obj, lookup.L.EnemyAbilities)
			res := nameToNamedAPIResource(E.EnemyAbilities, enemyAbility.Name, enemyAbility.Version)
			infliction.EnemyAbilities = append(infliction.EnemyAbilities, res)
		}
	}

	return infliction
}

func populateStatusRemoval(abilities []apiresource.TypedAPIResource) StatusRemoval {
	removal := StatusRemoval{
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
			res := nameToNamedAPIResource(E.PlayerAbilities, playerAbility.Name, playerAbility.Version)
			removal.PlayerAbilities = append(removal.PlayerAbilities, res)

		case string(database.AbilityTypeOverdriveAbility):
			overdriveAbility, _ := lookup.GetResource(obj, lookup.L.OverdriveAbilities)
			res := nameToNamedAPIResource(E.OverdriveAbilities, overdriveAbility.Name, overdriveAbility.Version)
			removal.OverdriveAbilities = append(removal.OverdriveAbilities, res)

		case string(database.AbilityTypeItemAbility):
			itemAbility, _ := lookup.GetResource(ability.Name, lookup.L.ItemAbilities)
			res := nameToNamedAPIResource(E.ItemAbilities, itemAbility.Name, itemAbility.Version)
			removal.ItemAbilities = append(removal.ItemAbilities, res)

		case string(database.AbilityTypeEnemyAbility):
			enemyAbility, _ := lookup.GetResource(obj, lookup.L.EnemyAbilities)
			res := nameToNamedAPIResource(E.EnemyAbilities, enemyAbility.Name, enemyAbility.Version)
			removal.EnemyAbilities = append(removal.EnemyAbilities, res)
		}
	}

	return removal
}
