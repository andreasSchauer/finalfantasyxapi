package seeding

import(
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Ability struct {
	Name				string
	Version				*int32
	Specification		*string
	Type				database.AbilityType
}


func(a Ability) ToHashFields() []any {
	return []any{
		a.Name,
		derefOrNil(a.Version),
		derefOrNil(a.Specification),
		a.Type,
	}
}