package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Accuracy struct {
	ID          int32
	AccSource   string   `json:"acc_source"`
	HitChance   *int32   `json:"hit_chance"`
	AccModifier *float32 `json:"acc_modifier"`
}

func (a Accuracy) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", a),
		a.AccSource,
		h.DerefOrNil(a.HitChance),
		h.DerefOrNil(a.AccModifier),
	}
}

func (a Accuracy) GetID() int32 {
	return a.ID
}

func (a Accuracy) Error() string {
	return fmt.Sprintf("accuracy with source: %s, hit chance: %v, modifier: %v", a.AccSource, h.PtrToString(a.HitChance), h.PtrToString(a.AccModifier))
}
