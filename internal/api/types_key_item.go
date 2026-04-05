package api


type KeyItem struct {
	ID                 	int32                	`json:"id"`
	Name              	string               	`json:"name"`
	UntypedItem			TypedAPIResource		`json:"untyped_item"`
	Category           	NamedAPIResource     	`json:"category"`
	Description        	string               	`json:"description"`
	Effect             	string               	`json:"effect"`
	CelestialWeapon		*NamedAPIResource		`json:"celestial_weapon,omitempty"`
	Primer				*NamedAPIResource		`json:"primer,omitempty"`
	Areas				[]AreaAPIResource		`json:"areas"`
	Treasures          	[]UnnamedAPIResource 	`json:"treasures"`
	Quests             	[]QuestAPIResource		`json:"quests"`
}