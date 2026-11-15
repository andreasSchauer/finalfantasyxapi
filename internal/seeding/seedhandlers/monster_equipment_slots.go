package seedhandlers

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Loop3MonsterEquipmentSlots(qtx *database.Queries, ctx context.Context) error {
	slots, err := extractMonsterEquipmentSlots()
	if err != nil {
		return err
	}

	params := database.CreateMonsterEquipmentSlotsBulkParams{
		DataHash:           make([]string, len(slots)),
		MonsterEquipmentID: make([]int32, len(slots)),
		MinAmount:          make([]int32, len(slots)),
		MaxAmount:          make([]int32, len(slots)),
		Type:               make([]database.EquipmentSlotsType, len(slots)),
	}

	for i, s := range slots {
		params.DataHash[i] = sh.GenerateDataHash(s)
		params.MonsterEquipmentID[i] = s.MonsterEquipmentID
		params.MinAmount[i] = s.MinAmount
		params.MaxAmount[i] = s.MaxAmount
		params.Type[i] = s.Type
	}

	dbRows, err := qtx.CreateMonsterEquipmentSlotsBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create monster equipment slots: %v", err)
	}

	for _, row := range dbRows {
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractMonsterEquipmentSlots() ([]seedtypes.MonsterEquipmentSlots, error) {
	slots := []seedtypes.MonsterEquipmentSlots{}
	var err error

	for i := range sh.RawJSON.Monsters {
		mon := &sh.RawJSON.Monsters[i]

		if mon.Equipment == nil {
			continue
		}

		abilitySlots := &mon.Equipment.AbilitySlots
		abilitySlots.MonsterEquipmentID, err = sh.GetHashID(mon.Equipment)
		if err != nil {
			return nil, err
		}

		abilitySlots.Type = database.EquipmentSlotsTypeAbilitySlots
		slots = append(slots, *abilitySlots)

		attachedAbilities := &mon.Equipment.AttachedAbilities
		attachedAbilities.MonsterEquipmentID, err = sh.GetHashID(mon.Equipment)
		if err != nil {
			return nil, err
		}

		attachedAbilities.Type = database.EquipmentSlotsTypeAttachedAbilities
		slots = append(slots, *attachedAbilities)
	}

	return sh.DedupeRows(slots, sh.Hashes), nil
}

func Loop1EquipmentSlotsChances(qtx *database.Queries, ctx context.Context) error {
	chances := extractEquipmentSlotsChances()

	params := database.CreateEquipmentSlotsChanceBulkParams{
		DataHash: make([]string, len(chances)),
		Amount:   make([]int32, len(chances)),
		Chance:   make([]int32, len(chances)),
	}

	for i, c := range chances {
		params.DataHash[i] = sh.GenerateDataHash(c)
		params.Amount[i] = c.Amount
		params.Chance[i] = c.Chance
	}

	dbRows, err := qtx.CreateEquipmentSlotsChanceBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create equipment slot chances: %v", err)
	}

	for _, row := range dbRows {
		sh.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func extractEquipmentSlotsChances() []seedtypes.EquipmentSlotsChance {
	chances := []seedtypes.EquipmentSlotsChance{}

	for _, m := range sh.RawJSON.Monsters {
		if m.Equipment != nil {
			chances = append(chances, m.Equipment.AbilitySlots.Chances...)
			chances = append(chances, m.Equipment.AttachedAbilities.Chances...)
		}
	}

	return sh.DedupeRows(chances, sh.Hashes)
}
