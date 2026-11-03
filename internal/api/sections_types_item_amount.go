package api

import (
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

type ItemAmountSimple struct {
	ia            seedtypes.ItemAmount `json:"-"`
	ItemAmountStr string               `json:"item"`
	ItemType      database.ItemType    `json:"item_type"`
}

func convertItemAmountSimple(cfg *Config, ia seedtypes.ItemAmount) string {
	return h.NameAmountString(ia.ItemName, nil, nil, ia.Amount)
}

func posItemToItemAmtSimple(cfg *Config, posItem seedtypes.PossibleItem) string {
	return convertItemAmountSimple(cfg, posItem.ItemAmount)
}

func sortItemAmountsByID(s []seedtypes.ItemAmount) []seedtypes.ItemAmount {
	slices.SortStableFunc(s, func(a, b seedtypes.ItemAmount) int {
		A := a.MasterItem.ID
		B := b.MasterItem.ID

		if A < B {
			return -1
		}

		if A > B {
			return 1
		}

		if A == B {
			if a.Amount < b.Amount {
				return -1
			}

			if a.Amount > b.Amount {
				return 1
			}
		}

		return 0
	})

	return s
}
