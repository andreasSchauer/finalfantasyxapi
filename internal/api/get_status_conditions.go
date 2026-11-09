package api

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getStatusCondition(r *http.Request, i handlerinput.Main[seedtypes.StatusCondition, apitypes.StatusCondition, apiresource.Named], id int32) (apitypes.StatusCondition, error) {
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
		AddedElemResist:         ah.ConvertObjPtr(status.AddedElemResist, ConvertElemResist),
		StatChanges:             ah.ConvertObjSlice(status.StatChanges, ConvertStatChange),
		ModifierChanges:         ah.ConvertObjSlice(status.ModifierChanges, ConvertModifierChange),
		AutoAbilities:           rel.AutoAbilities,
		InflictedBy:             rel.InflictedBy,
		RemovedBy:               rel.RemovedBy,
		MonstersResistance:      rel.MonstersResistance,
	}

	return response, nil
}

func retrieveStatusConditions(r *http.Request, i handlerinput.Main[seedtypes.StatusCondition, apitypes.StatusCondition, apiresource.Named]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryEnumList(r, enums.T.StatusConditionCategory, ids, ts.QpnCategory, config.Cfg.DB.GetStatusConditionIDsByCategory),
	})
}
