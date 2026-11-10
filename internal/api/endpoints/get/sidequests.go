package get

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apiconv"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/general/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func GetSidequest(r *http.Request, i handlerinput.Main[seedtypes.Sidequest, apitypes.Sidequest, apiresource.Quest], id int32) (apitypes.Sidequest, error) {
	sidequest, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Sidequest{}, err
	}

	rel, err := getSidequestRelationships(r, sidequest)
	if err != nil {
		return apitypes.Sidequest{}, err
	}

	response := apitypes.Sidequest{
		ID:           sidequest.ID,
		Name:         sidequest.Name,
		UntypedQuest: handlerinput.IDToQuestAPIResource(apilookups.E.Quests, sidequest.Quest.ID),
		Availability: sidequest.Availability,
		IsRepeatable: sidequest.IsRepeatable,
		Completion:   ah.ConvertObjPtr(sidequest.Completion, apiconv.ConvertQuestCompletion),
		Subquests:    rel.Subquests,
	}

	return response, nil
}

func getSidequestRelationships(r *http.Request, sidequest seedtypes.Sidequest) (apitypes.Sidequest, error) {
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
