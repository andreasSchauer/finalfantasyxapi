package sections

import (
	"slices"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func convertItemAmountSimple(ia seedtypes.ItemAmount) string {
	return h.NameAmountString(ia.ItemName, nil, nil, ia.Amount)
}

func posItemToItemAmtSimple(posItem seedtypes.PossibleItem) string {
	return convertItemAmountSimple(posItem.ItemAmount)
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
