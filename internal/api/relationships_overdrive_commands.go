package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func GetOverdriveCommandRelationships(r *http.Request, command seedtypes.OverdriveCommand) (apitypes.OverdriveCommand, error) {
	var rel apitypes.OverdriveCommand
	var err error

	rel.Overdrives, err = apilookups.E.Overdrives.DBGetRelations(r.Context(), command, handlerinput.ToIntManyNull(config.Cfg.DB.GetOverdriveCommandOverdriveIDs))
	if err != nil {
		return apitypes.OverdriveCommand{}, err
	}

	return rel, nil
}