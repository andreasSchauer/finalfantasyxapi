package assemble

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiconv"
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func GetMixRelationships(r *http.Request, mix seedtypes.Mix) (apitypes.Mix, error) {
	overdrive, _ := lookup.GetResourceByID(mix.OverdriveID, lookup.L.OverdrivesID)
	var err error

	rel := apitypes.Mix{
		Overdrive:   handlerinput.NameToNamedAPIResource(apilookups.E.Overdrives, overdrive.Name, nil),
		Description: overdrive.Description,
		Effect:      overdrive.Effect,
	}

	rel.Combinations, err = getMixCombinations(r, mix)
	if err != nil {
		return apitypes.Mix{}, err
	}

	return rel, nil
}

func getMixCombinations(r *http.Request, mix seedtypes.Mix) ([]apitypes.MixCombination, error) {
	combinations := mix.PossibleCombinations

	best, err := apilookups.Q.Mixes[ts.QpnBest].ParseBool(r)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}
	if best {
		combinations = mix.BestCombinations
	}

	itemID, err := apilookups.Q.Mixes[ts.QpnContainsItem].ParseNameID(r, apilookups.E.Items.ResTypeSingle, apilookups.E.Items.ObjLookup)
	if errs.ErrExceptEmptyQuery(err) {
		return nil, err
	}
	if itemID != 0 {
		item, _ := lookup.GetResourceByID(itemID, lookup.L.ItemsID)

		combinations = h.Filter(combinations, func(c seedtypes.MixCombination) bool {
			return item.Name == c.FirstItem || item.Name == c.SecondItem
		})
	}

	return ah.ConvertObjSlice(combinations, apiconv.ConvertMixCombination), nil
}
