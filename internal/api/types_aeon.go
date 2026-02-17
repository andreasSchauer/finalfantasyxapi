package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/seeding"


type Aeon struct {
	ID					int32				`json:"id"`
	Name				string				`json:"name"`
	UnlockCondition		string				`json:"unlock_condition"`
	Area				AreaAPIResource		`json:"area"`
	IsOptional			bool				`json:"is_optional"`
	BattlesToRegenerate	int32				`json:"battles_to_regenerate"`
	PhysAtkDmgConstant  *int32          	`json:"phys_atk_damage_constant"`
	PhysAtkRange        *int32          	`json:"phys_atk_range"`
	PhysAtkShatterRate  *int32          	`json:"phys_atk_shatter_rate"`
	PhysAtkAccuracy     *Accuracy       	`json:"phys_atk_accuracy"`
	CelestialWeapon		*NamedAPIResource	`json:"celestial_weapon"`
	CharacterClasses	[]NamedAPIResource	`json:"character_classes"`
	BaseStats			[]BaseStat			`json:"base_stats"`
	AeonCommands		[]NamedAPIResource	`json:"aeon_commands"`
	DefaultAbilities	[]NamedAPIResource	`json:"default_abilities"`
	Overdrives			[]NamedAPIResource	`json:"overdrives"`
	WeaponAbilities		[]AeonEquipment		`json:"weapon_abilities"`
	ArmorAbilities		[]AeonEquipment		`json:"armor_abilities"`
}


type AeonEquipment struct {
	AutoAbility		NamedAPIResource	`json:"auto_ability"`
	CelestialWeapon	bool				`json:"celestial_weapon"`
}

func convertAeonEquipment(cfg *Config, ae seeding.AeonEquipment) AeonEquipment {
	return AeonEquipment{
		AutoAbility: 		nameToNamedAPIResource(cfg, cfg.e.autoAbilities, ae.AutoAbility, nil),
		CelestialWeapon: 	ae.CelestialWeapon,
	}
}

