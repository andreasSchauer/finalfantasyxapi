package api


type CharacterClass struct {
	ID					int32				`json:"id"`
	Name				string				`json:"name"`
	Category			string				`json:"category"`
	Units				[]NamedAPIResource	`json:"units"`
	DefaultAbilities	[]NamedAPIResource	`json:"default_abilities"`
	LearnableAbilities	[]NamedAPIResource	`json:"learnable_abilities"`
	DefaultOverdrives	[]NamedAPIResource	`json:"default_overdrives"`
	LearnableOverdrives	[]NamedAPIResource	`json:"learnable_overdrives"`
	Submenus			[]NamedAPIResource	`json:"submenus"`
}