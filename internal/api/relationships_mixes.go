package api

import (
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getMixCombinations(cfg *Config, r *http.Request, mix seedtypes.Mix) ([]MixCombination, error) {
	combinations := mix.PossibleCombinations

	best, err := cfg.Q.mixes[ts.QpnBest].parseBool(r)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if best {
		combinations = mix.BestCombinations
	}

	itemID, err := cfg.Q.mixes[ts.QpnContainsItem].parseNameID(r, cfg.E.Items.resTypeSingle, cfg.E.Items.objLookup)
	if errExceptEmptyQuery(err) {
		return nil, err
	}
	if itemID != 0 {
		item, _ := lookup.GetResourceByID(itemID, lookup.L.ItemsID)

		combinations = h.Filter(combinations, func(c seedtypes.MixCombination) bool {
			return item.Name == c.FirstItem || item.Name == c.SecondItem
		})
	}

	return convertObjSlice(cfg, combinations, convertMixCombination), nil
}
