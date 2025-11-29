package main

import (
	"net/http"
	"net/url"
	"strconv"

	"github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

// deals with a single segment path that is either a name or an id and returns the id
func parseSingleSegmentResource[T helpers.HasID](segment string, lookup map[string]T) (int32, error) {
	decoded, err := url.PathUnescape(segment)
	if err != nil {
		return 0, newHTTPError(http.StatusBadRequest, "Invalid URL encoding", err)
	}

	parsedID, err := strconv.Atoi(decoded)
	if err == nil {
		return int32(parsedID), nil
	}

	resource, err := seeding.GetResource(segment, lookup)
	if err != nil {
		return 0, newHTTPError(http.StatusNotFound, err.Error(), err)
	}

	return resource.GetID(), nil
}
