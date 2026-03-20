package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getCharClassRelationships(cfg *Config, r *http.Request, class seeding.CharacterClass) (CharacterClass, error) {
	units, err := getResourcesDbItem(cfg, r, cfg.e.playerUnits, class, cfg.db.GetCharacterClassUnitIDs)
	if err != nil {
		return CharacterClass{}, err
	}

	defaultAbilities, err := getResourcesDbItem(cfg, r, cfg.e.abilities, class, cfg.db.GetCharacterClassDefaultAbilityIDs)
	if err != nil {
		return CharacterClass{}, err
	}

	learnableAbilities, err := getClassLearnableAbilities(cfg, r, class, defaultAbilities)
	if err != nil {
		return CharacterClass{}, err
	}

	defaultOverdrives, err := getResourcesDbItem(cfg, r, cfg.e.overdrives, class, cfg.db.GetCharacterClassDefaultOverdriveIDs)
	if err != nil {
		return CharacterClass{}, err
	}

	learnableOverdrives, err := getResourcesDbItem(cfg, r, cfg.e.overdrives, class, cfg.db.GetCharacterClassLearnableOverdriveIDs)
	if err != nil {
		return CharacterClass{}, err
	}

	submenus, err := getResourcesDbItem(cfg, r, cfg.e.submenus, class, cfg.db.GetCharacterClassSubmenuIDs)
	if err != nil {
		return CharacterClass{}, err
	}

	charClass := CharacterClass{
		Members:             units,
		DefaultAbilities:    defaultAbilities,
		LearnableAbilities:  learnableAbilities,
		DefaultOverdrives:   defaultOverdrives,
		LearnableOverdrives: learnableOverdrives,
		Submenus:            submenus,
	}

	return charClass, nil
}

func getClassUnits(cfg *Config, r *http.Request, class seeding.CharacterClass) ([]NamedAPIResource, error) {
	unitIDs, err := cfg.db.GetCharacterClassUnitIDs(r.Context(), class.ID)
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get units of class '%s'", class.Name), err)
	}

	resources := []NamedAPIResource{}

	for _, id := range unitIDs {
		unit, _ := seeding.GetResourceByID(id, cfg.l.PlayerUnitsID)
		unitRes := createPlayerUnitResource(cfg, unit.Name, unit.Type)
		resources = append(resources, unitRes)
	}

	return resources, nil
}

func getClassLearnableAbilities(cfg *Config, r *http.Request, class seeding.CharacterClass, defaultAbilities []TypedAPIResource) ([]TypedAPIResource, error) {
	allAbilities, err := getResourcesDbItem(cfg, r, cfg.e.abilities, class, cfg.db.GetCharacterClassLearnableAbilityIDs)
	if err != nil {
		return nil, err
	}

	learnableAbilities := removeResourcesURL(allAbilities, defaultAbilities)
	return learnableAbilities, nil
}
