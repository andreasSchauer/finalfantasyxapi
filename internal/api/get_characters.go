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
		StdSphereGrid:          rel.StdSphereGrid,
		ExpSphereGrid:          rel.ExpSphereGrid,
		OverdriveModes:         rel.OverdriveModes,
	}

	response.BaseStats, err = applyOsgStats(cfg, r, response)
	if err != nil {
		return Character{}, err
	}

	response.Stats = createStats(response.BaseStats)

	return response, nil
}

func (cfg *Config) retrieveCharacters(r *http.Request, i handlerInput[seeding.Character, Character, NamedAPIResource, NamedApiResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []IdFilter{
		boolQuery(r, i, ids, qpnStoryBased, cfg.db.GetCharacterIDsStoryBased),
		boolQuery(r, i, ids, qpnUnderwater, cfg.db.GetCharacterIDsCanFightUnderwater),
	})
}
