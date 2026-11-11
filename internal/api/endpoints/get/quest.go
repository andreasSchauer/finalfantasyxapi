package get

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apiconvert"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Quest(r *http.Request, i handlerinput.Main[seedtypes.Quest, apitypes.Quest, apiresource.Quest], id int32) (apitypes.Quest, error) {
	quest, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Quest{}, err
	}

	response := apitypes.Quest{
		ID:           quest.ID,
		Name:         quest.Name,
		Type:         quest.Type,
		TypedQuest:   apiconvert.QuestToQuestAPIResource(quest),
		Availability: quest.Availability,
		IsRepeatable: quest.IsRepeatable,
		Completion:   ah.ConvertObjPtr(quest.Completion, apiconvert.QuestCompletion),
	}

	return response, nil
}
