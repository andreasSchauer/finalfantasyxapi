package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

type MonsterItems struct {
	DropChance          int32                             `json:"drop_chance"`
	DropCondition       *string                           `json:"drop_condition,omitempty"`
	StealCommon         *ResourceAmount[TypedAPIResource] `json:"steal_common"`
	StealRare           *ResourceAmount[TypedAPIResource] `json:"steal_rare"`
	DropCommon          *ResourceAmount[TypedAPIResource] `json:"drop_common"`
	DropRare            *ResourceAmount[TypedAPIResource] `json:"drop_rare"`
	SecondaryDropCommon *ResourceAmount[TypedAPIResource] `json:"secondary_drop_common"`
	SecondaryDropRare   *ResourceAmount[TypedAPIResource] `json:"secondary_drop_rare"`
	Bribe               *ResourceAmount[TypedAPIResource] `json:"bribe"`
	OtherItemsCondition *string                           `json:"other_items_condition,omitempty"`
	OtherItems          []PossibleItem                    `json:"other_items"`
}

func (mi MonsterItems) IsZero() bool {
	return mi.DropChance == 0 && mi.OtherItemsCondition == nil
}

func convertMonsterItems(cfg *Config, items seedtypes.MonsterItems) MonsterItems {
	monItems := MonsterItems{
		DropChance:          items.DropChance,
		DropCondition:       items.DropCondition,
		OtherItemsCondition: items.OtherItemsCondition,
		OtherItems:          convertObjSlice(cfg, items.OtherItems, convertPossibleItem),
	}

	if items.StealCommon != nil {
		stealCommon := E.AllItems.toResAmtFromName(cfg, *items.StealCommon)
		monItems.StealCommon = &stealCommon
	}

	if items.StealRare != nil {
		stealRare := E.AllItems.toResAmtFromName(cfg, *items.StealRare)
		monItems.StealRare = &stealRare
	}

	if items.DropCommon != nil {
		dropCommon := E.AllItems.toResAmtFromName(cfg, *items.DropCommon)
		monItems.DropCommon = &dropCommon
	}

	if items.DropRare != nil {
		dropRare := E.AllItems.toResAmtFromName(cfg, *items.DropRare)
		monItems.DropRare = &dropRare
	}

	if items.SecondaryDropCommon != nil {
		secDropCommon := E.AllItems.toResAmtFromName(cfg, *items.SecondaryDropCommon)
		monItems.SecondaryDropCommon = &secDropCommon
	}

	if items.SecondaryDropRare != nil {
		secDropRare := E.AllItems.toResAmtFromName(cfg, *items.SecondaryDropRare)
		monItems.SecondaryDropRare = &secDropRare
	}

	if items.Bribe != nil {
		bribe := E.AllItems.toResAmtFromName(cfg, *items.Bribe)
		monItems.Bribe = &bribe
	}

	return monItems
}
