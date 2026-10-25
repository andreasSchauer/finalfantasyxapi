package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getPlayerUnit(r *http.Request, i handlerInput[seeding.PlayerUnit, PlayerUnit, TypedAPIResource, TypedAPIResourceList], id int32) (PlayerUnit, error) {
	unit, err := verifyParamsAndGet(r, i, id)
	if err != nil {
		return PlayerUnit{}, err
	}

	var response PlayerUnit

	switch unit.Type {
	case database.UnitTypeCharacter:
		character, _ := seeding.GetResource(unit, cfg.L.Characters)

		rel, err := getCharacterRelationships(cfg, r, character)
		if err != nil {
			return PlayerUnit{}, err
		}

		response = PlayerUnit{
			ID:               unit.ID,
			Name:             unit.Name,
			Type:             unit.Type,
			TypedUnit:        nameToNamedAPIResource(cfg, cfg.E.Characters, unit.Name, nil),
			Area:             locAreaToAreaAPIResource(cfg, cfg.E.Areas, character.LocationArea),
			CelestialWeapon:  rel.CelestialWeapon,
			CharacterClasses: rel.CharacterClasses,
		}

	case database.UnitTypeAeon:
		aeon, _ := seeding.GetResource(unit, cfg.L.Aeons)

		rel, err := getAeonRelationships(cfg, r, aeon)
		if err != nil {
			return PlayerUnit{}, err
		}

		response = PlayerUnit{
			ID:               unit.ID,
			Name:             unit.Name,
			Type:             unit.Type,
			TypedUnit:        nameToNamedAPIResource(cfg, cfg.E.Aeons, unit.Name, nil),
			Area:             locAreaToAreaAPIResource(cfg, cfg.E.Areas, aeon.LocationArea),
			CelestialWeapon:  rel.CelestialWeapon,
			CharacterClasses: rel.CharacterClasses,
		}
	}

	return response, nil
}

func (cfg *Config) retrievePlayerUnits(r *http.Request, i handlerInput[seeding.PlayerUnit, PlayerUnit, TypedAPIResource, TypedAPIResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []IdFilter{
		enumQuery(r, i, cfg.T.UnitType, ids, QpnType, cfg.db.GetPlayerUnitIDsByType),
	})
}
