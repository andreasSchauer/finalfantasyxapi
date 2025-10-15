package seeding

import "github.com/andreasSchauer/finalfantasyxapi/internal/database"

type HasID interface {
	GetID() int32
}

func assignFK[T any, R HasID](key T, lookup func(T) (R, error)) (int32, error) {
	result, err := lookup(key)
	if err != nil {
		return 0, err
	}

	id := result.GetID()
	return id, nil
}

func assignFKPtr[T any, R HasID](key *T, lookup func(T) (R, error)) (*int32, error) {
	if key == nil {
		return nil, nil
	}

	result, err := lookup(*key)
	if err != nil {
		return nil, err
	}

	id := result.GetID()
	return &id, nil
}

func seedObjAssignFK[T HasID](qtx *database.Queries, obj T, seed func(*database.Queries, T) (T, error)) (T, error) {
	var object T

	object, err := seed(qtx, obj)
	if err != nil {
		return object, err
	}

	return object, nil
}

func seedObjPtrAssignFK[T HasID](qtx *database.Queries, obj *T, seed func(*database.Queries, T) (T, error)) (*T, error) {
	if obj == nil {
		return nil, nil
	}

	object, err := seed(qtx, *obj)
	if err != nil {
		return nil, err
	}

	return &object, nil
}
