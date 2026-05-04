package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func (l *Lookup) loop2SeedMonsterEquipment(qtx *database.Queries, ctx context.Context) error {
	equipments := l.extractMonsterEquipment()

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

func (l *Lookup) extractMonsterEquipment() []MonsterEquipment {
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
