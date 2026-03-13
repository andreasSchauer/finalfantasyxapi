package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/seeding"

type AeonCommand struct {
	ID                int32                 `json:"id"`
	Name              string                `json:"name"`
	Description       string                `json:"description"`
	Effect            string                `json:"effect"`
	Cursor            *string               `json:"cursor"`
	User			  NamedAPIResource		`json:"user"`
	Topmenu           *NamedAPIResource     `json:"topmenu"`
	OpenSubmenu       *NamedAPIResource     `json:"open_submenu,omitempty"`
	PossibleAbilities []PossibleAbilityList `json:"possible_abilities"`
}

type PossibleAbilityList struct {
	User      NamedAPIResource   	`json:"user"`
	Abilities []AbilityAPIResource 	`json:"abilities"`
}

func convertPossibleAbilityList(cfg *Config, pa seeding.PossibleAbilityList) PossibleAbilityList {
	return PossibleAbilityList{
		User:      nameToNamedAPIResource(cfg, cfg.e.characterClasses, pa.User, nil),
		Abilities: refsToAbilityAPIResources(cfg, cfg.e.abilities, pa.Abilities),
	}
}
