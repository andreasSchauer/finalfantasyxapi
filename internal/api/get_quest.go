package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getQuest(r *http.Request, i handlerInput[seeding.Quest, Quest, QuestAPIResource, QuestApiResourceList], id int32) (Quest, error) {
	quest, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return Quest{}, err
	}

	response := Quest{
		ID:           quest.ID,
		Name:         quest.Name,
		Type:         enumToNamedAPIResource(cfg, cfg.e.questType.endpoint, string(quest.Type), cfg.t.QuestType),
		TypedQuest:   questToQuestAPIResource(cfg, quest),
		Availability: enumToNamedAPIResource(cfg, cfg.e.availabilityType.endpoint, quest.Availability, cfg.t.AvailabilityType),
		IsRepeatable: quest.IsRepeatable,
		Completion:   convertObjPtr(cfg, quest.Completion, convertQuestCompletion),
	}

	return response, nil
}

func (cfg *Config) retrieveQuests(r *http.Request, i handlerInput[seeding.Quest, Quest, QuestAPIResource, QuestApiResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []filteredIdList{
		fidl(enumQuery(r, i, cfg.t.QuestType, ids, qpnType, cfg.db.GetQuestIDsByType)),
		fidl(boolQuery(r, i, ids, qpnRepeatable, cfg.db.GetQuestIDsByRepeatable)),
	})
}
