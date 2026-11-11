package get

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apiconvert"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"

	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func Mix(r *http.Request, i handlerinput.Main[seedtypes.Mix, apitypes.Mix, apiresource.Named], id int32) (apitypes.Mix, error) {
	mix, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Mix{}, err
	}

	rel, err := getMixRelationships(r, mix)
	if err != nil {
		return apitypes.Mix{}, err
	}

	response := apitypes.Mix{
		ID:           mix.ID,
		Name:         mix.Name,
		Category:     mix.Category,
		Overdrive:    rel.Overdrive,
		Description:  rel.Description,
		Effect:       rel.Effect,
		Combinations: rel.Combinations,
	}

	return response, nil
}

func getMixRelationships(r *http.Request, mix seedtypes.Mix) (apitypes.Mix, error) {
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

	return ah.ConvertObjSlice(combinations, apiconvert.MixCombination), nil
}
