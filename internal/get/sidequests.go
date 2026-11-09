package get

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiconv"
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/assemble"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func GetSidequest(r *http.Request, i handlerinput.Main[seedtypes.Sidequest, apitypes.Sidequest, apiresource.Quest], id int32) (apitypes.Sidequest, error) {
	sidequest, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Sidequest{}, err
	}

	rel, err := assemble.GetSidequestRelationships(r, sidequest)
	if err != nil {
		return apitypes.Sidequest{}, err
	}

	response := apitypes.Sidequest{
		ID:           sidequest.ID,
		Name:         sidequest.Name,
		UntypedQuest: handlerinput.IDToQuestAPIResource(apilookups.E.Quests, sidequest.Quest.ID),
		Availability: sidequest.Availability,
		IsRepeatable: sidequest.IsRepeatable,
		Completion:   ah.ConvertObjPtr(sidequest.Completion, apiconv.ConvertQuestCompletion),
		Subquests:    rel.Subquests,
	}

	return response, nil
}

func RetrieveSidequests(r *http.Request, i handlerinput.Main[seedtypes.Sidequest, apitypes.Sidequest, apiresource.Quest]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{})
}
