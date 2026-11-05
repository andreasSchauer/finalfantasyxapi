package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

type PossibleItem struct {
	ResourceAmount[TypedAPIResource]
	Chance int32 `json:"chance"`
}

func (ps PossibleItem) GetAPIResource() APIResource {
	return ps.Resource.GetAPIResource()
}

func newPossibleItem(item seedtypes.ItemAmount, chance int32) PossibleItem {
	return PossibleItem{
		ResourceAmount: E.AllItems.toResAmtFromName(item),
		Chance:         chance,
	}
}

func convertPossibleItem(item seedtypes.PossibleItem) PossibleItem {
	return newPossibleItem(item.ItemAmount, item.Chance)
}
