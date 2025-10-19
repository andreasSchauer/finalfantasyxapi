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


func createJunctionSeed[P HasID, C HasID](qtx *database.Queries, parent P, child C, seed func(*database.Queries, C) (C, error)) (Junction, error) {
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