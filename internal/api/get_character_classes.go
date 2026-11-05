package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getCharacterClass(r *http.Request, i handlerInput[seedtypes.CharacterClass, CharacterClass, NamedAPIResource], id int32) (CharacterClass, error) {
	class, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return CharacterClass{}, err
	}

	rel, err := getCharClassRelationships(r, class)
	if err != nil {
		return CharacterClass{}, err
	}

	response := CharacterClass{
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

func retrieveCharacterClasses(r *http.Request, i handlerInput[seedtypes.CharacterClass, CharacterClass, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(r, ids, []IDFilter{
		i.queryEnumList(r, enums.T.CharacterClassCategory, ids, ts.QpnCategory, config.Cfg.DB.GetCharacterClassesIDsByCategory),
	})
}
