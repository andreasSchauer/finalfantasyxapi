package seedtypes

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

type Lookupable interface {
	LookupableID
	LookupableKey
}

type LookupableID interface {
	h.HasID
	fmt.Stringer
}

type LookupableKey interface {
	ToKeyFields() []any
	fmt.Stringer
}

func Key(l LookupableKey) string {
	fields := l.ToKeyFields()
	return h.CombineFields(fields)
}

type Hashable interface {
	ToHashFields() []any
}

type HasItemAmount interface {
	Lookupable
	GetItemAmount() ItemAmount
}

type HasItemAmounts interface {
	Lookupable
	GetItemAmounts() []ItemAmount
}

type HasLocArea interface {
	GetLocationArea() LocationArea
}
