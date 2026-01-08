package main

import (
	"fmt"
	"net/http"
	"strconv"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type parseResponse struct {
	ID   int32
	Name string // if Name != "", there are multiple resources with that name
}

func newParseResponse(id int32, name string) parseResponse {
	return parseResponse{
		ID:   id,
		Name: name,
	}
}

func parseID(idStr, resourceType string, maxID int) (parseResponse, error) {
	parsedID, err := strconv.Atoi(idStr)
	if err != nil {
		return parseResponse{}, errNotAnID
	}

	if parsedID > maxID || parsedID <= 0 {
		return parseResponse{}, newHTTPError(http.StatusNotFound, fmt.Sprintf("provided %s ID is out of range. Max ID: %d", resourceType, maxID), err)
	}

	return newParseResponse(int32(parsedID), ""), nil
}


func parseUniqueName[T h.HasID](name string, lookup map[string]T) (parseResponse, error) {
	response, err := checkUniqueName(name, lookup)
	if err == nil {
		return response, nil
	}

	nameWithSpaces := h.GetNameWithSpaces(name)

	response, err = checkUniqueName(nameWithSpaces, lookup)
	if err == nil {
		return response, nil
	}

	return parseResponse{}, errNoResource
}

func checkUniqueName[T h.HasID](name string, lookup map[string]T) (parseResponse, error) {
	lookupObj := seeding.LookupObject{
		Name: name,
	}

	// check name/version resources with version = null
	response, err := seeding.GetResource(lookupObj, lookup)
	if err == nil {
		return newParseResponse(response.GetID(), ""), nil
	}

	// check for unique names
	response, err = seeding.GetResource(name, lookup)
	if err == nil {
		return newParseResponse(response.GetID(), ""), nil
	}

	return parseResponse{}, errNoResource
}

func checkNameMultiple[T h.HasID](name string, lookup map[string]T) (parseResponse, error) {
	var testVersion int32 = 1
	lookupObj := seeding.LookupObject{
		Name:    name,
		Version: &testVersion,
	}

	// check for multi-versioned names with dashes
	_, err := seeding.GetResource(lookupObj, lookup)
	if err == nil {
		return newParseResponse(0, lookupObj.Name), nil
	}

	lookupObj.Name = h.GetNameWithSpaces(name)

	// check for multi-versioned names with spaces
	_, err = seeding.GetResource(lookupObj, lookup)
	if err == nil {
		return newParseResponse(0, lookupObj.Name), nil
	}

	return parseResponse{}, errNoResource
}

func checkNameVersion[T h.HasID](name string, version *int32, lookup map[string]T) (parseResponse, error) {
	key := seeding.LookupObject{
		Name:    name,
		Version: version,
	}

	// check for names with dashes
	resource, err := seeding.GetResource(key, lookup)
	if err == nil {
		return newParseResponse(resource.GetID(), ""), nil
	}

	nameWithSpaces := h.GetNameWithSpaces(name)
	key.Name = nameWithSpaces

	// check for names with spaces
	resource, err = seeding.GetResource(key, lookup)
	if err == nil {
		return newParseResponse(resource.GetID(), ""), nil
	}

	return parseResponse{}, errNoResource
}

func parseVersionStr(versionStr string) (*int32, error) {
	var versionPtr *int32

	if versionStr == "" {
		return nil, nil
	}

	version, err := strconv.Atoi(versionStr)
	if err != nil {
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid version number: %s", versionStr), err)
	}
	versionInt32 := int32(version)
	versionPtr = &versionInt32

	return versionPtr, nil
}

