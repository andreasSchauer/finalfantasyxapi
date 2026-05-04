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

func (cc CharacterClass) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", cc),
		cc.Name,
	}
}

func (cc CharacterClass) GetID() int32 {
	return cc.ID
}

func (cc CharacterClass) Error() string {
	return fmt.Sprintf("character class %s", cc.Name)
}

func (cc CharacterClass) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   cc.ID,
		Name: cc.Name,
	}
}
