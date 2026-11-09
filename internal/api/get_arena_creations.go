package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getArenaCreation(r *http.Request, i handlerinput.Main[seedtypes.ArenaCreation, apitypes.ArenaCreation, apiresource.Named], id int32) (apitypes.ArenaCreation, error) {
	creation, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.ArenaCreation{}, err
	}

	subquest, _ := lookup.GetResourceByID(creation.SubquestID, lookup.L.SubquestsID)
	ia := subquest.Completion.Reward

	monsters, err := getArenaCreationMonsters(r, creation)
	if err != nil {
		return apitypes.ArenaCreation{}, err
	}

	response := apitypes.ArenaCreation{
		ID:                        creation.ID,
		Name:                      creation.Name,
		Category:                  creation.Category,
		Monster:                   handlerinput.IDToNamedAPIResource(apilookups.E.Monsters, *creation.MonsterID),
		ParentSubquest:            handlerinput.IDToQuestAPIResource(apilookups.E.Subquests, creation.SubquestID),
		Reward:                    apilookups.E.AllItems.ToResAmtFromName(ia),
		RequiredCatchAmount:       creation.Amount,
		UnlockedCreationsCategory: creation.CreationsUnlockedCategory,
		RequiredMonsters:          monsters,
	}

	return response, nil
}

func retrieveArenaCreations(r *http.Request, i handlerinput.Main[seedtypes.ArenaCreation, apitypes.ArenaCreation, apiresource.Named]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryEnumList(r, enums.T.ArenaCreationCategory, ids, ts.QpnCategory, config.Cfg.DB.GetArenaCreationIDsByCategory),
	})
}
