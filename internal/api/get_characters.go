package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (cfg *Config) getCharacter(r *http.Request, i handlerInput[seedtypes.Character, Character, NamedAPIResource], id int32) (Character, error) {
	character, err := i.queryParamsVerifyGet(r, id)
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
		UntypedUnit:            idToTypedAPIResource(cfg, cfg.E.PlayerUnits, character.PlayerUnit.ID),
		Area:                   locAreaToAreaAPIResource(cfg, cfg.E.Areas, character.LocationArea),
		IsStoryBased:           character.IsStoryBased,
		CanFightUnderwater:     character.CanFightUnderwater,
		PhysAtkRange:           character.PhysAtkRange,
		WeaponType:             character.WeaponType,
		ArmorType:              character.ArmorType,
		CelestialWeapon:        rel.CelestialWeapon,
		OverdriveCommand:       rel.OverdriveCommand,
		CharacterClasses:       rel.CharacterClasses,
		BaseStats:              cfg.E.Stats.toResAmtType(cfg, character.BaseStats, newBaseStat),
		DefaultPlayerAbilities: rel.DefaultPlayerAbilities,
		StdSphereGrid:          rel.StdSphereGrid,
		ExpSphereGrid:          rel.ExpSphereGrid,
		OverdriveModes:         rel.OverdriveModes,
	}

	response.BaseStats, err = applyOsgStats(cfg, r, response)
	if err != nil {
		return Character{}, err
	}

	response.RegenAmounts = calcRegenAmounts(response.BaseStats)
	response.CelestialFormula = getCharacterClstlFormula(response)

	response.AgilityParameters, err = getUnitAgilityParams(cfg, r, response.BaseStats)
	if err != nil {
		return Character{}, err
	}

	return response, nil
}

func (cfg *Config) retrieveCharacters(r *http.Request, i handlerInput[seedtypes.Character, Character, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(cfg, r, ids, []IDFilter{
		i.queryBool(r, ids, ts.QpnStoryBased, cfg.db.GetCharacterIDsStoryBased),
		i.queryBool(r, ids, ts.QpnUnderwater, cfg.db.GetCharacterIDsCanFightUnderwater),
	})
}
