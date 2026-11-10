package get

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func GetOverdriveMode(r *http.Request, i handlerinput.Main[seedtypes.OverdriveMode, apitypes.OverdriveMode, apiresource.Named], id int32) (apitypes.OverdriveMode, error) {
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
