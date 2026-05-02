package seeding

import (
	"fmt"

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
	SetID(int32)
	Hashable
	error
}

func (l *Lookup) assignID(obj needsID) error {
	id, ok := l.Hashes[generateDataHash(obj)]
	if !ok {
		return fmt.Errorf("no id found for %s", obj)
	}
	
	obj.SetID(id)
	return nil
}

func assignIDs[T any, P interface {*T; needsID}](l *Lookup, items []T) error {
	for i := range items {
		err := l.assignID(P(&items[i]))
		if err != nil {
			return err
		}
	}

	return nil
}


type HasItemAmount interface {
	h.HasID
	GetItemAmount() ItemAmount
}

type HasItemAmounts interface {
	h.HasID
	GetItemAmounts() []ItemAmount
}
