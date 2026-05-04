package seeding

import (
	"context"
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

func (l *Lookup) getMonsterItems() []MonsterItems {
	monsterItems := []MonsterItems{}

	for _, mon := range l.json.monsters {
		if mon.Items != nil {
			monsterItems = append(monsterItems, *mon.Items)
		}
	}

	return monsterItems
}

func (l *Lookup) getMonsterItemsOtherItems(mi MonsterItems) ([]PossibleItem, error) {
	return mi.OtherItems, nil
}

func (l *Lookup) seedJuncMonsterItemsOtherItems(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monster items + other items"
	jParams, err := processJunctions(l, desc, l.getMonsterItems(), l.getMonsterItemsOtherItems)
	if err != nil {
		return err
	}

	return qtx.CreateMonsterItemsOtherItemsJunctionBulk(ctx, database.CreateMonsterItemsOtherItemsJunctionBulkParams{
		DataHash:       jParams.DataHashes,
		MonsterItemsID: jParams.ParentIDs,
		PossibleItemID: jParams.ChildIDs,
	})
}

func (l *Lookup) getMonsterEquipment() []MonsterEquipment {
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
	jParams, err := processJunctions(l, desc, l.getMonsterEquipment(), l.getMonsterEquipmentEquipmentDrops)
	if err != nil {
		return err
	}

	return qtx.CreateMonsterEquipmentAbilitiesJunctionBulk(ctx, database.CreateMonsterEquipmentAbilitiesJunctionBulkParams{
		DataHash:           jParams.DataHashes,
		MonsterEquipmentID: jParams.ParentIDs,
		EquipmentDropID:    jParams.ChildIDs,
	})
}

func (l *Lookup) getMonsterEquipmentEquipmentSlots(me MonsterEquipment) ([]MonsterEquipmentSlots, error) {
	return []MonsterEquipmentSlots{me.AbilitySlots, me.AttachedAbilities}, nil
}

func (l *Lookup) getMonsterEquipmentSlotsSlotChances(mes MonsterEquipmentSlots) ([]EquipmentSlotsChance, error) {
	return mes.Chances, nil
}

func (l *Lookup) seedJuncMonsterEquipmentSlotChances(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monster equipment + equipment slot chances"
	jParams, err := processThreewayJunctions(l, desc, l.getMonsterEquipment(), l.getMonsterEquipmentEquipmentSlots, l.getMonsterEquipmentSlotsSlotChances)
	if err != nil {
		return err
	}

	return qtx.CreateMonsterEquipmentSlotsChancesJunctionBulk(ctx, database.CreateMonsterEquipmentSlotsChancesJunctionBulkParams{
		DataHash:           jParams.DataHashes,
		MonsterEquipmentID: jParams.GrandParentIDs,
		EquipmentSlotsID:   jParams.ParentIDs,
		SlotsChanceID:      jParams.ChildIDs,
	})
}

func (l *Lookup) getEquipmentDropCharacters(ed EquipmentDrop) ([]Character, error) {
	return getResources(ed.Characters, l.Characters)
}

func (l *Lookup) seedJuncEquipmentDropsCharacters(qtx *database.Queries, ctx context.Context) error {
	const desc string = "equipment drops + characters"
	jParams, err := processThreewayJunctions(l, desc, l.getMonsterEquipment(), l.getMonsterEquipmentEquipmentDrops, l.getEquipmentDropCharacters)
	if err != nil {
		return err
	}

	return qtx.CreateEquipmentDropsCharactersJunctionBulk(ctx, database.CreateEquipmentDropsCharactersJunctionBulkParams{
		DataHash:           jParams.DataHashes,
		MonsterEquipmentID: jParams.GrandParentIDs,
		EquipmentDropID:    jParams.ParentIDs,
		CharacterID:        jParams.ChildIDs,
	})
}
