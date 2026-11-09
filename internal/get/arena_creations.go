package get

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/assemble"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func GetArenaCreation(r *http.Request, i handlerinput.Main[seedtypes.ArenaCreation, apitypes.ArenaCreation, apiresource.Named], id int32) (apitypes.ArenaCreation, error) {
	creation, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.ArenaCreation{}, err
	}

	rel, err := assemble.GetArenaCreationRelationships(r, creation)
	if err != nil {
		return apitypes.ArenaCreation{}, err
	}

	response := apitypes.ArenaCreation{
		ID:                        creation.ID,
		Name:                      creation.Name,
		Category:                  creation.Category,
		Monster:                   handlerinput.IDToNamedAPIResource(apilookups.E.Monsters, *creation.MonsterID),
		ParentSubquest:            handlerinput.IDToQuestAPIResource(apilookups.E.Subquests, creation.SubquestID),
		Reward:                    rel.Reward,
		RequiredCatchAmount:       creation.Amount,
		UnlockedCreationsCategory: creation.CreationsUnlockedCategory,
		RequiredMonsters:          rel.RequiredMonsters,
	}

	return response, nil
}

func RetrieveArenaCreations(r *http.Request, i handlerinput.Main[seedtypes.ArenaCreation, apitypes.ArenaCreation, apiresource.Named]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryEnumList(r, enums.T.ArenaCreationCategory, ids, ts.QpnCategory, config.Cfg.DB.GetArenaCreationIDsByCategory),
	})
}
