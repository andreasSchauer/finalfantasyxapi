package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type QuestCompletion struct {
	ID        int32
	QuestID	  int32
	Condition *string          `json:"condition"`
	Areas     []CompletionArea `json:"areas"`
	Reward    ItemAmount       `json:"reward"`
}

func (qc QuestCompletion) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", qc),
		qc.QuestID,
		h.DerefOrNil(qc.Condition),
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
		fmt.Sprintf("%T", cl),
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
