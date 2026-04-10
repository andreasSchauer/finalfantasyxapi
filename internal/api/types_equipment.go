package api

type EquipmentName struct {
	ID                      int32              		`json:"id"`
	Name					string					`json:"name"`
	Character				NamedAPIResource		`json:"character"`
	EquipmentTable			UnnamedAPIResource		`json:"equipment_table"`
	Type                    string             		`json:"type"`
	Classification          string             		`json:"classification"`
	Priority                *int32             		`json:"priority"`
	CelestialWeapon			*NamedAPIResource		`json:"celestial_weapon,omitempty"`
	RequiredAutoAbilities   []NamedAPIResource 		`json:"required_auto_abilities"`
	SelectableAutoAbilities []AbilityPool      		`json:"selectable_auto_abilities"`
	EmptySlotsAmt           int32              		`json:"empty_slots_amount"`
	Treasures				[]UnnamedAPIResource	`json:"treasures"`
	Shops					[]UnnamedAPIResource	`json:"shops"`
}