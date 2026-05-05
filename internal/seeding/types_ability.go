package seeding

import (
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Ability struct {
	ID            int32
	Name          string `json:"name"`
	Version       *int32 `json:"version"`
	Type          database.AbilityType
	Specification *string `json:"specification"`
	Attributes
}

func (a Ability) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", a),
		a.Name,
		h.DerefOrNil(a.Version),
		h.DerefOrNil(a.Specification),
		a.Type,
		a.Attributes,
	}
}

func (a Ability) ToKeyFields() []any {
	return []any{
		a.Name,
		h.DerefOrNil(a.Version),
		a.Type,
	}
}

func (a Ability) GetID() int32 {
	return a.ID
}

func (a Ability) Error() string {
	return fmt.Sprintf("ability '%s', type %s", h.NameToString(a.Name, a.Version, a.Specification), a.Type)
}

func (a Ability) GetAbilityRef() AbilityReference {
	return AbilityReference{
		Name:        a.Name,
		Version:     a.Version,
		AbilityType: string(a.Type),
	}
}

func (a Ability) GetResParamsTyped() ResParamsTyped {
	return ResParamsTyped{
		ID:            a.ID,
		Name:          a.Name,
		Version:       a.Version,
		Specification: a.Specification,
		Type:          string(a.Type),
	}
}

type Attributes struct {
	ID               int32
	Rank             *int32 `json:"rank"`
	AppearsInHelpBar bool   `json:"appears_in_help_bar"`
	CanCopycat       bool   `json:"can_copycat"`
}

func (a Attributes) ToHashFields() []any {
	return []any{
		fmt.Sprintf("%T", a),
		h.DerefOrNil(a.Rank),
		a.AppearsInHelpBar,
		a.CanCopycat,
	}
}

func (a Attributes) GetID() int32 {
	return a.ID
}

func (a Attributes) Error() string {
	return fmt.Sprintf("ability attributes with rank: %v, help bar: %t, copycat: %t", h.PtrToString(a.Rank), a.AppearsInHelpBar, a.CanCopycat)
}
