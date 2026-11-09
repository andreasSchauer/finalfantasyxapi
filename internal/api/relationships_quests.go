package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func GetSidequestRelationships(r *http.Request, sidequest seedtypes.Sidequest) (apitypes.Sidequest, error) {
	var rel apitypes.Sidequest
	var err error

	rel.Subquests, err = apilookups.E.Subquests.DBGetRelations(r.Context(), sidequest, config.Cfg.DB.GetSidequestSubquestIDs)
	if err != nil {
		return apitypes.Sidequest{}, err
	}

	return rel, nil
}


func CompleteSubquestResponse(subquest apitypes.Subquest) apitypes.Subquest {
	if subquest.ParentSidequest.Sidequest == "monster arena" {
		subquest.ArenaCreation = handlerinput.NamePtrToNamedAPIResPtr(apilookups.E.ArenaCreations, &subquest.Name, nil)
	}

	return subquest
}