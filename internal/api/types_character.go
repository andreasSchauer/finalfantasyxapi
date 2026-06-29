package api

import "github.com/andreasSchauer/finalfantasyxapi/internal/seeding"

type Character struct {
	ID                     int32                              `json:"id"`
	Name                   string                             `json:"name"`
	UntypedUnit            TypedAPIResource                   `json:"untyped_unit"`
	Area                   AreaAPIResource                    `json:"area"`
	IsStoryBased           bool                               `json:"is_story_based"`
	CanFightUnderwater     bool                               `json:"can_fight_underwater"`
	PhysAtkRange           int32                              `json:"physical_attack_range"`
	WeaponType             string                             `json:"weapon_type"`
	ArmorType              string                             `json:"armor_type"`
	CelestialWeapon        *NamedAPIResource                  `json:"celestial_weapon"`
	OverdriveCommand       *NamedAPIResource                  `json:"overdrive_command"`
	CharacterClasses       []NamedAPIResource                 `json:"character_classes"`
	BaseStats              []BaseStat         				  `json:"-"`
	Stats				   Stats			  				  `json:"stats"`
	DefaultPlayerAbilities []NamedAPIResource                 `json:"default_player_abilities"`
	StdSphereGrid		   *SphereGrid						  `json:"std_sphere_grid"`
	ExpSphereGrid		   *SphereGrid						  `json:"exp_sphere_grid"`
	OverdriveModes         []ResourceAmount[NamedAPIResource] `json:"overdrive_modes"`
}


type SphereGrid struct {
	seeding.StatTable
	Lv1Locks	 			int32				`json:"lv_1_locks"`
	Lv2Locks	 			int32				`json:"lv_2_locks"`
	Lv3Locks	 			int32				`json:"lv_3_locks"`
	Lv4Locks	 			int32				`json:"lv_4_locks"`
	EmptyNodes	 			int32				`json:"empty_nodes"`
	PlayerAbilities   		[]NamedAPIResource  `json:"player_abilities"`
}

func convertSphereGrid(_ *Config, sg seeding.SphereGrid) SphereGrid {
	return SphereGrid{
		StatTable: 	sg.StatTable,
		Lv1Locks: 	sg.Lv1Locks,
		Lv2Locks: 	sg.Lv2Locks,
		Lv3Locks: 	sg.Lv3Locks,
		Lv4Locks: 	sg.Lv4Locks,
		EmptyNodes: sg.EmptyNodes,
	}
}