package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type TreasureList struct {
	LocationArea LocationArea `json:"location_area"`
	Treasures    []Treasure   `json:"treasures"`
}

func (tl TreasureList) Error() string {
	return fmt.Sprintf("treasures at %s", tl.LocationArea)
}

type Treasure struct {
	ID              int32
	Version         int32
	AreaID          int32
	TreasureType    string             `json:"treasure_type"`
	LootType        string             `json:"loot_type"`
	Availability    string             `json:"availability"`
	IsAnimaTreasure bool               `json:"is_anima_treasure"`
	Notes           *string            `json:"notes"`
	GilAmount       *int32             `json:"gil_amount"`
	Items           []ItemAmount       `json:"items"`
	Equipment       *TreasureEquipment `json:"equipment"`
}

func (t Treasure) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", t),
		t.Version,
		t.AreaID,
		t.TreasureType,
		t.LootType,
		t.Availability,
		t.IsAnimaTreasure,
		h.DerefOrNil(t.Notes),
		h.DerefOrNil(t.GilAmount),
	}
}

func (t Treasure) ToKeyFields() []any {
	return []any{
		t.AreaID,
		t.Version,
	}
}

func (t Treasure) GetID() int32 {
	return t.ID
}

func (t *Treasure) SetID(id int32) {
	t.ID = id
}

func (t Treasure) Error() string {
	return fmt.Sprintf("treasure number: %d", t.Version)
}

func (t Treasure) GetResParamsUnnamed() h.ResParamsUnnamed {
	return h.ResParamsUnnamed{
		ID: t.ID,
	}
}

func (t Treasure) GetItemAmounts() []ItemAmount {
	return t.Items
}
