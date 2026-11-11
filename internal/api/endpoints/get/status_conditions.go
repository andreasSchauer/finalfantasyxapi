package get

import (
	"fmt"
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apiconvert"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
	"golang.org/x/sync/errgroup"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func StatusCondition(r *http.Request, i handlerinput.Main[seedtypes.StatusCondition, apitypes.StatusCondition, apiresource.Named], id int32) (apitypes.StatusCondition, error) {
	status, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.StatusCondition{}, err
	}

	rel, err := getStatusConditionRelationships(r, status)
	if err != nil {
		return apitypes.StatusCondition{}, err
	}

	response := apitypes.StatusCondition{
		ID:                      status.ID,
		Name:                    status.Name,
		Category:                status.Category,
		IsPermanent:             status.IsPermanent,
		Visualization:           status.Visualization,
		Effect:                  status.Effect,
		NullifyArmored:          status.NullifyArmored,
		RelatedStats:            handlerinput.NamesToNamedAPIResources(apilookups.E.Stats, status.RelatedStats),
		RemovedStatusConditions: handlerinput.NamesToNamedAPIResources(apilookups.E.StatusConditions, status.RemovedStatusConditions),
		AddedElemResist:         ah.ConvertObjPtr(status.AddedElemResist, apiconvert.ElemResist),
		StatChanges:             ah.ConvertObjSlice(status.StatChanges, apiconvert.StatChange),
		ModifierChanges:         ah.ConvertObjSlice(status.ModifierChanges, apiconvert.ModifierChange),
		AutoAbilities:           rel.AutoAbilities,
		InflictedBy:             rel.InflictedBy,
		RemovedBy:               rel.RemovedBy,
		MonstersResistance:      rel.MonstersResistance,
	}

	return response, nil
}

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
		rel.MonstersResistance, err = apilookups.E.Monsters.DBGetRelations(ctx, status, getStatusConditionResistingMonsterIDs(res.resistance))
		return err
	})

	g.Go(func() error {
		var err error
		rel.InflictedBy, err = getStatusInfliction(ctx, status, statusInteractionQueries{
			abilities: getStatusConditionAbilityIDsInflicted(res.minRate, res.maxRate),
		})
		return err
	})

	g.Go(func() error {
		var err error
		rel.RemovedBy, err = getStatusRemoval(ctx, status, statusInteractionQueries{
			abilities:        config.Cfg.DB.GetStatusConditionAbilityIDsRemoved,
			statusConditions: config.Cfg.DB.GetStatusConditionRemovedConditionIDs,
		})
		return err
	})

	err = g.Wait()
	if err != nil {
		return apitypes.StatusCondition{}, err
	}

	return rel, nil
}

type statusQueryResistance struct {
	resistance int32
	minRate    int32
	maxRate    int32
}

func getStatusQueryResistance(r *http.Request) (statusQueryResistance, error) {
	var res statusQueryResistance

	queryParamResistance := apilookups.Q.StatusConditions[ts.QpnResistance]
	resistance, err := queryParamResistance.ParseInt(r)
	if errs.ErrExceptEmptyQuery(err) {
		return statusQueryResistance{}, err
	}
	res.resistance = int32(resistance)

	queryParamMinRate := apilookups.Q.StatusConditions[ts.QpnInflictMin]
	minRate, err := queryParamMinRate.ParseInt(r)
	if errs.ErrExceptEmptyQuery(err) {
		return statusQueryResistance{}, err
	}
	res.minRate = int32(minRate)

	queryParamMaxRate := apilookups.Q.StatusConditions[ts.QpnInflictMax]
	maxRate, err := queryParamMaxRate.ParseInt(r)
	if errs.ErrExceptEmptyQuery(err) {
		return statusQueryResistance{}, err
	}
	res.maxRate = int32(maxRate)

	if res.minRate > res.maxRate {
		return statusQueryResistance{}, errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid use of parameter '%s'. '%s' can't be higher than '%s'", queryParamMinRate.Name, queryParamMinRate.Name, queryParamMaxRate.Name), nil)
	}

	return res, nil
}
