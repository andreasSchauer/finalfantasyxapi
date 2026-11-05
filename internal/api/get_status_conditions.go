package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getStatusCondition(r *http.Request, i handlerInput[seedtypes.StatusCondition, StatusCondition, NamedAPIResource], id int32) (StatusCondition, error) {
	status, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return StatusCondition{}, err
	}

	rel, err := getStatusConditionRelationships(r, status)
	if err != nil {
		return StatusCondition{}, err
	}

	response := StatusCondition{
		ID:                      status.ID,
		Name:                    status.Name,
		Category:                status.Category,
		IsPermanent:             status.IsPermanent,
		Visualization:           status.Visualization,
		Effect:                  status.Effect,
		NullifyArmored:          status.NullifyArmored,
		RelatedStats:            namesToNamedAPIResources(E.Stats, status.RelatedStats),
		RemovedStatusConditions: namesToNamedAPIResources(E.StatusConditions, status.RemovedStatusConditions),
		AddedElemResist:         ConvertObjPtr(status.AddedElemResist, convertElemResist),
		StatChanges:             ConvertObjSlice(status.StatChanges, convertStatChange),
		ModifierChanges:         ConvertObjSlice(status.ModifierChanges, convertModifierChange),
		AutoAbilities:           rel.AutoAbilities,
		InflictedBy:             rel.InflictedBy,
		RemovedBy:               rel.RemovedBy,
		MonstersResistance:      rel.MonstersResistance,
	}

	return response, nil
}

func retrieveStatusConditions(r *http.Request, i handlerInput[seedtypes.StatusCondition, StatusCondition, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(r, ids, []IDFilter{
		i.queryEnumList(r, enums.T.StatusConditionCategory, ids, ts.QpnCategory, config.Cfg.DB.GetStatusConditionIDsByCategory),
	})
}
