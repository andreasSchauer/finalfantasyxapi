package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Character struct {
	ID int32
	PlayerUnit
	LocationArea       LocationArea `json:"location_area"`
	AreaID             *int32
	IsStoryBased       bool       `json:"is_story_based"`
	WeaponType         string     `json:"weapon_type"`
	ArmorType          string     `json:"armor_type"`
	PhysAtkRange       int32      `json:"physical_attack_range"`
	CanFightUnderwater bool       `json:"can_fight_underwater"`
	BaseStats          []BaseStat `json:"base_stats"`
}

func (c Character) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", c),
		c.PlayerUnit.ID,
		h.DerefOrNil(c.AreaID),
		c.IsStoryBased,
		c.WeaponType,
		c.ArmorType,
		c.PhysAtkRange,
		c.CanFightUnderwater,
	}
}

func (c Character) ToKeyFields() []any {
	return []any{
		c.Name,
	}
}

func (c Character) GetID() int32 {
	return c.ID
}

func (c Character) GetResParamsNamed() ResParamsNamed {
	return ResParamsNamed{
		ID:   c.ID,
		Name: c.Name,
	}
}

func (c Character) Error() string {
	return fmt.Sprintf("character %s", c.Name)
}
