package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)


type AbilitySimple struct {
	ID             		int32                	   	`json:"id"`
	URL            		string               	   	`json:"url"`
	Name           		string               	   	`json:"name"`
	Version        		*int32             	    	`json:"version,omitempty"`
	Specification  		*string            	    	`json:"specification,omitempty"`
	Type				*string						`json:"type,omitempty"`
	Rank				*int32						`json:"rank"`
	BattleInteractions	[]BattleInteractionSimple	`json:"battle_interactions"`
}

func (a AbilitySimple) GetURL() string {
	return a.URL
}

func createAbilitySimple(cfg *Config, r *http.Request, id int32) (SimpleResource, error) {
	i := cfg.e.abilities
	ability, _ := seeding.GetResourceByID(id, i.objLookupID)
	typeStr := string(ability.Type)
	var rank *int32

	if ability.Type == database.AbilityTypeOverdriveAbility {
		attributes, err := cfg.db.GetOverdriveAbilityAttributes(r.Context(), id)
		if err != nil {
			var zeroType SimpleResource
			return zeroType, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get attributes for %s", ability), err)
		}
		rank = h.NullInt32ToPtr(attributes.Rank)
	} else {
		rank = ability.Rank
	}

	abilitySimple := AbilitySimple{
		ID: 				ability.ID,
		URL: 				createResourceURL(cfg, i.endpoint, id),
		Name: 				ability.Name,
		Version: 			ability.Version,
		Specification: 		ability.Specification,
		Rank: 				rank,
		Type: 				&typeStr,
		BattleInteractions: getAbilityBattleInteractionsSimple(cfg, ability),
	}

	return abilitySimple, nil
}


func createEnemyAbilitySimple(cfg *Config, _ *http.Request, id int32) (SimpleResource, error) {
	i := cfg.e.enemyAbilities
	ability, _ := seeding.GetResourceByID(id, i.objLookupID)

	abilitySimple := AbilitySimple{
		ID: 				ability.ID,
		URL: 				createResourceURL(cfg, i.endpoint, id),
		Name: 				ability.Name,
		Version: 			ability.Version,
		Specification: 		ability.Specification,
		Rank: 				ability.Rank,
		BattleInteractions: convertObjSlice(cfg, ability.BattleInteractions, convertBattleInteractionSimple),
	}

	return abilitySimple, nil
}


func createItemAbilitySimple(cfg *Config, _ *http.Request, id int32) (SimpleResource, error) {
	i := cfg.e.itemAbilities
	ability, _ := seeding.GetResourceByID(id, i.objLookupID)

	abilitySimple := AbilitySimple{
		ID: 				ability.ID,
		URL: 				createResourceURL(cfg, i.endpoint, id),
		Name: 				ability.Name,
		Version: 			ability.Version,
		Specification: 		ability.Specification,
		Rank: 				ability.Rank,
		BattleInteractions: convertObjSlice(cfg, ability.BattleInteractions, convertBattleInteractionSimple),
	}

	return abilitySimple, nil
}


func createOverdriveAbilitySimple(cfg *Config, r *http.Request, id int32) (SimpleResource, error) {
	i := cfg.e.overdriveAbilities
	ability, _ := seeding.GetResourceByID(id, i.objLookupID)

	attributes, err := cfg.db.GetOverdriveAbilityAttributes(r.Context(), id)
	if err != nil {
		var zeroType SimpleResource
		return zeroType, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get attributes for %s", ability), err)
	}

	abilitySimple := AbilitySimple{
		ID: 				ability.ID,
		URL: 				createResourceURL(cfg, i.endpoint, id),
		Name: 				ability.Name,
		Version: 			ability.Version,
		Specification: 		ability.Specification,
		Rank: 				h.NullInt32ToPtr(attributes.Rank),
		BattleInteractions: convertObjSlice(cfg, ability.BattleInteractions, convertBattleInteractionSimple),
	}

	return abilitySimple, nil
}


func createPlayerAbilitySimple(cfg *Config, _ *http.Request, id int32) (SimpleResource, error) {
	i := cfg.e.playerAbilities
	ability, _ := seeding.GetResourceByID(id, i.objLookupID)

	abilitySimple := AbilitySimple{
		ID: 				ability.ID,
		URL: 				createResourceURL(cfg, i.endpoint, id),
		Name: 				ability.Name,
		Version: 			ability.Version,
		Specification: 		ability.Specification,
		Rank: 				ability.Rank,
		BattleInteractions: convertObjSlice(cfg, ability.BattleInteractions, convertBattleInteractionSimple),
	}

	return abilitySimple, nil
}


func createTriggerCommandSimple(cfg *Config, _ *http.Request, id int32) (SimpleResource, error) {
	i := cfg.e.triggerCommands
	ability, _ := seeding.GetResourceByID(id, i.objLookupID)

	abilitySimple := AbilitySimple{
		ID: 				ability.ID,
		URL: 				createResourceURL(cfg, i.endpoint, id),
		Name: 				ability.Name,
		Version: 			ability.Version,
		Specification: 		ability.Specification,
		Rank: 				ability.Rank,
		BattleInteractions: convertObjSlice(cfg, ability.BattleInteractions, convertBattleInteractionSimple),
	}

	return abilitySimple, nil
}


func createUnspecifiedAbilitySimple(cfg *Config, _ *http.Request, id int32) (SimpleResource, error) {
	i := cfg.e.unspecifiedAbilities
	ability, _ := seeding.GetResourceByID(id, i.objLookupID)

	abilitySimple := AbilitySimple{
		ID: 				ability.ID,
		URL: 				createResourceURL(cfg, i.endpoint, id),
		Name: 				ability.Name,
		Version: 			ability.Version,
		Specification: 		ability.Specification,
		Rank: 				ability.Rank,
		BattleInteractions: convertObjSlice(cfg, ability.BattleInteractions, convertBattleInteractionSimple),
	}

	return abilitySimple, nil
}