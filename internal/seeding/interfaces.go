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

func Key(l Lookupable) string {
	fields := l.ToKeyFields()
	return combineFields(fields)
}

type Hashable interface {
	ToHashFields() []any
}

type needsID interface {
	Hashable
	SetID(int32)
}

func (l *Lookup) assignID(obj needsID) {
	obj.SetID(l.Hashes[generateDataHash(obj)])
}

type HasItemAmount interface {
	h.HasID
	GetItemAmount() ItemAmount
}

type HasItemAmounts interface {
	h.HasID
	GetItemAmounts() []ItemAmount
}
