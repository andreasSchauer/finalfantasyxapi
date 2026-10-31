package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getQuest(r *http.Request, i handlerInput[seeding.Quest, Quest, QuestAPIResource, QuestAPIResourceList], id int32) (Quest, error) {
	quest, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return Quest{}, err
	}

	response := Quest{
		ID:           quest.ID,
		Name:         quest.Name,
		Type:         quest.Type,
		TypedQuest:   questToQuestAPIResource(cfg, quest),
		Availability: quest.Availability,
		IsRepeatable: quest.IsRepeatable,
		Completion:   convertObjPtr(cfg, quest.Completion, convertQuestCompletion),
	}

	return response, nil
}

func (cfg *Config) retrieveQuests(r *http.Request, i handlerInput[seeding.Quest, Quest, QuestAPIResource, QuestAPIResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []IDFilter{
		enumQuery(r, i, cfg.T.QuestType, ids, QpnType, cfg.db.GetQuestIDsByType),
		boolQuery(r, i, ids, QpnRepeatable, cfg.db.GetQuestIDsByRepeatable),
	})
}
