package seeding

import (
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type MonsterEquipment struct {
	ID                int32
	MonsterID         int32
	DropChance        int32                 `json:"drop_chance"`
	Power             int32                 `json:"power"`
	CriticalPlus      int32                 `json:"critical_plus"`
	AbilitySlots      MonsterEquipmentSlots `json:"ability_slots"`
	AttachedAbilities MonsterEquipmentSlots `json:"attached_abilities"`
	WeaponAbilities   []EquipmentDrop       `json:"weapon_abilities"`
	ArmorAbilities    []EquipmentDrop       `json:"armor_abilities"`
}

func (m MonsterEquipment) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", m),
		m.MonsterID,
		m.DropChance,
		m.Power,
		m.CriticalPlus,
	}
}

func (m MonsterEquipment) GetID() int32 {
	return m.ID
}

func (m *MonsterEquipment) SetID(id int32) {
	m.ID = id
}

func (m MonsterEquipment) Error() string {
	return fmt.Sprintf("monster equipment of monster with id %d", m.MonsterID)
}

type MonsterEquipmentSlots struct {
	ID                 int32
	MonsterEquipmentID int32
	MinAmount          int32                  `json:"min_amount"`
	MaxAmount          int32                  `json:"max_amount"`
	Chances            []EquipmentSlotsChance `json:"chances"`
	Type               database.EquipmentSlotsType
}

func (m MonsterEquipmentSlots) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", m),
		m.MonsterEquipmentID,
		m.MinAmount,
		m.MaxAmount,
		m.Type,
	}
}

func (m MonsterEquipmentSlots) GetID() int32 {
	return m.ID
}

func (m *MonsterEquipmentSlots) SetID(id int32) {
	m.ID = id
}

func (m MonsterEquipmentSlots) Error() string {
	return fmt.Sprintf("monster equipment slots with monster equipment id: %d, type: %s, min amount: %d, max amount: %d", m.ID, m.Type, m.MinAmount, m.MaxAmount)
}

type EquipmentSlotsChance struct {
	ID     int32
	Amount int32 `json:"amount"`
	Chance int32 `json:"chance"`
}

func (e EquipmentSlotsChance) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", e),
		e.Amount,
		e.Chance,
	}
}

func (e EquipmentSlotsChance) GetID() int32 {
	return e.ID
}

func (e *EquipmentSlotsChance) SetID(id int32) {
	e.ID = id
}

func (e EquipmentSlotsChance) Error() string {
	return fmt.Sprintf("equipment slots chance with amount: %d, chance: %d", e.Amount, e.Chance)
}

type EquipmentDrop struct {
	ID            int32
	AutoAbilityID int32
	Ability       string   `json:"ability"`
	Characters    []string `json:"characters"`
	IsForced      bool     `json:"is_forced"`
	Probability   *int32   `json:"probability"`
	Type          database.EquipType
}

func (e EquipmentDrop) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", e),
		e.AutoAbilityID,
		e.IsForced,
		h.DerefOrNil(e.Probability),
		e.Type,
	}
}

func (e EquipmentDrop) GetID() int32 {
	return e.ID
}

func (e *EquipmentDrop) SetID(id int32) {
	e.ID = id
}

func (e EquipmentDrop) Error() string {
	return fmt.Sprintf("equipment drop with auto-ability id: %d, type: %s, is forced: %t, probability: %v", e.AutoAbilityID, e.Type, e.IsForced, h.PtrToString(e.Probability))
}
