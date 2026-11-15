package seeding

import (
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func assignFK[T any, R seedtypes.Lookupable](key T, l map[string]R) (int32, error) {
	result, err := lookup.GetResource(key, l)
	if err != nil {
		return 0, fmt.Errorf("couldn't assign foreign key: %v", err)
	}

	id := result.GetID()
	return id, nil
}

func assignFKPtr[T any, R seedtypes.Lookupable](key *T, l map[string]R) (*int32, error) {
	if key == nil {
		return nil, nil
	}

	result, err := lookup.GetResource(*key, l)
	if err != nil {
		return nil, fmt.Errorf("couldn't assign foreign key ptr: %v", err)
	}

	id := result.GetID()
	return &id, nil
}
