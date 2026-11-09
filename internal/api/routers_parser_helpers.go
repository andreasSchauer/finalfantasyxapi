package api

import (
	"strconv"
)

func isValidInt(idStr string) bool {
	_, err := strconv.Atoi(idStr)
	return err == nil
}

func getSegmentCases(segment1, segment2 string) (bool, bool, bool) {
	firstIsInt := isValidInt(segment1)
	secondIsInt := isValidInt(segment2)

	isSubsection := firstIsInt && !secondIsInt   // endpoint/1/section
	isNameVersion := !firstIsInt && secondIsInt  // endpoint/name/1
	subsectionIsInt := firstIsInt && secondIsInt // endpoint/1/2

	return isSubsection, isNameVersion, subsectionIsInt
}
