package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Property struct {
	ID              int32
	Name            string           `json:"name"`
	Effect          string           `json:"effect"`
	RelatedStats    []string         `json:"related_stats"`
	NullifyArmored  *string          `json:"nullify_armored"`
	StatChanges     []StatChange     `json:"stat_changes"`
	ModifierChanges []ModifierChange `json:"modifier_changes"`
}

func (p Property) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", p),
		p.Name,
		p.Effect,
		h.DerefOrNil(p.NullifyArmored),
	}
}

func (p Property) GetID() int32 {
	return p.ID
}

func (p Property) Error() string {
	return fmt.Sprintf("property %s", p.Name)
}

func (p Property) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   p.ID,
		Name: p.Name,
	}
}
