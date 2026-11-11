package get

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func PlayerUnit(r *http.Request, i handlerinput.Main[seedtypes.PlayerUnit, apitypes.PlayerUnit, apiresource.Typed], id int32) (apitypes.PlayerUnit, error) {
	unit, err := i.VerifyQueriesGet(r, id)
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
			TypedUnit:        handlerinput.NameToNamedAPIResource(apilookups.E.Characters, unit.Name, nil),
			Area:             handlerinput.LocAreaToAreaAPIResource(apilookups.E.Areas, character.LocationArea),
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
			TypedUnit:        handlerinput.NameToNamedAPIResource(apilookups.E.Aeons, unit.Name, nil),
			Area:             handlerinput.LocAreaToAreaAPIResource(apilookups.E.Areas, aeon.LocationArea),
			CelestialWeapon:  rel.CelestialWeapon,
			CharacterClasses: rel.CharacterClasses,
		}
	}

	return response, nil
}
