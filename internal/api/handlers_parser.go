package api

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
)

// deals with a single segment path that is either a name or an id and returns a parseResponse with the id, if a single match is found, or a name, if multiple matches with that name were found.
func (i handlerInput[T, R, A]) parseSingleSegmentResource(segment string) (parseResponse, error) {
	decoded, err := url.PathUnescape(segment)
	if err != nil {
		return parseResponse{}, newHTTPError(http.StatusBadRequest, "invalid url encoding", err)
	}

	response, err := checkID(decoded, i.resTypeSingle, len(i.objLookup))
	if err == nil {
		return response, nil
	}
	if !errors.Is(err, errs.ErrNotAnID) {
		return parseResponse{}, err
	}

	response, err = checkUniqueName(decoded, i.objLookup)
	if err == nil {
		return response, nil
	}

	response, err = checkNameMultiple(decoded, i.objLookup)
	if err == nil {
		return response, nil
	}

	return parseResponse{}, newHTTPError(http.StatusNotFound, fmt.Sprintf("%s not found: '%s'", i.resTypeSingle, segment), err)
}

// deals with a name/version path and returns a parseResponse with the id, if a match is found.
func (i handlerInput[T, R, A]) parseNameVersionResource(name, versionStr string) (parseResponse, error) {
	nameDecoded, err := url.PathUnescape(name)
	if err != nil {
		return parseResponse{}, newHTTPError(http.StatusBadRequest, "invalid url encoding", err)
	}

	versionPtr, err := parseVersionStr(versionStr)
	if err != nil {
		return parseResponse{}, err
	}

	response, err := checkNameVersion(nameDecoded, versionPtr, i.objLookup)
	if err == nil {
		return response, nil
	}

	return parseResponse{}, newHTTPError(http.StatusNotFound, fmt.Sprintf("%s not found: '%s', version '%s'", i.resTypeSingle, name, versionStr), err)
}

func (i handlerInputEnums) parseSegment(cfg *Config, segment string) (enums.EnumResponse, string, error) {
	key, err := url.PathUnescape(segment)
	if err != nil {
		return enums.EnumResponse{}, "", newHTTPError(http.StatusBadRequest, "invalid url encoding", err)
	}

	enum, ok := i.lookup[key]
	if !ok {
		return enums.EnumResponse{}, "", newHTTPError(http.StatusNotFound, fmt.Sprintf("enum '%s' not found", key), nil)
	}
	enum.UsedByEndpoints = endpointsToURLs(cfg, enum.UsedByEndpointsInt)

	return enum, key, nil
}
