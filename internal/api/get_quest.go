package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getQuest(r *http.Request, i handlerInput[seeding.Quest, Quest, QuestAPIResource, QuestApiResourceList], id int32) (Quest, error) {
	quest, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return Quest{}, err
	}

	response := Quest{
		ID:         	quest.ID,
		Name:       	quest.Name,
		Type: 			newNamedAPIResourceFromEnum(cfg, cfg.e.questType.endpoint, string(quest.Type), cfg.t.QuestType),
		TypedQuest: 	questToQuestAPIResource(cfg, quest),
		IsPostAirship: 	quest.IsPostAirship,
	}

	return response, nil
}

func (cfg *Config) retrieveQuests(r *http.Request, i handlerInput[seeding.Quest, Quest, QuestAPIResource, QuestApiResourceList]) (QuestApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return QuestApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[QuestAPIResource]{
		frl(typeQuery(cfg, r, i, cfg.t.QuestType, resources, "type", cfg.db.GetQuestIDsByType)),
		frl(boolQuery(cfg, r, i, resources, "post_airship", cfg.db.GetQuestIDsByPostAirship)),
	})
}
