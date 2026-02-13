package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type SubquestSub struct {
	ID          int32                `json:"id"`
	URL         string               `json:"url"`
	Name        string               `json:"name"`
	Completions []QuestCompletionSub `json:"completions"`
}

func (s SubquestSub) GetURL() string {
	return s.URL
}

type QuestCompletionSub struct {
	Condition string        `json:"condition"`
	Areas     []string      `json:"areas"`
	Reward    ItemAmountSub `json:"reward"`
}

func convertQuestCompletionSub(cfg *Config, qc seeding.QuestCompletion) QuestCompletionSub {
	return QuestCompletionSub{
		Condition: qc.Condition,
		Areas:     locAreaStrings(cfg, qc.Areas),
		Reward:    convertSubItemAmount(cfg, qc.Reward),
	}
}

func createSubquestSub(cfg *Config, _ *http.Request, id int32) (SubResource, error) {
	i := cfg.e.subquests
	subquest, _ := seeding.GetResourceByID(id, i.objLookupID)

	subquestSub := SubquestSub{
		ID:          subquest.ID,
		URL:         createResourceURL(cfg, i.endpoint, id),
		Name:        subquest.Name,
		Completions: convertObjSlice(cfg, subquest.Completions, convertQuestCompletionSub),
	}

	return subquestSub, nil
}
