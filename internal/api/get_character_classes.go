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

func (cfg *Config) retrieveCharacterClasses(r *http.Request, i handlerInput[seeding.CharacterClass, CharacterClass, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[NamedAPIResource]{
		frl(typeQuery(cfg, r, i, cfg.t.CharacterClassCategory, resources, "category", cfg.db.GetCharacterClassesIDsByCategory)),
	})
}
