package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)


func getStatusConditionRelationships(cfg *Config, r *http.Request, status seeding.StatusCondition) (StatusCondition, error) {
	queryParamResistance := cfg.q.statusConditions["resistance"]
	resistance, err := parseIntQuery(r, queryParamResistance)
	if errIsNotEmptyQuery(err) {
		return StatusCondition{}, err
	}
	resistance32 := int32(resistance)
	
	queryParamMinRate := cfg.q.statusConditions["inflict_min"]
	minRate, err := parseIntQuery(r, queryParamMinRate)
	if errIsNotEmptyQuery(err) {
		return StatusCondition{}, err
	}
	minRate32 := int32(minRate)
	
	queryParamMaxRate := cfg.q.statusConditions["inflict_max"]
	maxRate, err := parseIntQuery(r, queryParamMaxRate)
	if errIsNotEmptyQuery(err) {
		return StatusCondition{}, err
	}
	maxRate32 := int32(maxRate)
	
	if minRate > maxRate {
		return StatusCondition{}, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid use of parameter '%s'. '%s' can't be higher than '%s'.", queryParamMinRate.Name, queryParamMinRate.Name, queryParamMaxRate.Name), nil)
	}
	
	autoAbilities, err := getResourcesDbItem(cfg, r, cfg.e.autoAbilities, status, cfg.db.GetStatusConditionAutoAbilityIDs)
	if err != nil {
		return StatusCondition{}, err
	}

	monsters, err := getResourcesDbItem(cfg, r, cfg.e.monsters, status, convGetStatusConditionResistingMonsterIDs(cfg, resistance32))
	if err != nil {
		return StatusCondition{}, err
	}

	inflictedBy, err := getStatusInteractions(cfg, r, status, StatusInteractionQueries{
		Abilities: 			convGetStatusConditionAbilityIDsInflicted(cfg, minRate32, maxRate32),
		StatusConditions: 	cfg.db.GetStatusConditionInflictedDelayConditionIDs,
	})
	if err != nil {
		return StatusCondition{}, err
	}

	removedBy, err := getStatusInteractions(cfg, r, status, StatusInteractionQueries{
		Abilities: 			cfg.db.GetStatusConditionAbilityIDsRemoved,
		StatusConditions: 	cfg.db.GetStatusConditionRemovedConditionIDs,
	})
	if err != nil {
		return StatusCondition{}, err
	}

	rel := StatusCondition{
		AutoAbilities: 		autoAbilities,
		InflictedBy: 		inflictedBy,
		RemovedBy: 			removedBy,
		MonstersResistance: monsters,
	}

	return rel, nil
}



type StatusInteractions struct {
	PlayerAbilities			[]NamedAPIResource	`json:"player_abilities"`
	OverdriveAbilities		[]NamedAPIResource	`json:"overdrive_abilities"`
	ItemAbilities			[]NamedAPIResource	`json:"item_abilities"`
	UnspecifiedAbilities	[]NamedAPIResource	`json:"unspecified_abilities"`
	EnemyAbilities			[]NamedAPIResource	`json:"enemy_abilities"`
	StatusConditions		[]NamedAPIResource	`json:"status_conditions"`
}

func (s StatusInteractions) IsZero() bool {
	return 	len(s.PlayerAbilities) == 0 &&
			len(s.OverdriveAbilities) == 0 &&
			len(s.ItemAbilities) == 0 &&
			len(s.UnspecifiedAbilities) == 0 &&
			len(s.EnemyAbilities) == 0 &&
			len(s.StatusConditions) == 0
}

type StatusInteractionQueries struct {
	Abilities			DbQueryIntMany
	StatusConditions 		DbQueryIntMany
}

func getStatusInteractions(cfg *Config, r *http.Request, status seeding.StatusCondition, queries StatusInteractionQueries) (*StatusInteractions, error) {
	abilities, err := getResourcesDbItem(cfg, r, cfg.e.abilities, status, queries.Abilities)
	if err != nil {
		return nil, err
	}
	statusInteractions := populateStatusInteractions(cfg, abilities)

	statusInteractions.StatusConditions, err = getResourcesDbItem(cfg, r, cfg.e.statusConditions, status, queries.StatusConditions)
	if err != nil {
		return nil, err
	}

	if statusInteractions.IsZero() {
		return nil, nil
	}

	return &statusInteractions, nil
}

func populateStatusInteractions(cfg *Config, abilities []TypedAPIResource) StatusInteractions {
	ints := StatusInteractions{
		PlayerAbilities: []NamedAPIResource{},
		OverdriveAbilities: []NamedAPIResource{},
		ItemAbilities: []NamedAPIResource{},
		UnspecifiedAbilities: []NamedAPIResource{},
		EnemyAbilities: []NamedAPIResource{},
	}

	for _, ability := range abilities {
		obj := seeding.LookupObject{
			Name: 	ability.Name,
			Version: ability.Version,
		}
		switch ability.Type {
		case string(database.AbilityTypePlayerAbility):
			playerAbility, _ := seeding.GetResource(obj, cfg.l.PlayerAbilities)
			res := nameToNamedAPIResource(cfg, cfg.e.playerAbilities, playerAbility.Name, playerAbility.Version)
			ints.PlayerAbilities = append(ints.PlayerAbilities, res)

		case string(database.AbilityTypeOverdriveAbility):
			overdriveAbility, _ := seeding.GetResource(obj, cfg.l.OverdriveAbilities)
			res := nameToNamedAPIResource(cfg, cfg.e.overdriveAbilities, overdriveAbility.Name, overdriveAbility.Version)
			ints.OverdriveAbilities = append(ints.OverdriveAbilities, res)

		case string(database.AbilityTypeItemAbility):
			itemAbility, _ := seeding.GetResource(ability.Name, cfg.l.ItemAbilities)
			res := nameToNamedAPIResource(cfg, cfg.e.itemAbilities, itemAbility.Name, itemAbility.Version)
			ints.ItemAbilities = append(ints.ItemAbilities, res)

		case string(database.AbilityTypeUnspecifiedAbility):
			unspecifiedAbility, _ := seeding.GetResource(obj, cfg.l.UnspecifiedAbilities)
			res := nameToNamedAPIResource(cfg, cfg.e.unspecifiedAbilities, unspecifiedAbility.Name, unspecifiedAbility.Version)
			ints.UnspecifiedAbilities = append(ints.UnspecifiedAbilities, res)

		case string(database.AbilityTypeEnemyAbility):
			enemyAbility, _ := seeding.GetResource(obj, cfg.l.EnemyAbilities)
			res := nameToNamedAPIResource(cfg, cfg.e.enemyAbilities, enemyAbility.Name, enemyAbility.Version)
			ints.EnemyAbilities = append(ints.EnemyAbilities, res)
		}
	}
	
	return ints
}