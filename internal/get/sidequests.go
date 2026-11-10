package get

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiconv"
	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
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
