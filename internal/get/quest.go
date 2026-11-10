package get

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiconv"
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func GetQuest(r *http.Request, i handlerinput.Main[seedtypes.Quest, apitypes.Quest, apiresource.Quest], id int32) (apitypes.Quest, error) {
	quest, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Quest{}, err
	}

	response := apitypes.Quest{
		ID:           quest.ID,
		Name:         quest.Name,
		Type:         quest.Type,
		TypedQuest:   apiconv.QuestToQuestAPIResource(quest),
		Availability: quest.Availability,
		IsRepeatable: quest.IsRepeatable,
		Completion:   ah.ConvertObjPtr(quest.Completion, apiconv.ConvertQuestCompletion),
	}

	return response, nil
}
