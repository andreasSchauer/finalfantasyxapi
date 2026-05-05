package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Modifier struct {
	ID           int32
	Name         string   `json:"name"`
	Effect       string   `json:"effect"`
	Category     string   `json:"type"`
	DefaultValue *float32 `json:"default_value"`
}

func (m Modifier) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", m),
		m.Name,
		m.Effect,
		m.Category,
		h.DerefOrNil(m.DefaultValue),
	}
}

func (m Modifier) ToKeyFields() []any {
	return []any{
		m.Name,
	}
}

func (m Modifier) GetID() int32 {
	return m.ID
}

func (m Modifier) Error() string {
	return fmt.Sprintf("modifier %s", m.Name)
}

func (m Modifier) GetResParamsNamed() ResParamsNamed {
	return ResParamsNamed{
		ID:   m.ID,
		Name: m.Name,
	}
}
