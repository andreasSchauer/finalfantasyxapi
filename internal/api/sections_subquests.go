package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type SubquestSimple struct {
	ID          int32                   `json:"id"`
	URL         string                  `json:"url"`
	Name        string                  `json:"name"`
	Completions []QuestCompletionSimple `json:"completions"`
}

func (s SubquestSimple) GetURL() string {
	return s.URL
}

type QuestCompletionSimple struct {
	Condition string           `json:"condition"`
	Areas     []string         `json:"areas"`
	Reward    ItemAmountSimple `json:"reward"`
}

func convertQuestCompletionSimple(cfg *Config, qc seeding.QuestCompletion) QuestCompletionSimple {
	return QuestCompletionSimple{
		Condition: qc.Condition,
		Areas:     locAreaStrings(cfg, qc.Areas),
		Reward:    convertItemAmountSimple(cfg, qc.Reward),
	}
}

func createSubquestSimple(cfg *Config, _ *http.Request, id int32) (SimpleResource, error) {
	i := cfg.e.subquests
	subquest, _ := seeding.GetResourceByID(id, i.objLookupID)

	subquestSimple := SubquestSimple{
		ID:          subquest.ID,
		URL:         createResourceURL(cfg, i.endpoint, id),
		Name:        subquest.Name,
		Completions: convertObjSlice(cfg, subquest.Completions, convertQuestCompletionSimple),
	}

	return subquestSimple, nil
}
