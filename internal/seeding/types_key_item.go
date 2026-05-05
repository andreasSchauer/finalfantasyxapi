package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type KeyItem struct {
	ID int32
	MasterItem
	Category    string `json:"category"`
	Description string `json:"description"`
	Effect      string `json:"effect"`
}

func (k KeyItem) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", k),
		k.MasterItem.ID,
		k.Category,
		k.Description,
		k.Effect,
	}
}

func (k KeyItem) ToKeyFields() []any {
	return []any{
		k.Name,
	}
}

func (k KeyItem) GetID() int32 {
	return k.ID
}

func (k KeyItem) Error() string {
	return fmt.Sprintf("key item %s", k.Name)
}

func (k KeyItem) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   k.ID,
		Name: k.Name,
	}
}
