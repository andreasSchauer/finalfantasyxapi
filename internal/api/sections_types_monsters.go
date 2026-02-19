package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/seeding"

type MonsterItemsSimple struct {
	StealCommon         *ItemAmountSimple  `json:"steal_common"`
	StealRare           *ItemAmountSimple  `json:"steal_rare"`
	DropCommon          *ItemAmountSimple  `json:"drop_common"`
	DropRare            *ItemAmountSimple  `json:"drop_rare"`
	SecondaryDropCommon *ItemAmountSimple  `json:"secondary_drop_common"`
	SecondaryDropRare   *ItemAmountSimple  `json:"secondary_drop_rare"`
	Bribe               *ItemAmountSimple  `json:"bribe"`
	OtherItems          []ItemAmountSimple `json:"other_items"`
}

func convertMonsterItemsSimple(cfg *Config, items seeding.MonsterItems) MonsterItemsSimple {
	return MonsterItemsSimple{
		StealCommon:         convertObjPtr(cfg, items.StealCommon, convertItemAmountSimple),
		StealRare:           convertObjPtr(cfg, items.StealRare, convertItemAmountSimple),
		DropCommon:          convertObjPtr(cfg, items.DropCommon, convertItemAmountSimple),
		DropRare:            convertObjPtr(cfg, items.DropRare, convertItemAmountSimple),
		SecondaryDropCommon: convertObjPtr(cfg, items.SecondaryDropCommon, convertItemAmountSimple),
		SecondaryDropRare:   convertObjPtr(cfg, items.SecondaryDropRare, convertItemAmountSimple),
		Bribe:               convertObjPtr(cfg, items.Bribe, convertItemAmountSimple),
		OtherItems:          convertObjSlice(cfg, items.OtherItems, posItemToItemAmtSimple),
	}
}

type MonsterEquipmentSimple struct {
	WeaponAbilities []string `json:"weapon_abilities"`
	ArmorAbilities  []string `json:"armor_abilities"`
}

func convertMonsterEquipmentSimple(cfg *Config, equipment seeding.MonsterEquipment) MonsterEquipmentSimple {
	return MonsterEquipmentSimple{
		WeaponAbilities: convertObjSlice(cfg, equipment.WeaponAbilities, monsterAutoAbilityString),
		ArmorAbilities:  convertObjSlice(cfg, equipment.ArmorAbilities, monsterAutoAbilityString),
	}
}
