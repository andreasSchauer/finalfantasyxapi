package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (srv *Server) getOverdriveMode(r *http.Request, i handlerInput[seedtypes.OverdriveMode, OverdriveMode, NamedAPIResource], id int32) (OverdriveMode, error) {
	mode, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return OverdriveMode{}, err
	}

	response := OverdriveMode{
		ID:          mode.ID,
		Name:        mode.Name,
		Description: mode.Description,
		Effect:      mode.Effect,
		Type:        mode.Type,
		FillRate:    mode.FillRate,
		Actions:     E.Characters.toResAmtsFromNames(srv, mode.ActionsToLearn),
	}

	return response, nil
}

func (srv *Server) retrieveOverdriveModes(r *http.Request, i handlerInput[seedtypes.OverdriveMode, OverdriveMode, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(srv, r, ids, []IDFilter{
		i.queryEnum(r, enums.T.OverdriveModeType, ids, ts.QpnType, config.Cfg.DB.GetOverdriveModeIDsByType),
	})
}
