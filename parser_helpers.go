package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

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
	response, err := checkID(idStr, resourceType, maxID)
	if errors.Is(err, errNotAnID) {
		return parseResponse{}, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid id '%s'.", idStr), err)
	}
	if err != nil {
		return parseResponse{}, err
	}

	return response, nil
}


func checkID(idStr, resourceType string, maxID int) (parseResponse, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return parseResponse{}, errNotAnID
	}

	if id > maxID || id <= 0 {
		return parseResponse{}, newHTTPError(http.StatusNotFound, fmt.Sprintf("%s with provided id '%d' doesn't exist. max id: %d.", resourceType, id, maxID), err)
	}

	return newParseResponse(int32(id), ""), nil
}

func checkUniqueName[T h.HasID](name string, lookup map[string]T) (parseResponse, error) {
	response, err := checkNameSpelling(name, lookup)
	if err == nil {
		return response, nil
	}

	nameWithSpaces := h.GetNameWithSpaces(name)

	response, err = checkNameSpelling(nameWithSpaces, lookup)
	if err == nil {
		return response, nil
	}

	return parseResponse{}, errNoResource
}

func checkNameSpelling[T h.HasID](name string, lookup map[string]T) (parseResponse, error) {
	nameLower := strings.ToLower(name)
	lookupObj := seeding.LookupObject{
		Name: nameLower,
	}

	// check name/version resources with version = null
	response, err := seeding.GetResource(lookupObj, lookup)
	if err == nil {
		return newParseResponse(response.GetID(), ""), nil
	}

	// check for unique names
	response, err = seeding.GetResource(nameLower, lookup)
	if err == nil {
		return newParseResponse(response.GetID(), ""), nil
	}

	return parseResponse{}, errNoResource
}

func checkNameMultiple[T h.HasID](name string, lookup map[string]T) (parseResponse, error) {
	nameLower := strings.ToLower(name)
	var testVersion int32 = 1
	lookupObj := seeding.LookupObject{
		Name:    nameLower,
		Version: &testVersion,
	}

	// check for multi-versioned names with dashes
	_, err := seeding.GetResource(lookupObj, lookup)
	if err == nil {
		return newParseResponse(0, lookupObj.Name), nil
	}

	lookupObj.Name = h.GetNameWithSpaces(nameLower)

	// check for multi-versioned names with spaces
	_, err = seeding.GetResource(lookupObj, lookup)
	if err == nil {
		return newParseResponse(0, lookupObj.Name), nil
	}

	return parseResponse{}, errNoResource
}

func checkNameVersion[T h.HasID](name string, version *int32, lookup map[string]T) (parseResponse, error) {
	nameLower := strings.ToLower(name)

	key := seeding.LookupObject{
		Name:    nameLower,
		Version: version,
	}

	// check for names with dashes
	resource, err := seeding.GetResource(key, lookup)
	if err == nil {
		return newParseResponse(resource.GetID(), ""), nil
	}

	nameWithSpaces := h.GetNameWithSpaces(nameLower)
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
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid version number: '%s'.", versionStr), err)
	}
	versionInt32 := int32(version)
	versionPtr = &versionInt32

	return versionPtr, nil
}
