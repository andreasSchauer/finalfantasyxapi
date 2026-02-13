package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/seeding"

type MonsterItemsSub struct {
	StealCommon         *ItemAmountSub  `json:"steal_common"`
	StealRare           *ItemAmountSub  `json:"steal_rare"`
	DropCommon          *ItemAmountSub  `json:"drop_common"`
	DropRare            *ItemAmountSub  `json:"drop_rare"`
	SecondaryDropCommon *ItemAmountSub  `json:"secondary_drop_common"`
	SecondaryDropRare   *ItemAmountSub  `json:"secondary_drop_rare"`
	Bribe               *ItemAmountSub  `json:"bribe"`
	OtherItems          []ItemAmountSub `json:"other_items"`
}

func convertMonsterSubItems(cfg *Config, items seeding.MonsterItems) MonsterItemsSub {
	return MonsterItemsSub{
		StealCommon:         convertObjPtr(cfg, items.StealCommon, convertSubItemAmount),
		StealRare:           convertObjPtr(cfg, items.StealRare, convertSubItemAmount),
		DropCommon:          convertObjPtr(cfg, items.DropCommon, convertSubItemAmount),
		DropRare:            convertObjPtr(cfg, items.DropRare, convertSubItemAmount),
		SecondaryDropCommon: convertObjPtr(cfg, items.SecondaryDropCommon, convertSubItemAmount),
		SecondaryDropRare:   convertObjPtr(cfg, items.SecondaryDropRare, convertSubItemAmount),
		Bribe:               convertObjPtr(cfg, items.Bribe, convertSubItemAmount),
		OtherItems:          convertObjSlice(cfg, items.OtherItems, posItemToItemAmtSub),
	}
}

type MonsterEquipmentSub struct {
	WeaponAbilities []string `json:"weapon_abilities"`
	ArmorAbilities  []string `json:"armor_abilities"`
}

func convertMonsterSubEquipment(cfg *Config, equipment seeding.MonsterEquipment) MonsterEquipmentSub {
	return MonsterEquipmentSub{
		WeaponAbilities: convertObjSlice(cfg, equipment.WeaponAbilities, monsterAutoAbilityString),
		ArmorAbilities:  convertObjSlice(cfg, equipment.ArmorAbilities, monsterAutoAbilityString),
	}
}