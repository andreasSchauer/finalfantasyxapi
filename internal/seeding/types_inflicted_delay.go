package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type InflictedDelay struct {
	ID             int32
	Condition      *string `json:"condition"`
	CTBAttackType  string  `json:"ctb_attack_type"`
	DelayType      string  `json:"delay_type"`
	DamageConstant int32   `json:"damage_constant"`
}

func (id InflictedDelay) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", id),
		h.DerefOrNil(id.Condition),
		id.CTBAttackType,
		id.DelayType,
		id.DamageConstant,
	}
}

func (id InflictedDelay) GetID() int32 {
	return id.ID
}

func (id InflictedDelay) Error() string {
	return fmt.Sprintf("inflicted delay with ctb attack type: %s, delay type: %s, constant: %d, condition: %v", id.CTBAttackType, id.DelayType, id.DamageConstant, h.PtrToString(id.Condition))
}
