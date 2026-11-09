package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getOverdriveMode(r *http.Request, i handlerinput.Main[seedtypes.OverdriveMode, apitypes.OverdriveMode, apiresource.Named], id int32) (apitypes.OverdriveMode, error) {
	mode, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.OverdriveMode{}, err
	}

	response := apitypes.OverdriveMode{
		ID:          mode.ID,
		Name:        mode.Name,
		Description: mode.Description,
		Effect:      mode.Effect,
		Type:        mode.Type,
		FillRate:    mode.FillRate,
		Actions:     apilookups.E.Characters.ToResAmtsFromNames(mode.ActionsToLearn),
	}

	return response, nil
}

func retrieveOverdriveModes(r *http.Request, i handlerinput.Main[seedtypes.OverdriveMode, apitypes.OverdriveMode, apiresource.Named]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryEnum(r, enums.T.OverdriveModeType, ids, ts.QpnType, config.Cfg.DB.GetOverdriveModeIDsByType),
	})
}
