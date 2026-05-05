package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type CharacterClass struct {
	ID       int32
	Name     string   `json:"name"`
	Category string   `json:"category"`
	Members  []string `json:"members"`
}

func (c CharacterClass) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", c),
		c.Name,
	}
}

func (c CharacterClass) ToKeyFields() []any {
	return []any{
		c.Name,
	}
}

func (c CharacterClass) GetID() int32 {
	return c.ID
}

func (c CharacterClass) Error() string {
	return fmt.Sprintf("character class %s", c.Name)
}

func (c CharacterClass) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   c.ID,
		Name: c.Name,
	}
}
