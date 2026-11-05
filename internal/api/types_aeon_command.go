package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"

type AeonCommand struct {
	ID                int32                 `json:"id"`
	Name              string                `json:"name"`
	Description       string                `json:"description"`
	Effect            string                `json:"effect"`
	Cursor            *string               `json:"cursor"`
	User              NamedAPIResource      `json:"user"`
	Topmenu           *NamedAPIResource     `json:"topmenu"`
	OpenSubmenu       *NamedAPIResource     `json:"open_submenu,omitempty"`
	PossibleAbilities []PossibleAbilityList `json:"possible_abilities"`
}

type PossibleAbilityList struct {
	User      NamedAPIResource   `json:"user"`
	Abilities []TypedAPIResource `json:"abilities"`
}

func convertPossibleAbilityList(pa seedtypes.PossibleAbilityList) PossibleAbilityList {
	return PossibleAbilityList{
		User:      nameToNamedAPIResource(E.CharacterClasses, pa.User, nil),
		Abilities: keysToTypedAPIResources(E.Abilities, pa.Abilities),
	}
}
