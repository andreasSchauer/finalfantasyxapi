package seedhandlers

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop2SeedMonsterEquipment(qtx *database.Queries, ctx context.Context) error {
	equipments := extractMonsterEquipment()

	params := database.CreateMonsterEquipmentBulkParams{
		DataHash:     make([]string, len(equipments)),
		MonsterID:    make([]int32, len(equipments)),
		DropChance:   make([]int32, len(equipments)),
		Power:        make([]int32, len(equipments)),
		CriticalPlus: make([]int32, len(equipments)),
	}

	for i, e := range equipments {
		params.DataHash[i] = sh.GenerateDataHash(e)
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
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractMonsterEquipment() []seedtypes.MonsterEquipment {
	equipments := []seedtypes.MonsterEquipment{}

	for i := range sh.RawJSON.Monsters {
		mon := &sh.RawJSON.Monsters[i]

		if mon.Equipment != nil {
			mon.Equipment.MonsterID = mon.ID
			equipments = append(equipments, *mon.Equipment)
		}
	}

	return sh.DedupeRows(equipments, sh.Hashes)
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
