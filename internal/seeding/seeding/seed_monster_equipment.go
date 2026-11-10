package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func loop2SeedMonsterEquipment(qtx *database.Queries, ctx context.Context) error {
	equipments := extractMonsterEquipment()

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
		hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractMonsterEquipment() []seedtypes.MonsterEquipment {
	equipments := []seedtypes.MonsterEquipment{}

	for i := range rawJSON.monsters {
		mon := &rawJSON.monsters[i]

		if mon.Equipment != nil {
			mon.Equipment.MonsterID = mon.ID
			equipments = append(equipments, *mon.Equipment)
		}
	}

	return dedupeRows(equipments, hashes)
}

func completeMonsterEquipment(equipment *seedtypes.MonsterEquipment) error {
	if equipment == nil {
		return nil
	}

	err := assignID(equipment)
	if err != nil {
		return nil
	}

	err = assignID(&equipment.AbilitySlots)
	if err != nil {
		return nil
	}

	err = assignID(&equipment.AttachedAbilities)
	if err != nil {
		return nil
	}

	err = assignIDs(equipment.AbilitySlots.Chances)
	if err != nil {
		return nil
	}

	err = assignIDs(equipment.AttachedAbilities.Chances)
	if err != nil {
		return nil
	}

	err = assignIDs(equipment.WeaponAbilities)
	if err != nil {
		return nil
	}

	err = assignIDs(equipment.ArmorAbilities)
	if err != nil {
		return nil
	}

	return nil
}
