package api

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

func handleEndpointList[T h.HasID, R any, A APIResource, L APIResourceList](w http.ResponseWriter, r *http.Request, i handlerInput[T, R, A, L]) {
	resourceList, err := i.retrieveFunc(r, i)
	if handleHTTPError(w, err) {
		return
	}
	respondWithJSON(w, http.StatusOK, resourceList)
}



func handleEndpointIDOnly[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInput[T, R, A, L], segments []string) {
	segment := segments[0]

	if segment == "parameters" {
		handleParameters(cfg, w, r, i)
		return
	}

	if segment == "sections" {
		handleSections(cfg, w, r, i)
		return
	}

	parseRes, err := parseID(segment, i.resourceType, len(i.objLookup))
	if handleHTTPError(w, err) {
		return
	}

	resource, err := i.getSingleFunc(r, i, parseRes.ID)
	if handleHTTPError(w, err) {
		return
	}

	respondWithJSON(w, http.StatusOK, resource)
}



func handleEndpointNameOrID[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInput[T, R, A, L], segments []string) {
	segment := segments[0]

	if segment == "parameters" {
		handleParameters(cfg, w, r, i)
		return
	}

	if segment == "sections" {
		handleSections(cfg, w, r, i)
		return
	}

	parseRes, err := parseSingleSegmentResource(i.resourceType, segment, i.objLookup)
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



func handleEndpointNameVersion[T h.HasID, R any, A APIResource, L APIResourceList](w http.ResponseWriter, r *http.Request, i handlerInput[T, R, A, L], segments []string) {
	name := segments[0]
	versionStr := segments[1]

	parseRes, err := parseNameVersionResource(i.resourceType, name, versionStr, i.objLookup)
	if handleHTTPError(w, err) {
		return
	}

	resource, err := i.getSingleFunc(r, i, parseRes.ID)
	if handleHTTPError(w, err) {
		return
	}

	respondWithJSON(w, http.StatusOK, resource)
}



func handleEndpointSubsections[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInput[T, R, A, L], segments []string) {
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



func handleEndpointSubOrNameVer[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInput[T, R, A, L], segments []string) {
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



func handleSubsection[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInput[T, R, A, L], segments []string) {
	idStr := segments[0]
	subsection := segments[1]

	parseRes, err := parseID(idStr, i.resourceType, len(i.objLookup))
	if handleHTTPError(w, err) {
		return
	}

	fns, ok := i.subsections[subsection]
	if !ok {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("subsection '%s' does not exist for endpoint /%s. supported subsections: %s.", subsection, i.endpoint, h.GetMapKeyStr(i.subsections)), nil)
		return
	}

	list, err := newSubResourceList(cfg, r, i, parseRes.ID, subsection, fns)
	if handleHTTPError(w, err) {
		return
	}
	respondWithJSON(w, http.StatusOK, list)
}



func handleParameters[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInput[T, R, A, L]) {
	parameterList, err := getQueryParamList(cfg, r, i)
	if handleHTTPError(w, err) {
		return
	}
	respondWithJSON(w, http.StatusOK, parameterList)
}



func handleSections[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInput[T, R, A, L]) {
	sectionList, err := getSectionList(cfg, r, i)
	if handleHTTPError(w, err) {
		return
	}
	respondWithJSON(w, http.StatusOK, sectionList)
}
