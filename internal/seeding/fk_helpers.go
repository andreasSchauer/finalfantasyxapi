package seeding

import "github.com/andreasSchauer/finalfantasyxapi/internal/database"


type HasID interface {
	GetID() *int32
}


func assignFK[T any, R HasID](val *T, lookup func(T) (R, error)) (*int32, error) {
	if val == nil {
		return nil, nil
	}

	result, err := lookup(*val)
	if err != nil {
		return nil, err
	}

	return result.GetID(), nil
}


func assignFKSeed[T HasID](qtx *database.Queries, val *T, seed func(*database.Queries, T) (T, error)) (*T, error) {
	if val == nil {
		return nil, nil
	}

	object, err := seed(qtx, *val)
	if err != nil {
		return nil, err
	}

	return &object, nil
}