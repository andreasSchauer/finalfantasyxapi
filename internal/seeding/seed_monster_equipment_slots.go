package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func (l *Lookup) loop3SeedMonsterEquipmentSlots(qtx *database.Queries, ctx context.Context) error {
	slots, err := l.extractMonsterEquipmentSlots()
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
		params.DataHash[i] = generateDataHash(s)
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
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractMonsterEquipmentSlots() ([]MonsterEquipmentSlots, error) {
	slots := []MonsterEquipmentSlots{}
	var err error

	for i := range l.json.monsters {
		mon := &l.json.monsters[i]

		if mon.Equipment == nil {
			continue
		}

		abilitySlots := &mon.Equipment.AbilitySlots
		abilitySlots.MonsterEquipmentID, err = l.GetHashID(mon.Equipment)
		if err != nil {
			return nil, err
		}

		abilitySlots.Type = database.EquipmentSlotsTypeAbilitySlots
		slots = append(slots, *abilitySlots)

		attachedAbilities := &mon.Equipment.AttachedAbilities
		attachedAbilities.MonsterEquipmentID, err = l.GetHashID(mon.Equipment)
		if err != nil {
			return nil, err
		}

		attachedAbilities.Type = database.EquipmentSlotsTypeAttachedAbilities
		slots = append(slots, *attachedAbilities)
	}

	return dedupeRows(slots, l.Hashes), nil
}

func (l *Lookup) loop1SeedEquipmentSlotsChances(qtx *database.Queries, ctx context.Context) error {
	chances := l.extractEquipmentSlotsChances()

	params := database.CreateEquipmentSlotsChanceBulkParams{
		DataHash: make([]string, len(chances)),
		Amount:   make([]int32, len(chances)),
		Chance:   make([]int32, len(chances)),
	}

	for i, c := range chances {
		params.DataHash[i] = generateDataHash(c)
		params.Amount[i] = c.Amount
		params.Chance[i] = c.Chance
	}

	dbRows, err := qtx.CreateEquipmentSlotsChanceBulk(ctx, params)
	if err != nil {
		return fmt.Errorf("couldn't create equipment slot chances: %v", err)
	}

	for _, row := range dbRows {
		l.Hashes[row.DataHash] = row.ID
	}

	return nil
}

func (l *Lookup) extractEquipmentSlotsChances() []EquipmentSlotsChance {
	chances := []EquipmentSlotsChance{}

	for _, m := range l.json.monsters {
		if m.Equipment != nil {
			chances = append(chances, m.Equipment.AbilitySlots.Chances...)
			chances = append(chances, m.Equipment.AttachedAbilities.Chances...)
		}
	}

	return dedupeRows(chances, l.Hashes)
}
