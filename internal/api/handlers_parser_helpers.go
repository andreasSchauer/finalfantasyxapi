package api

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
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

func parseID(idStr string, resourceType ts.ResTypeSingle, maxID int) (parseResponse, error) {
	response, err := checkID(idStr, resourceType, maxID)
	if errors.Is(err, errs.ErrNotAnID) {
		return parseResponse{}, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid id '%s'", idStr), err)
	}
	if err != nil {
		return parseResponse{}, err
	}

	return response, nil
}

func checkID(idStr string, resourceType ts.ResTypeSingle, maxID int) (parseResponse, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return parseResponse{}, errs.ErrNotAnID
	}

	if id > maxID || id <= 0 {
		return parseResponse{}, newHTTPError(http.StatusNotFound, fmt.Sprintf("%s with provided id '%d' doesn't exist. max id: %d", resourceType, id, maxID), err)
	}

	return newParseResponse(int32(id), ""), nil
}

func checkUniqueName[T seedtypes.Lookupable](name string, l map[string]T) (parseResponse, error) {
	nameWithSpaces := h.GetNameWithSpaces(name, "_")

	lookupObj := seedtypes.LookupObject{
		Name: nameWithSpaces,
	}

	// check name/version resources with version = null
	response, err := lookup.GetResource(lookupObj, l)
	if err == nil {
		return newParseResponse(response.GetID(), ""), nil
	}

	// check for unique names
	response, err = lookup.GetResource(nameWithSpaces, l)
	if err == nil {
		return newParseResponse(response.GetID(), ""), nil
	}

	return parseResponse{}, errs.ErrNoResource
}

func checkNameMultiple[T seedtypes.Lookupable](name string, l map[string]T) (parseResponse, error) {
	nameWithSpaces := h.GetNameWithSpaces(name, "_")
	var testVersion int32 = 1

	lookupObj := seedtypes.LookupObject{
		Name:    nameWithSpaces,
		Version: &testVersion,
	}

	_, err := lookup.GetResource(lookupObj, l)
	if err == nil {
		return newParseResponse(0, lookupObj.Name), nil
	}
	return parseResponse{}, errs.ErrNoResource
}

func checkNameVersion[T seedtypes.Lookupable](name string, version *int32, l map[string]T) (parseResponse, error) {
	nameWithSpaces := h.GetNameWithSpaces(name, "_")

	key := seedtypes.LookupObject{
		Name:    nameWithSpaces,
		Version: version,
	}

	resource, err := lookup.GetResource(key, l)
	if err == nil {
		return newParseResponse(resource.GetID(), ""), nil
	}
	return parseResponse{}, errs.ErrNoResource
}

func parseVersionStr(versionStr string) (*int32, error) {
	var versionPtr *int32

	if versionStr == "" {
		return nil, nil
	}

	version, err := strconv.Atoi(versionStr)
	if err != nil {
		return nil, newHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid version number: '%s'", versionStr), err)
	}
	versionInt32 := int32(version)
	versionPtr = &versionInt32

	return versionPtr, nil
}
