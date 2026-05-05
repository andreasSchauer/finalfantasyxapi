package seeding

import (
	"fmt"
)

type BlitzballPosition struct {
	ID       int32
	Category string          `json:"category"`
	Slot     string          `json:"slot"`
	Items    []BlitzballItem `json:"items"`
}

func (b BlitzballPosition) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", b),
		b.Category,
		b.Slot,
	}
}

func (b BlitzballPosition) GetID() int32 {
	return b.ID
}

func (b BlitzballPosition) ToKeyFields() []any {
	return []any{
		b.Category,
		b.Slot,
	}
}

func (b BlitzballPosition) Error() string {
	return fmt.Sprintf("blitzball position %s in %s", b.Slot, b.Category)
}

func (b BlitzballPosition) GetResParamsNamed() ResParamsNamed {
	return ResParamsNamed{
		ID:   b.ID,
		Name: fmt.Sprintf("%s - %s", b.Category, b.Slot),
	}
}

func (b BlitzballPosition) GetItemAmounts() []ItemAmount {
	ias := []ItemAmount{}

	for _, item := range b.Items {
		ias = append(ias, item.ItemAmount)
	}

	return ias
}

type BlitzballItem struct {
	ID         int32
	PositionID int32
	PossibleItem
}

func (b BlitzballItem) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", b),
		b.PositionID,
		b.PossibleItem.ID,
	}
}

func (b BlitzballItem) GetID() int32 {
	return b.ID
}

func (b *BlitzballItem) SetID(id int32) {
	b.ID = id
}

func (b BlitzballItem) Error() string {
	return fmt.Sprintf("blitzball item %s, chance %d, position id %d", b.ItemAmount.ItemName, b.Chance, b.PositionID)
}
