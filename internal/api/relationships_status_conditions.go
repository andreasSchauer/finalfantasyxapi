package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)


func getStatusConditionRelationships(cfg *Config, r *http.Request, status seeding.StatusCondition) (StatusCondition, error) {
	queryParamResistance := cfg.q.statusConditions["resistance"]
	queryParamMinRate := cfg.q.statusConditions["inflict_min"]
	queryParamMaxRate := cfg.q.statusConditions["inflict_max"]
	
	resistance, err := parseIntQuery(r, queryParamResistance)
	if errIsNotEmptyQuery(err) {
		return StatusCondition{}, err
	}
	resistance32 := int32(resistance)

	minRate, err := parseIntQuery(r, queryParamMinRate)
	if errIsNotEmptyQuery(err) {
		return StatusCondition{}, err
	}
	minRate32 := int32(minRate)

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
		PlayerAbilities: convGetStatusConditionPlayerAbilityIDsInflicted(cfg, minRate32, maxRate32),
		OverdriveAbilities: convGetStatusConditionOverdriveAbilityIDsInflicted(cfg, minRate32, maxRate32),
		ItemAbilities: convGetStatusConditionItemAbilityIDsInflicted(cfg, minRate32, maxRate32),
		UnspecifiedAbilities: convGetStatusConditionUnspecifiedAbilityIDsInflicted(cfg, minRate32, maxRate32),
		EnemyAbilities: convGetStatusConditionEnemyAbilityIDsInflicted(cfg, minRate32, maxRate32),
	})
	if err != nil {
		return StatusCondition{}, err
	}

	removedBy, err := getStatusInteractions(cfg, r, status, StatusInteractionQueries{
		PlayerAbilities: cfg.db.GetStatusConditionPlayerAbilityIDsRemoved,
		OverdriveAbilities: cfg.db.GetStatusConditionOverdriveAbilityIDsRemoved,
		ItemAbilities: cfg.db.GetStatusConditionItemAbilityIDsRemoved,
		UnspecifiedAbilities: cfg.db.GetStatusConditionUnspecifiedAbilityIDsRemoved,
		EnemyAbilities: cfg.db.GetStatusConditionEnemyAbilityIDsRemoved,
		StatusConditions: cfg.db.GetStatusConditionRemovedConditionIDs,
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
	StatusConditions		[]NamedAPIResource	`json:"status_conditions,omitempty"`
}

type StatusInteractionQueries struct {
	PlayerAbilities			DbQueryIntMany
	OverdriveAbilities		DbQueryIntMany
	ItemAbilities			DbQueryIntMany
	UnspecifiedAbilities	DbQueryIntMany
	EnemyAbilities			DbQueryIntMany
	StatusConditions 		DbQueryIntMany
}

func getStatusInteractions(cfg *Config, r *http.Request, status seeding.StatusCondition, queries StatusInteractionQueries) (StatusInteractions, error) {
	var err error
	statusInteractions := StatusInteractions{}

	statusInteractions.PlayerAbilities, err = getResourcesDbItem(cfg, r, cfg.e.playerAbilities, status, queries.PlayerAbilities)
	if err != nil {
		return StatusInteractions{}, err
	}

	statusInteractions.OverdriveAbilities, err = getResourcesDbItem(cfg, r, cfg.e.overdriveAbilities, status, queries.OverdriveAbilities)
	if err != nil {
		return StatusInteractions{}, err
	}

	statusInteractions.ItemAbilities, err = getResourcesDbItem(cfg, r, cfg.e.itemAbilities, status, queries.ItemAbilities)
	if err != nil {
		return StatusInteractions{}, err
	}

	statusInteractions.UnspecifiedAbilities, err = getResourcesDbItem(cfg, r, cfg.e.unspecifiedAbilities, status, queries.UnspecifiedAbilities)
	if err != nil {
		return StatusInteractions{}, err
	}

	statusInteractions.EnemyAbilities, err = getResourcesDbItem(cfg, r, cfg.e.enemyAbilities, status, queries.EnemyAbilities)
	if err != nil {
		return StatusInteractions{}, err
	}

	if queries.StatusConditions != nil {
		statusInteractions.StatusConditions, err = getResourcesDbItem(cfg, r, cfg.e.statusConditions, status, queries.StatusConditions)
		if err != nil {
			return StatusInteractions{}, err
		}
	}

	return statusInteractions, nil
}