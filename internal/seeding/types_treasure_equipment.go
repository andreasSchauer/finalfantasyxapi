package seeding

import "fmt"

type TreasureEquipment struct {
	ID               int32		`json:"treasure_equipment_id"`
	TreasureID       int32		`json:"treasure_id"`
	EquipmentNameID  int32		`json:"equipment_name_id"`
	Name             string   	`json:"name"`
	Abilities        []string 	`json:"abilities"`
	EmptySlotsAmount int32    	`json:"empty_slots_amount"`
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
