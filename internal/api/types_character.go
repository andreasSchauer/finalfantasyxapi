package api

type Character struct {
	ID                     int32              `json:"id"`
	Name                   string             `json:"name"`
	Area                   AreaAPIResource    `json:"area"`
	IsStoryBased           bool               `json:"is_story_based"`
	CanFightUnderwater     bool               `json:"can_fight_underwater"`
	PhysAtkRange           int32              `json:"physical_attack_range"`
	WeaponType             string             `json:"weapon_type"`
	ArmorType              string             `json:"armor_type"`
	CelestialWeapon        *NamedAPIResource  `json:"celestial_weapon"`
	OverdriveCommand       *NamedAPIResource  `json:"overdrive_command"`
	CharacterClasses       []NamedAPIResource `json:"character_classes"`
	BaseStats              []BaseStat         `json:"base_stats"`
	DefaultPlayerAbilities []NamedAPIResource `json:"default_player_abilities"`
	StdSgPlayerAbilities   []NamedAPIResource `json:"std_sg_player_abilities"`
	ExpSgPlayerAbilities   []NamedAPIResource `json:"exp_sg_player_abilities"`
	OverdriveModes         []ModeAmount       `json:"overdrive_modes"`
}
