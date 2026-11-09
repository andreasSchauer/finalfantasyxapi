package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/assemble"
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

	rel, err := assemble.GetMixRelationships(r, mix)
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
