package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	th "github.com/andreasSchauer/finalfantasyxapi/internal/typehelpers"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getCharacter(r *http.Request, i handlerinput.Main[seedtypes.Character, apitypes.Character, apiresource.Named], id int32) (apitypes.Character, error) {
	character, err := i.VerifyQueriesGet(r, id)
	if err != nil {
		return apitypes.Character{}, err
	}

	rel, err := getCharacterRelationships(r, character)
	if err != nil {
		return apitypes.Character{}, err
	}

	response := apitypes.Character{
		ID:                     character.ID,
		Name:                   character.Name,
		UntypedUnit:            handlerinput.IDToTypedAPIResource(apilookups.E.PlayerUnits, character.PlayerUnit.ID),
		Area:                   handlerinput.LocAreaToAreaAPIResource(apilookups.E.Areas, character.LocationArea),
		IsStoryBased:           character.IsStoryBased,
		CanFightUnderwater:     character.CanFightUnderwater,
		PhysAtkRange:           character.PhysAtkRange,
		WeaponType:             character.WeaponType,
		ArmorType:              character.ArmorType,
		CelestialWeapon:        rel.CelestialWeapon,
		OverdriveCommand:       rel.OverdriveCommand,
		CharacterClasses:       rel.CharacterClasses,
		BaseStats:              apilookups.E.Stats.ToResAmtType(character.BaseStats, th.NewBaseStat),
		DefaultPlayerAbilities: rel.DefaultPlayerAbilities,
		StdSphereGrid:          rel.StdSphereGrid,
		ExpSphereGrid:          rel.ExpSphereGrid,
		OverdriveModes:         rel.OverdriveModes,
	}

	response.BaseStats, err = applyOsgStats(r, response)
	if err != nil {
		return apitypes.Character{}, err
	}

	response.RegenAmounts = calcRegenAmounts(response.BaseStats)
	response.CelestialFormula = getCharacterClstlFormula(response)

	response.AgilityParameters, err = getUnitAgilityParams(r, response.BaseStats)
	if err != nil {
		return apitypes.Character{}, err
	}

	return response, nil
}

func retrieveCharacters(r *http.Request, i handlerinput.Main[seedtypes.Character, apitypes.Character, apiresource.Named]) ([]int32, error) {
	ids, err := i.VerifyQueriesRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []handlerinput.IDFilter{
		i.QueryBool(r, ids, ts.QpnStoryBased, config.Cfg.DB.GetCharacterIDsStoryBased),
		i.QueryBool(r, ids, ts.QpnUnderwater, config.Cfg.DB.GetCharacterIDsCanFightUnderwater),
	})
}
