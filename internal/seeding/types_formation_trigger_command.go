package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type FormationTriggerCommand struct {
	ID int32
	AbilityReference
	TriggerCommandID int32
	Condition        *string  `json:"condition"`
	UseAmount        *int32   `json:"use_amount"`
	Users            []string `json:"users"`
}

func (tc FormationTriggerCommand) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", tc),
		tc.TriggerCommandID,
		h.DerefOrNil(tc.Condition),
		h.DerefOrNil(tc.UseAmount),
	}
}

func (tc FormationTriggerCommand) GetID() int32 {
	return tc.ID
}

func (tc *FormationTriggerCommand) SetID(id int32) {
	tc.ID = id
}

func (tc FormationTriggerCommand) Error() string {
	return fmt.Sprintf("formation trigger command with %s", tc.AbilityReference)
}
