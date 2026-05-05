package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Overdrive struct {
	ID          int32
	ODCommandID *int32
	CharClassID *int32
	TopmenuID   *int32
	Ability
	Description        string             `json:"description"`
	Effect             string             `json:"effect"`
	Topmenu            *string            `json:"topmenu"`
	OverdriveCommand   *string            `json:"overdrive_command"`
	User               string             `json:"user"`
	UnlockCondition    *string            `json:"unlock_condition"`
	CountdownInSec     *int32             `json:"countdown_in_sec"`
	Cursor             *string            `json:"cursor"`
	OverdriveAbilities []AbilityReference `json:"overdrive_abilities"`
}

func (o Overdrive) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", o),
		h.DerefOrNil(o.ODCommandID),
		h.DerefOrNil(o.CharClassID),
		o.Name,
		h.DerefOrNil(o.Version),
		o.Description,
		o.Effect,
		h.DerefOrNil(o.TopmenuID),
		o.Attributes,
		h.DerefOrNil(o.UnlockCondition),
		h.DerefOrNil(o.CountdownInSec),
		h.DerefOrNil(o.Cursor),
	}
}

func (o Overdrive) ToKeyFields() []any {
	return []any{
		o.Name,
		h.DerefOrNil(o.Version),
	}
}

func (o Overdrive) GetID() int32 {
	return o.ID
}

func (o Overdrive) Error() string {
	return fmt.Sprintf("overdrive '%s'", h.NameToString(o.Name, o.Version, o.Specification))
}

func (o Overdrive) GetResParamsNamed() ResParamsNamed {
	return ResParamsNamed{
		ID:      o.ID,
		Name:    o.Name,
		Version: o.Version,
	}
}
