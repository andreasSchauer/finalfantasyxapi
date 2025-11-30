package main

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type MonsterItems struct {
	DropChance          int32          `json:"drop_chance"`
	DropCondition       *string        `json:"drop_condition"`
	StealCommon         *ItemAmount    `json:"steal_common"`
	StealRare           *ItemAmount    `json:"steal_rare"`
	DropCommon          *ItemAmount    `json:"drop_common"`
	DropRare            *ItemAmount    `json:"drop_rare"`
	SecondaryDropCommon *ItemAmount    `json:"secondary_drop_common"`
	SecondaryDropRare   *ItemAmount    `json:"secondary_drop_rare"`
	Bribe               *ItemAmount    `json:"bribe"`
	OtherItemsCondition *string        `json:"other_items_condition"`
	OtherItems          []PossibleItem `json:"other_items"`
}

func (mi MonsterItems) IsZero() bool {
	return mi.DropChance == 0 && mi.OtherItemsCondition == nil
}

func (cfg *apiConfig) getMonsterItems(r *http.Request, mon database.Monster) (MonsterItems, error) {
	dbItems, err := cfg.db.GetMonsterItems(r.Context(), mon.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return MonsterItems{}, nil
		}
		return MonsterItems{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("Couldn't retrieve items of Monster %s Version %d", mon.Name, *h.NullInt32ToPtr(mon.Version)), err)
	}

	stealCommon := cfg.newItemAmount(
		dbItems.StealCommonItemType.ItemType,
		dbItems.StealCommonItem.String,
		dbItems.StealCommonItemID.Int32,
		dbItems.StealCommonAmount.Int32,
	)

	stealRare := cfg.newItemAmount(
		dbItems.StealRareItemType.ItemType,
		dbItems.StealRareItem.String,
		dbItems.StealRareItemID.Int32,
		dbItems.StealRareAmount.Int32,
	)

	dropCommon := cfg.newItemAmount(
		dbItems.DropCommonItemType.ItemType,
		dbItems.DropCommonItem.String,
		dbItems.DropCommonItemID.Int32,
		dbItems.DropCommonAmount.Int32,
	)

	dropRare := cfg.newItemAmount(
		dbItems.DropRareItemType.ItemType,
		dbItems.DropRareItem.String,
		dbItems.DropRareItemID.Int32,
		dbItems.DropRareAmount.Int32,
	)

	secDropCommon := cfg.newItemAmount(
		dbItems.SecDropCommonItemType.ItemType,
		dbItems.SecDropCommonItem.String,
		dbItems.SecDropCommonItemID.Int32,
		dbItems.SecDropCommonAmount.Int32,
	)

	secDropRare := cfg.newItemAmount(
		dbItems.SecDropRareItemType.ItemType,
		dbItems.SecDropRareItem.String,
		dbItems.SecDropRareItemID.Int32,
		dbItems.SecDropRareAmount.Int32,
	)

	bribe := cfg.newItemAmount(
		dbItems.BribeItemType.ItemType,
		dbItems.BribeItem.String,
		dbItems.BribeItemID.Int32,
		dbItems.BribeAmount.Int32,
	)

	otherItems, err := cfg.getMonsterOtherItems(r, mon, dbItems.ID)
	if err != nil {
		return MonsterItems{}, err
	}

	return MonsterItems{
		DropChance:          anyToInt32(dbItems.DropChance),
		DropCondition:       h.NullStringToPtr(dbItems.DropCondition),
		StealCommon:         h.NilOrPtr(stealCommon),
		StealRare:           h.NilOrPtr(stealRare),
		DropCommon:          h.NilOrPtr(dropCommon),
		DropRare:            h.NilOrPtr(dropRare),
		SecondaryDropCommon: h.NilOrPtr(secDropCommon),
		SecondaryDropRare:   h.NilOrPtr(secDropRare),
		Bribe:               h.NilOrPtr(bribe),
		OtherItemsCondition: h.NullStringToPtr(dbItems.OtherItemsCondition),
		OtherItems:          otherItems,
	}, nil
}

func (cfg *apiConfig) getMonsterOtherItems(r *http.Request, mon database.Monster, monItemsID int32) ([]PossibleItem, error) {
	dbOtherItems, err := cfg.db.GetMonsterOtherItems(r.Context(), monItemsID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("Couldn't retrieve other items of Monster %s Version %d", mon.Name, *h.NullInt32ToPtr(mon.Version)), err)
	}

	otherItems := []PossibleItem{}

	for _, item := range dbOtherItems {
		possibleItem := cfg.newPossibleItem(item.ItemType.ItemType, item.Item.String, item.ItemID.Int32, item.Amount.Int32, anyToInt32(item.Chance))

		otherItems = append(otherItems, possibleItem)
	}

	return otherItems, nil
}
