package api

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getStatusCondition(r *http.Request, i handlerInput[seedtypes.StatusCondition, apitypes.StatusCondition, apiresource.NamedAPIResource], id int32) (apitypes.StatusCondition, error) {
	status, err := i.queryParamsVerifyGet(r, id)
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
		RelatedStats:            namesToNamedAPIResources(E.Stats, status.RelatedStats),
		RemovedStatusConditions: namesToNamedAPIResources(E.StatusConditions, status.RemovedStatusConditions),
		AddedElemResist:         ah.ConvertObjPtr(status.AddedElemResist, convertElemResist),
		StatChanges:             ah.ConvertObjSlice(status.StatChanges, convertStatChange),
		ModifierChanges:         ah.ConvertObjSlice(status.ModifierChanges, convertModifierChange),
		AutoAbilities:           rel.AutoAbilities,
		InflictedBy:             rel.InflictedBy,
		RemovedBy:               rel.RemovedBy,
		MonstersResistance:      rel.MonstersResistance,
	}

	return response, nil
}

func retrieveStatusConditions(r *http.Request, i handlerInput[seedtypes.StatusCondition, apitypes.StatusCondition, apiresource.NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(r, ids, []IDFilter{
		i.queryEnumList(r, enums.T.StatusConditionCategory, ids, ts.QpnCategory, config.Cfg.DB.GetStatusConditionIDsByCategory),
	})
}
