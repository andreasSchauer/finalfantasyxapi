package main

import (
	"fmt"
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
