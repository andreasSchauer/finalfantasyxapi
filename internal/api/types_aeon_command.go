package api

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

type AeonCommand struct {
	ID                int32                         `json:"id"`
	Name              string                        `json:"name"`
	Description       string                        `json:"description"`
	Effect            string                        `json:"effect"`
	Cursor            *string                       `json:"cursor"`
	User              apiresource.NamedAPIResource  `json:"user"`
	Topmenu           *apiresource.NamedAPIResource `json:"topmenu"`
	OpenSubmenu       *apiresource.NamedAPIResource `json:"open_submenu,omitempty"`
	PossibleAbilities []PossibleAbilityList         `json:"possible_abilities"`
}

type PossibleAbilityList struct {
	User      apiresource.NamedAPIResource   `json:"user"`
	Abilities []apiresource.TypedAPIResource `json:"abilities"`
}

func convertPossibleAbilityList(pa seedtypes.PossibleAbilityList) PossibleAbilityList {
	return PossibleAbilityList{
		User:      nameToNamedAPIResource(E.CharacterClasses, pa.User, nil),
		Abilities: keysToTypedAPIResources(E.Abilities, pa.Abilities),
	}
}
