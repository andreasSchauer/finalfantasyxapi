package api

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/sectiontypes"
)

func createAbilitySimple(r *http.Request, id int32, subsection Subsection) (apiresource.SimpleResource, error) {
	i := E.Abilities
	ability, _ := lookup.GetResourceByID(id, i.objLookupID)
	typeStr := string(ability.Type)

	abilitySimple := sectiontypes.AbilitySimple{
		ID:                 ability.ID,
		URL:                ah.CreateResourceURL(i.Endpoint, id),
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		Rank:               ability.Rank,
		Type:               &typeStr,
		BattleInteractions: ah.ConvertObjSlice(ability.BattleInteractions, convertBattleInteractionSimple),
	}

	return abilitySimple, nil
}

func createEnemyAbilitySimple(_ *http.Request, id int32, _ Subsection) (apiresource.SimpleResource, error) {
	i := E.EnemyAbilities
	ability, _ := lookup.GetResourceByID(id, i.objLookupID)

	abilitySimple := sectiontypes.AbilitySimple{
		ID:                 ability.ID,
		URL:                ah.CreateResourceURL(i.Endpoint, id),
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		Rank:               ability.Rank,
		BattleInteractions: ah.ConvertObjSlice(ability.BattleInteractions, convertBattleInteractionSimple),
	}

	return abilitySimple, nil
}

func createItemAbilitySimple(_ *http.Request, id int32, _ Subsection) (apiresource.SimpleResource, error) {
	i := E.ItemAbilities
	ability, _ := lookup.GetResourceByID(id, i.objLookupID)

	abilitySimple := sectiontypes.AbilitySimple{
		ID:                 ability.ID,
		URL:                ah.CreateResourceURL(i.Endpoint, id),
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		Rank:               ability.Rank,
		BattleInteractions: ah.ConvertObjSlice(ability.BattleInteractions, convertBattleInteractionSimple),
	}

	return abilitySimple, nil
}

func createOverdriveAbilitySimple(r *http.Request, id int32, subsection Subsection) (apiresource.SimpleResource, error) {
	i := E.OverdriveAbilities
	ability, _ := lookup.GetResourceByID(id, i.objLookupID)

	abilitySimple := sectiontypes.AbilitySimple{
		ID:                 ability.ID,
		URL:                ah.CreateResourceURL(i.Endpoint, id),
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		Rank:               ability.Rank,
		BattleInteractions: ah.ConvertObjSlice(ability.BattleInteractions, convertBattleInteractionSimple),
	}

	return abilitySimple, nil
}

func createPlayerAbilitySimple(_ *http.Request, id int32, _ Subsection) (apiresource.SimpleResource, error) {
	i := E.PlayerAbilities
	ability, _ := lookup.GetResourceByID(id, i.objLookupID)

	playerAbilitySimple := sectiontypes.PlayerAbilitySimple{
		ID:                 ability.ID,
		URL:                ah.CreateResourceURL(i.Endpoint, id),
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		Rank:               ability.Rank,
		MpCost:             ability.MPCost,
		BattleInteractions: ah.ConvertObjSlice(ability.BattleInteractions, convertBattleInteractionSimple),
	}

	return playerAbilitySimple, nil
}

func createTriggerCommandSimple(_ *http.Request, id int32, _ Subsection) (apiresource.SimpleResource, error) {
	i := E.TriggerCommands
	ability, _ := lookup.GetResourceByID(id, i.objLookupID)

	abilitySimple := sectiontypes.AbilitySimple{
		ID:                 ability.ID,
		URL:                ah.CreateResourceURL(i.Endpoint, id),
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		Rank:               ability.Rank,
		BattleInteractions: ah.ConvertObjSlice(ability.BattleInteractions, convertBattleInteractionSimple),
	}

	return abilitySimple, nil
}

func createMiscAbilitySimple(_ *http.Request, id int32, _ Subsection) (apiresource.SimpleResource, error) {
	i := E.MiscAbilities
	ability, _ := lookup.GetResourceByID(id, i.objLookupID)

	abilitySimple := sectiontypes.AbilitySimple{
		ID:                 ability.ID,
		URL:                ah.CreateResourceURL(i.Endpoint, id),
		Name:               ability.Name,
		Version:            ability.Version,
		Specification:      ability.Specification,
		Rank:               ability.Rank,
		BattleInteractions: ah.ConvertObjSlice(ability.BattleInteractions, convertBattleInteractionSimple),
	}

	return abilitySimple, nil
}
