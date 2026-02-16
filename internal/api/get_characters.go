package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getCharacter(r *http.Request, i handlerInput[seeding.Character, Character, NamedAPIResource, NamedApiResourceList], id int32) (Character, error) {
	character, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return Character{}, err
	}

	rel, err := getCharacterRelationships(cfg, r, character)
	if err != nil {
		return Character{}, err
	}

	response := Character{
		ID: 					character.ID,
		Name: 					character.Name,
		Area: 					idPtrToAreaAPIResPtr(cfg, cfg.e.areas, character.AreaID),
		StoryOnly: 				character.StoryOnly,
		CanFightUnderwater: 	character.CanFightUnderwater,
		PhysAtkRange: 			character.PhysAtkRange,
		WeaponType: 			character.WeaponType,
		ArmorType: 				character.ArmorType,
		CelestialWeapon: 		rel.CelestialWeapon,
		OverdriveCommand: 		rel.OverdriveCommand,
		CharacterClasses: 		rel.CharacterClasses,
		BaseStats: 				namesToResourceAmounts(cfg, cfg.e.stats, character.BaseStats, newBaseStat),
		DefaultAbilities: 		rel.DefaultAbilities,
		StdSphereGridAbilities: rel.StdSphereGridAbilities,
		ExpSphereGridAbilities: rel.ExpSphereGridAbilities,
		OverdriveModes: 		rel.OverdriveModes,
	}

	return response, nil
}

func (cfg *Config) retrieveCharacters(r *http.Request, i handlerInput[seeding.Character, Character, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[NamedAPIResource]{
		frl(boolQuery(cfg, r, i, resources, "story_based", cfg.db.GetCharacterIDsStoryOnly)),
	})
}
