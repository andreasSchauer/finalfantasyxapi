// Package sections asidjasdijasdij
package sections

import (
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/sectiontypes"
)

func CreateAbilitySimple(r *http.Request, id int32, subsection handlerinput.Subsection) (apiresource.SimpleResource, error) {
	i := apilookups.E.Abilities
	ability, _ := lookup.GetResourceByID(id, i.ObjLookupID)
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

func CreateEnemyAbilitySimple(_ *http.Request, id int32, _ handlerinput.Subsection) (apiresource.SimpleResource, error) {
	i := apilookups.E.EnemyAbilities
	ability, _ := lookup.GetResourceByID(id, i.ObjLookupID)

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

func CreateItemAbilitySimple(_ *http.Request, id int32, _ handlerinput.Subsection) (apiresource.SimpleResource, error) {
	i := apilookups.E.ItemAbilities
	ability, _ := lookup.GetResourceByID(id, i.ObjLookupID)

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

func CreateOverdriveAbilitySimple(r *http.Request, id int32, subsection handlerinput.Subsection) (apiresource.SimpleResource, error) {
	i := apilookups.E.OverdriveAbilities
	ability, _ := lookup.GetResourceByID(id, i.ObjLookupID)

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

func CreatePlayerAbilitySimple(_ *http.Request, id int32, _ handlerinput.Subsection) (apiresource.SimpleResource, error) {
	i := apilookups.E.PlayerAbilities
	ability, _ := lookup.GetResourceByID(id, i.ObjLookupID)

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

func CreateTriggerCommandSimple(_ *http.Request, id int32, _ handlerinput.Subsection) (apiresource.SimpleResource, error) {
	i := apilookups.E.TriggerCommands
	ability, _ := lookup.GetResourceByID(id, i.ObjLookupID)

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

func CreateMiscAbilitySimple(_ *http.Request, id int32, _ handlerinput.Subsection) (apiresource.SimpleResource, error) {
	i := apilookups.E.MiscAbilities
	ability, _ := lookup.GetResourceByID(id, i.ObjLookupID)

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
