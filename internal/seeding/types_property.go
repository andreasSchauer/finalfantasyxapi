package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Property struct {
	ID             int32
	Name           string           `json:"name"`
	Effect         string           `json:"effect"`
	RelatedStats   []string         `json:"related_stats"`
	NullifyArmored *string          `json:"nullify_armored"`
	ModifierChange *ModifierChange `json:"modifier_change"`
}

func (p Property) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", p),
		p.Name,
		p.Effect,
		h.DerefOrNil(p.NullifyArmored),
		h.ObjPtrToInt32ID(p.ModifierChange),
	}
}

func (p Property) ToKeyFields() []any {
	return []any{
		p.Name,
	}
}

func (p Property) GetID() int32 {
	return p.ID
}

func (p Property) Error() string {
	return fmt.Sprintf("property %s", p.Name)
}

func (p Property) GetResParamsNamed() ResParamsNamed {
	return ResParamsNamed{
		ID:   p.ID,
		Name: p.Name,
	}
}
