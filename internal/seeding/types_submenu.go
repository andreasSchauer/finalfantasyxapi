package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Submenu struct {
	ID          int32
	Name        string  `json:"name"`
	Description *string `json:"description"`
	Effect      string  `json:"effect"`
	Topmenu     *string `json:"topmenu"`
	TopmenuID   *int32
	Users       []string `json:"users"`
}

func (s Submenu) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", s),
		s.Name,
		s.Description,
		s.Effect,
		h.DerefOrNil(s.Topmenu),
	}
}

func (s Submenu) ToKeyFields() []any {
	return []any{
		s.Name,
	}
}

func (s Submenu) GetID() int32 {
	return s.ID
}

func (s Submenu) Error() string {
	return fmt.Sprintf("submenu %s", s.Name)
}

func (s Submenu) GetResParamsNamed() ResParamsNamed {
	return ResParamsNamed{
		ID:   s.ID,
		Name: s.Name,
	}
}
