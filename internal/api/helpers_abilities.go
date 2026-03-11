package api

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func createAbilityResource(cfg *Config, ref seeding.AbilityReference) NamedAPIResource {
	var res NamedAPIResource

	switch database.AbilityType(ref.AbilityType) {
	case database.AbilityTypeUnspecifiedAbility:
		res = nameToNamedAPIResource(cfg, cfg.e.unspecifiedAbilities, ref.Name, ref.Version)

	case database.AbilityTypePlayerAbility:
		res = nameToNamedAPIResource(cfg, cfg.e.playerAbilities, ref.Name, ref.Version)

	case database.AbilityTypeEnemyAbility:
		res = nameToNamedAPIResource(cfg, cfg.e.enemyAbilities, ref.Name, ref.Version)

	case database.AbilityTypeOverdriveAbility:
		res = nameToNamedAPIResource(cfg, cfg.e.overdriveAbilities, ref.Name, ref.Version)

	case database.AbilityTypeItemAbility:
		res = nameToNamedAPIResource(cfg, cfg.e.itemAbilities, ref.Name, ref.Version)

	case database.AbilityTypeTriggerCommand:
		res = nameToNamedAPIResource(cfg, cfg.e.triggerCommands, ref.Name, ref.Version)

	}
	return res
}

func getAbilityResourcesDB[T seeding.LookupableID](cfg *Config, r *http.Request, item T, dbQuery func(context.Context, int32) ([]int32, error)) ([]NamedAPIResource, error) {
	abilityIDs, err := dbQuery(r.Context(), item.GetID())
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get abilities of '%s'", item), err)
	}

	abilities := idsToAbilityResources(cfg, abilityIDs)

	return abilities, nil
}

func getAbilityResPtrDB[T seeding.LookupableID](cfg *Config, r *http.Request, item T, dbQuery func(context.Context, sql.NullInt32) (int32, error)) (*NamedAPIResource, error) {
	abilityID, err := convertDbQueryOne(dbQuery)(r.Context(), item.GetID())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get abilities of '%s'", item), err)
	}

	ability := idToAbilityResource(cfg, abilityID)
	return &ability, nil
}

func getAbilityResourcesDbNullable[T seeding.LookupableID](cfg *Config, r *http.Request, item T, dbQuery func(context.Context, sql.NullInt32) ([]int32, error)) ([]NamedAPIResource, error) {
	abilityIDs, err := convertDbQueryMany(dbQuery)(r.Context(), item.GetID())
	if err != nil {
		return nil, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get abilities of '%s'", item), err)
	}

	abilities := idsToAbilityResources(cfg, abilityIDs)
	return abilities, nil
}

func idsToAbilityResources(cfg *Config, ids []int32) []NamedAPIResource {
	abilities := []NamedAPIResource{}

	for _, id := range ids {
		abilityRes := idToAbilityResource(cfg, id)
		abilities = append(abilities, abilityRes)
	}

	return abilities
}

func idToAbilityResource(cfg *Config, id int32) NamedAPIResource {
	ability, _ := seeding.GetResourceByID(id, cfg.l.AbilitiesID)
	return createAbilityResource(cfg, ability.GetAbilityRef())
}

func refsToAbilityResources(cfg *Config, refs []seeding.AbilityReference) []NamedAPIResource {
	abilities := []NamedAPIResource{}

	for _, ref := range refs {
		abilityRes := createAbilityResource(cfg, ref)
		abilities = append(abilities, abilityRes)
	}

	return abilities
}
