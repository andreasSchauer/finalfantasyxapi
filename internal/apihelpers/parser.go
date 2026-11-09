package apihelpers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)


func GetPathSegments(path string, endpoint ts.EndpointName) []string {
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


type ParseResponse struct {
	ID   int32
	Name string // if Name != "", there are multiple resources with that name
}

func newParseResponse(id int32, name string) ParseResponse {
	return ParseResponse{
		ID:   id,
		Name: name,
	}
}

func CheckID(idStr string, resourceType ts.ResTypeSingle, maxID int) (ParseResponse, error) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return ParseResponse{}, errs.ErrNotAnID
	}

	if id > maxID || id <= 0 {
		return ParseResponse{}, errs.NewHTTPError(http.StatusNotFound, fmt.Sprintf("%s with provided id '%d' doesn't exist. max id: %d", resourceType, id, maxID), err)
	}

	return newParseResponse(int32(id), ""), nil
}

func CheckUniqueName[T seedtypes.Lookupable](name string, l map[string]T) (ParseResponse, error) {
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

	return ParseResponse{}, errs.ErrNoResource
}

func CheckNameMultiple[T seedtypes.Lookupable](name string, l map[string]T) (ParseResponse, error) {
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
	return ParseResponse{}, errs.ErrNoResource
}

func CheckNameVersion[T seedtypes.Lookupable](name string, version *int32, l map[string]T) (ParseResponse, error) {
	nameWithSpaces := h.GetNameWithSpaces(name, "_")

	key := seedtypes.LookupObject{
		Name:    nameWithSpaces,
		Version: version,
	}

	resource, err := lookup.GetResource(key, l)
	if err == nil {
		return newParseResponse(resource.GetID(), ""), nil
	}
	return ParseResponse{}, errs.ErrNoResource
}
