package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
	"golang.org/x/sync/errgroup"
)

func getStatusConditionRelationships(cfg *Config, r *http.Request, status seeding.StatusCondition) (StatusCondition, error) {
	var rel StatusCondition
	g, ctx := errgroup.WithContext(r.Context())
	
	res, err := getStatusQueryResistance(cfg, r)
	if err != nil {
		return StatusCondition{}, err
	}

	g.Go(func() error {
		var err error
		rel.AutoAbilities, err = getResourcesDbItem(cfg, ctx, cfg.e.autoAbilities, status, cfg.db.GetStatusConditionAutoAbilityIDs)
		return err
	})
	
	g.Go(func() error {
		var err error
		rel.MonstersResistance, err = getResourcesDbItem(cfg, ctx, cfg.e.monsters, status, convGetStatusConditionResistingMonsterIDs(cfg, res.Resistance))
		return err
	})
	
	g.Go(func() error {
		var err error
		rel.InflictedBy, err = getStatusInfliction(cfg, ctx, status, StatusInteractionQueries{
			Abilities:        convGetStatusConditionAbilityIDsInflicted(cfg, res.MinRate, res.MaxRate),
			StatusConditions: cfg.db.GetStatusConditionInflictedDelayConditionIDs,
		})
		return err
	})
	
	g.Go(func() error {
		var err error
		rel.RemovedBy, err = getStatusRemoval(cfg, ctx, status, StatusInteractionQueries{
			Abilities:        cfg.db.GetStatusConditionAbilityIDsRemoved,
			StatusConditions: cfg.db.GetStatusConditionRemovedConditionIDs,
		})
		return err
	})
	
	err = g.Wait()
	if err != nil {
		return StatusCondition{}, err
	}

	return rel, nil
}


type StatusQueryResistance struct {
	Resistance 	int32
	MinRate		int32
	MaxRate		int32
}

func getStatusQueryResistance(cfg *Config, r *http.Request) (StatusQueryResistance, error) {
	var res StatusQueryResistance
	
	queryParamResistance := cfg.q.statusConditions[qpnResistance]
	resistance, err := parseIntQuery(r, queryParamResistance)
	if errExceptEmptyQuery(err) {
		return StatusQueryResistance{}, err
	}
	res.Resistance = int32(resistance)

	queryParamMinRate := cfg.q.statusConditions[qpnInflictMin]
	minRate, err := parseIntQuery(r, queryParamMinRate)
	if errExceptEmptyQuery(err) {
		return StatusQueryResistance{}, err
	}
	res.MinRate = int32(minRate)

	queryParamMaxRate := cfg.q.statusConditions[qpnInflictMax]
	maxRate, err := parseIntQuery(r, queryParamMaxRate)
	if errExceptEmptyQuery(err) {
		return StatusQueryResistance{}, err
	}
	res.MaxRate = int32(maxRate)

	if res.MinRate > res.MaxRate {
		return StatusQueryResistance{}, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid use of parameter '%s'. '%s' can't be higher than '%s'.", queryParamMinRate.Name, queryParamMinRate.Name, queryParamMaxRate.Name), nil)
	}

	return res, nil
}