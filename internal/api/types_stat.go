package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"

type Stat struct {
	ID                 int32                          `json:"id"`
	Name               string                         `json:"name"`
	Effect             string                         `json:"effect"`
	MinVal             int32                          `json:"min_val"`
	MaxVal             int32                          `json:"max_val"`
	MaxVal2            *int32                         `json:"max_val_2,omitempty"`
	ActivationSphere   apiresource.NamedAPIResource   `json:"activation_sphere"`
	Spheres            []apiresource.NamedAPIResource `json:"spheres"`
	AutoAbilities      []apiresource.NamedAPIResource `json:"auto_abilities"`
	PlayerAbilities    []apiresource.NamedAPIResource `json:"player_abilities"`
	OverdriveAbilities []apiresource.NamedAPIResource `json:"overdrive_abilities"`
	ItemAbilities      []apiresource.NamedAPIResource `json:"item_abilities"`
	TriggerCommands    []apiresource.NamedAPIResource `json:"trigger_commands"`
	StatusConditions   []apiresource.NamedAPIResource `json:"status_conditions"`
	Properties         []apiresource.NamedAPIResource `json:"properties"`
}
