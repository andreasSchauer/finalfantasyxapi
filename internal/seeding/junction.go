package seeding

import (
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)


type Junction struct {
	ParentID 	int32
	ChildID  	int32
}


func (j Junction) ToHashFields() []any {
	return []any{
		j.ParentID,
		j.ChildID,
	}
}


type ThreeWayJunction struct {
	GrandparentID	int32
	Junction
}


func (j ThreeWayJunction) ToHashFields() []any {
	return []any{
		j.GrandparentID,
		j.ParentID,
		j.ChildID,
	}
}


func createJunction[T any, P, C HasID](parent P, childKey T, lookup func(T) (C, error)) (Junction, error) {
	child, err := lookup(childKey)
	if err != nil {
		return Junction{}, fmt.Errorf("couldn't create junction: %v", err)
	}

	junction := Junction{
		ParentID: 	parent.GetID(),
		ChildID: 	child.GetID(),
	}

	return junction, nil
}


func createJunctionSeed[P, C HasID](qtx *database.Queries, parent P, child C, seed func(*database.Queries, C) (C, error)) (Junction, error) {
	child, err := seed(qtx, child)
	if err != nil {
		return Junction{}, fmt.Errorf("couldn't seed object and create junction: %v", err)
	}

	junction := Junction{
		ParentID: 	parent.GetID(),
		ChildID: 	child.GetID(),
	}

	return junction, nil
}


// first check, if they are needed
func createThreeWayJunction[T any, GP, P, C HasID](grandParent GP, parent P, childKey T, lookup func(T) (C, error)) (ThreeWayJunction, error) {
	junction, err := createJunction(parent, childKey, lookup)
	if err != nil {
		return ThreeWayJunction{}, fmt.Errorf("couldn't create three way junction: %v", err)
	}

	threeWay := ThreeWayJunction{
		GrandparentID:	grandParent.GetID(),
		Junction: 		junction,	
	}

	return threeWay, nil
}


func createThreeWayJunctionSeed[GP, P, C HasID](qtx *database.Queries, grandParent GP, parent P, child C, seed func(*database.Queries, C) (C, error)) (ThreeWayJunction, error) {
	junction, err := createJunctionSeed(qtx, parent, child, seed)
	if err != nil {
		return ThreeWayJunction{}, fmt.Errorf("couldn't seed object and create three way junction: %v", err)
	}

	threeWay := ThreeWayJunction{
		GrandparentID:	grandParent.GetID(),
		Junction: 		junction,	
	}

	return threeWay, nil
}