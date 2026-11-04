package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (s *Server) getStatusCondition(r *http.Request, i handlerInput[seedtypes.StatusCondition, StatusCondition, NamedAPIResource], id int32) (StatusCondition, error) {
	status, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return StatusCondition{}, err
	}

	rel, err := getStatusConditionRelationships(s, r, status)
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
		RelatedStats:            namesToNamedAPIResources(s, E.Stats, status.RelatedStats),
		RemovedStatusConditions: namesToNamedAPIResources(s, E.StatusConditions, status.RemovedStatusConditions),
		AddedElemResist:         convertObjPtr(s, status.AddedElemResist, convertElemResist),
		StatChanges:             convertObjSlice(s, status.StatChanges, convertStatChange),
		ModifierChanges:         convertObjSlice(s, status.ModifierChanges, convertModifierChange),
		AutoAbilities:           rel.AutoAbilities,
		InflictedBy:             rel.InflictedBy,
		RemovedBy:               rel.RemovedBy,
		MonstersResistance:      rel.MonstersResistance,
	}

	return response, nil
}

func (s *Server) retrieveStatusConditions(r *http.Request, i handlerInput[seedtypes.StatusCondition, StatusCondition, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(s, r, ids, []IDFilter{
		i.queryEnumList(s, r, enums.T.StatusConditionCategory, ids, ts.QpnCategory, s.Cfg.DB.GetStatusConditionIDsByCategory),
	})
}
