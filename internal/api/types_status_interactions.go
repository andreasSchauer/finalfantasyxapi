package api

import (
	"context"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type StatusInfliction struct {
	PlayerAbilities    []NamedAPIResource `json:"player_abilities"`
	OverdriveAbilities []NamedAPIResource `json:"overdrive_abilities"`
	ItemAbilities      []NamedAPIResource `json:"item_abilities"`
	MiscAbilities      []NamedAPIResource `json:"misc_abilities"`
	EnemyAbilities     []NamedAPIResource `json:"enemy_abilities"`
}

func (s StatusInfliction) IsZero() bool {
	return len(s.PlayerAbilities) == 0 &&
		len(s.OverdriveAbilities) == 0 &&
		len(s.ItemAbilities) == 0 &&
		len(s.MiscAbilities) == 0 &&
		len(s.EnemyAbilities) == 0
}

type StatusRemoval struct {
	PlayerAbilities    []NamedAPIResource `json:"player_abilities"`
	OverdriveAbilities []NamedAPIResource `json:"overdrive_abilities"`
	ItemAbilities      []NamedAPIResource `json:"item_abilities"`
	EnemyAbilities     []NamedAPIResource `json:"enemy_abilities"`
	StatusConditions   []NamedAPIResource `json:"status_conditions"`
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

func getStatusInfliction(cfg *Config, ctx context.Context, status seeding.StatusCondition, queries StatusInteractionQueries) (*StatusInfliction, error) {
	var statusInfliction StatusInfliction

	abilities, err := getResourcesDBItem(cfg, ctx, cfg.E.Abilities, status, queries.Abilities)
	if err != nil {
		return nil, err
	}
	statusInfliction = populateStatusInfliction(cfg, abilities)

	if statusInfliction.IsZero() {
		return nil, nil
	}

	return &statusInfliction, nil
}

func getStatusRemoval(cfg *Config, ctx context.Context, status seeding.StatusCondition, queries StatusInteractionQueries) (*StatusRemoval, error) {
	abilities, err := getResourcesDBItem(cfg, ctx, cfg.E.Abilities, status, queries.Abilities)
	if err != nil {
		return nil, err
	}
	statusRemoval := populateStatusRemoval(cfg, abilities)

	statusRemoval.StatusConditions, err = getResourcesDBItem(cfg, ctx, cfg.E.StatusConditions, status, queries.StatusConditions)
	if err != nil {
		return nil, err
	}

	if statusRemoval.IsZero() {
		return nil, nil
	}

	return &statusRemoval, nil
}

func populateStatusInfliction(cfg *Config, abilities []TypedAPIResource) StatusInfliction {
	infliction := StatusInfliction{
		PlayerAbilities:    []NamedAPIResource{},
		OverdriveAbilities: []NamedAPIResource{},
		ItemAbilities:      []NamedAPIResource{},
		MiscAbilities:      []NamedAPIResource{},
		EnemyAbilities:     []NamedAPIResource{},
	}

	for _, ability := range abilities {
		obj := seeding.LookupObject{
			Name:    ability.Name,
			Version: ability.Version,
		}
		switch ability.Type {
		case string(database.AbilityTypePlayerAbility):
			playerAbility, _ := seeding.GetResource(obj, cfg.L.PlayerAbilities)
			res := nameToNamedAPIResource(cfg, cfg.E.PlayerAbilities, playerAbility.Name, playerAbility.Version)
			infliction.PlayerAbilities = append(infliction.PlayerAbilities, res)

		case string(database.AbilityTypeOverdriveAbility):
			overdriveAbility, _ := seeding.GetResource(obj, cfg.L.OverdriveAbilities)
			res := nameToNamedAPIResource(cfg, cfg.E.OverdriveAbilities, overdriveAbility.Name, overdriveAbility.Version)
			infliction.OverdriveAbilities = append(infliction.OverdriveAbilities, res)

		case string(database.AbilityTypeItemAbility):
			itemAbility, _ := seeding.GetResource(ability.Name, cfg.L.ItemAbilities)
			res := nameToNamedAPIResource(cfg, cfg.E.ItemAbilities, itemAbility.Name, itemAbility.Version)
			infliction.ItemAbilities = append(infliction.ItemAbilities, res)

		case string(database.AbilityTypeMiscAbility):
			miscAbility, _ := seeding.GetResource(obj, cfg.L.MiscAbilities)
			res := nameToNamedAPIResource(cfg, cfg.E.MiscAbilities, miscAbility.Name, miscAbility.Version)
			infliction.MiscAbilities = append(infliction.MiscAbilities, res)

		case string(database.AbilityTypeEnemyAbility):
			enemyAbility, _ := seeding.GetResource(obj, cfg.L.EnemyAbilities)
			res := nameToNamedAPIResource(cfg, cfg.E.EnemyAbilities, enemyAbility.Name, enemyAbility.Version)
			infliction.EnemyAbilities = append(infliction.EnemyAbilities, res)
		}
	}

	return infliction
}

func populateStatusRemoval(cfg *Config, abilities []TypedAPIResource) StatusRemoval {
	removal := StatusRemoval{
		PlayerAbilities:    []NamedAPIResource{},
		OverdriveAbilities: []NamedAPIResource{},
		ItemAbilities:      []NamedAPIResource{},
		EnemyAbilities:     []NamedAPIResource{},
	}

	for _, ability := range abilities {
		obj := seeding.LookupObject{
			Name:    ability.Name,
			Version: ability.Version,
		}
		switch ability.Type {
		case string(database.AbilityTypePlayerAbility):
			playerAbility, _ := seeding.GetResource(obj, cfg.L.PlayerAbilities)
			res := nameToNamedAPIResource(cfg, cfg.E.PlayerAbilities, playerAbility.Name, playerAbility.Version)
			removal.PlayerAbilities = append(removal.PlayerAbilities, res)

		case string(database.AbilityTypeOverdriveAbility):
			overdriveAbility, _ := seeding.GetResource(obj, cfg.L.OverdriveAbilities)
			res := nameToNamedAPIResource(cfg, cfg.E.OverdriveAbilities, overdriveAbility.Name, overdriveAbility.Version)
			removal.OverdriveAbilities = append(removal.OverdriveAbilities, res)

		case string(database.AbilityTypeItemAbility):
			itemAbility, _ := seeding.GetResource(ability.Name, cfg.L.ItemAbilities)
			res := nameToNamedAPIResource(cfg, cfg.E.ItemAbilities, itemAbility.Name, itemAbility.Version)
			removal.ItemAbilities = append(removal.ItemAbilities, res)

		case string(database.AbilityTypeEnemyAbility):
			enemyAbility, _ := seeding.GetResource(obj, cfg.L.EnemyAbilities)
			res := nameToNamedAPIResource(cfg, cfg.E.EnemyAbilities, enemyAbility.Name, enemyAbility.Version)
			removal.EnemyAbilities = append(removal.EnemyAbilities, res)
		}
	}

	return removal
}
