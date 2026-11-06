package api

import (
	"context"

	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

type Junction struct {
	ParentID int32
	ChildID  int32
}

// queries db for junctionPairs and converts them into a []Junction
func (i handlerInput[T, R, A]) GetDBJunctions[row any](ctx context.Context, ids []int32, childResType ts.ResTypeSingle, dbQuery func(context.Context, []int32) ([]row, error), converter func(row) Junction) ([]Junction, error) {
	dbJunctions, err := dbQuery(ctx, ids)
	if err != nil {
		return nil, i.ErrorDBPairs(childResType, err)
	}

	junctions := []Junction{}

	for _, dbJunction := range dbJunctions {
		junctions = append(junctions, converter(dbJunction))
	}

	return junctions, nil
}

// extracts the child IDs from a pre-sorted []Junction into a []int32 and removes the extracted pairs from the input slice
func getJunctionIDs(parentID int32, junctions []Junction) ([]int32, []Junction) {
	ids := []int32{}

	for i, junction := range junctions {
		if junction.ParentID != parentID {
			return ids, junctions[i:]
		}
		ids = append(ids, junction.ChildID)
	}

	return ids, nil
}
