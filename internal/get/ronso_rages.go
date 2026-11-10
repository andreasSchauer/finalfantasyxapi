package get

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func GetRonsoRage(r *http.Request, i handlerinput.Main[seedtypes.RonsoRage, apitypes.RonsoRage, apiresource.Named], id int32) (apitypes.RonsoRage, error) {
	ronsoRage, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.RonsoRage{}, err
	}

	rel, err := getRonsoRageRelationships(r, ronsoRage)
	if err != nil {
		return apitypes.RonsoRage{}, err
	}

	response := apitypes.RonsoRage{
		ID:          ronsoRage.ID,
		Name:        ronsoRage.Name,
		Description: ronsoRage.Description,
		Effect:      ronsoRage.Effect,
		Overdrive:   handlerinput.NameToNamedAPIResource(apilookups.E.Overdrives, ronsoRage.Name, nil),
		Monsters:    rel.Monsters,
	}

	return response, nil
}

func getRonsoRageRelationships(r *http.Request, ronsoRage seedtypes.RonsoRage) (apitypes.RonsoRage, error) {
	var rel apitypes.RonsoRage
	var err error

	rel.Monsters, err = apilookups.E.Monsters.DBGetRelations(r.Context(), ronsoRage, config.Cfg.DB.GetRonsoRageMonsterIDs)
	if err != nil {
		return apitypes.RonsoRage{}, err
	}

	return rel, nil
}
