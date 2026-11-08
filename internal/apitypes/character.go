package apitypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

type Character struct {
	ID                     int32                                           `json:"id"`
	Name                   string                                          `json:"name"`
	UntypedUnit            apiresource.Typed                               `json:"untyped_unit"`
	Area                   apiresource.Area                                `json:"area"`
	IsStoryBased           bool                                            `json:"is_story_based"`
	CanFightUnderwater     bool                                            `json:"can_fight_underwater"`
	PhysAtkRange           int32                                           `json:"physical_attack_range"`
	AgilityParameters      AgilityParams                                   `json:"agility_parameters"`
	WeaponType             string                                          `json:"weapon_type"`
	ArmorType              string                                          `json:"armor_type"`
	CelestialWeapon        *apiresource.Named                              `json:"celestial_weapon"`
	CelestialFormula       *string                                         `json:"celestial_formula,omitempty"`
	OverdriveCommand       *apiresource.Named                              `json:"overdrive_command"`
	CharacterClasses       []apiresource.Named                             `json:"character_classes"`
	BaseStats              []BaseStat                                      `json:"base_stats"`
	RegenAmounts           []RegenAmount                                   `json:"regen_amounts"`
	DefaultPlayerAbilities []apiresource.Named                             `json:"default_player_abilities"`
	StdSphereGrid          *SphereGrid                                     `json:"std_sphere_grid"`
	ExpSphereGrid          *SphereGrid                                     `json:"exp_sphere_grid"`
	OverdriveModes         []apiresource.ResourceAmount[apiresource.Named] `json:"overdrive_modes"`
}

type SphereGrid struct {
	seedtypes.StatTable
	Lv1Locks        int32               `json:"lv_1_locks"`
	Lv2Locks        int32               `json:"lv_2_locks"`
	Lv3Locks        int32               `json:"lv_3_locks"`
	Lv4Locks        int32               `json:"lv_4_locks"`
	EmptyNodes      int32               `json:"empty_nodes"`
	PlayerAbilities []apiresource.Named `json:"player_abilities"`
}
