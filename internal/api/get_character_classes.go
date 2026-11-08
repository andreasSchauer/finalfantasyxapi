package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getCharacterClass(r *http.Request, i handlerinput.Main[seedtypes.CharacterClass, apitypes.CharacterClass, apiresource.Named], id int32) (apitypes.CharacterClass, error) {
	class, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.CharacterClass{}, err
	}

	rel, err := getCharClassRelationships(r, class)
	if err != nil {
		return apitypes.CharacterClass{}, err
	}

	response := apitypes.CharacterClass{
		ID:                  class.ID,
		Name:                class.Name,
		Category:            class.Category,
		Members:             rel.Members,
		DefaultAbilities:    rel.DefaultAbilities,
		DefaultOverdrives:   rel.DefaultOverdrives,
		LearnableAbilities:  rel.LearnableAbilities,
		LearnableOverdrives: rel.LearnableOverdrives,
		Submenus:            rel.Submenus,
	}

	return response, nil
}

func retrieveCharacterClasses(r *http.Request, i handlerinput.Main[seedtypes.CharacterClass, apitypes.CharacterClass, apiresource.Named]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryEnumList(r, enums.T.CharacterClassCategory, ids, ts.QpnCategory, config.Cfg.DB.GetCharacterClassesIDsByCategory),
	})
}
