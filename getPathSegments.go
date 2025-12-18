package main

import (
	"fmt"
	"strings"
)

func getPathSegments(path, endpoint string) []string {
	prefix := fmt.Sprintf("/api/%s/", endpoint)
	pathTrimmed := strings.TrimPrefix(path, prefix)
	pathTrimmed = strings.TrimSuffix(pathTrimmed, "/")
	segments := []string{}

	if pathTrimmed != "" {
		segments = strings.Split(pathTrimmed, "/")
	}

	return segments
}
