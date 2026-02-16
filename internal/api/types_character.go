package api

type Character struct {
	ID						int32				`json:"id"`
	Name					string				`json:"name"`
	Area					*AreaAPIResource	`json:"area"`
	StoryOnly				bool				`json:"story_only"`
	CanFightUnderwater		bool				`json:"can_fight_underwater"`
	PhysAtkRange			int32				`json:"physical_attack_range"`
	WeaponType				string				`json:"weapon_type"`
	ArmorType				string				`json:"armor_type"`
	CelestialWeapon			*NamedAPIResource	`json:"celestial_weapon"`
	OverdriveCommand		*NamedAPIResource	`json:"overdrive_command"`
	CharacterClasses		[]NamedAPIResource	`json:"character_classes"`
	BaseStats				[]BaseStat			`json:"base_stats"`
	DefaultAbilities		[]NamedAPIResource	`json:"default_abilities"`
	StdSphereGridAbilities	[]NamedAPIResource	`json:"std_sphere_grid_abilities"`
	ExpSphereGridAbilities	[]NamedAPIResource	`json:"exp_sphere_grid_abilities"`
	OverdriveModes			[]ModeAmount		`json:"overdrive_modes"`
}