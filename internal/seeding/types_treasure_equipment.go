package seeding

import "fmt"

type TreasureEquipment struct {
	ID               int32
	TreasureID       int32
	EquipmentNameID  int32
	Name             string   `json:"name"`
	Abilities        []string `json:"abilities"`
	EmptySlotsAmount int32    `json:"empty_slots_amount"`
}

func (te TreasureEquipment) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", te),
		te.TreasureID,
		te.EquipmentNameID,
		te.EmptySlotsAmount,
	}
}

func (te TreasureEquipment) GetID() int32 {
	return te.ID
}

func (te *TreasureEquipment) SetID(id int32) {
	te.ID = id
}

func (te TreasureEquipment) Error() string {
	return fmt.Sprintf("treasure equipment with name: %s, empty slots: %d", te.Name, te.EmptySlotsAmount)
}
