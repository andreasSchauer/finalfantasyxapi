package apitypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
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
