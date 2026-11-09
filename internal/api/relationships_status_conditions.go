package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
	"golang.org/x/sync/errgroup"
)

func getStatusConditionRelationships(r *http.Request, status seedtypes.StatusCondition) (apitypes.StatusCondition, error) {
	var rel apitypes.StatusCondition
	g, ctx := errgroup.WithContext(r.Context())

	res, err := getStatusQueryResistance(r)
	if err != nil {
		return apitypes.StatusCondition{}, err
	}

	g.Go(func() error {
		var err error
		rel.AutoAbilities, err = apilookups.E.AutoAbilities.DBGetRelations(ctx, status, config.Cfg.DB.GetStatusConditionAutoAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.MonstersResistance, err = apilookups.E.Monsters.DBGetRelations(ctx, status, GetStatusConditionResistingMonsterIDs(res.Resistance))
		return err
	})

	g.Go(func() error {
		var err error
		rel.InflictedBy, err = getStatusInfliction(ctx, status, StatusInteractionQueries{
			Abilities: GetStatusConditionAbilityIDsInflicted(res.MinRate, res.MaxRate),
		})
		return err
	})

	g.Go(func() error {
		var err error
		rel.RemovedBy, err = getStatusRemoval(ctx, status, StatusInteractionQueries{
			Abilities:        config.Cfg.DB.GetStatusConditionAbilityIDsRemoved,
			StatusConditions: config.Cfg.DB.GetStatusConditionRemovedConditionIDs,
		})
		return err
	})

	err = g.Wait()
	if err != nil {
		return apitypes.StatusCondition{}, err
	}

	return rel, nil
}

type StatusQueryResistance struct {
	Resistance int32
	MinRate    int32
	MaxRate    int32
}

func getStatusQueryResistance(r *http.Request) (StatusQueryResistance, error) {
	var res StatusQueryResistance

	queryParamResistance := apilookups.Q.StatusConditions[ts.QpnResistance]
	resistance, err := queryParamResistance.ParseInt(r)
	if errs.ErrExceptEmptyQuery(err) {
		return StatusQueryResistance{}, err
	}
	res.Resistance = int32(resistance)

	queryParamMinRate := apilookups.Q.StatusConditions[ts.QpnInflictMin]
	minRate, err := queryParamMinRate.ParseInt(r)
	if errs.ErrExceptEmptyQuery(err) {
		return StatusQueryResistance{}, err
	}
	res.MinRate = int32(minRate)

	queryParamMaxRate := apilookups.Q.StatusConditions[ts.QpnInflictMax]
	maxRate, err := queryParamMaxRate.ParseInt(r)
	if errs.ErrExceptEmptyQuery(err) {
		return StatusQueryResistance{}, err
	}
	res.MaxRate = int32(maxRate)

	if res.MinRate > res.MaxRate {
		return StatusQueryResistance{}, errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid use of parameter '%s'. '%s' can't be higher than '%s'", queryParamMinRate.Name, queryParamMinRate.Name, queryParamMaxRate.Name), nil)
	}

	return res, nil
}
