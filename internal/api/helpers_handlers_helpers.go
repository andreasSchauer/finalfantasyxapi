package api

import (
	"fmt"
	"strconv"
	"strings"
)

func getPathSegments(path, endpoint string) []string {
	prefix := fmt.Sprintf("/api/%s", endpoint)
	pathLower := strings.ToLower(path)
	pathTrimmed := strings.TrimPrefix(pathLower, prefix)
	pathTrimmed = strings.TrimPrefix(pathTrimmed, "/")
	pathTrimmed = strings.TrimSuffix(pathTrimmed, "/")
	segments := []string{}

	if pathTrimmed != "" {
		segments = strings.Split(pathTrimmed, "/")
	}

	return segments
}

func isValidInt(idStr string) bool {
	_, err := strconv.Atoi(idStr)
	return err == nil
}

func getSegmentCases(segments []string) (bool, bool, bool) {
	firstIsInt := isValidInt(segments[0])
	secondIsInt := isValidInt(segments[1])

	isSubsection := firstIsInt && !secondIsInt
	isNameVersion := !firstIsInt && secondIsInt
	subsectionIsInt := firstIsInt && secondIsInt

	return isSubsection, isNameVersion, subsectionIsInt
}
