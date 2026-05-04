package seeding

import (
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type TriggerCommand struct {
	ID int32
	Ability
	TopmenuID          *int32
	Description        string              `json:"description"`
	Effect             string              `json:"effect"`
	Topmenu            *string             `json:"topmenu"`
	RelatedStats       []string            `json:"related_stats"`
	Cursor             string              `json:"cursor"`
	BattleInteractions []BattleInteraction `json:"battle_interactions"`
}

func (t TriggerCommand) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", t),
		t.Ability.ID,
		t.Description,
		t.Effect,
		h.DerefOrNil(t.TopmenuID),
		t.Cursor,
	}
}

func (t TriggerCommand) ToKeyFields() []any {
	return []any{
		t.Ability.Name,
		h.DerefOrNil(t.Ability.Version),
	}
}

func (t TriggerCommand) GetID() int32 {
	return t.ID
}

func (t TriggerCommand) GetAbilityRef() AbilityReference {
	return AbilityReference{
		Name:        t.Name,
		Version:     t.Version,
		AbilityType: string(database.AbilityTypeTriggerCommand),
	}
}

func (t TriggerCommand) Error() string {
	return fmt.Sprintf("trigger command '%s'", h.NameToString(t.Name, t.Version, t.Specification))
}

func (t TriggerCommand) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:            t.ID,
		Name:          t.Name,
		Version:       t.Version,
		Specification: t.Specification,
	}
}
