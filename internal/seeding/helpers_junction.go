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

type JunctionParams struct {
	DataHashes          []string
	GreatGrandParentIDs []int32 // probably not needed
	GrandParentIDs      []int32
	ParentIDs           []int32
	ChildIDs            []int32
}

func processJunctions[P, C Hashable](l *Lookup, desc string, parents []P, getChildren func(P) ([]C, error)) (JunctionParams, error) {
	params := JunctionParams{
		DataHashes: make([]string, 0),
		ParentIDs:  make([]int32, 0),
		ChildIDs:   make([]int32, 0),
	}

	for _, p := range parents {
		pID, err := l.getHashID(p)
		if err != nil {
			return JunctionParams{}, err
		}

		children, err := getChildren(p)
		if err != nil {
			return JunctionParams{}, err
		}

		for _, c := range children {
			cID, err := l.getHashID(c)
			if err != nil {
				return JunctionParams{}, err
			}

			j := StdJunction{
				ParentID: pID,
				ChildID:  cID,
			}
			dataHash := generateJunctionHash(j, desc)

			params.DataHashes = append(params.DataHashes, dataHash)
			params.ParentIDs = append(params.ParentIDs, pID)
			params.ChildIDs = append(params.ChildIDs, cID)
		}
	}

	return params, nil
}

func processThreewayJunctions[GP, P, C Hashable](l *Lookup, desc string, grandParents []GP, getParents func(GP) ([]P, error), getChildren func(P) ([]C, error)) (JunctionParams, error) {
	params := JunctionParams{
		DataHashes: 	make([]string, 0),
		GrandParentIDs: make([]int32, 0),
		ParentIDs:  	make([]int32, 0),
		ChildIDs:   	make([]int32, 0),
	}
	for _, gp := range grandParents {
		gpID, err := l.getHashID(gp)
		if err != nil {
			return JunctionParams{}, err
		}

		parents, err := getParents(gp)
		if err != nil {
			return JunctionParams{}, err
		}

		for _, p := range parents {
			pID, err := l.getHashID(p)
			if err != nil {
				return JunctionParams{}, err
			}
			
			children, err := getChildren(p)
			if err != nil {
				return JunctionParams{}, err
			}
			
			for _, c := range children {
				cID, err := l.getHashID(c)
				if err != nil {
					return JunctionParams{}, err
				}
				
				j := ThreeWayJunction{}
				j.GrandparentID = gpID
				j.ParentID = pID
				j.ChildID = cID
				dataHash := generateJunctionHash(j, desc)
				
				params.DataHashes = append(params.DataHashes, dataHash)
				params.GrandParentIDs = append(params.GrandParentIDs, gpID)
				params.ParentIDs = append(params.ParentIDs, pID)
				params.ChildIDs = append(params.ChildIDs, cID)
			}
		}
	}
		
	return params, nil
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
