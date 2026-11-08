package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getMix(r *http.Request, i handlerinput.Main[seedtypes.Mix, apitypes.Mix, apiresource.Named], id int32) (apitypes.Mix, error) {
	mix, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Mix{}, err
	}

	overdrive, _ := lookup.GetResourceByID(mix.OverdriveID, lookup.L.OverdrivesID)

	combinations, err := getMixCombinations(r, mix)
	if err != nil {
		return apitypes.Mix{}, err
	}

	response := apitypes.Mix{
		ID:           mix.ID,
		Name:         mix.Name,
		Category:     mix.Category,
		Overdrive:    handlerinput.NameToNamedAPIResource(E.Overdrives, overdrive.Name, nil),
		Description:  overdrive.Description,
		Effect:       overdrive.Effect,
		Combinations: combinations,
	}

	return response, nil
}

func retrieveMixes(r *http.Request, i handlerinput.Main[seedtypes.Mix, apitypes.Mix, apiresource.Named]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryEnumList(r, enums.T.MixCategory, ids, ts.QpnCategory, config.Cfg.DB.GetMixIDsByCategory),
		i.QueryIDWrapper(r, ids, ts.QpnReqItem, lookup.L.Items, getMixesByItem),
	})
}
