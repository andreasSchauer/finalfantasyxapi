package apitypes

import "github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"

type Modifier struct {
	ID                 int32               `json:"id"`
	Name               string              `json:"name"`
	Effect             string              `json:"effect"`
	Category           string              `json:"category"`
	DefaultValue       *float32            `json:"default_value"`
	AutoAbilities      []apiresource.Named `json:"auto_abilities"`
	PlayerAbilities    []apiresource.Named `json:"player_abilities"`
	OverdriveAbilities []apiresource.Named `json:"overdrive_abilities"`
	ItemAbilities      []apiresource.Named `json:"item_abilities"`
	TriggerCommands    []apiresource.Named `json:"trigger_commands"`
	EnemyAbilities     []apiresource.Named `json:"enemy_abilities"`
	StatusConditions   []apiresource.Named `json:"status_conditions"`
	Properties         []apiresource.Named `json:"properties"`
}
