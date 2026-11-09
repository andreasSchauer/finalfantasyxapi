package assemble

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/config"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func GetAbilityRelationships(r *http.Request, ability seedtypes.Ability) (apitypes.Ability, error) {
	var err error
	var rel apitypes.Ability

	rel.Monsters, err = apilookups.E.Monsters.DBGetRelations(r.Context(), ability, config.Cfg.DB.GetAbilityMonsterIDs)
	if err != nil {
		return apitypes.Ability{}, err
	}

	return rel, nil
}

func GetEnemyAbilityRelationships(r *http.Request, ability seedtypes.EnemyAbility) (apitypes.EnemyAbility, error) {
	var err error
	var rel apitypes.EnemyAbility

	rel.Monsters, err = apilookups.E.Monsters.DBGetRelations(r.Context(), ability, config.Cfg.DB.GetEnemyAbilityMonsterIDs)
	if err != nil {
		return apitypes.EnemyAbility{}, err
	}

	return rel, nil
}

func GetItemAbilityRelationships(r *http.Request, ability seedtypes.ItemAbility) apitypes.ItemAbility {
	item, _ := lookup.GetResourceByID(ability.ID, lookup.L.ItemsID)

	rel := apitypes.ItemAbility{
		Item:         handlerinput.IDToNamedAPIResource(apilookups.E.Items, item.ID),
		Description:  item.Description,
		Effect:       item.Effect,
		Category:     item.Category,
		RelatedStats: handlerinput.NamesToNamedAPIResources(apilookups.E.Stats, item.RelatedStats),
		Cursor:       item.Cursor,
	}

	if item.Usability == string(database.ItemUsabilityOutsideBattle) || item.Usability == string(database.ItemUsabilityAlways) {
		rel.CanUseOutsideBattle = true
	}

	return rel
}

func GetOverdriveAbilityRelationships(r *http.Request, ability seedtypes.OverdriveAbility) (apitypes.OverdriveAbility, error) {
	overdriveIDs, err := config.Cfg.DB.GetOverdriveAbilityOverdriveIDs(r.Context(), ability.ID)
	if err != nil {
		return apitypes.OverdriveAbility{}, errs.NewHTTPError(http.StatusInternalServerError, "couldn't get parent overdrive", err)
	}
	overdrive, _ := lookup.GetResourceByID(overdriveIDs[0], lookup.L.OverdrivesID)

	rel := apitypes.OverdriveAbility{
		Overdrives:       apilookups.E.Overdrives.ToAPIResources(overdriveIDs),
		OverdriveCommand: handlerinput.NamePtrToNamedAPIResPtr(apilookups.E.OverdriveCommands, overdrive.OverdriveCommand, nil),
		User:             handlerinput.NameToNamedAPIResource(apilookups.E.CharacterClasses, overdrive.User, nil),
	}

	return rel, nil
}

func GetPlayerAbilityRelationships(r *http.Request, ability seedtypes.PlayerAbility) (apitypes.PlayerAbility, error) {
	var rel apitypes.PlayerAbility
	var err error

	rel.Monsters, err = apilookups.E.Monsters.DBGetRelations(r.Context(), ability, config.Cfg.DB.GetPlayerAbilityMonsterIDs)
	if err != nil {
		return apitypes.PlayerAbility{}, err
	}

	return rel, nil
}
