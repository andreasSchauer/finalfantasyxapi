package api

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/api/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

// parseID deals with a single segment path that can only be an id and returns a ah.ParseResponse with the id, if it is valid and in range.
func parseID(idStr string, resourceType ts.ResTypeSingle, maxID int) (ah.ParseResponse, error) {
	response, err := ah.CheckID(idStr, resourceType, maxID)
	if errors.Is(err, errs.ErrNotAnID) {
		return ah.ParseResponse{}, errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid id '%s'", idStr), err)
	}
	if err != nil {
		return ah.ParseResponse{}, err
	}

	return response, nil
}

// parseSingleSegmentResource deals with a single segment path that is either a name or an id and returns a ah.ParseResponse with the id, if a single match is found, or a name, if multiple matches with that name were found.
func parseSingleSegmentResource[T seedtypes.Lookupable, R any, A apiresource.APIResource](i handlerinput.Main[T, R, A], segment string) (ah.ParseResponse, error) {
	decoded, err := url.PathUnescape(segment)
	if err != nil {
		return ah.ParseResponse{}, errs.NewHTTPError(http.StatusBadRequest, "invalid url encoding", err)
	}

	response, err := ah.CheckID(decoded, i.ResTypeSingle, len(i.ObjLookup))
	if err == nil {
		return response, nil
	}
	if !errors.Is(err, errs.ErrNotAnID) {
		return ah.ParseResponse{}, err
	}

	response, err = ah.CheckUniqueName(decoded, i.ObjLookup)
	if err == nil {
		return response, nil
	}

	response, err = ah.CheckNameMultiple(decoded, i.ObjLookup)
	if err == nil {
		return response, nil
	}

	return ah.ParseResponse{}, errs.NewHTTPError(http.StatusNotFound, fmt.Sprintf("%s not found: '%s'", i.ResTypeSingle, segment), err)
}

// parseNameVersionResource deals with a name/version path and returns a ah.ParseResponse with the id, if a match is found.
func parseNameVersionResource[T seedtypes.Lookupable, R any, A apiresource.APIResource](i handlerinput.Main[T, R, A], name, versionStr string) (ah.ParseResponse, error) {
	nameDecoded, err := url.PathUnescape(name)
	if err != nil {
		return ah.ParseResponse{}, errs.NewHTTPError(http.StatusBadRequest, "invalid url encoding", err)
	}

	versionPtr, err := parseVersionStr(versionStr)
	if err != nil {
		return ah.ParseResponse{}, err
	}

	response, err := ah.CheckNameVersion(nameDecoded, versionPtr, i.ObjLookup)
	if err == nil {
		return response, nil
	}

	return ah.ParseResponse{}, errs.NewHTTPError(http.StatusNotFound, fmt.Sprintf("%s not found: '%s', version '%s'", i.ResTypeSingle, name, versionStr), err)
}

// parseVersionStr evaluates the expected version segment and either returns nil, if the segment wasn't used, or a pointer to an version integer, if a valid integer was given.
func parseVersionStr(versionStr string) (*int32, error) {
	var versionPtr *int32

	if versionStr == "" {
		return nil, nil
	}

	version, err := strconv.Atoi(versionStr)
	if err != nil {
		return nil, errs.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("invalid version number: '%s'", versionStr), err)
	}
	versionInt32 := int32(version)
	versionPtr = &versionInt32

	return versionPtr, nil
}

func parseSegment(i handlerinput.Enums, segment string) (enums.EnumResponse, string, error) {
	key, err := url.PathUnescape(segment)
	if err != nil {
		return enums.EnumResponse{}, "", errs.NewHTTPError(http.StatusBadRequest, "invalid url encoding", err)
	}

	enum, ok := i.Lookup[key]
	if !ok {
		return enums.EnumResponse{}, "", errs.NewHTTPError(http.StatusNotFound, fmt.Sprintf("enum '%s' not found", key), nil)
	}
	enum.UsedByEndpoints = ah.EndpointsToURLs(enum.UsedByEndpointsInt)

	return enum, key, nil
}
