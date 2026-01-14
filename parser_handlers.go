package main

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

// deals with a single segment path that is either a name or an id and returns a parseResponse with the id, if a single match is found, or a name, if multiple matches with that name were found.
func parseSingleSegmentResource[T h.HasID](resourceType, segment string, lookup map[string]T) (parseResponse, error) {
	decoded, err := url.PathUnescape(segment)
	if err != nil {
		return parseResponse{}, newHTTPError(http.StatusBadRequest, "invalid url encoding.", err)
	}

	response, err := checkID(decoded, resourceType, len(lookup))
	if err == nil {
		return response, nil
	}
	if !errors.Is(err, errNotAnID) {
		return parseResponse{}, err
	}

	response, err = checkUniqueName(decoded, lookup)
	if err == nil {
		return response, nil
	}

	response, err = checkNameMultiple(decoded, lookup)
	if err == nil {
		return response, nil
	}

	return parseResponse{}, newHTTPError(http.StatusNotFound, fmt.Sprintf("%s not found: '%s'.", resourceType, segment), err)
}

// deals with a name/version path and returns a parseResponse with the id, if a match is found.
func parseNameVersionResource[T h.HasID](resourceType, name, versionStr string, lookup map[string]T) (parseResponse, error) {
	nameDecoded, err := url.PathUnescape(name)
	if err != nil {
		return parseResponse{}, newHTTPError(http.StatusBadRequest, "invalid url encoding.", err)
	}

	versionPtr, err := parseVersionStr(versionStr)
	if err != nil {
		return parseResponse{}, err
	}

	response, err := checkNameVersion(nameDecoded, versionPtr, lookup)
	if err == nil {
		return response, nil
	}

	return parseResponse{}, newHTTPError(http.StatusNotFound, fmt.Sprintf("%s not found: '%s', version '%s'", resourceType, name, versionStr), err)
}
