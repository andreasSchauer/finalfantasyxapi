package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getCharacter(r *http.Request, i handlerInput[seeding.Character, Character, NamedAPIResource, NamedApiResourceList], id int32) (Character, error) {
	character, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return Character{}, err
	}

	rel, err := getCharacterRelationships(cfg, r, character)
	if err != nil {
		return Character{}, err
	}

	response := Character{
		ID:                     character.ID,
		Name:                   character.Name,
		UntypedUnit:            idToTypedAPIResource(cfg, cfg.e.playerUnits, character.PlayerUnit.ID),
		Area:                   locAreaToAreaAPIResource(cfg, cfg.e.areas, character.LocationArea),
		IsStoryBased:           character.IsStoryBased,
		CanFightUnderwater:     character.CanFightUnderwater,
		PhysAtkRange:           character.PhysAtkRange,
		WeaponType:             character.WeaponType,
		ArmorType:              character.ArmorType,
		CelestialWeapon:        rel.CelestialWeapon,
		OverdriveCommand:       rel.OverdriveCommand,
		CharacterClasses:       rel.CharacterClasses,
		BaseStats:              toResAmtType(cfg, cfg.e.stats, character.BaseStats, newBaseStat),
		DefaultPlayerAbilities: rel.DefaultPlayerAbilities,
		StdSgPlayerAbilities:   rel.StdSgPlayerAbilities,
		ExpSgPlayerAbilities:   rel.ExpSgPlayerAbilities,
		OverdriveModes:         rel.OverdriveModes,
	}

	response.Stats = createStats(response.BaseStats)

	return response, nil
}

func (cfg *Config) retrieveCharacters(r *http.Request, i handlerInput[seeding.Character, Character, NamedAPIResource, NamedApiResourceList]) (NamedApiResourceList, error) {
	ids, err := verifyParamsAndRetrieve(cfg, r, i)
	if err != nil {
		return NamedApiResourceList{}, err
	}

	return filterIDs(cfg, r, i, ids, []filteredIdList{
		fidl(boolQuery(r, i, ids, "story_based", cfg.db.GetCharacterIDsStoryBased)),
		fidl(boolQuery(r, i, ids, "underwater", cfg.db.GetCharacterIDsCanFightUnderwater)),
	})
}
