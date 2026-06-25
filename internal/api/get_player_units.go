package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func (cfg *Config) getPlayerUnit(r *http.Request, i handlerInput[seeding.PlayerUnit, PlayerUnit, TypedAPIResource, TypedAPIResourceList], id int32) (PlayerUnit, error) {
	unit, err := verifyParamsAndGet(cfg, r, i, id)
	if err != nil {
		return PlayerUnit{}, err
	}

	var response PlayerUnit

	switch unit.Type {
	case database.UnitTypeCharacter:
		character, _ := seeding.GetResource(unit, cfg.l.Characters)

		rel, err := getCharacterRelationships(cfg, r, character)
		if err != nil {
			return PlayerUnit{}, err
		}

		response = PlayerUnit{
			ID:               unit.ID,
			Name:             unit.Name,
			Type:             enumToNamedAPIResource(cfg, cfg.e.unitType.endpoint, string(unit.Type), cfg.t.UnitType),
			TypedUnit:        nameToNamedAPIResource(cfg, cfg.e.characters, unit.Name, nil),
			Area:             locAreaToAreaAPIResource(cfg, cfg.e.areas, character.LocationArea),
			CelestialWeapon:  rel.CelestialWeapon,
			CharacterClasses: rel.CharacterClasses,
		}

	case database.UnitTypeAeon:
		aeon, _ := seeding.GetResource(unit, cfg.l.Aeons)

		rel, err := getAeonRelationships(cfg, r, aeon)
		if err != nil {
			return PlayerUnit{}, err
		}

		response = PlayerUnit{
			ID:               unit.ID,
			Name:             unit.Name,
			Type:             enumToNamedAPIResource(cfg, cfg.e.unitType.endpoint, string(unit.Type), cfg.t.UnitType),
			TypedUnit:        nameToNamedAPIResource(cfg, cfg.e.aeons, unit.Name, nil),
			Area:             locAreaToAreaAPIResource(cfg, cfg.e.areas, aeon.LocationArea),
			CelestialWeapon:  rel.CelestialWeapon,
			CharacterClasses: rel.CharacterClasses,
		}
	}

	return response, nil
}

func (cfg *Config) retrievePlayerUnits(r *http.Request, i handlerInput[seeding.PlayerUnit, PlayerUnit, TypedAPIResource, TypedAPIResourceList]) ([]int32, error) {
	ids, err := verifyParamsAndRetrieve(cfg, r, i)
	if err != nil {
		return nil, err
	}

	return filterIDs(cfg, r, i, ids, []filteredIdList{
		fidl(enumQuery(r, i, cfg.t.UnitType, ids, qpnType, cfg.db.GetPlayerUnitIDsByType)),
	})
}
