package api

import (
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type ItemAmountSimple struct {
	ia            seeding.ItemAmount `json:"-"`
	ItemAmountStr string             `json:"item"`
	ItemType      database.ItemType  `json:"item_type"`
}

func convertItemAmountSimple(cfg *Config, ia seeding.ItemAmount) string {
	return h.NameAmountString(ia.ItemName, nil, nil, ia.Amount)
}

func posItemToItemAmtSimple(cfg *Config, posItem seeding.PossibleItem) string {
	return convertItemAmountSimple(cfg, posItem.ItemAmount)
}

func sortItemAmountsByID(s []seeding.ItemAmount) []seeding.ItemAmount {
	slices.SortStableFunc(s, func(a, b seeding.ItemAmount) int {
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
