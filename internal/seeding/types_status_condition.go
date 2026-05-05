package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type StatusCondition struct {
	ID                      int32
	Name                    string           `json:"name"`
	Category                string           `json:"category"`
	IsPermanent             bool             `json:"is_permanent"`
	Effect                  string           `json:"effect"`
	Visualization           *string          `json:"visualization"`
	RelatedStats            []string         `json:"related_stats"`
	RemovedStatusConditions []string         `json:"removed_status_conditions"`
	AddedElemResist         *ElementalResist `json:"added_elem_resist"`
	CtbOnInfliction         *InflictedDelay  `json:"ctb_on_infliction"`
	NullifyArmored          *string          `json:"nullify_armored"`
	StatChanges             []StatChange     `json:"stat_changes"`
	ModifierChanges         []ModifierChange `json:"modifier_changes"`
}

func (s StatusCondition) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", s),
		s.Name,
		s.Category,
		s.IsPermanent,
		s.Effect,
		h.DerefOrNil(s.Visualization),
		h.ObjPtrToID(s.AddedElemResist),
		h.DerefOrNil(s.NullifyArmored),
	}
}

func (s StatusCondition) ToKeyFields() []any {
	return []any{
		s.Name,
	}
}

func (s StatusCondition) GetID() int32 {
	return s.ID
}

func (s StatusCondition) Error() string {
	return fmt.Sprintf("status condition %s", s.Name)
}

func (s StatusCondition) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   s.ID,
		Name: s.Name,
	}
}
