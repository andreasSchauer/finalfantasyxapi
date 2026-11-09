package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func GetRonsoRageRelationships(r *http.Request, ronsoRage seedtypes.RonsoRage) (apitypes.RonsoRage, error) {
	var rel apitypes.RonsoRage
	var err error

	rel.Monsters, err = apilookups.E.Monsters.DBGetRelations(r.Context(), ronsoRage, config.Cfg.DB.GetRonsoRageMonsterIDs)
	if err != nil {
		return apitypes.RonsoRage{}, err
	}

	return rel, nil
}