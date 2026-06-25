package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getCharacterClass(r *http.Request, i handlerInput[seeding.CharacterClass, CharacterClass, NamedAPIResource, NamedApiResourceList], id int32) (CharacterClass, error) {
	class, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return CharacterClass{}, err
	}

	rel, err := getCharClassRelationships(cfg, r, class)
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

func (cfg *Config) retrieveCharacterClasses(r *http.Request, i handlerInput[seeding.CharacterClass, CharacterClass, NamedAPIResource, NamedApiResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []filteredIdList{
		fidl(enumListQuery(cfg, r, i, cfg.t.CharacterClassCategory, ids, qpnCategory, cfg.db.GetCharacterClassesIDsByCategory)),
	})
}
