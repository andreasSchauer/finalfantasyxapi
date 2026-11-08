package apitypes

import "github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"

type Topmenu struct {
	ID                int32               `json:"id"`
	Name              string              `json:"name"`
	Submenus          []apiresource.Named `json:"submenus"`
	Abilities         []apiresource.Typed `json:"abilities"`
	OverdriveCommands []apiresource.Named `json:"overdrive_commands"`
	Overdrives        []apiresource.Named `json:"overdrives"`
	AeonCommands      []apiresource.Named `json:"aeon_commands"`
}

type Submenu struct {
	ID          int32               `json:"id"`
	Name        string              `json:"name"`
	Description *string             `json:"description"`
	Effect      string              `json:"effect"`
	Topmenu     *apiresource.Named  `json:"topmenu"`
	Users       []apiresource.Named `json:"users"`
	Abilities   []apiresource.Typed `json:"abilities"`
	OpenedBy    *MenuOpen           `json:"opened_by"`
}

type MenuOpen struct {
	Ability           *apiresource.Typed  `json:"ability,omitempty"`
	AeonCommand       *apiresource.Named  `json:"aeon_command,omitempty"`
	OverdriveCommands []apiresource.Named `json:"overdrive_commands,omitempty"`
}

func (mo MenuOpen) IsZero() bool {
	return mo.Ability == nil && mo.AeonCommand == nil && len(mo.OverdriveCommands) == 0
}
