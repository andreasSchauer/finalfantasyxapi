package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
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
	return CombineFields(fields)
}

type Hashable interface {
	ToHashFields() []any
}

type needsID interface {
	SetID(int32)
	Hashable
	fmt.Stringer
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
