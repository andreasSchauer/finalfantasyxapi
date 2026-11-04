package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (srv *Server) getSubquest(r *http.Request, i handlerInput[seedtypes.Subquest, Subquest, QuestAPIResource], id int32) (Subquest, error) {
	subquest, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Subquest{}, err
	}

	response := Subquest{
		ID:              subquest.ID,
		Name:            subquest.Name,
		UntypedQuest:    idToQuestAPIResource(srv, E.Quests, subquest.Quest.ID),
		ParentSidequest: idToQuestAPIResource(srv, E.Sidequests, subquest.SidequestID),
		Availability:    subquest.Availability,
		IsRepeatable:    subquest.IsRepeatable,
		Completion:      convertQuestCompletion(srv, *subquest.Completion),
	}

	if response.ParentSidequest.Sidequest == "monster arena" {
		response.ArenaCreation = namePtrToNamedAPIResPtr(srv, E.ArenaCreations, &subquest.Name, nil)
	}

	return response, nil
}

func (srv *Server) retrieveSubquests(r *http.Request, i handlerInput[seedtypes.Subquest, Subquest, QuestAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(srv, r, ids, []IDFilter{
		i.queryBool(r, ids, ts.QpnRepeatable, srv.Cfg.DB.GetSubquestIDsByRepeatable),
	})
}
