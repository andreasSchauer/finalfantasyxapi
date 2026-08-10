package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func handleEndpointList[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInput[T, R, A, L]) {
	ids, err := i.retrieveFunc(r, i)
	if handleHTTPError(w, err) {
		return
	}
	resources := idsToAPIResources(cfg, i, ids)

	resourceList, err := i.resToListFunc(cfg, r, resources)
	if handleHTTPError(w, err) {
		return
	}
	respondWithJSON(w, http.StatusOK, resourceList)
}

func handleEndpointIDOnly[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInput[T, R, A, L], segments []string) {
	segment := segments[0]

	if segment == string(snParameters) {
		handleParameters(cfg, w, r, i.endpoint, i.queryLookup)
		return
	}

	if segment == string(snSections) {
		handleSections(cfg, w, r, i)
		return
	}

	if segment == string(snSimple) {
		handleSimple(cfg, w, r, i)
		return
	}

	parseRes, err := parseID(segment, i.resTypeSingle, len(i.objLookupID))
	if handleHTTPError(w, err) {
		return
	}

	resource, err := i.getSingleFunc(r, i, parseRes.ID)
	if handleHTTPError(w, err) {
		return
	}

	respondWithJSON(w, http.StatusOK, resource)
}

func handleEndpointNameOrID[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInput[T, R, A, L], segments []string) {
	segment := segments[0]

	if segment == string(snParameters) {
		handleParameters(cfg, w, r, i.endpoint, i.queryLookup)
		return
	}

	if segment == string(snSections) {
		handleSections(cfg, w, r, i)
		return
	}

	if segment == string(snSimple) {
		handleSimple(cfg, w, r, i)
		return
	}

	parseRes, err := parseSingleSegmentResource(i.resTypeSingle, segment, i.objLookup)
	if handleHTTPError(w, err) {
		return
	}

	if i.getMultipleQuery != nil && parseRes.Name != "" {
		resources, err := getMultipleAPIResources(cfg, r, i, parseRes.Name)
		if handleHTTPError(w, err) {
			return
		}
		respondWithJSON(w, http.StatusMultipleChoices, resources)
		return
	}

	resource, err := i.getSingleFunc(r, i, parseRes.ID)
	if handleHTTPError(w, err) {
		return
	}

	respondWithJSON(w, http.StatusOK, resource)
}

func handleEndpointNameVersion[T seeding.Lookupable, R any, A APIResource, L APIResourceList](w http.ResponseWriter, r *http.Request, i handlerInput[T, R, A, L], segments []string) {
	name := segments[0]
	versionStr := segments[1]

	parseRes, err := parseNameVersionResource(i.resTypeSingle, name, versionStr, i.objLookup)
	if handleHTTPError(w, err) {
		return
	}

	resource, err := i.getSingleFunc(r, i, parseRes.ID)
	if handleHTTPError(w, err) {
		return
	}

	respondWithJSON(w, http.StatusOK, resource)
}

func handleEndpointSubsections[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInput[T, R, A, L], segments []string) {
	posIDStr := segments[0]
	idIsValid := isValidInt(posIDStr)
	posSection := segments[1]
	sectionIsInt := isValidInt(posSection)

	switch {
	// /ep/a/2 + /ep/a/a
	case !idIsValid:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", getUsageString(i)), nil)
		return

	// /ep/2/2
	case sectionIsInt:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("invalid subsection '%s'. subsection can't be an integer. use /api/%s/sections for valid subsections.", posSection, i.endpoint), nil)
		return

	// /ep/2/a (no subsections)
	case i.subsections == nil:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("endpoint /%s doesn't have any subsections.", i.endpoint), nil)
		return

	// /ep/2/a (subsections)
	case i.subsections != nil:
		handleSubsection(cfg, w, r, i, segments)
		return
	}
}

func handleEndpointSubOrNameVer[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInput[T, R, A, L], segments []string) {
	isSubsection, isNameVersion, subsectionIsInt := getSegmentCases(segments)

	switch {
	case isSubsection:
		handleSubsection(cfg, w, r, i, segments)
		return

	case isNameVersion:
		handleEndpointNameVersion(w, r, i, segments)
		return

	case subsectionIsInt:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("invalid subsection '%s'. subsection can't be an integer. use /api/%s/sections for available subsections.", segments[1], i.endpoint), nil)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", getUsageString(i)), nil)
		return
	}
}

// if len(segments) is 0, this function crafts the response for /enums
func handleEnumsEndpointList(cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInputEnums) {
	err := verifyQueryParams[string](r, i.endpoint, i.queryLookup, nil, nil)
	if handleHTTPError(w, err) {
		return
	}

	enums := enumLookupToSlice(i.lookup)
	resources := enumsToNamedAPIResources(cfg, enums)

	resourceList, err := newNamedAPIResourceList(cfg, r, resources)
	if handleHTTPError(w, err) {
		return
	}
	respondWithJSON(w, http.StatusOK, resourceList)
}

// if len(segments) is 1, this function crafts the response for /enums/{enum_name} and /enums/parameters
func handleEnumsEndpointSingle(cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInputEnums, segments []string) {
	segment := segments[0]

	if segment == string(snParameters) {
		handleParameters(cfg, w, r, i.endpoint, i.queryLookup)
		return
	}

	enum, key, err := parseEnumsEndpointSegment(cfg, i, segment)
	if handleHTTPError(w, err) {
		return
	}

	// happens after enum parsing, just for error priority reasons. if I prioritize wrong query param errors over wrong enum errors, I can simply swap the two functions
	err = verifyQueryParams(r, i.endpoint, i.queryLookup, &key, nil)
	if handleHTTPError(w, err) {
		return
	}

	respondWithJSON(w, http.StatusOK, enum)
}

// if len(segments) is 0, this function crafts the response for /endpoints
func handleEndpointsEndpointList(cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInputEndpoints) {
	err := verifyQueryParams[string](r, i.endpoint, i.queryLookup, nil, nil)
	if handleHTTPError(w, err) {
		return
	}

	endpoints := i.slice
	resources := endpointsToNamedAPIResources(cfg, endpoints)
	setLimitMax(cfg, r, r.URL.Query())

	resourceList, err := newNamedAPIResourceList(cfg, r, resources)
	if handleHTTPError(w, err) {
		return
	}
	respondWithJSON(w, http.StatusOK, resourceList)
}

// if len(segments) is 1, this function checks if /endpoints was combined with 'parameters'. returns an error, if it wasn't
func handleEndpointsEndpointSingle(cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInputEndpoints, segments []string) {
	segment := segments[0]

	if segment == string(snParameters) {
		handleParameters(cfg, w, r, i.endpoint, i.queryLookup)
		return
	}

	respondWithError(w, http.StatusBadRequest, "wrong format. '/endpoints' doesn't support single-resource requests.", nil)
}

func handleServiceEndpointSingle[R any](cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInputService[R], segments []string) {
	segment := segments[0]

	if segment == string(snParameters) {
		handleParameters(cfg, w, r, i.endpoint, i.queryLookup)
		return
	}

	respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. '/%s' doesn't support single-resource requests.", i.endpoint), nil)
}

func handleEndpointService[R any](cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInputService[R]) {
	result, err := i.serviceFn(cfg, r, i)
	if handleHTTPError(w, err) {
		return
	}

	respondWithJSON(w, http.StatusOK, result)
}