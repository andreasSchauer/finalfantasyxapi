package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type PossibleItem struct {
	ResourceAmount[TypedAPIResource]
	Chance int32 `json:"chance"`
}

func (ps PossibleItem) GetAPIResource() APIResource {
	return ps.Resource.GetAPIResource()
}

func newPossibleItem(cfg *Config, item seeding.ItemAmount, chance int32) PossibleItem {
	return PossibleItem{
		ResourceAmount: nameAmountToResourceAmount(cfg, cfg.e.allItems, item),
		Chance:         chance,
	}
}

func convertPossibleItem(cfg *Config, item seeding.PossibleItem) PossibleItem {
	return newPossibleItem(cfg, item.ItemAmount, item.Chance)
}
