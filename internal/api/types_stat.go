package api

type Stat struct {
	ID       			int32				`json:"id"`
	Name     			string 				`json:"name"`
	Effect   			string 				`json:"effect"`
	MinVal   			int32  				`json:"min_val"`
	MaxVal   			int32  				`json:"max_val"`
	MaxVal2  			*int32 				`json:"max_val_2,omitempty"`
	ActivationSphere  	NamedAPIResource	`json:"activation_sphere"`
	Spheres				[]NamedAPIResource	`json:"spheres"`
	AutoAbilities		[]NamedAPIResource	`json:"auto_abilities"`
	PlayerAbilities		[]NamedAPIResource	`json:"player_abilities"`
	OverdriveAbilities	[]NamedAPIResource	`json:"overdrive_abilities"`
	ItemAbilities		[]NamedAPIResource	`json:"item_abilities"`
	TriggerCommands		[]NamedAPIResource	`json:"trigger_commands"`
	StatusConditions	[]NamedAPIResource	`json:"status_conditions"`
	Properties			[]NamedAPIResource	`json:"properties"`
}