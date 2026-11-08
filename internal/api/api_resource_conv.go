package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func refsToNamedAPIResources(refs []seedtypes.AbilityReference) []apiresource.Named {
	abilities := []apiresource.Named{}

	for _, ref := range refs {
		abilityRes := refToNamedAPIResource(ref)
		abilities = append(abilities, abilityRes)
	}

	return abilities
}

func refToNamedAPIResource(ref seedtypes.AbilityReference) apiresource.Named {
	var res apiresource.Named

	switch database.AbilityType(ref.AbilityType) {
	case database.AbilityTypeMiscAbility:
		res = handlerinput.NameToNamedAPIResource(E.MiscAbilities, ref.Name, ref.Version)

	case database.AbilityTypePlayerAbility:
		res = handlerinput.NameToNamedAPIResource(E.PlayerAbilities, ref.Name, ref.Version)

	case database.AbilityTypeEnemyAbility:
		res = handlerinput.NameToNamedAPIResource(E.EnemyAbilities, ref.Name, ref.Version)

	case database.AbilityTypeOverdriveAbility:
		res = handlerinput.NameToNamedAPIResource(E.OverdriveAbilities, ref.Name, ref.Version)

	case database.AbilityTypeItemAbility:
		res = handlerinput.NameToNamedAPIResource(E.ItemAbilities, ref.Name, ref.Version)

	case database.AbilityTypeTriggerCommand:
		res = handlerinput.NameToNamedAPIResource(E.TriggerCommands, ref.Name, ref.Version)

	}
	return res
}

func questToQuestAPIResource(quest seedtypes.Quest) apiresource.Quest {
	switch quest.Type {
	case database.QuestTypeSidequest:
		return handlerinput.NameToQuestAPIResource(E.Sidequests, quest.Name)

	case database.QuestTypeSubquest:
		return handlerinput.NameToQuestAPIResource(E.Subquests, quest.Name)
	}

	return apiresource.Quest{}
}
