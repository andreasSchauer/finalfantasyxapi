package api

import (
	"errors"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getAeonRelationships(cfg *Config, r *http.Request, ae seeding.Aeon) (Aeon, error) {
	celestialWeapon, err := getResPtrDB(cfg, r, cfg.e.celestialWeapons, ae, cfg.db.GetAeonCelestialWeaponID)
	if err != nil {
		return Aeon{}, err
	}

	characterClasses, err := getResourcesDB(cfg, r, cfg.e.characterClasses, ae, cfg.db.GetAeonCharClassIDs)
	if err != nil {
		return Aeon{}, err
	}

	aeonCommands, err := getResourcesDB(cfg, r, cfg.e.aeonCommands, ae, cfg.db.GetAeonAeonCommandIDs)
	if err != nil {
		return Aeon{}, err
	}

	overdrives, err := getResourcesDB(cfg, r, cfg.e.overdrives, ae, cfg.db.GetAeonOverdriveIDs)
	if err != nil {
		return Aeon{}, err
	}

	defaultAbilities, err := getResourcesDB(cfg, r, cfg.e.playerAbilities, ae, cfg.db.GetAeonDefaultAbilityIDs)
	if err != nil {
		return Aeon{}, err
	}

	aeon := Aeon{
		CelestialWeapon:  celestialWeapon,
		CharacterClasses: characterClasses,
		AeonCommands:     aeonCommands,
		Overdrives:       overdrives,
		DefaultAbilities: defaultAbilities,
	}

	return aeon, nil
}


func applyAeonStatsBattles(cfg *Config, r *http.Request, aeon Aeon, queryName string) ([]BaseStat, error) {
	queryParam := cfg.q.aeons[queryName]
	battles, err := parseIntQuery(r, queryParam)
	if errors.Is(err, errEmptyQuery) {
		return aeon.BaseStats, nil
	}
	if err != nil {
		return nil, err
	}

	i := battles / 30 - 1

	if battles < 60 {
		i = 0
	}

	seedAeon, _ := seeding.GetResourceByID(aeon.ID, cfg.l.AeonsID)
	newBaseStats := seedAeon.BaseStats.XVals[i].BaseStats
	baseStats := namesToResourceAmounts(cfg, cfg.e.stats, newBaseStats, newBaseStat)

	return baseStats, nil
}
