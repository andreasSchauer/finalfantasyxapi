package api

import (
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getMixCombinations(cfg *Config, r *http.Request, mix seeding.Mix) ([]MixCombination, error) {
	combinations := mix.PossibleCombinations

	best, err := parseBooleanQuery(r, cfg.Q.mixes[QpnBest])
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if best {
		combinations = mix.BestCombinations
	}

	itemID, err := parseNameIdQuery(r, cfg.Q.mixes[QpnContainsItem], cfg.E.Items.resTypeSingle, cfg.E.Items.objLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if itemID != 0 {
		item, _ := seeding.GetResourceByID(itemID, cfg.L.ItemsID)

		combinations = h.Filter(combinations, func(c seeding.MixCombination) bool {
			return item.Name == c.FirstItem || item.Name == c.SecondItem
		})
	}

	return convertObjSlice(cfg, combinations, convertMixCombination), nil
}
