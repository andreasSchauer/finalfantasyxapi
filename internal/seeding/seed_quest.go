package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Quest struct {
	ID   			int32
	Name 			string 				`json:"name"`
	Availability	string			  	`json:"availability"`
	Completion 		*QuestCompletion 	`json:"completion"`
	Type 			database.QuestType
}

func (q Quest) ToHashFields() []any {
	return []any{
		q.Name,
		q.Type,
		h.ObjPtrToID(q.Completion),
		q.Availability,
	}
}

func (q Quest) ToKeyFields() []any {
	return []any{
		q.Name,
		q.Type,
	}
}

func (q Quest) GetID() int32 {
	return q.ID
}

func (q Quest) Error() string {
	return fmt.Sprintf("quest %s, type %s", q.Name, q.Type)
}

func (q Quest) GetResParamsQuest() h.ResParamsQuest {
	switch q.Type {
	case database.QuestTypeSidequest:
		return h.ResParamsQuest{
			ID: 		q.ID,
			Sidequest: 	&q.Name,
			Subquest: 	nil,
			Type: 		string(q.Type),
		}

	case database.QuestTypeSubquest:
		return h.ResParamsQuest{
			ID: 		q.ID,
			Sidequest: 	nil,
			Subquest: 	&q.Name,
			Type: 		string(q.Type),
		}
	}

	return h.ResParamsQuest{}
}

func (q Quest) GetItemAmount() ItemAmount {
	if q.Completion == nil {
		return ItemAmount{}
	}

	return q.Completion.Reward
}

func (l *Lookup) seedQuest(qtx *database.Queries, quest Quest) (Quest, error) {
	dbQuest, err := qtx.CreateQuest(context.Background(), database.CreateQuestParams{
		DataHash: generateDataHash(quest),
		Name:     		quest.Name,
		Type:     		quest.Type,
		Availability: 	database.AvailabilityType(quest.Availability),
	})
	if err != nil {
		return Quest{}, h.NewErr(quest.Error(), err, "couldn't create quest")
	}

	quest.ID = dbQuest.ID
	key := CreateLookupKey(quest)
	l.Quests[key] = quest
	l.QuestsID[quest.ID] = quest

	return quest, nil
}


type QuestCompletion struct {
	ID        		int32
	Condition 		*string          `json:"condition"`
	Areas     		[]CompletionArea `json:"areas"`
	Reward    		ItemAmount       `json:"reward"`
}

func (qc QuestCompletion) ToHashFields() []any {
	return []any{
		qc.Condition,
		qc.Reward.ID,
	}
}

func (qc QuestCompletion) GetID() int32 {
	return qc.ID
}

func (qc QuestCompletion) Error() string {
	return fmt.Sprintf("quest completion with reward item: %s, amount: %d, condition: %s", qc.Reward.ItemName, qc.Reward.Amount, h.DerefStringPtr(qc.Condition))
}

type CompletionArea struct {
	CompletionID int32
	AreaID       int32
	LocationArea LocationArea `json:"location_area"`
	Notes        *string      `json:"notes"`
}

func (cl CompletionArea) ToHashFields() []any {
	return []any{
		cl.CompletionID,
		cl.AreaID,
		h.DerefOrNil(cl.Notes),
	}
}

func (cl CompletionArea) Error() string {
	return fmt.Sprintf("completion location %s, with completion id: %d, notes: %v", cl.LocationArea, cl.CompletionID, h.PtrToString(cl.Notes))
}

func (cl CompletionArea) GetLocationArea() LocationArea {
	return cl.LocationArea
}


func (l *Lookup) seedQuestCompletion(qtx *database.Queries, completion QuestCompletion) (QuestCompletion, error) {
	var err error

	completion.Reward, err = seedObjAssignID(qtx, completion.Reward, l.seedItemAmount)
	if err != nil {
		return QuestCompletion{}, h.NewErr(completion.Error(), err)
	}

	dbCompletion, err := qtx.CreateQuestCompletion(context.Background(), database.CreateQuestCompletionParams{
		DataHash:     generateDataHash(completion),
		Condition:    h.GetNullString(completion.Condition),
		ItemAmountID: completion.Reward.ID,
	})
	if err != nil {
		return QuestCompletion{}, h.NewErr(completion.Error(), err, "couldn't create quest completion")
	}
	completion.ID = dbCompletion.ID

	err = l.seedCompletionAreas(qtx, completion)
	if err != nil {
		return QuestCompletion{}, h.NewErr(completion.Error(), err)
	}

	return completion, nil
}

func (l *Lookup) seedCompletionAreas(qtx *database.Queries, completion QuestCompletion) error {
	for _, location := range completion.Areas {
		var err error

		location.AreaID, err = assignFK(location.LocationArea, l.Areas)
		if err != nil {
			return err
		}
		location.CompletionID = completion.ID

		err = qtx.CreateCompletionArea(context.Background(), database.CreateCompletionAreaParams{
			DataHash:     generateDataHash(location),
			CompletionID: location.CompletionID,
			AreaID:       location.AreaID,
			Notes:        h.GetNullString(location.Notes),
		})
		if err != nil {
			return h.NewErr(location.Error(), err, "couldn't create completion location")
		}
	}

	return nil
}
