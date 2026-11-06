package apitypes

import "github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"

type Modifier struct {
	ID                 int32                          `json:"id"`
	Name               string                         `json:"name"`
	Effect             string                         `json:"effect"`
	Category           string                         `json:"category"`
	DefaultValue       *float32                       `json:"default_value"`
	AutoAbilities      []apiresource.NamedAPIResource `json:"auto_abilities"`
	PlayerAbilities    []apiresource.NamedAPIResource `json:"player_abilities"`
	OverdriveAbilities []apiresource.NamedAPIResource `json:"overdrive_abilities"`
	ItemAbilities      []apiresource.NamedAPIResource `json:"item_abilities"`
	TriggerCommands    []apiresource.NamedAPIResource `json:"trigger_commands"`
	EnemyAbilities     []apiresource.NamedAPIResource `json:"enemy_abilities"`
	StatusConditions   []apiresource.NamedAPIResource `json:"status_conditions"`
	Properties         []apiresource.NamedAPIResource `json:"properties"`
}
