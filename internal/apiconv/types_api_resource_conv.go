// Package apiconv asidjasdij
package apiconv

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func RefsToNamedAPIResources(refs []seedtypes.AbilityReference) []apiresource.Named {
	abilities := []apiresource.Named{}

	for _, ref := range refs {
		abilityRes := RefToNamedAPIResource(ref)
		abilities = append(abilities, abilityRes)
	}

	return abilities
}

func RefToNamedAPIResource(ref seedtypes.AbilityReference) apiresource.Named {
	var res apiresource.Named

	switch database.AbilityType(ref.AbilityType) {
	case database.AbilityTypeMiscAbility:
		res = handlerinput.NameToNamedAPIResource(apilookups.E.MiscAbilities, ref.Name, ref.Version)

	case database.AbilityTypePlayerAbility:
		res = handlerinput.NameToNamedAPIResource(apilookups.E.PlayerAbilities, ref.Name, ref.Version)

	case database.AbilityTypeEnemyAbility:
		res = handlerinput.NameToNamedAPIResource(apilookups.E.EnemyAbilities, ref.Name, ref.Version)

	case database.AbilityTypeOverdriveAbility:
		res = handlerinput.NameToNamedAPIResource(apilookups.E.OverdriveAbilities, ref.Name, ref.Version)

	case database.AbilityTypeItemAbility:
		res = handlerinput.NameToNamedAPIResource(apilookups.E.ItemAbilities, ref.Name, ref.Version)

	case database.AbilityTypeTriggerCommand:
		res = handlerinput.NameToNamedAPIResource(apilookups.E.TriggerCommands, ref.Name, ref.Version)

	}
	return res
}

func QuestToQuestAPIResource(quest seedtypes.Quest) apiresource.Quest {
	switch quest.Type {
	case database.QuestTypeSidequest:
		return handlerinput.NameToQuestAPIResource(apilookups.E.Sidequests, quest.Name)

	case database.QuestTypeSubquest:
		return handlerinput.NameToQuestAPIResource(apilookups.E.Subquests, quest.Name)
	}

	return apiresource.Quest{}
}
