package seeding

import (
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type MasterItem struct {
	ID   int32
	Name string `json:"name"`
	Type database.ItemType
}

func (i MasterItem) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", i),
		i.Name,
		i.Type,
	}
}

func (i MasterItem) ToKeyFields() []any {
	return []any{
		i.Name,
	}
}

func (i MasterItem) GetID() int32 {
	return i.ID
}

func (i MasterItem) Error() string {
	return fmt.Sprintf("master item %s, type %s", i.Name, i.Type)
}

func (i MasterItem) GetResParamsTyped() ResParamsTyped {
	return ResParamsTyped{
		ID:   i.ID,
		Name: i.Name,
		Type: string(i.Type),
	}
}
