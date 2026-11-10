package sections

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/sections/sectiontypes"
)

// GetJunctionIDs extracts the child IDs from a pre-sorted []Junction into a []int32 and removes the extracted pairs from the input slice
func GetJunctionIDs(parentID int32, junctions []sectiontypes.Junction) ([]int32, []sectiontypes.Junction) {
	ids := []int32{}

	for i, junction := range junctions {
		if junction.ParentID != parentID {
			return ids, junctions[i:]
		}
		ids = append(ids, junction.ChildID)
	}

	return ids, nil
}
