package seedhandlers

import (
	"fmt"

	sh "github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedhelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

type needsID interface {
	SetID(int32)
	seedtypes.Hashable
	fmt.Stringer
}

func assignID(obj needsID) error {
	id, ok := sh.Hashes[sh.GenerateDataHash(obj)]
	if !ok {
		return fmt.Errorf("no id found for %s", obj)
	}

	obj.SetID(id)
	return nil
}

func assignIDs[T any, P interface {
	*T
	needsID
}](items []T) error {
	for i := range items {
		err := assignID(P(&items[i]))
		if err != nil {
			return err
		}
	}

	return nil
}
