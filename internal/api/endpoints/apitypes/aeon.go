package apitypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
)

type Aeon struct {
	ID                     int32               `json:"id"`
	Name                   string              `json:"name"`
	UntypedUnit            apiresource.Typed   `json:"untyped_unit"`
	UnlockCondition        string              `json:"unlock_condition"`
	Area                   apiresource.Area    `json:"area"`
	IsOptional             bool                `json:"is_optional"`
	BattlesToRegenerate    int32               `json:"battles_to_regenerate"`
	PhysAtkDmgConstant     *int32              `json:"phys_atk_damage_constant"`
	PhysAtkRange           *int32              `json:"phys_atk_range"`
	PhysAtkShatterRate     *int32              `json:"phys_atk_shatter_rate"`
	PhysAtkAccuracy        *Accuracy           `json:"phys_atk_accuracy"`
	AgilityParameters      AgilityParams       `json:"agility_parameters"`
	CelestialWeapon        *apiresource.Named  `json:"celestial_weapon"`
	CharacterClasses       []apiresource.Named `json:"character_classes"`
	BaseStats              []BaseStat          `json:"base_stats"`
	RegenAmounts           []RegenAmount       `json:"regen_amounts"`
	AeonCommands           []apiresource.Named `json:"aeon_commands"`
	DefaultPlayerAbilities []apiresource.Named `json:"default_player_abilities"`
	Overdrives             []apiresource.Named `json:"overdrives"`
	WeaponAbilities        []AeonEquipment     `json:"weapon_abilities"`
	ArmorAbilities         []AeonEquipment     `json:"armor_abilities"`
}

type AeonEquipment struct {
	AutoAbility     apiresource.Named `json:"auto_ability"`
	CelestialWeapon bool              `json:"celestial_weapon"`
}
