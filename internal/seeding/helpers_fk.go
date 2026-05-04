package seeding

import (
	"fmt"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func assignFK[T any, R h.HasID](key T, lookup map[string]R) (int32, error) {
	result, err := GetResource(key, lookup)
	if err != nil {
		return 0, fmt.Errorf("couldn't assign foreign key: %v", err)
	}

	id := result.GetID()
	return id, nil
}

func assignFKPtr[T any, R h.HasID](key *T, lookup map[string]R) (*int32, error) {
	if key == nil {
		return nil, nil
	}

	result, err := GetResource(*key, lookup)
	if err != nil {
		return nil, fmt.Errorf("couldn't assign foreign key ptr: %v", err)
	}

	id := result.GetID()
	return &id, nil
}

func (l *Lookup) assignID(obj needsID) error {
	id, ok := l.Hashes[generateDataHash(obj)]
	if !ok {
		return fmt.Errorf("no id found for %s", obj)
	}

	obj.SetID(id)
	return nil
}

func assignIDs[T any, P interface {
	*T
	needsID
}](l *Lookup, items []T) error {
	for i := range items {
		err := l.assignID(P(&items[i]))
		if err != nil {
			return err
		}
	}

	return nil
}
