package apitypes

import "github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"

type Stat struct {
	ID                 int32               `json:"id"`
	Name               string              `json:"name"`
	Effect             string              `json:"effect"`
	MinVal             int32               `json:"min_val"`
	MaxVal             int32               `json:"max_val"`
	MaxVal2            *int32              `json:"max_val_2,omitempty"`
	ActivationSphere   apiresource.Named   `json:"activation_sphere"`
	Spheres            []apiresource.Named `json:"spheres"`
	AutoAbilities      []apiresource.Named `json:"auto_abilities"`
	PlayerAbilities    []apiresource.Named `json:"player_abilities"`
	OverdriveAbilities []apiresource.Named `json:"overdrive_abilities"`
	ItemAbilities      []apiresource.Named `json:"item_abilities"`
	TriggerCommands    []apiresource.Named `json:"trigger_commands"`
	StatusConditions   []apiresource.Named `json:"status_conditions"`
	Properties         []apiresource.Named `json:"properties"`
}
