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

func (cfg *Config) getMonsterItems(items *seeding.MonsterItems) *MonsterItems {
	if items == nil {
		return nil
	}
	monItems := MonsterItems{
		DropChance:          items.DropChance,
		DropCondition:       items.DropCondition,
		OtherItemsCondition: items.OtherItemsCondition,
		OtherItems:          cfg.getMonsterOtherItems(items.OtherItems),
	}

	if items.StealCommon != nil {
		stealCommon := cfg.createItemAmount(*items.StealCommon)
		monItems.StealCommon = &stealCommon
	}

	if items.StealRare != nil {
		stealRare := cfg.createItemAmount(*items.StealRare)
		monItems.StealRare = &stealRare
	}

	if items.DropCommon != nil {
		dropCommon := cfg.createItemAmount(*items.DropCommon)
		monItems.DropCommon = &dropCommon
	}

	if items.DropRare != nil {
		dropRare := cfg.createItemAmount(*items.DropRare)
		monItems.DropRare = &dropRare
	}

	if items.SecondaryDropCommon != nil {
		secDropCommon := cfg.createItemAmount(*items.SecondaryDropCommon)
		monItems.SecondaryDropCommon = &secDropCommon
	}

	if items.SecondaryDropRare != nil {
		secDropRare := cfg.createItemAmount(*items.SecondaryDropRare)
		monItems.SecondaryDropRare = &secDropRare
	}

	if items.Bribe != nil {
		bribe := cfg.createItemAmount(*items.Bribe)
		monItems.Bribe = &bribe
	}

	return &monItems
}


func (cfg *Config) getMonsterOtherItems(seedItems []seeding.PossibleItem) []PossibleItem {
	otherItems := []PossibleItem{}

	for _, item := range seedItems {
		possibleItem := cfg.newPossibleItem(item.ItemAmount, item.Chance)
		otherItems = append(otherItems, possibleItem)
	}

	return otherItems
}
