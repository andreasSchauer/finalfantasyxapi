package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func getPlayerUnit(r *http.Request, i HandlerInput[seedtypes.PlayerUnit, apitypes.PlayerUnit, apiresource.TypedAPIResource], id int32) (apitypes.PlayerUnit, error) {
	unit, err := i.QueryParamsVerifyGet(r, id)
	if err != nil {
		return apitypes.PlayerUnit{}, err
	}

	var response apitypes.PlayerUnit

	switch unit.Type {
	case database.UnitTypeCharacter:
		character, _ := lookup.GetResource(unit, lookup.L.Characters)

		rel, err := getCharacterRelationships(r, character)
		if err != nil {
			return apitypes.PlayerUnit{}, err
		}

		response = apitypes.PlayerUnit{
			ID:               unit.ID,
			Name:             unit.Name,
			Type:             unit.Type,
			TypedUnit:        nameToNamedAPIResource(E.Characters, unit.Name, nil),
			Area:             locAreaToAreaAPIResource(E.Areas, character.LocationArea),
			CelestialWeapon:  rel.CelestialWeapon,
			CharacterClasses: rel.CharacterClasses,
		}

	case database.UnitTypeAeon:
		aeon, _ := lookup.GetResource(unit, lookup.L.Aeons)

		rel, err := getAeonRelationships(r, aeon)
		if err != nil {
			return apitypes.PlayerUnit{}, err
		}

		response = apitypes.PlayerUnit{
			ID:               unit.ID,
			Name:             unit.Name,
			Type:             unit.Type,
			TypedUnit:        nameToNamedAPIResource(E.Aeons, unit.Name, nil),
			Area:             locAreaToAreaAPIResource(E.Areas, aeon.LocationArea),
			CelestialWeapon:  rel.CelestialWeapon,
			CharacterClasses: rel.CharacterClasses,
		}
	}

	return response, nil
}

func retrievePlayerUnits(r *http.Request, i HandlerInput[seedtypes.PlayerUnit, apitypes.PlayerUnit, apiresource.TypedAPIResource]) ([]int32, error) {
	ids, err := i.QueryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.FilterIDs(r, ids, []IDFilter{
		i.QueryEnum(r, enums.T.UnitType, ids, ts.QpnType, config.Cfg.DB.GetPlayerUnitIDsByType),
	})
}
