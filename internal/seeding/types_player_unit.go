package seeding

import (
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type PlayerUnit struct {
	ID   int32
	Name string `json:"name"`
	Type database.UnitType
}

func (pu PlayerUnit) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", pu),
		pu.Name,
		pu.Type,
	}
}

func (pu PlayerUnit) ToKeyFields() []any {
	return []any{
		pu.Name,
	}
}

func (pu PlayerUnit) GetID() int32 {
	return pu.ID
}

func (pu PlayerUnit) Error() string {
	return fmt.Sprintf("player unit %s, type %s", pu.Name, pu.Type)
}

func (pu PlayerUnit) GetResParamsTyped() ResParamsTyped {
	return ResParamsTyped{
		ID:   pu.ID,
		Name: pu.Name,
		Type: string(pu.Type),
	}
}
