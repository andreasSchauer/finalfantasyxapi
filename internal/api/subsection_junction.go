package api

import (
	"context"
	"net/http"
)

type Junction struct {
	ParentID int32
	ChildID  int32
}

// queries db for junctionPairs and converts them into a []Junction
func getDbJunctions[R any](r *http.Request, ids []int32, parentResType, childResType string, dbQuery func(context.Context, []int32) ([]R, error), converter func(R) Junction) ([]Junction, error) {
	dbJunctions, err := dbQuery(r.Context(), ids)
	if err != nil {
		return nil, newHTTPErrorDbPairs(parentResType, childResType, err)
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
