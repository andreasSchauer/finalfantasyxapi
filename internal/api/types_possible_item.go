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

func newPossibleItem(server *Server, item seedtypes.ItemAmount, chance int32) PossibleItem {
	return PossibleItem{
		ResourceAmount: E.AllItems.toResAmtFromName(server, item),
		Chance:         chance,
	}
}

func convertPossibleItem(server *Server, item seedtypes.PossibleItem) PossibleItem {
	return newPossibleItem(server, item.ItemAmount, item.Chance)
}
