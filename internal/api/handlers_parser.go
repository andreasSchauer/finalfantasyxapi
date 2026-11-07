package api

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
)

// ParseSingleSegmentResource deals with a single segment path that is either a name or an id and returns a ah.ParseResponse with the id, if a single match is found, or a name, if multiple matches with that name were found.
func (i HandlerInput[T, R, A]) ParseSingleSegmentResource(segment string) (ah.ParseResponse, error) {
	decoded, err := url.PathUnescape(segment)
	if err != nil {
		return ah.ParseResponse{}, errs.NewHTTPError(http.StatusBadRequest, "invalid url encoding", err)
	}

	response, err := ah.CheckID(decoded, i.resTypeSingle, len(i.objLookup))
	if err == nil {
		return response, nil
	}
	if !errors.Is(err, errs.ErrNotAnID) {
		return ah.ParseResponse{}, err
	}

	response, err = ah.CheckUniqueName(decoded, i.objLookup)
	if err == nil {
		return response, nil
	}

	response, err = ah.CheckNameMultiple(decoded, i.objLookup)
	if err == nil {
		return response, nil
	}

	return ah.ParseResponse{}, errs.NewHTTPError(http.StatusNotFound, fmt.Sprintf("%s not found: '%s'", i.resTypeSingle, segment), err)
}

// ParseNameVersionResource deals with a name/version path and returns a ah.ParseResponse with the id, if a match is found.
func (i HandlerInput[T, R, A]) ParseNameVersionResource(name, versionStr string) (ah.ParseResponse, error) {
	nameDecoded, err := url.PathUnescape(name)
	if err != nil {
		return ah.ParseResponse{}, errs.NewHTTPError(http.StatusBadRequest, "invalid url encoding", err)
	}

	versionPtr, err := ah.ParseVersionStr(versionStr)
	if err != nil {
		return ah.ParseResponse{}, err
	}

	response, err := ah.CheckNameVersion(nameDecoded, versionPtr, i.objLookup)
	if err == nil {
		return response, nil
	}

	return ah.ParseResponse{}, errs.NewHTTPError(http.StatusNotFound, fmt.Sprintf("%s not found: '%s', version '%s'", i.resTypeSingle, name, versionStr), err)
}

func (i HandlerInputEnums) ParseSegment(segment string) (enums.EnumResponse, string, error) {
	key, err := url.PathUnescape(segment)
	if err != nil {
		return enums.EnumResponse{}, "", errs.NewHTTPError(http.StatusBadRequest, "invalid url encoding", err)
	}

	enum, ok := i.lookup[key]
	if !ok {
		return enums.EnumResponse{}, "", errs.NewHTTPError(http.StatusNotFound, fmt.Sprintf("enum '%s' not found", key), nil)
	}
	enum.UsedByEndpoints = ah.EndpointsToURLs(enum.UsedByEndpointsInt)

	return enum, key, nil
}
