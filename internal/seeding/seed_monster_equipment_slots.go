package seeding

import (
	"context"
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

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

func (l *Lookup) loop1SeedEquipmentSlotsChances(qtx *database.Queries, ctx context.Context) error {
	chances := l.getEquipmentSlotsChances()

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

func (l *Lookup) getEquipmentSlotsChances() []EquipmentSlotsChance {
	chances := []EquipmentSlotsChance{}

	for _, m := range l.json.monsters {
		if m.Equipment != nil {
			chances = append(chances, m.Equipment.AbilitySlots.Chances...)
			chances = append(chances, m.Equipment.AttachedAbilities.Chances...)
		}
	}

	return dedupeRows(chances, l.Hashes)
}


func (l *Lookup) loop3SeedMonsterEquipmentSlots(qtx *database.Queries, ctx context.Context) error {
	slots, err := l.extractMonsterEquipmentSlots()
	if err != nil {
		return err
	}

	params := database.CreateMonsterEquipmentSlotsBulkParams{
		DataHash: 			make([]string, len(slots)),
		MonsterEquipmentID: make([]int32, len(slots)),
		MinAmount: 			make([]int32, len(slots)),
		MaxAmount: 			make([]int32, len(slots)),
		Type: 				make([]database.EquipmentSlotsType, len(slots)),
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
		abilitySlots.MonsterEquipmentID, err = l.getHashID(mon.Equipment)
		if err != nil {
			return nil, err
		}

		abilitySlots.Type = database.EquipmentSlotsTypeAbilitySlots
		slots = append(slots, *abilitySlots)

		attachedAbilities := &mon.Equipment.AttachedAbilities
		attachedAbilities.MonsterEquipmentID, err = l.getHashID(mon.Equipment)
		if err != nil {
			return nil, err
		}

		attachedAbilities.Type = database.EquipmentSlotsTypeAttachedAbilities
		slots = append(slots, *attachedAbilities)
	}

	return dedupeRows(slots, l.Hashes), nil
}

func (l *Lookup) getMonsterEquipmentEquipmentSlots(me MonsterEquipment) ([]MonsterEquipmentSlots, error) {
	return []MonsterEquipmentSlots{me.AbilitySlots, me.AttachedAbilities}, nil
}

func (l *Lookup) getMonsterEquipmentSlotsSlotChances(mes MonsterEquipmentSlots) ([]EquipmentSlotsChance, error) {
	return mes.Chances, nil
}

func (l *Lookup) seedJuncMonsterEquipmentSlotChances(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monster equipment + equipment slot chances"
	jParams, err := processThreewayJunctions(l, desc, l.getMonsterEquipments(), l.getMonsterEquipmentEquipmentSlots, l.getMonsterEquipmentSlotsSlotChances)
	if err != nil {
		return err
	}

	return qtx.CreateMonsterEquipmentSlotsChancesJunctionBulk(ctx, database.CreateMonsterEquipmentSlotsChancesJunctionBulkParams{
		DataHash:       	jParams.DataHashes,
		MonsterEquipmentID: jParams.GrandParentIDs,
		EquipmentSlotsID: 	jParams.ParentIDs,
		SlotsChanceID:  	jParams.ChildIDs,
	})
}