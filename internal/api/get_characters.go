package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (s *Server) getCharacter(r *http.Request, i handlerInput[seedtypes.Character, Character, NamedAPIResource], id int32) (Character, error) {
	character, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return Character{}, err
	}

	rel, err := getCharacterRelationships(s, r, character)
	if err != nil {
		return Character{}, err
	}

	response := Character{
		ID:                     character.ID,
		Name:                   character.Name,
		UntypedUnit:            idToTypedAPIResource(s, E.PlayerUnits, character.PlayerUnit.ID),
		Area:                   locAreaToAreaAPIResource(s, E.Areas, character.LocationArea),
		IsStoryBased:           character.IsStoryBased,
		CanFightUnderwater:     character.CanFightUnderwater,
		PhysAtkRange:           character.PhysAtkRange,
		WeaponType:             character.WeaponType,
		ArmorType:              character.ArmorType,
		CelestialWeapon:        rel.CelestialWeapon,
		OverdriveCommand:       rel.OverdriveCommand,
		CharacterClasses:       rel.CharacterClasses,
		BaseStats:              E.Stats.toResAmtType(s, character.BaseStats, newBaseStat),
		DefaultPlayerAbilities: rel.DefaultPlayerAbilities,
		StdSphereGrid:          rel.StdSphereGrid,
		ExpSphereGrid:          rel.ExpSphereGrid,
		OverdriveModes:         rel.OverdriveModes,
	}

	response.BaseStats, err = applyOsgStats(r, response)
	if err != nil {
		return Character{}, err
	}

	response.RegenAmounts = calcRegenAmounts(response.BaseStats)
	response.CelestialFormula = getCharacterClstlFormula(response)

	response.AgilityParameters, err = getUnitAgilityParams(s, r, response.BaseStats)
	if err != nil {
		return Character{}, err
	}

	return response, nil
}

func (s *Server) retrieveCharacters(r *http.Request, i handlerInput[seedtypes.Character, Character, NamedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(s, r, ids, []IDFilter{
		i.queryBool(r, ids, ts.QpnStoryBased, s.Cfg.DB.GetCharacterIDsStoryBased),
		i.queryBool(r, ids, ts.QpnUnderwater, s.Cfg.DB.GetCharacterIDsCanFightUnderwater),
	})
}
