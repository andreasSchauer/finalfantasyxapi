package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


// maaaaaybe could trim down the logic into helper functions (also for the other funcs in parser.go)
// I can also simply check in the OG function, if queryParam is "" and if not, return the badRequest error
func parseSingleSegmentResourceQuery[T h.HasID](resourceType, segment, queryParam string, lookup map[string]T) (parseResponse, error) {
	decoded, err := url.PathUnescape(segment)
	if err != nil {
		return parseResponse{}, newHTTPError(http.StatusBadRequest, "Invalid URL encoding", err)
	}

	// check, if the segment is an id
	parsedID, err := strconv.Atoi(decoded)
	if err == nil {
		if parsedID > len(lookup) || parsedID <= 0 {
			return parseResponse{}, newHTTPError(http.StatusBadRequest, fmt.Sprintf("provided %s ID %d is out of range in %s. Max ID: %d", resourceType, parsedID, queryParam, len(lookup)), err)
		}
		return newParseResponse(int32(parsedID), ""), nil
	}

	lookupObj := seeding.LookupObject{
		Name: decoded,
	}

	// check for unique names with dashes (obj input)
	resource, err := seeding.GetResource(lookupObj, lookup)
	if err == nil {
		return newParseResponse(resource.GetID(), ""), nil
	}

	// check for unique names with dashes (string input)
	resource, err = seeding.GetResource(lookupObj.Name, lookup)
	if err == nil {
		return newParseResponse(resource.GetID(), ""), nil
	}

	var testVersion int32 = 1
	lookupObjVer := seeding.LookupObject{
		Name:    decoded,
		Version: &testVersion,
	}

	// check for multi-versioned names with dashes
	_, err = seeding.GetResource(lookupObjVer, lookup)
	if err == nil {
		return newParseResponse(0, lookupObj.Name), nil
	}

	nameWithSpaces := strings.ReplaceAll(decoded, "-", " ")
	nameWithSpaces = strings.ReplaceAll(nameWithSpaces, " >", "->")
	lookupObj.Name = nameWithSpaces
	lookupObjVer.Name = nameWithSpaces

	// check for unique names with spaces (obj input)
	resource, err = seeding.GetResource(lookupObj, lookup)
	if err == nil {
		return newParseResponse(resource.GetID(), ""), nil
	}

	// check for unique names with spaces (str input)
	resource, err = seeding.GetResource(lookupObj.Name, lookup)
	if err == nil {
		return newParseResponse(resource.GetID(), ""), nil
	}

	// check for multi-versioned names with spaces
	_, err = seeding.GetResource(lookupObjVer, lookup)
	if err == nil {
		return newParseResponse(0, lookupObjVer.Name), nil
	}

	return parseResponse{}, newHTTPError(http.StatusBadRequest, fmt.Sprintf("unknown %s '%s' in %s.", resourceType, segment, queryParam), err)
}


func parseNameVersionResourceQuery[T h.HasID](resourceType, name, versionStr, queryParam string, lookup map[string]T) (parseResponse, error) {
	var versionPtr *int32

	// parse the version (int or null)
	switch versionStr {
	case "":
		versionPtr = nil
	default:
		version, err := strconv.Atoi(versionStr)
		if err != nil {
			return parseResponse{}, newHTTPError(http.StatusBadRequest, "Invalid version number", err)
		}
		versionInt32 := int32(version)
		versionPtr = &versionInt32
	}

	nameDecoded, err := url.PathUnescape(name)
	if err != nil {
		return parseResponse{}, newHTTPError(http.StatusBadRequest, "Invalid URL encoding", err)
	}

	key := seeding.LookupObject{
		Name:    nameDecoded,
		Version: versionPtr,
	}

	// check for names with dashes
	resource, err := seeding.GetResource(key, lookup)
	if err == nil {
		return newParseResponse(resource.GetID(), ""), nil
	}

	nameWithSpaces := strings.ReplaceAll(name, "-", " ")
	key.Name = nameWithSpaces

	// check for names with spaces
	resource, err = seeding.GetResource(key, lookup)
	if err == nil {
		return newParseResponse(resource.GetID(), ""), nil
	}

	return parseResponse{}, newHTTPError(http.StatusBadRequest, fmt.Sprintf("unknown %s '%s', version %s in %s.", resourceType, name, versionStr, queryParam), err)
}


func parseBooleanQuery(r *http.Request, queryParam string) (bool, bool, error) {
	query := r.URL.Query().Get(queryParam)
	isEmpty := false

	if query == "" {
		isEmpty = true
		return false, isEmpty, nil
	}

	b, err := strconv.ParseBool(query)
	if err != nil {
		return false, false, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value. usage: %s={boolean}", queryParam), err)
	}

	return b, isEmpty, nil
}

func parseTypeQuery(r *http.Request, queryParam string, lookup map[string]TypedAPIResource) (TypedAPIResource, bool, error) {
	query := r.URL.Query().Get(queryParam)
	isEmpty := false

	if query == "" {
		isEmpty = true
		return TypedAPIResource{}, isEmpty, nil
	}

	enum, err := GetEnumType(query, lookup)
	if err != nil {
		return TypedAPIResource{}, false, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid value: '%s', use /api/%s to see valid values", query, queryParam), err)
	}

	return enum, isEmpty, nil
}

func parseUniqueNameQuery[T h.HasID](r *http.Request, queryParam string, lookup map[string]T) (parseResponse, bool, error) {
	query := r.URL.Query().Get(queryParam)
	isEmpty := false

	if query == "" {
		isEmpty = true
		return parseResponse{}, isEmpty, nil
	}

	resource, err := parseSingleSegmentResourceQuery(queryParam, query, queryParam, lookup)
	if err != nil {
		return parseResponse{}, false, err
	}

	return resource, isEmpty, nil
}

func parseIDBasedQuery(r *http.Request, queryParam string, maxID int) (int32, bool, error) {
	query := r.URL.Query().Get(queryParam)
	isEmpty := false

	if query == "" {
		isEmpty = true
		return 0, isEmpty, nil
	}

	id, err := strconv.Atoi(query)
	if err != nil {
		return 0, false, newHTTPError(http.StatusBadRequest, "invalid id", err)
	}

	if id > maxID || id <= 0 {
		return 0, false, newHTTPError(http.StatusBadRequest, fmt.Sprintf("provided %s ID %d is out of range. Max ID: %d", queryParam, id, maxID), err)
	}

	return int32(id), isEmpty, nil
}

func queryStrToInt(s string, defaultVal int) (int, error) {
	if s == "" {
		return defaultVal, nil
	}

	val, err := strconv.Atoi(s)
	if err != nil {
		return 0, err
	}

	return val, nil
}
