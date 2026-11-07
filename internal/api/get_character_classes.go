package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getCharacterClass(r *http.Request, i HandlerInput[seedtypes.CharacterClass, apitypes.CharacterClass, apiresource.NamedAPIResource], id int32) (apitypes.CharacterClass, error) {
	class, err := i.QueryParamsVerifyGet(r, id)
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

func retrieveCharacterClasses(r *http.Request, i HandlerInput[seedtypes.CharacterClass, apitypes.CharacterClass, apiresource.NamedAPIResource]) ([]int32, error) {
	ids, err := i.QueryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []IDFilter{
		i.QueryEnumList(r, enums.T.CharacterClassCategory, ids, ts.QpnCategory, config.Cfg.DB.GetCharacterClassesIDsByCategory),
	})
}
