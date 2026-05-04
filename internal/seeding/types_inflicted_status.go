package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type InflictedStatus struct {
	ID                int32
	StatusConditionID int32
	StatusCondition   string `json:"name"`
	Probability       int32  `json:"probability"`
	DurationType      string `json:"duration_type"`
	Amount            *int32 `json:"amount"`
}

func (is InflictedStatus) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", is),
		is.StatusConditionID,
		is.Probability,
		is.DurationType,
		h.DerefOrNil(is.Amount),
	}
}

func (is InflictedStatus) GetID() int32 {
	return is.ID
}

func (is *InflictedStatus) SetID(id int32) {
	is.ID = id
}

func (is InflictedStatus) Error() string {
	return fmt.Sprintf("inflicted status with condition: %s, probability: %d, duration type: %s, amount: %v", is.StatusCondition, is.Probability, is.DurationType, is.Amount)
}
