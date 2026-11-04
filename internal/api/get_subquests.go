package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (s *Server) getSubquest(r *http.Request, i handlerInput[seedtypes.Subquest, Subquest, QuestAPIResource], id int32) (Subquest, error) {
	subquest, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Subquest{}, err
	}

	response := Subquest{
		ID:              subquest.ID,
		Name:            subquest.Name,
		UntypedQuest:    idToQuestAPIResource(s, E.Quests, subquest.Quest.ID),
		ParentSidequest: idToQuestAPIResource(s, E.Sidequests, subquest.SidequestID),
		Availability:    subquest.Availability,
		IsRepeatable:    subquest.IsRepeatable,
		Completion:      convertQuestCompletion(s, *subquest.Completion),
	}

	if response.ParentSidequest.Sidequest == "monster arena" {
		response.ArenaCreation = namePtrToNamedAPIResPtr(s, E.ArenaCreations, &subquest.Name, nil)
	}

	return response, nil
}

func (s *Server) retrieveSubquests(r *http.Request, i handlerInput[seedtypes.Subquest, Subquest, QuestAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(s, r, ids, []IDFilter{
		i.queryBool(r, ids, ts.QpnRepeatable, s.Cfg.DB.GetSubquestIDsByRepeatable),
	})
}
