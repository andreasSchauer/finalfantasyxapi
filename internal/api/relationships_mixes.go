package api

import (
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getMixCombinations(cfg *Config, r *http.Request, mix seeding.Mix) ([]MixCombination, error) {
	combinations := mix.PossibleCombinations

	best, err := parseBooleanQuery(r, cfg.q.mixes["best"])
	if errIsNotEmptyQuery(err) {
		return nil, err
	}
	if best {
		combinations = mix.BestCombinations
	}

	itemID, err := parseNameIdQuery(r, cfg.q.mixes["contains_item"], cfg.e.items.resourceType, cfg.e.items.objLookup)
	if errIsNotEmptyQuery(err) {
		return nil, err
	}
	if itemID != 0 {
		item, _ := seeding.GetResourceByID(itemID, cfg.l.ItemsID)

		combinations = h.Filter(combinations, func(c seeding.MixCombination) bool {
			return item.Name == c.FirstItem || item.Name == c.SecondItem
		})
	}

	return convertObjSlice(cfg, combinations, convertMixCombination), nil
}