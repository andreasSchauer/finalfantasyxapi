package seedtypes


import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type LookupObject struct {
	Name    string `json:"name"`
	Version *int32 `json:"version"`
}

func (l LookupObject) ToKeyFields() []any {
	return []any{
		l.Name,
		h.DerefOrNil(l.Version),
	}
}

func (l LookupObject) String() string {
	return fmt.Sprintf("lookup object '%s'", h.NameToString(l.Name, l.Version, nil))
}
