package seedhelpers

import (
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

type needsID interface {
	SetID(int32)
	seedtypes.Hashable
	fmt.Stringer
}

func AssignID(obj needsID) error {
	id, ok := Hashes[GenerateDataHash(obj)]
	if !ok {
		return fmt.Errorf("no id found for %s", obj)
	}

	obj.SetID(id)
	return nil
}

func AssignIDs[T any, P interface {
	*T
	needsID
}](items []T) error {
	for i := range items {
		err := AssignID(P(&items[i]))
		if err != nil {
			return err
		}
	}

	return nil
}
