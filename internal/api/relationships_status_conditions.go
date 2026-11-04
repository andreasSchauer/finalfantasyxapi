package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
	"golang.org/x/sync/errgroup"
)

func getStatusConditionRelationships(server *Server, r *http.Request, status seedtypes.StatusCondition) (StatusCondition, error) {
	var rel StatusCondition
	g, ctx := errgroup.WithContext(r.Context())

	res, err := getStatusQueryResistance(r)
	if err != nil {
		return StatusCondition{}, err
	}

	g.Go(func() error {
		var err error
		rel.AutoAbilities, err = E.AutoAbilities.dbGetRelations(server, ctx, status, server.Cfg.DB.GetStatusConditionAutoAbilityIDs)
		return err
	})

	g.Go(func() error {
		var err error
		rel.MonstersResistance, err = E.Monsters.dbGetRelations(server, ctx, status, convGetStatusConditionResistingMonsterIDs(server, res.Resistance))
		return err
	})

	g.Go(func() error {
		var err error
		rel.InflictedBy, err = getStatusInfliction(server, ctx, status, StatusInteractionQueries{
			Abilities: convGetStatusConditionAbilityIDsInflicted(server, res.MinRate, res.MaxRate),
		})
		return err
	})

	g.Go(func() error {
		var err error
		rel.RemovedBy, err = getStatusRemoval(server, ctx, status, StatusInteractionQueries{
			Abilities:        server.Cfg.DB.GetStatusConditionAbilityIDsRemoved,
			StatusConditions: server.Cfg.DB.GetStatusConditionRemovedConditionIDs,
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
	Resistance int32
	MinRate    int32
	MaxRate    int32
}

func getStatusQueryResistance(r *http.Request) (StatusQueryResistance, error) {
	var res StatusQueryResistance

	queryParamResistance := Q.statusConditions[ts.QpnResistance]
	resistance, err := queryParamResistance.parseInt(r)
	if errExceptEmptyQuery(err) {
		return StatusQueryResistance{}, err
	}
	res.Resistance = int32(resistance)

	queryParamMinRate := Q.statusConditions[ts.QpnInflictMin]
	minRate, err := queryParamMinRate.parseInt(r)
	if errExceptEmptyQuery(err) {
		return StatusQueryResistance{}, err
	}
	res.MinRate = int32(minRate)

	queryParamMaxRate := Q.statusConditions[ts.QpnInflictMax]
	maxRate, err := queryParamMaxRate.parseInt(r)
	if errExceptEmptyQuery(err) {
		return StatusQueryResistance{}, err
	}
	res.MaxRate = int32(maxRate)

	if res.MinRate > res.MaxRate {
		return StatusQueryResistance{}, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid use of parameter '%s'. '%s' can't be higher than '%s'", queryParamMinRate.Name, queryParamMinRate.Name, queryParamMaxRate.Name), nil)
	}

	return res, nil
}
