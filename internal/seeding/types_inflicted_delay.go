package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type InflictedDelay struct {
	ID            int32   `json:"inflicted_delay_id,omitempty"`
	Condition     *string `json:"condition"`
	DelayStrength string  `json:"delay_strength"`
}

func (id InflictedDelay) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", id),
		h.DerefOrNil(id.Condition),
		id.DelayStrength,
	}
}

func (id InflictedDelay) GetID() int32 {
	return id.ID
}

func (id InflictedDelay) Error() string {
	return fmt.Sprintf("inflicted delay with delay strength: %s, condition: %v", id.DelayStrength, h.PtrToString(id.Condition))
}
