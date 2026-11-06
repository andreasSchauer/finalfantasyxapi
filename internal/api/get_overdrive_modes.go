package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getOverdriveMode(r *http.Request, i handlerInput[seedtypes.OverdriveMode, apitypes.OverdriveMode, apiresource.NamedAPIResource], id int32) (apitypes.OverdriveMode, error) {
	mode, err := i.queryParamsVerifyGet(r, id)
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
		Actions:     E.Characters.toResAmtsFromNames(mode.ActionsToLearn),
	}

	return response, nil
}

func retrieveOverdriveModes(r *http.Request, i handlerInput[seedtypes.OverdriveMode, apitypes.OverdriveMode, apiresource.NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(r, ids, []IDFilter{
		i.queryEnum(r, enums.T.OverdriveModeType, ids, ts.QpnType, config.Cfg.DB.GetOverdriveModeIDsByType),
	})
}
