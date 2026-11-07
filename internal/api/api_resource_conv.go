package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

func refsToNamedAPIResources(refs []seedtypes.AbilityReference) []apiresource.NamedAPIResource {
	abilities := []apiresource.NamedAPIResource{}

	for _, ref := range refs {
		abilityRes := refToNamedAPIResource(ref)
		abilities = append(abilities, abilityRes)
	}

	return abilities
}

func refToNamedAPIResource(ref seedtypes.AbilityReference) apiresource.NamedAPIResource {
	var res apiresource.NamedAPIResource

	switch database.AbilityType(ref.AbilityType) {
	case database.AbilityTypeMiscAbility:
		res = NameToNamedAPIResource(E.MiscAbilities, ref.Name, ref.Version)

	case database.AbilityTypePlayerAbility:
		res = NameToNamedAPIResource(E.PlayerAbilities, ref.Name, ref.Version)

	case database.AbilityTypeEnemyAbility:
		res = NameToNamedAPIResource(E.EnemyAbilities, ref.Name, ref.Version)

	case database.AbilityTypeOverdriveAbility:
		res = NameToNamedAPIResource(E.OverdriveAbilities, ref.Name, ref.Version)

	case database.AbilityTypeItemAbility:
		res = NameToNamedAPIResource(E.ItemAbilities, ref.Name, ref.Version)

	case database.AbilityTypeTriggerCommand:
		res = NameToNamedAPIResource(E.TriggerCommands, ref.Name, ref.Version)

	}
	return res
}

func questToQuestAPIResource(quest seedtypes.Quest) apiresource.QuestAPIResource {
	switch quest.Type {
	case database.QuestTypeSidequest:
		return NameToQuestAPIResource(E.Sidequests, quest.Name)

	case database.QuestTypeSubquest:
		return NameToQuestAPIResource(E.Subquests, quest.Name)
	}

	return apiresource.QuestAPIResource{}
}
