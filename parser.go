package main

import (
	"net/http"
	"net/url"
	"strconv"
	"strings"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

// deals with a single segment path that is either a name or an id and returns the id
// when I have the complete function, everything that returns a 404 returns nil and just at the end, when I've tried everything else will I return a 404


type parseResponse struct {
	ID		int32
	Name	string
}

func newParseResponse (id int32, name string) parseResponse {
	return parseResponse{
		ID: id,
		Name: name, // if Name != "", there are multiple resources with that name
	}
}


// location area parsing might be a whole different level of annoying
// maybe add a switch case LocationArea to GetResource
func parseSingleSegmentResource[T h.HasID](segment string, lookup map[string]T) (parseResponse, error) {
	decoded, err := url.PathUnescape(segment)
	if err != nil {
		return parseResponse{}, newHTTPError(http.StatusBadRequest, "Invalid URL encoding", err)
	}

	// check, if the segment is an id
	parsedID, err := strconv.Atoi(decoded)
	if err == nil {
		return newParseResponse(int32(parsedID), ""), nil
	}

	lookupName := seeding.LookupObject{
		Name: decoded,
	}

	// check for unique names with dashes (obj input)
	resource, err := seeding.GetResource(lookupName, lookup)
	if err == nil {
		return newParseResponse(resource.GetID(), ""), nil
	}

	// check for unique names with dashes (string input)
	resource, err = seeding.GetResource(lookupName.Name, lookup)
	if err == nil {
		return newParseResponse(resource.GetID(), ""), nil
	}

	var testVersion int32 = 1
	lookupNameVer := seeding.LookupObject{
		Name: decoded,
		Version: &testVersion,
	}
	
	// check for multi-versioned names with dashes
	_, err = seeding.GetResource(lookupNameVer, lookup)
	if err == nil {
		return newParseResponse(0, lookupName.Name), nil
	}

	nameWithSpaces := strings.ReplaceAll(decoded, "-", " ")
	lookupName.Name = nameWithSpaces
	lookupNameVer.Name = nameWithSpaces

	// check for unique names with spaces (obj input)
	resource, err = seeding.GetResource(lookupName, lookup)
	if err == nil {
		return newParseResponse(resource.GetID(), ""), nil
	}

	// check for unique names with spaces (str input)
	resource, err = seeding.GetResource(lookupName.Name, lookup)
	if err == nil {
		return newParseResponse(resource.GetID(), ""), nil
	}

	// check for multi-versioned names with spaces
	_, err = seeding.GetResource(lookupNameVer, lookup)
	if err == nil {
		return newParseResponse(0, lookupNameVer.Name), nil
	}

	return parseResponse{}, newHTTPError(http.StatusNotFound, "Resource not found", err)
}


func parseNameVersionResource[T h.HasID](name, versionStr string, lookup map[string]T) (parseResponse, error) {
	version, err := (strconv.Atoi(versionStr))
	if err != nil {
		return parseResponse{}, newHTTPError(http.StatusBadRequest, "Invalid version number", err)
	}
	versionInt32 := int32(version)

	nameDecoded, err := url.PathUnescape(name)
	if err != nil {
		return parseResponse{}, newHTTPError(http.StatusBadRequest, "Invalid URL encoding", err)
	}

	lookupObj := seeding.LookupObject{
		Name: 		nameDecoded,
		Version: 	&versionInt32,
	}

	// check for names with dashes
	resource, err := seeding.GetResource(lookupObj, lookup)
	if err == nil {
		return newParseResponse(resource.GetID(), ""), nil
	}

	nameWithSpaces := strings.ReplaceAll(name, "-", " ")
	lookupObj.Name = nameWithSpaces

	// check for names with spaces
	resource, err = seeding.GetResource(lookupObj, lookup)
	if err == nil {
		return newParseResponse(resource.GetID(), ""), nil
	}

	return parseResponse{}, newHTTPError(http.StatusNotFound, "Resource not found", err)
}