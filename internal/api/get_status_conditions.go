package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (srv *Server) getStatusCondition(r *http.Request, i handlerInput[seedtypes.StatusCondition, StatusCondition, NamedAPIResource], id int32) (StatusCondition, error) {
	status, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return StatusCondition{}, err
	}

	rel, err := getStatusConditionRelationships(srv, r, status)
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
		RelatedStats:            namesToNamedAPIResources(srv, E.Stats, status.RelatedStats),
		RemovedStatusConditions: namesToNamedAPIResources(srv, E.StatusConditions, status.RemovedStatusConditions),
		AddedElemResist:         convertObjPtr(srv, status.AddedElemResist, convertElemResist),
		StatChanges:             convertObjSlice(srv, status.StatChanges, convertStatChange),
		ModifierChanges:         convertObjSlice(srv, status.ModifierChanges, convertModifierChange),
		AutoAbilities:           rel.AutoAbilities,
		InflictedBy:             rel.InflictedBy,
		RemovedBy:               rel.RemovedBy,
		MonstersResistance:      rel.MonstersResistance,
	}

	return response, nil
}

func (srv *Server) retrieveStatusConditions(r *http.Request, i handlerInput[seedtypes.StatusCondition, StatusCondition, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(srv, r, ids, []IDFilter{
		i.queryEnumList(srv, r, enums.T.StatusConditionCategory, ids, ts.QpnCategory, srv.Cfg.DB.GetStatusConditionIDsByCategory),
	})
}
