package seeding

import (
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type ItemAbility struct {
	ID int32
	Ability
	ItemID             int32
	Cursor             string              `json:"cursor"`
	BattleInteractions []BattleInteraction `json:"battle_interactions"`
}

func (a ItemAbility) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", a),
		a.ItemID,
		a.Ability.ID,
		a.Cursor,
	}
}

func (a ItemAbility) ToKeyFields() []any {
	return []any{
		a.Name,
	}
}

func (a ItemAbility) GetID() int32 {
	return a.ID
}

func (a *ItemAbility) SetID(id int32) {
	a.ID = id
}

func (a ItemAbility) GetAbilityRef() AbilityReference {
	return AbilityReference{
		Name:        a.Name,
		Version:     a.Version,
		AbilityType: string(database.AbilityTypeItemAbility),
	}
}

func (a ItemAbility) Error() string {
	return fmt.Sprintf("item ability %s", a.Name)
}

func (a ItemAbility) GetResParamsNamed() h.ResParamsNamed {
	return h.ResParamsNamed{
		ID:   a.ID,
		Name: a.Name,
	}
}
