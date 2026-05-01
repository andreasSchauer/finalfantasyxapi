package seeding

import (
	"context"
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
		return MonsterEquipment{}, h.NewErr(monsterEquipment.Error(), err, "couldn't create monster equipment")
	}

	monsterEquipment.ID = dbMonsterEquipment.ID
	l.currentME = monsterEquipment

	err = l.seedMonsterEquipmentRelationships(qtx, monsterEquipment)
	if err != nil {
		return MonsterEquipment{}, h.NewErr(monsterEquipment.Error(), err)
	}

	return monsterEquipment, nil
}

func (l *Lookup) seedMonsterEquipmentRelationships(qtx *database.Queries, monsterEquipment MonsterEquipment) error {
	err := l.seedMonsterEquipmentSlotsWrapper(qtx, monsterEquipment, monsterEquipment.AbilitySlots, database.EquipmentSlotsTypeAbilitySlots)
	if err != nil {
		return h.NewErr(monsterEquipment.Error(), err)
	}

	err = l.seedMonsterEquipmentSlotsWrapper(qtx, monsterEquipment, monsterEquipment.AttachedAbilities, database.EquipmentSlotsTypeAttachedAbilities)
	if err != nil {
		return h.NewErr(monsterEquipment.Error(), err)
	}

	err = l.seedEquipmentDrops(qtx, monsterEquipment, monsterEquipment.WeaponAbilities, database.EquipTypeWeapon)
	if err != nil {
		return h.NewErr(monsterEquipment.Error(), err)
	}

	err = l.seedEquipmentDrops(qtx, monsterEquipment, monsterEquipment.ArmorAbilities, database.EquipTypeArmor)
	if err != nil {
		return h.NewErr(monsterEquipment.Error(), err)
	}

	return nil
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