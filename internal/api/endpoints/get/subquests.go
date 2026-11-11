package get

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apiconvert"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Subquest(r *http.Request, i handlerinput.Main[seedtypes.Subquest, apitypes.Subquest, apiresource.Quest], id int32) (apitypes.Subquest, error) {
	subquest, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Subquest{}, err
	}

	response := apitypes.Subquest{
		ID:              subquest.ID,
		Name:            subquest.Name,
		UntypedQuest:    handlerinput.IDToQuestAPIResource(apilookups.E.Quests, subquest.Quest.ID),
		ParentSidequest: handlerinput.IDToQuestAPIResource(apilookups.E.Sidequests, subquest.SidequestID),
		Availability:    subquest.Availability,
		IsRepeatable:    subquest.IsRepeatable,
		Completion:      apiconvert.QuestCompletion(*subquest.Completion),
	}

	return completeSubquestResponse(response), nil
}
