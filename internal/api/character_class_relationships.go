package api

import (
	"context"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func getCharClassRelationships(cfg *Config, r *http.Request, class seeding.CharacterClass) (CharacterClass, error) {
	units, err := getClassUnits(cfg, r, class)
	if err != nil {
		return CharacterClass{}, err
	}

	defaultAbilities, err := createAbilityResourceSlice(cfg, r, class, cfg.db.GetCharacterClassDefaultAbilityIDs)
	if err != nil {
		return CharacterClass{}, err
	}

	learnableAbilities, err := getClassLearnableAbilities(cfg, r, class, defaultAbilities)
	if err != nil {
		return CharacterClass{}, err
	}

	defaultOverdrives, err := getResourcesDB(cfg, r, cfg.e.overdrives, class, cfg.db.GetCharacterClassDefaultOverdriveIDs)
	if err != nil {
		return CharacterClass{}, err
	}

	learnableOverdrives, err := getResourcesDB(cfg, r, cfg.e.overdrives, class, cfg.db.GetCharacterClassLearnableOverdriveIDs)
	if err != nil {
		return CharacterClass{}, err
	}

	submenus, err := getResourcesDB(cfg, r, cfg.e.submenus, class, cfg.db.GetCharacterClassSubmenuIDs)
	if err != nil {
		return CharacterClass{}, err
	}

	charClass := CharacterClass{
		Units: 				units,
		DefaultAbilities: 	defaultAbilities,
		LearnableAbilities: learnableAbilities,
		DefaultOverdrives: defaultOverdrives,
		LearnableOverdrives: learnableOverdrives,
		Submenus:			submenus,
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


func getClassLearnableAbilities(cfg *Config, r *http.Request, class seeding.CharacterClass, defaultAbilities []NamedAPIResource) ([]NamedAPIResource, error) {
	allAbilities, err := createAbilityResourceSlice(cfg, r, class, cfg.db.GetCharacterClassLearnableAbilityIDs)
	if err != nil {
		return nil, err
	}

	learnableAbilities := removeResources(allAbilities, defaultAbilities)
	return learnableAbilities, nil
}



func createAbilityResourceSlice[T seeding.LookupableID](cfg *Config, r *http.Request, item T, dbQuery func (context.Context, int32) ([]int32, error)) ([]NamedAPIResource, error) {
	abilityIDs, err := dbQuery(r.Context(), item.GetID())
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get abilities of '%s'", item), err)
	}

	abilities := []NamedAPIResource{}

	for _, id := range abilityIDs {
		ability, _ := seeding.GetResourceByID(id, cfg.l.AbilitiesID)
		abilityRes := createAbilityResource(cfg, ability.Name, ability.Version, ability.Type)
		abilities = append(abilities, abilityRes)
	}

	return abilities, nil
}