package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type AbilityReference struct {
	Name        string `json:"name"`
	Version     *int32 `json:"version"`
	AbilityType string `json:"ability_type"`
}

func (a AbilityReference) ToKeyFields() []any {
	return []any{
		a.Name,
		h.DerefOrNil(a.Version),
		a.AbilityType,
	}
}

func (a AbilityReference) Error() string {
	return fmt.Sprintf("ability reference '%s', type %s", h.NameToString(a.Name, a.Version, nil), a.AbilityType)
}

func (a AbilityReference) Untyped() UntypedAbilityRef {
	return UntypedAbilityRef{
		Name:    a.Name,
		Version: a.Version,
	}
}

type UntypedAbilityRef struct {
	Name    string
	Version *int32
}

func (a UntypedAbilityRef) ToKeyFields() []any {
	return []any{
		a.Name,
		h.DerefOrNil(a.Version),
	}
}

func (a UntypedAbilityRef) Error() string {
	return fmt.Sprintf("untyped ability reference '%s'", h.NameToString(a.Name, a.Version, nil))
}
