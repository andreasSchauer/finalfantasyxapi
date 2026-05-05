package seeding

import (
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Quest struct {
	ID           int32
	Name         string           `json:"name"`
	Availability string           `json:"availability"`
	IsRepeatable bool             `json:"is_repeatable"`
	Completion   *QuestCompletion `json:"completion"`
	Type         database.QuestType
}

func (q Quest) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", q),
		q.Name,
		q.Type,
		h.ObjPtrToID(q.Completion),
		q.Availability,
		q.IsRepeatable,
	}
}

func (q Quest) ToKeyFields() []any {
	return []any{
		q.Name,
	}
}

func (q Quest) GetID() int32 {
	return q.ID
}

func (q Quest) Error() string {
	return fmt.Sprintf("quest %s, type %s", q.Name, q.Type)
}

func (q Quest) GetResParamsQuest() ResParamsQuest {
	switch q.Type {
	case database.QuestTypeSidequest:
		return ResParamsQuest{
			ID:        q.ID,
			Sidequest: &q.Name,
			Subquest:  nil,
			Type:      string(q.Type),
		}

	case database.QuestTypeSubquest:
		return ResParamsQuest{
			ID:        q.ID,
			Sidequest: nil,
			Subquest:  &q.Name,
			Type:      string(q.Type),
		}
	}

	return ResParamsQuest{}
}

func (q Quest) GetItemAmount() ItemAmount {
	if q.Completion == nil {
		return ItemAmount{}
	}

	return q.Completion.Reward
}
