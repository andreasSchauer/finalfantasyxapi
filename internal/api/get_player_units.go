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
		character, _ := seeding.GetResource(unit.Name, cfg.l.Characters)

		rel, err := getCharacterRelationships(cfg, r, character)
		if err != nil {
			return PlayerUnit{}, err
		}

		response = PlayerUnit{
			ID:               unit.ID,
			Name:             unit.Name,
			Type:             newNamedAPIResourceFromEnum(cfg, cfg.e.unitType.endpoint, string(unit.Type), cfg.t.UnitType),
			TypedUnit:        nameToNamedAPIResource(cfg, cfg.e.characters, unit.Name, nil),
			Area:             locAreaToAreaAPIResource(cfg, cfg.e.areas, character.LocationArea),
			CelestialWeapon:  rel.CelestialWeapon,
			CharacterClasses: rel.CharacterClasses,
		}

	case database.UnitTypeAeon:
		aeon, _ := seeding.GetResource(unit.Name, cfg.l.Aeons)

		rel, err := getAeonRelationships(cfg, r, aeon)
		if err != nil {
			return PlayerUnit{}, err
		}

		response = PlayerUnit{
			ID:               unit.ID,
			Name:             unit.Name,
			Type:             newNamedAPIResourceFromEnum(cfg, cfg.e.unitType.endpoint, string(unit.Type), cfg.t.UnitType),
			TypedUnit:        nameToNamedAPIResource(cfg, cfg.e.aeons, unit.Name, nil),
			Area:             locAreaToAreaAPIResource(cfg, cfg.e.areas, aeon.LocationArea),
			CelestialWeapon:  rel.CelestialWeapon,
			CharacterClasses: rel.CharacterClasses,
		}
	}

	return response, nil
}

func (cfg *Config) retrievePlayerUnits(r *http.Request, i handlerInput[seeding.PlayerUnit, PlayerUnit, TypedAPIResource, TypedAPIResourceList]) (TypedAPIResourceList, error) {
	resources, err := retrieveAPIResources(cfg, r, i)
	if err != nil {
		return TypedAPIResourceList{}, err
	}

	return filterAPIResources(cfg, r, i, resources, []filteredResList[TypedAPIResource]{
		frl(enumQuery(cfg, r, i, cfg.t.UnitType, resources, "type", cfg.db.GetPlayerUnitIDsByType)),
	})
}
