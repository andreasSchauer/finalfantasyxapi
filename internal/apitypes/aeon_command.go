package apitypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
)

type AeonCommand struct {
	ID                int32                 `json:"id"`
	Name              string                `json:"name"`
	Description       string                `json:"description"`
	Effect            string                `json:"effect"`
	Cursor            *string               `json:"cursor"`
	User              apiresource.Named     `json:"user"`
	Topmenu           *apiresource.Named    `json:"topmenu"`
	OpenSubmenu       *apiresource.Named    `json:"open_submenu,omitempty"`
	PossibleAbilities []PossibleAbilityList `json:"possible_abilities"`
}

type PossibleAbilityList struct {
	User      apiresource.Named   `json:"user"`
	Abilities []apiresource.Typed `json:"abilities"`
}
