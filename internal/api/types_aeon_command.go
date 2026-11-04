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

func convertPossibleAbilityList(cfg *Config, pa seedtypes.PossibleAbilityList) PossibleAbilityList {
	return PossibleAbilityList{
		User:      nameToNamedAPIResource(cfg, E.CharacterClasses, pa.User, nil),
		Abilities: keysToTypedAPIResources(cfg, E.Abilities, pa.Abilities),
	}
}
