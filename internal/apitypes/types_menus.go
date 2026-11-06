package apitypes

import "github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"

type Topmenu struct {
	ID                int32                          `json:"id"`
	Name              string                         `json:"name"`
	Submenus          []apiresource.NamedAPIResource `json:"submenus"`
	Abilities         []apiresource.TypedAPIResource `json:"abilities"`
	OverdriveCommands []apiresource.NamedAPIResource `json:"overdrive_commands"`
	Overdrives        []apiresource.NamedAPIResource `json:"overdrives"`
	AeonCommands      []apiresource.NamedAPIResource `json:"aeon_commands"`
}

type Submenu struct {
	ID          int32                          `json:"id"`
	Name        string                         `json:"name"`
	Description *string                        `json:"description"`
	Effect      string                         `json:"effect"`
	Topmenu     *apiresource.NamedAPIResource  `json:"topmenu"`
	Users       []apiresource.NamedAPIResource `json:"users"`
	Abilities   []apiresource.TypedAPIResource `json:"abilities"`
	OpenedBy    *MenuOpen                      `json:"opened_by"`
}

type MenuOpen struct {
	Ability           *apiresource.TypedAPIResource  `json:"ability,omitempty"`
	AeonCommand       *apiresource.NamedAPIResource  `json:"aeon_command,omitempty"`
	OverdriveCommands []apiresource.NamedAPIResource `json:"overdrive_commands,omitempty"`
}

func (mo MenuOpen) IsZero() bool {
	return mo.Ability == nil && mo.AeonCommand == nil && len(mo.OverdriveCommands) == 0
}
