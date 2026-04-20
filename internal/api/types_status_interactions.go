package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)


type StatusInfliction struct {
	PlayerAbilities      []NamedAPIResource `json:"player_abilities"`
	OverdriveAbilities   []NamedAPIResource `json:"overdrive_abilities"`
	ItemAbilities        []NamedAPIResource `json:"item_abilities"`
	UnspecifiedAbilities []NamedAPIResource `json:"unspecified_abilities"`
	EnemyAbilities       []NamedAPIResource `json:"enemy_abilities"`
	StatusConditions     []NamedAPIResource `json:"status_conditions"`
}

func (s StatusInfliction) IsZero() bool {
	return len(s.PlayerAbilities) == 0 &&
		len(s.OverdriveAbilities) == 0 &&
		len(s.ItemAbilities) == 0 &&
		len(s.UnspecifiedAbilities) == 0 &&
		len(s.EnemyAbilities) == 0 &&
		len(s.StatusConditions) == 0
}

type StatusRemoval struct {
	PlayerAbilities      []NamedAPIResource `json:"player_abilities"`
	OverdriveAbilities   []NamedAPIResource `json:"overdrive_abilities"`
	ItemAbilities        []NamedAPIResource `json:"item_abilities"`
	EnemyAbilities       []NamedAPIResource `json:"enemy_abilities"`
	StatusConditions     []NamedAPIResource `json:"status_conditions"`
}

func (s StatusRemoval) IsZero() bool {
	return len(s.PlayerAbilities) == 0 &&
		len(s.OverdriveAbilities) == 0 &&
		len(s.ItemAbilities) == 0 &&
		len(s.EnemyAbilities) == 0 &&
		len(s.StatusConditions) == 0
}

type StatusInteractionQueries struct {
	Abilities        DbQueryIntMany
	StatusConditions DbQueryIntMany
}

func getStatusInfliction(cfg *Config, r *http.Request, status seeding.StatusCondition, queries StatusInteractionQueries) (*StatusInfliction, error) {
	abilities, err := getResourcesDbItem(cfg, r, cfg.e.abilities, status, queries.Abilities)
	if err != nil {
		return nil, err
	}
	statusInfliction := populateStatusInfliction(cfg, abilities)

	statusInfliction.StatusConditions, err = getResourcesDbItem(cfg, r, cfg.e.statusConditions, status, queries.StatusConditions)
	if err != nil {
		return nil, err
	}

	if statusInfliction.IsZero() {
		return nil, nil
	}

	return &statusInfliction, nil
}

func getStatusRemoval(cfg *Config, r *http.Request, status seeding.StatusCondition, queries StatusInteractionQueries) (*StatusRemoval, error) {
	abilities, err := getResourcesDbItem(cfg, r, cfg.e.abilities, status, queries.Abilities)
	if err != nil {
		return nil, err
	}
	statusRemoval := populateStatusRemoval(cfg, abilities)

	statusRemoval.StatusConditions, err = getResourcesDbItem(cfg, r, cfg.e.statusConditions, status, queries.StatusConditions)
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
		PlayerAbilities:      []NamedAPIResource{},
		OverdriveAbilities:   []NamedAPIResource{},
		ItemAbilities:        []NamedAPIResource{},
		UnspecifiedAbilities: []NamedAPIResource{},
		EnemyAbilities:       []NamedAPIResource{},
	}

	for _, ability := range abilities {
		obj := seeding.LookupObject{
			Name:    ability.Name,
			Version: ability.Version,
		}
		switch ability.Type {
		case string(database.AbilityTypePlayerAbility):
			playerAbility, _ := seeding.GetResource(obj, cfg.l.PlayerAbilities)
			res := nameToNamedAPIResource(cfg, cfg.e.playerAbilities, playerAbility.Name, playerAbility.Version)
			infliction.PlayerAbilities = append(infliction.PlayerAbilities, res)

		case string(database.AbilityTypeOverdriveAbility):
			overdriveAbility, _ := seeding.GetResource(obj, cfg.l.OverdriveAbilities)
			res := nameToNamedAPIResource(cfg, cfg.e.overdriveAbilities, overdriveAbility.Name, overdriveAbility.Version)
			infliction.OverdriveAbilities = append(infliction.OverdriveAbilities, res)

		case string(database.AbilityTypeItemAbility):
			itemAbility, _ := seeding.GetResource(ability.Name, cfg.l.ItemAbilities)
			res := nameToNamedAPIResource(cfg, cfg.e.itemAbilities, itemAbility.Name, itemAbility.Version)
			infliction.ItemAbilities = append(infliction.ItemAbilities, res)

		case string(database.AbilityTypeUnspecifiedAbility):
			unspecifiedAbility, _ := seeding.GetResource(obj, cfg.l.UnspecifiedAbilities)
			res := nameToNamedAPIResource(cfg, cfg.e.unspecifiedAbilities, unspecifiedAbility.Name, unspecifiedAbility.Version)
			infliction.UnspecifiedAbilities = append(infliction.UnspecifiedAbilities, res)

		case string(database.AbilityTypeEnemyAbility):
			enemyAbility, _ := seeding.GetResource(obj, cfg.l.EnemyAbilities)
			res := nameToNamedAPIResource(cfg, cfg.e.enemyAbilities, enemyAbility.Name, enemyAbility.Version)
			infliction.EnemyAbilities = append(infliction.EnemyAbilities, res)
		}
	}

	return infliction
}


func populateStatusRemoval(cfg *Config, abilities []TypedAPIResource) StatusRemoval {
	removal := StatusRemoval{
		PlayerAbilities:      []NamedAPIResource{},
		OverdriveAbilities:   []NamedAPIResource{},
		ItemAbilities:        []NamedAPIResource{},
		EnemyAbilities:       []NamedAPIResource{},
	}

	for _, ability := range abilities {
		obj := seeding.LookupObject{
			Name:    ability.Name,
			Version: ability.Version,
		}
		switch ability.Type {
		case string(database.AbilityTypePlayerAbility):
			playerAbility, _ := seeding.GetResource(obj, cfg.l.PlayerAbilities)
			res := nameToNamedAPIResource(cfg, cfg.e.playerAbilities, playerAbility.Name, playerAbility.Version)
			removal.PlayerAbilities = append(removal.PlayerAbilities, res)

		case string(database.AbilityTypeOverdriveAbility):
			overdriveAbility, _ := seeding.GetResource(obj, cfg.l.OverdriveAbilities)
			res := nameToNamedAPIResource(cfg, cfg.e.overdriveAbilities, overdriveAbility.Name, overdriveAbility.Version)
			removal.OverdriveAbilities = append(removal.OverdriveAbilities, res)

		case string(database.AbilityTypeItemAbility):
			itemAbility, _ := seeding.GetResource(ability.Name, cfg.l.ItemAbilities)
			res := nameToNamedAPIResource(cfg, cfg.e.itemAbilities, itemAbility.Name, itemAbility.Version)
			removal.ItemAbilities = append(removal.ItemAbilities, res)

		case string(database.AbilityTypeEnemyAbility):
			enemyAbility, _ := seeding.GetResource(obj, cfg.l.EnemyAbilities)
			res := nameToNamedAPIResource(cfg, cfg.e.enemyAbilities, enemyAbility.Name, enemyAbility.Version)
			removal.EnemyAbilities = append(removal.EnemyAbilities, res)
		}
	}

	return removal
}
