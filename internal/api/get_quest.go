package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getQuest(r *http.Request, i handlerInput[seeding.Quest, Quest, QuestAPIResource], id int32) (Quest, error) {
	quest, err := i.queryParamsVerifyGet(r, id)
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

func (cfg *Config) retrieveQuests(r *http.Request, i handlerInput[seeding.Quest, Quest, QuestAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(cfg, r, ids, []IDFilter{
		i.queryEnum(r, cfg.T.QuestType, ids, QpnType, cfg.db.GetQuestIDsByType),
		i.queryBool(r, ids, QpnRepeatable, cfg.db.GetQuestIDsByRepeatable),
	})
}
