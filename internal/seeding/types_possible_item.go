package seeding

import (
	"fmt"
)

type PossibleItem struct {
	ID         int32
	ItemAmount ItemAmount `json:"item"`
	Chance     int32      `json:"chance"`
}

func (i PossibleItem) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", i),
		i.ItemAmount.ID,
		i.Chance,
	}
}

func (i PossibleItem) GetID() int32 {
	return i.ID
}

func (i *PossibleItem) SetID(id int32) {
	i.ID = id
}

func (i PossibleItem) Error() string {
	return fmt.Sprintf("possible item %s, amount: %d, chance: %d", i.ItemAmount.ItemName, i.ItemAmount.Amount, i.Chance)
}
