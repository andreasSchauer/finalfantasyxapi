package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (s *Server) getPlayerUnit(r *http.Request, i handlerInput[seedtypes.PlayerUnit, PlayerUnit, TypedAPIResource], id int32) (PlayerUnit, error) {
	unit, err := i.queryParamsVerifyGet(r, id)
	if err != nil {
		return PlayerUnit{}, err
	}

	var response PlayerUnit

	switch unit.Type {
	case database.UnitTypeCharacter:
		character, _ := lookup.GetResource(unit, lookup.L.Characters)

		rel, err := getCharacterRelationships(s, r, character)
		if err != nil {
			return PlayerUnit{}, err
		}

		response = PlayerUnit{
			ID:               unit.ID,
			Name:             unit.Name,
			Type:             unit.Type,
			TypedUnit:        nameToNamedAPIResource(s, E.Characters, unit.Name, nil),
			Area:             locAreaToAreaAPIResource(s, E.Areas, character.LocationArea),
			CelestialWeapon:  rel.CelestialWeapon,
			CharacterClasses: rel.CharacterClasses,
		}

	case database.UnitTypeAeon:
		aeon, _ := lookup.GetResource(unit, lookup.L.Aeons)

		rel, err := getAeonRelationships(s, r, aeon)
		if err != nil {
			return PlayerUnit{}, err
		}

		response = PlayerUnit{
			ID:               unit.ID,
			Name:             unit.Name,
			Type:             unit.Type,
			TypedUnit:        nameToNamedAPIResource(s, E.Aeons, unit.Name, nil),
			Area:             locAreaToAreaAPIResource(s, E.Areas, aeon.LocationArea),
			CelestialWeapon:  rel.CelestialWeapon,
			CharacterClasses: rel.CharacterClasses,
		}
	}

	return response, nil
}

func (s *Server) retrievePlayerUnits(r *http.Request, i handlerInput[seedtypes.PlayerUnit, PlayerUnit, TypedAPIResource]) ([]int32, error) {
	ids, err := i.queryParamsVerifyRetrieve(r)
	if err != nil {
		return nil, err
	}

	return i.filterIDs(s, r, ids, []IDFilter{
		i.queryEnum(r, enums.T.UnitType, ids, ts.QpnType, s.Cfg.DB.GetPlayerUnitIDsByType),
	})
}
