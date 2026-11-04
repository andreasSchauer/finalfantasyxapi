package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func (s *Server) getSidequest(r *http.Request, i handlerInput[seedtypes.Sidequest, Sidequest, QuestAPIResource], id int32) (Sidequest, error) {
	sidequest, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Sidequest{}, err
	}

	subquests, err := E.Subquests.dbGetRelations(s, r.Context(), sidequest, s.Cfg.DB.GetSidequestSubquestIDs)
	if err != nil {
		return Sidequest{}, err
	}

	response := Sidequest{
		ID:           sidequest.ID,
		Name:         sidequest.Name,
		UntypedQuest: idToQuestAPIResource(s, E.Quests, sidequest.Quest.ID),
		Availability: sidequest.Availability,
		IsRepeatable: sidequest.IsRepeatable,
		Completion:   convertObjPtr(s, sidequest.Completion, convertQuestCompletion),
		Subquests:    subquests,
	}

	return response, nil
}

func (s *Server) retrieveSidequests(r *http.Request, i handlerInput[seedtypes.Sidequest, Sidequest, QuestAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(s, r, ids, []IDFilter{})
}
