package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type Quest struct {
	ID		int32
	Name 	string 				`json:"name"`
	Type 	database.QuestType
}

func (q Quest) ToHashFields() []any {
	return []any{
		q.Name,
		q.Type,
	}
}


func (q Quest) ToKeyFields() []any {
	return []any{
		q.Name,
		q.Type,
	}
}


func (l *lookup) seedQuest(qtx *database.Queries, quest Quest) (database.Quest, error) {
	dbQuest, err := qtx.CreateQuest(context.Background(), database.CreateQuestParams{
		DataHash: generateDataHash(quest),
		Name:     quest.Name,
		Type:     quest.Type,
	})
	if err != nil {
		return database.Quest{}, fmt.Errorf("couldn't create Quest: %s: %v", quest.Name, err)
	}


	return dbQuest, nil
}