package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Stat struct {
	ID               int32
	Name             string `json:"name"`
	Effect           string `json:"effect"`
	MinVal           int32  `json:"min_val"`
	MaxVal           int32  `json:"max_val"`
	MaxVal2          *int32 `json:"max_val_2"`
	SphereID         *int32
	ActivationSphere string `json:"activation_sphere"`
}

func (s Stat) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", s),
		s.Name,
		s.Effect,
		s.MinVal,
		s.MaxVal,
		h.DerefOrNil(s.MaxVal2),
		h.DerefOrNil(s.SphereID),
	}
}

func (s Stat) GetID() int32 {
	return s.ID
}

func (s Stat) Error() string {
	return fmt.Sprintf("stat %s", s.Name)
}

func (s Stat) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   s.ID,
		Name: s.Name,
	}
}
