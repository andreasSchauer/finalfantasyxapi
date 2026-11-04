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

func newPossibleItem(srv *Server, item seedtypes.ItemAmount, chance int32) PossibleItem {
	return PossibleItem{
		ResourceAmount: E.AllItems.toResAmtFromName(srv, item),
		Chance:         chance,
	}
}

func convertPossibleItem(srv *Server, item seedtypes.PossibleItem) PossibleItem {
	return newPossibleItem(srv, item.ItemAmount, item.Chance)
}
