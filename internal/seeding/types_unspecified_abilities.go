package seeding

import (
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type UnspecifiedAbility struct {
	ID int32
	Ability
	TopmenuID          *int32
	SubmenuID          *int32
	OpenSubmenuID      *int32
	Description        string              `json:"description"`
	Effect             string              `json:"effect"`
	RelatedStats       []string            `json:"related_stats"`
	Topmenu            *string             `json:"topmenu"`
	Submenu            *string             `json:"submenu"`
	OpenSubmenu        *string             `json:"open_submenu"`
	LearnedBy          []string            `json:"learned_by"`
	Cursor             *string             `json:"cursor"`
	BattleInteractions []BattleInteraction `json:"battle_interactions"`
}

func (u UnspecifiedAbility) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", u),
		u.Ability.ID,
		u.Description,
		u.Effect,
		h.DerefOrNil(u.TopmenuID),
		h.DerefOrNil(u.Cursor),
		h.DerefOrNil(u.SubmenuID),
		h.DerefOrNil(u.OpenSubmenuID),
	}
}

func (u UnspecifiedAbility) ToKeyFields() []any {
	return []any{
		u.Ability.Name,
		h.DerefOrNil(u.Ability.Version),
	}
}

func (u UnspecifiedAbility) GetID() int32 {
	return u.ID
}

func (u UnspecifiedAbility) GetAbilityRef() AbilityReference {
	return AbilityReference{
		Name:        u.Name,
		Version:     u.Version,
		AbilityType: string(database.AbilityTypeUnspecifiedAbility),
	}
}

func (u UnspecifiedAbility) Error() string {
	return fmt.Sprintf("unspecified ability '%s'", h.NameToString(u.Name, u.Version, u.Specification))
}

func (u UnspecifiedAbility) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:            u.ID,
		Name:          u.Name,
		Version:       u.Version,
		Specification: u.Specification,
	}
}
