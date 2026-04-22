package seeding

import (
	"fmt"
	"slices"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type Junction interface {
	ToHashFieldsJ(string) []any
	Hashable
}

type StdJunction struct {
	ParentID int32
	ChildID  int32
}

func (j StdJunction) ToHashFields() []any {
	return []any{
		j.ParentID,
		j.ChildID,
	}
}

func (j StdJunction) ToHashFieldsJ(name string) []any {
	return slices.Concat([]any{name}, j.ToHashFields())
}


type ThreeWayJunction struct {
	GrandparentID int32
	StdJunction
}

func (j ThreeWayJunction) ToHashFields() []any {
	return []any{
		j.GrandparentID,
		j.ParentID,
		j.ChildID,
	}
}

func (j ThreeWayJunction) ToHashFieldsJ(name string) []any {
	return slices.Concat([]any{name}, j.ToHashFields())
}

type FourWayJunction struct {
	GreatGrandparentID int32
	ThreeWayJunction
}

func (j FourWayJunction) ToHashFields() []any {
	return []any{
		j.GreatGrandparentID,
		j.GrandparentID,
		j.ParentID,
		j.ChildID,
	}
}

func (j FourWayJunction) ToHashFieldsJ(name string) []any {
	return slices.Concat([]any{name}, j.ToHashFields())
}

func createJunction[T any, P, C h.HasID](parent P, childKey T, lookup map[string]C) (StdJunction, error) {
	child, err := GetResource(childKey, lookup)
	if err != nil {
		return StdJunction{}, fmt.Errorf("couldn't create junction: %v", err)
	}

	junction := StdJunction{
		ParentID: parent.GetID(),
		ChildID:  child.GetID(),
	}

	return junction, nil
}

func createJunctionSeed[P, C h.HasID](qtx *database.Queries, parent P, child C, seed func(*database.Queries, C) (C, error)) (StdJunction, error) {
	child, err := seed(qtx, child)
	if err != nil {
		return StdJunction{}, fmt.Errorf("couldn't seed object and create junction: %v", err)
	}

	junction := StdJunction{
		ParentID: parent.GetID(),
		ChildID:  child.GetID(),
	}

	return junction, nil
}

func createThreeWayJunction[T any, GP, P, C h.HasID](grandParent GP, parent P, childKey T, lookup map[string]C) (ThreeWayJunction, error) {
	junction, err := createJunction(parent, childKey, lookup)
	if err != nil {
		return ThreeWayJunction{}, fmt.Errorf("couldn't create three way junction: %v", err)
	}

	threeWay := ThreeWayJunction{
		GrandparentID: grandParent.GetID(),
		StdJunction:   junction,
	}

	return threeWay, nil
}

func createThreeWayJunctionSeed[GP, P, C h.HasID](qtx *database.Queries, grandParent GP, parent P, child C, seed func(*database.Queries, C) (C, error)) (ThreeWayJunction, error) {
	junction, err := createJunctionSeed(qtx, parent, child, seed)
	if err != nil {
		return ThreeWayJunction{}, fmt.Errorf("couldn't seed object and create three way junction: %v", err)
	}

	threeWay := ThreeWayJunction{
		GrandparentID: grandParent.GetID(),
		StdJunction:   junction,
	}

	return threeWay, nil
}

func createFourWayJunctionSeed[GGP, GP, P, C h.HasID](qtx *database.Queries, greatGrandParent GGP, grandParent GP, parent P, child C, seed func(*database.Queries, C) (C, error)) (FourWayJunction, error) {
	threeWay, err := createThreeWayJunctionSeed(qtx, grandParent, parent, child, seed)
	if err != nil {
		return FourWayJunction{}, fmt.Errorf("couldn't seed object and create four way junction: %v", err)
	}

	fourWay := FourWayJunction{
		GreatGrandparentID: greatGrandParent.GetID(),
		ThreeWayJunction:   threeWay,
	}

	return fourWay, nil
}
