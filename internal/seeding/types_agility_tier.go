package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type AgilityTier struct {
	ID               int32
	MinAgility       int32            `json:"min_agility"`
	MaxAgility       int32            `json:"max_agility"`
	TickSpeed        int32            `json:"tick_speed"`
	MonsterMinICV    *int32           `json:"monster_min_icv"`
	MonsterMaxICV    *int32           `json:"monster_max_icv"`
	CharacterMaxICV  *int32           `json:"character_max_icv"`
	CharacterMinICVs []AgilitySubtier `json:"character_min_icvs"`
}

func (a AgilityTier) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", a),
		a.MinAgility,
		a.MaxAgility,
		a.TickSpeed,
		h.DerefOrNil(a.MonsterMinICV),
		h.DerefOrNil(a.MonsterMaxICV),
		h.DerefOrNil(a.CharacterMaxICV),
	}
}

func (a AgilityTier) ToKeyFields() []any {
	return []any{
		a.MinAgility,
		a.MaxAgility,
	}
}

func (a AgilityTier) GetID() int32 {
	return a.ID
}

func (a AgilityTier) Error() string {
	return fmt.Sprintf("agility tier with min agility: %d, max agility: %d", a.MinAgility, a.MaxAgility)
}

func (a AgilityTier) GetResParamsUnnamed() ResParamsUnnamed {
	return ResParamsUnnamed{
		ID: a.ID,
	}
}

type AgilitySubtier struct {
	ID              int32
	AgilityTierID   int32
	MinAgility      int32  `json:"subtier_min_agility"`
	MaxAgility      int32  `json:"subtier_max_agility"`
	CharacterMinICV *int32 `json:"character_min_icv"`
}

func (a AgilitySubtier) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", a),
		a.AgilityTierID,
		a.MinAgility,
		a.MaxAgility,
		h.DerefOrNil(a.CharacterMinICV),
	}
}

func (a AgilitySubtier) GetID() int32 {
	return a.ID
}

func (a *AgilitySubtier) SetID(id int32) {
	a.ID = id
}

func (a AgilitySubtier) Error() string {
	return fmt.Sprintf("agility subtier with min agility: %d, max agility: %d", a.MinAgility, a.MaxAgility)
}
