package seeding

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Lookupable interface {
	LookupableID
	LookupableKey
}

type LookupableID interface {
	h.HasID
	error
}

type LookupableKey interface {
	ToKeyFields() []any
	error
}

func Key(l LookupableKey) string {
	fields := l.ToKeyFields()
	return combineFields(fields)
}

type Hashable interface {
	ToHashFields() []any
}

type needsID interface {
	SetID(int32)
	Hashable
	error
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
