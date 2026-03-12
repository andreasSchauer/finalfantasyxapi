package api

type Topmenu struct {
	ID          		int32				`json:"id"`
	Name        		string   			`json:"name"`
	Submenus			[]NamedAPIResource	`json:"submenus"`
	Abilities			[]NamedAPIResource	`json:"abilities"`
	OverdriveCommands	[]NamedAPIResource	`json:"overdrive_commands"`
	Overdrives			[]NamedAPIResource	`json:"overdrives"`
	AeonCommands		[]NamedAPIResource	`json:"aeon_commands"`
}

type Submenu struct {
	ID          int32				`json:"id"`
	Name        string   			`json:"name"`
	Description *string   			`json:"description"`
	Effect      string   			`json:"effect"`
	Topmenu     *NamedAPIResource  	`json:"topmenu"`
	Users       []NamedAPIResource 	`json:"users"`
	Abilities	[]NamedAPIResource	`json:"abilities"`
	OpenedBy	*MenuOpen			`json:"opened_by"`
}

type MenuOpen struct {
	Ability				*NamedAPIResource	`json:"ability"`
	AeonCommand			*NamedAPIResource	`json:"aeon_command"`
	OverdriveCommands	[]NamedAPIResource	`json:"overdrive_commands"`
}

func (mo MenuOpen) IsZero() bool {
	return mo.Ability == nil && mo.AeonCommand == nil && len(mo.OverdriveCommands) == 0
}