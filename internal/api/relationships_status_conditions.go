package api

import (
	"fmt"
	"net/http"

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

	inflictedBy, err := getStatusInfliction(cfg, r, status, StatusInteractionQueries{
		Abilities:        convGetStatusConditionAbilityIDsInflicted(cfg, minRate32, maxRate32),
		StatusConditions: cfg.db.GetStatusConditionInflictedDelayConditionIDs,
	})
	if err != nil {
		return StatusCondition{}, err
	}

	removedBy, err := getStatusRemoval(cfg, r, status, StatusInteractionQueries{
		Abilities:        cfg.db.GetStatusConditionAbilityIDsRemoved,
		StatusConditions: cfg.db.GetStatusConditionRemovedConditionIDs,
	})
	if err != nil {
		return StatusCondition{}, err
	}

	rel := StatusCondition{
		AutoAbilities:      autoAbilities,
		InflictedBy:        inflictedBy,
		RemovedBy:          removedBy,
		MonstersResistance: monsters,
	}

	return rel, nil
}

