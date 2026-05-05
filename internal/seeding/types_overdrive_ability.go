package seeding

import (
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type OverdriveAbility struct {
	ID int32
	Ability
	Overdrive          LookupObject        `json:"overdrive"` // not meant for seeding (still needed?)
	RelatedStats       []string            `json:"related_stats"`
	BattleInteractions []BattleInteraction `json:"battle_interactions"`
}

func (o OverdriveAbility) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", o),
		o.Ability.ID,
	}
}

func (o OverdriveAbility) ToKeyFields() []any {
	return []any{
		o.Ability.Name,
		h.DerefOrNil(o.Ability.Version),
	}
}

func (o OverdriveAbility) GetID() int32 {
	return o.ID
}

func (o OverdriveAbility) GetAbilityRef() AbilityReference {
	return AbilityReference{
		Name:        o.Name,
		Version:     o.Version,
		AbilityType: string(database.AbilityTypeOverdriveAbility),
	}
}

func (o OverdriveAbility) Error() string {
	return fmt.Sprintf("overdrive ability '%s'", h.NameToString(o.Name, o.Version, o.Specification))
}

func (o OverdriveAbility) GetResParamsNamed() ResParamsNamed {
	return ResParamsNamed{
		ID:            o.ID,
		Name:          o.Name,
		Version:       o.Version,
		Specification: o.Specification,
	}
}
