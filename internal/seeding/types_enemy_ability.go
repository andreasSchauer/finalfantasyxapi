package seeding

import (
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type EnemyAbility struct {
	ID int32
	Ability
	Effect             *string             `json:"effect"`
	BattleInteractions []BattleInteraction `json:"battle_interactions"`
}

func (e EnemyAbility) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", e),
		e.Ability.ID,
		h.DerefOrNil(e.Effect),
	}
}

func (e EnemyAbility) ToKeyFields() []any {
	return []any{
		e.Ability.Name,
		h.DerefOrNil(e.Ability.Version),
	}
}

func (e EnemyAbility) GetID() int32 {
	return e.ID
}

func (e EnemyAbility) GetAbilityRef() AbilityReference {
	return AbilityReference{
		Name:        e.Name,
		Version:     e.Version,
		AbilityType: string(database.AbilityTypeEnemyAbility),
	}
}

func (e EnemyAbility) Error() string {
	return fmt.Sprintf("enemy ability '%s'", h.NameToString(e.Name, e.Version, e.Specification))
}

func (e EnemyAbility) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:            e.ID,
		Name:          e.Name,
		Version:       e.Version,
		Specification: e.Specification,
	}
}
