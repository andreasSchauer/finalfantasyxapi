package api

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getMixCombinations(r *http.Request, mix seedtypes.Mix) ([]apitypes.MixCombination, error) {
	combinations := mix.PossibleCombinations

	best, err := Q.Mixes[ts.QpnBest].ParseBool(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}
	if best {
		combinations = mix.BestCombinations
	}

	itemID, err := Q.Mixes[ts.QpnContainsItem].ParseNameID(r, E.Items.resTypeSingle, E.Items.objLookup)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}
	if itemID != 0 {
		item, _ := lookup.GetResourceByID(itemID, lookup.L.ItemsID)

		combinations = h.Filter(combinations, func(c seedtypes.MixCombination) bool {
			return item.Name == c.FirstItem || item.Name == c.SecondItem
		})
	}

	return ah.ConvertObjSlice(combinations, convertMixCombination), nil
}
