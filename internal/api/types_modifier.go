package api

type Modifier struct {
	ID           		int32				`json:"id"`
	Name         		string   			`json:"name"`
	Effect       		string   			`json:"effect"`
	Category         	string   			`json:"category"`
	DefaultValue 		*float32 			`json:"default_value"`
	AutoAbilities		[]NamedAPIResource	`json:"auto_abilities"`
	PlayerAbilities		[]NamedAPIResource	`json:"player_abilities"`
	OverdriveAbilities	[]NamedAPIResource	`json:"overdrive_abilities"`
	ItemAbilities		[]NamedAPIResource	`json:"item_abilities"`
	TriggerCommands		[]NamedAPIResource	`json:"trigger_commands"`
	EnemyAbilities		[]NamedAPIResource	`json:"enemy_abilities"`
	StatusConditions	[]NamedAPIResource	`json:"status_conditions"`
	Properties			[]NamedAPIResource	`json:"properties"`
}