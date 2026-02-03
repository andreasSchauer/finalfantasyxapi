package main

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type MonsterItems struct {
	DropChance          int32          `json:"drop_chance"`
	DropCondition       *string        `json:"drop_condition,omitempty"`
	StealCommon         *ItemAmount    `json:"steal_common"`
	StealRare           *ItemAmount    `json:"steal_rare"`
	DropCommon          *ItemAmount    `json:"drop_common"`
	DropRare            *ItemAmount    `json:"drop_rare"`
	SecondaryDropCommon *ItemAmount    `json:"secondary_drop_common"`
	SecondaryDropRare   *ItemAmount    `json:"secondary_drop_rare"`
	Bribe               *ItemAmount    `json:"bribe"`
	OtherItemsCondition *string        `json:"other_items_condition,omitempty"`
	OtherItems          []PossibleItem `json:"other_items"`
}

func (mi MonsterItems) IsZero() bool {
	return mi.DropChance == 0 && mi.OtherItemsCondition == nil
}

func convertMonsterItems(cfg *Config, items seeding.MonsterItems) MonsterItems {
	monItems := MonsterItems{
		DropChance:          items.DropChance,
		DropCondition:       items.DropCondition,
		OtherItemsCondition: items.OtherItemsCondition,
		OtherItems:          convertObjSlice(cfg, items.OtherItems, convertPossibleItem),
	}

	if items.StealCommon != nil {
		stealCommon := convertItemAmount(cfg, *items.StealCommon)
		monItems.StealCommon = &stealCommon
	}

	if items.StealRare != nil {
		stealRare := convertItemAmount(cfg, *items.StealRare)
		monItems.StealRare = &stealRare
	}

	if items.DropCommon != nil {
		dropCommon := convertItemAmount(cfg, *items.DropCommon)
		monItems.DropCommon = &dropCommon
	}

	if items.DropRare != nil {
		dropRare := convertItemAmount(cfg, *items.DropRare)
		monItems.DropRare = &dropRare
	}

	if items.SecondaryDropCommon != nil {
		secDropCommon := convertItemAmount(cfg, *items.SecondaryDropCommon)
		monItems.SecondaryDropCommon = &secDropCommon
	}

	if items.SecondaryDropRare != nil {
		secDropRare := convertItemAmount(cfg, *items.SecondaryDropRare)
		monItems.SecondaryDropRare = &secDropRare
	}

	if items.Bribe != nil {
		bribe := convertItemAmount(cfg, *items.Bribe)
		monItems.Bribe = &bribe
	}

	return monItems
}