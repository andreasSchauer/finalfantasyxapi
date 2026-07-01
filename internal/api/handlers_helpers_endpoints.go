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
		handleParameters(cfg, w, r, i)
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
		handleParameters(cfg, w, r, i)
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

func handleEnumsEndpointList(cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInputEnums) {
	err := verifyQueryParamsKey(r, i.endpoint, i.queryLookup, nil)
	if handleHTTPError(w, err) {
		return
	}

	enums := typeLookupToSlice(i.enumLookup)
	resources := enumsToNamedAPIResources(cfg, enums)

	resourceList, err := newNamedAPIResourceList(cfg, r, resources)
	if handleHTTPError(w, err) {
		return
	}
	respondWithJSON(w, http.StatusOK, resourceList)
}

func handleEnumsEndpointSingle(cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInputEnums, segments []string) {
	segment := segments[0]

	if segment == string(snParameters) {
		handleParametersEnums(cfg, w, r, i)
		return
	}

	enum, err := parseEnumsEndpointSegment(cfg, i, segment)
	if handleHTTPError(w, err) {
		return
	}

	enumName := string(enum.Name)
	err = verifyQueryParamsKey(r, i.endpoint, i.queryLookup, &enumName)
	if handleHTTPError(w, err) {
		return
	}

	respondWithJSON(w, http.StatusOK, enum)
}
