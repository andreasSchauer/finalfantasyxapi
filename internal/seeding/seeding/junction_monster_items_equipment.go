package seeding

import (
	"context"
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func getMonsterItems() []seedtypes.MonsterItems {
	monsterItems := []seedtypes.MonsterItems{}

	for _, mon := range rawJSON.monsters {
		if mon.Items != nil {
			monsterItems = append(monsterItems, *mon.Items)
		}
	}

	return monsterItems
}

func getMonsterItemsOtherItems(mi seedtypes.MonsterItems) ([]seedtypes.PossibleItem, error) {
	return mi.OtherItems, nil
}

func seedJuncMonsterItemsOtherItems(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monster items + other items"
	jParams, err := processJunctions(desc, getMonsterItems(), getMonsterItemsOtherItems)
	if err != nil {
		return err
	}

	return qtx.CreateMonsterItemsOtherItemsJunctionBulk(ctx, database.CreateMonsterItemsOtherItemsJunctionBulkParams{
		DataHash:       jParams.DataHashes,
		MonsterItemsID: jParams.ParentIDs,
		PossibleItemID: jParams.ChildIDs,
	})
}

func getMonsterEquipment() []seedtypes.MonsterEquipment {
	monsterEquipment := []seedtypes.MonsterEquipment{}

	for _, mon := range rawJSON.monsters {
		if mon.Equipment != nil {
			monsterEquipment = append(monsterEquipment, *mon.Equipment)
		}
	}

	return monsterEquipment
}

func getMonsterEquipmentEquipmentDrops(me seedtypes.MonsterEquipment) ([]seedtypes.EquipmentDrop, error) {
	return slices.Concat(me.WeaponAbilities, me.ArmorAbilities), nil
}

func seedJuncMonsterEquipmentEquipmentDrops(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monster equipment + equipment drops"
	jParams, err := processJunctions(desc, getMonsterEquipment(), getMonsterEquipmentEquipmentDrops)
	if err != nil {
		return err
	}

	return qtx.CreateMonsterEquipmentAbilitiesJunctionBulk(ctx, database.CreateMonsterEquipmentAbilitiesJunctionBulkParams{
		DataHash:           jParams.DataHashes,
		MonsterEquipmentID: jParams.ParentIDs,
		EquipmentDropID:    jParams.ChildIDs,
	})
}

func getMonsterEquipmentEquipmentSlots(me seedtypes.MonsterEquipment) ([]seedtypes.MonsterEquipmentSlots, error) {
	return []seedtypes.MonsterEquipmentSlots{me.AbilitySlots, me.AttachedAbilities}, nil
}

func getMonsterEquipmentSlotsSlotChances(mes seedtypes.MonsterEquipmentSlots) ([]seedtypes.EquipmentSlotsChance, error) {
	return mes.Chances, nil
}

func seedJuncMonsterEquipmentSlotChances(qtx *database.Queries, ctx context.Context) error {
	const desc string = "monster equipment + equipment slot chances"
	jParams, err := processThreewayJunctions(desc, getMonsterEquipment(), getMonsterEquipmentEquipmentSlots, getMonsterEquipmentSlotsSlotChances)
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

func getEquipmentDropCharacters(ed seedtypes.EquipmentDrop) ([]seedtypes.Character, error) {
	return lookup.GetResources(ed.Characters, lookup.L.Characters)
}

func seedJuncEquipmentDropsCharacters(qtx *database.Queries, ctx context.Context) error {
	const desc string = "equipment drops + characters"
	jParams, err := processThreewayJunctions(desc, getMonsterEquipment(), getMonsterEquipmentEquipmentDrops, getEquipmentDropCharacters)
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
