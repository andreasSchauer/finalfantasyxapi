package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

// can be used for various other functions related to abilities
func createPlayerUnitResource(cfg *Config, name string, unitType database.UnitType) NamedAPIResource {
	var res NamedAPIResource

	switch unitType {
	case database.UnitTypeCharacter:
		res = nameToNamedAPIResource(cfg, cfg.e.characters, name, nil)
		
	case database.UnitTypeAeon:
		res = nameToNamedAPIResource(cfg, cfg.e.aeons, name, nil)
	}

	return res
}
