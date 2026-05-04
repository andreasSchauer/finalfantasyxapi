package seeding

import (
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type PlayerAbility struct {
	ID int32
	Ability
	TopmenuID           *int32
	SubmenuID           *int32
	OpenSubmenuID       *int32
	StandardGridCharID  *int32
	ExpertGridCharID    *int32
	Description         *string             `json:"description"`
	Effect              string              `json:"effect"`
	RelatedStats        []string            `json:"related_stats"`
	Category            string              `json:"category"`
	Topmenu             *string             `json:"topmenu"`
	Submenu             *string             `json:"submenu"`
	OpenSubmenu         *string             `json:"open_submenu"`
	LearnedBy           []string            `json:"learned_by"`
	StandardGridPos     *string             `json:"standard_grid_pos"`
	ExpertGridPos       *string             `json:"expert_grid_pos"`
	CanUseOutsideBattle bool                `json:"can_use_outside_battle"`
	AeonLearnItem       *ItemAmount         `json:"aeon_learn_item"`
	MPCost              int32               `json:"mp_cost"`
	Cursor              *string             `json:"cursor"`
	BattleInteractions  []BattleInteraction `json:"battle_interactions"`
}

func (p PlayerAbility) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", p),
		p.Ability.ID,
		h.DerefOrNil(p.Description),
		p.Effect,
		p.Category,
		h.DerefOrNil(p.TopmenuID),
		p.CanUseOutsideBattle,
		p.MPCost,
		h.DerefOrNil(p.Cursor),
		h.DerefOrNil(p.SubmenuID),
		h.DerefOrNil(p.OpenSubmenuID),
		h.DerefOrNil(p.StandardGridCharID),
		h.DerefOrNil(p.ExpertGridCharID),
		h.ObjPtrToID(p.AeonLearnItem),
	}
}

func (p PlayerAbility) ToKeyFields() []any {
	return []any{
		p.Ability.Name,
		h.DerefOrNil(p.Ability.Version),
	}
}

func (p PlayerAbility) GetID() int32 {
	return p.ID
}

func (p PlayerAbility) GetAbilityRef() AbilityReference {
	return AbilityReference{
		Name:        p.Name,
		Version:     p.Version,
		AbilityType: string(database.AbilityTypePlayerAbility),
	}
}

func (p PlayerAbility) Error() string {
	return fmt.Sprintf("player ability '%s'", h.NameToString(p.Name, p.Version, p.Specification))
}

func (p PlayerAbility) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:            p.ID,
		Name:          p.Name,
		Version:       p.Version,
		Specification: p.Specification,
	}
}

func (p PlayerAbility) GetItemAmount() ItemAmount {
	itemAmtPtr := p.AeonLearnItem

	if itemAmtPtr == nil {
		return ItemAmount{}
	}

	return *itemAmtPtr
}
