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

func newPossibleItem(cfg *Config, item seedtypes.ItemAmount, chance int32) PossibleItem {
	return PossibleItem{
		ResourceAmount: cfg.E.AllItems.toResAmtFromName(cfg, item),
		Chance:         chance,
	}
}

func convertPossibleItem(cfg *Config, item seedtypes.PossibleItem) PossibleItem {
	return newPossibleItem(cfg, item.ItemAmount, item.Chance)
}
