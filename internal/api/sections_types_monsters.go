package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/seeding"

type MonsterItemsSimple struct {
	StealCommon         *string  `json:"steal_common,omitempty"`
	StealRare           *string  `json:"steal_rare,omitempty"`
	DropCommon          *string  `json:"drop_common,omitempty"`
	DropRare            *string  `json:"drop_rare,omitempty"`
	SecondaryDropCommon *string  `json:"secondary_drop_common,omitempty"`
	SecondaryDropRare   *string  `json:"secondary_drop_rare,omitempty"`
	Bribe               *string  `json:"bribe,omitempty"`
	OtherItems          []string `json:"other_items,omitempty"`
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
		OtherItems:          convertObjSliceNullable(cfg, items.OtherItems, posItemToItemAmtSimple),
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
