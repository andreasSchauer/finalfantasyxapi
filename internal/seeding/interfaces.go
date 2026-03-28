package seeding

import (
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type HasLocArea interface {
	GetLocationArea() LocationArea
}


type LookupableID interface {
	h.HasID
	error
}

type Lookupable interface {
	ToKeyFields() []any
	error
}

func CreateLookupKey(l Lookupable) string {
	fields := l.ToKeyFields()
	return combineFields(fields)
}


type Hashable interface {
	ToHashFields() []any
}

type HasItemAmount interface {
	h.HasID
	GetItemAmount() ItemAmount
}

type HasItemAmounts interface {
	h.HasID
	GetItemAmounts() []ItemAmount
}