package main

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)


type SubquestSub struct {
	ID			int32					`json:"id"`
	URL			string					`json:"url"`
	Name		string					`json:"name"`
	Completions	[]QuestCompletionSub 	`json:"completions"`
}

func (s SubquestSub) GetSectionName() string {
	return "monster-formations"
}

func (s SubquestSub) GetURL() string {
	return s.URL
}

type QuestCompletionSub struct {
	Condition	string
	Areas		[]string
	Reward		ItemAmountSub
}

func convertQuestCompletionSub(cfg *Config, qc seeding.QuestCompletion) QuestCompletionSub {
	return QuestCompletionSub{
		Condition: 	qc.Condition,
		Areas: 		locAreaStrings(cfg, qc.Locations),
		Reward: 	convertSubItemAmount(cfg, qc.Reward),
	}
}


func handleSubquestsSection(cfg *Config, _ *http.Request, dbIDs []int32) ([]SubResource, error) {
	i := cfg.e.subquests
	subquests := []SubquestSub{}

	for _, subquestID := range dbIDs {
		subquest, _ := seeding.GetResourceByID(subquestID, i.objLookupID)

		subquestSub := SubquestSub{
			ID:          subquest.ID,
			URL:         createResourceURL(cfg, i.endpoint, subquestID),
			Name:		 subquest.Name,
			Completions: convertObjSlice(cfg, subquest.Completions, convertQuestCompletionSub),
		}

		subquests = append(subquests, subquestSub)
	}

	return toSubResourceSlice(subquests), nil
}
