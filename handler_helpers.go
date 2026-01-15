package main

import (
	"fmt"
	"net/http"
	"strconv"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)


func handleEndpointList[T h.HasID, R any, L APIResourceList](w http.ResponseWriter, r *http.Request, i handlerInput[T, R, L]) {
	resourceList, err := i.retrieveFunc(r)
	if handleHTTPError(w, err) {
		return
	}
	respondWithJSON(w, http.StatusOK, resourceList)
}



func handleEndpointIDOnly[T h.HasID, R any, L APIResourceList](cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInput[T, R, L], segments []string) {
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

	resource, err := i.getSingleFunc(r, parseRes.ID)
	if handleHTTPError(w, err) {
		return
	}

	respondWithJSON(w, http.StatusOK, resource)
}


func handleEndpointNameOrID[T h.HasID, R any, L APIResourceList](cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInput[T, R, L], segments []string) {
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

	if i.getMultipleFunc != nil && parseRes.Name != "" {
		resources, err := i.getMultipleFunc(r, parseRes.Name)
		if handleHTTPError(w, err) {
			return
		}
		respondWithJSON(w, http.StatusMultipleChoices, resources)
		return
	}

	resource, err := i.getSingleFunc(r, parseRes.ID)
	if handleHTTPError(w, err) {
		return
	}

	respondWithJSON(w, http.StatusOK, resource)
}


func handleEndpointNameVersion[T h.HasID, R any, L APIResourceList](w http.ResponseWriter, r *http.Request, i handlerInput[T, R, L], segments []string) {
	name := segments[0]
	versionStr := segments[1]

	parseRes, err := parseNameVersionResource(i.resourceType, name, versionStr, i.objLookup)
	if handleHTTPError(w, err) {
		return
	}

	resource, err := i.getSingleFunc(r, parseRes.ID)
	if handleHTTPError(w, err) {
		return
	}

	respondWithJSON(w, http.StatusOK, resource)
}


func handleParameters[T h.HasID, R any, L APIResourceList](cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInput[T, R, L]) {
	parameterList, err := getQueryParamList(cfg, r, i)
	if handleHTTPError(w, err) {
		return
	}
	respondWithJSON(w, http.StatusOK, parameterList)
}

func handleSections[T h.HasID, R any, L APIResourceList](cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInput[T, R, L]) {
	sectionList, err := getSectionList(cfg, r, i)
	if handleHTTPError(w, err) {
		return
	}
	respondWithJSON(w, http.StatusOK, sectionList)
}


func handleEndpointSubsections[T h.HasID, R any, L APIResourceList](w http.ResponseWriter, i handlerInput[T, R, L], segments []string) {
	posIDStr := segments[0]
	isValidID := isValidInt(posIDStr)

	if isValidID {
		handleSubsection(w, i, segments)
		return
	}
	respondWithError(w, http.StatusBadRequest, fmt.Sprintf("invalid id: '%s'.", posIDStr), nil)
}


func handleEndpointSubOrNameVer[T h.HasID, R any, L APIResourceList](w http.ResponseWriter, r *http.Request, i handlerInput[T, R, L], segments []string) {
	posIDStr := segments[0]
	posVerStr := segments[1]
	idIsInt := isValidInt(posIDStr)
	versionIsInt := isValidInt(posVerStr)

	switch {
	case idIsInt:
		handleSubsection(w, i, segments)
		return

	case !idIsInt && versionIsInt:
		handleEndpointNameVersion(w, r, i, segments)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage for two segments: /api/%s/{name}/{version}, or /api/%s/{id}/{subsection}. available subsections: %s.", i.endpoint, i.endpoint, h.GetMapKeyStr(i.subsections)), nil)
		return
	}
}


func handleSubsection[T h.HasID, R any, L APIResourceList](w http.ResponseWriter, i handlerInput[T, R, L], segments []string) {
	idStr := segments[0]
	subsection := segments[1]

	id, _ := strconv.Atoi(idStr)
	if id < 1 || id > len(i.objLookup) {
		respondWithError(w, http.StatusNotFound, fmt.Sprintf("%s with id '%d' doesn't exist. max id: %d.", i.resourceType, id, len(i.objLookup)), nil)
		return
	}

	fn, ok := i.subsections[subsection]
	if !ok {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("subsection '%s' is not supported. supported subsections: %s.", subsection, h.GetMapKeyStr(i.subsections)), nil)
		return
	}

	list, err := fn(subsection)
	if handleHTTPError(w, err) {
		return
	}
	respondWithJSON(w, http.StatusOK, list)
}



func isValidInt(idStr string) bool {
	_, err := strconv.Atoi(idStr)
	return err == nil
}