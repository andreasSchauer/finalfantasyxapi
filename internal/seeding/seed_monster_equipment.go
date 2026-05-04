package seeding

import (
	"context"
	"fmt"
	"slices"

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

func (l *Lookup) loop2SeedMonsterEquipments(qtx *database.Queries, ctx context.Context) error {
	equipments := l.extractMonsterEquipments()

	params := database.CreateMonsterEquipmentBulkParams{
		DataHash:     make([]string, len(equipments)),
		MonsterID:    make([]int32, len(equipments)),
		DropChance:   make([]int32, len(equipments)),
		Power:        make([]int32, len(equipments)),
		CriticalPlus: make([]int32, len(equipments)),
	}

	for i, e := range equipments {
		params.DataHash[i] = generateDataHash(e)
		params.MonsterID[i] = e.MonsterID
		params.DropChance[i] = e.DropChance
		params.Power[i] = e.Power
		params.CriticalPlus[i] = e.CriticalPlus
	}

	dbRows, err := qtx.CreateMonsterEquipmentBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create monster equipments: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractMonsterEquipments() []MonsterEquipment {
	equipments := []MonsterEquipment{}

	for i := range l.json.monsters {
		mon := &l.json.monsters[i]

		if mon.Equipment != nil {
			mon.Equipment.MonsterID = mon.ID
			equipments = append(equipments, *mon.Equipment)
		}
	}

	return dedupeRows(equipments, l.Hashes)
}

func (l *Lookup) completeMonsterEquipment(equipment *MonsterEquipment) error {
	if equipment == nil {
		return nil
	}

	err := l.assignID(equipment)
	if err != nil {
		return nil
	}

	err = l.assignID(&equipment.AbilitySlots)
	if err != nil {
		return nil
	}

	err = l.assignID(&equipment.AttachedAbilities)
	if err != nil {
		return nil
	}

	err = assignIDs(l, equipment.AbilitySlots.Chances)
	if err != nil {
		return nil
	}

	err = assignIDs(l, equipment.AttachedAbilities.Chances)
	if err != nil {
		return nil
	}

	err = assignIDs(l, equipment.WeaponAbilities)
	if err != nil {
		return nil
	}

	err = assignIDs(l, equipment.ArmorAbilities)
	if err != nil {
		return nil
	}

	return nil
}

func (l *Lookup) getMonsterEquipments() []MonsterEquipment {
	monsterEquipment := []MonsterEquipment{}

	for _, mon := range l.json.monsters {
		if mon.Equipment != nil {
			monsterEquipment = append(monsterEquipment, *mon.Equipment)
		}
	}

	return monsterEquipment
}

func (l *Lookup) getMonsterEquipmentEquipmentDrops(me MonsterEquipment) ([]EquipmentDrop, error) {
	return slices.Concat(me.WeaponAbilities, me.ArmorAbilities), nil
}

func (l *Lookup) seedJuncMonsterEquipmentEquipmentDrops(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monster equipment + equipment drops"
	jParams, err := processJunctions(l, desc, l.getMonsterEquipments(), l.getMonsterEquipmentEquipmentDrops)
	if err != nil {
		return err
	}

	return qtx.CreateMonsterEquipmentAbilitiesJunctionBulk(ctx, database.CreateMonsterEquipmentAbilitiesJunctionBulkParams{
		DataHash:       	jParams.DataHashes,
		MonsterEquipmentID: jParams.ParentIDs,
		EquipmentDropID:  	jParams.ChildIDs,
	})
}