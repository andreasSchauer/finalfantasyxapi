package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
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
		m.MonsterID,
		m.DropChance,
		m.Power,
		m.CriticalPlus,
	}
}

func (m MonsterEquipment) GetID() int32 {
	return m.ID
}

func (m MonsterEquipment) Error() string {
	return fmt.Sprintf("monster equipment of monster with id %d", m.MonsterID)
}

func (l *Lookup) seedMonsterEquipment(qtx *database.Queries, monsterEquipment MonsterEquipment) (MonsterEquipment, error) {
	var err error

	dbMonsterEquipment, err := qtx.CreateMonsterEquipment(context.Background(), database.CreateMonsterEquipmentParams{
		DataHash:     generateDataHash(monsterEquipment),
		MonsterID:    monsterEquipment.MonsterID,
		DropChance:   monsterEquipment.DropChance,
		Power:        monsterEquipment.Power,
		CriticalPlus: monsterEquipment.CriticalPlus,
	})
	if err != nil {
		return MonsterEquipment{}, getErr(monsterEquipment.Error(), err, "couldn't create monster equipment")
	}

	monsterEquipment.ID = dbMonsterEquipment.ID
	l.currentME = monsterEquipment

	err = l.seedMonsterEquipmentRelationships(qtx, monsterEquipment)
	if err != nil {
		return MonsterEquipment{}, getErr(monsterEquipment.Error(), err)
	}

	return monsterEquipment, nil
}

func (l *Lookup) seedMonsterEquipmentRelationships(qtx *database.Queries, monsterEquipment MonsterEquipment) error {
	err := l.seedMonsterEquipmentSlotsWrapper(qtx, monsterEquipment, monsterEquipment.AbilitySlots, database.EquipmentSlotsTypeAbilitySlots)
	if err != nil {
		return getErr(monsterEquipment.Error(), err)
	}

	err = l.seedMonsterEquipmentSlotsWrapper(qtx, monsterEquipment, monsterEquipment.AttachedAbilities, database.EquipmentSlotsTypeAttachedAbilities)
	if err != nil {
		return getErr(monsterEquipment.Error(), err)
	}

	err = l.seedEquipmentDrops(qtx, monsterEquipment, monsterEquipment.WeaponAbilities, database.EquipTypeWeapon)
	if err != nil {
		return getErr(monsterEquipment.Error(), err)
	}

	err = l.seedEquipmentDrops(qtx, monsterEquipment, monsterEquipment.ArmorAbilities, database.EquipTypeArmor)
	if err != nil {
		return getErr(monsterEquipment.Error(), err)
	}

	return nil
}
