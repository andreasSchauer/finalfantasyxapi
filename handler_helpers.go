package main

import (
	"fmt"
	"net/http"
	"strconv"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type handlerInput[T h.HasID, R any, L IsAPIResourceList] struct {
	endpoint        string
	resourceType    string
	objLookup       map[string]T
	queryLookup     map[string]QueryType
	getSingleFunc   func(*http.Request, string, int32) (R, error)
	getMultipleFunc func(*http.Request, string, string) (L, error)
	retrieveFunc    func(*http.Request, string) (L, error)
	subSections     map[string]func(string) (IsAPIResourceList, error)
}

func handleEndpointList[T h.HasID, R any, L IsAPIResourceList](w http.ResponseWriter, r *http.Request, i handlerInput[T, R, L]) {
	resourceList, err := i.retrieveFunc(r, i.endpoint)
	if handleHTTPError(w, err) {
		return
	}
	respondWithJSON(w, http.StatusOK, resourceList)
}

func handleEndpointIDOnly[T h.HasID, R any, L IsAPIResourceList](cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInput[T, R, L], segments []string) {
	segment := segments[0]

	handleEndpointParameters(cfg, w, r, i, segment)

	id, err := strconv.Atoi(segment)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("Wrong format. Usage: /api/%s/{id}.", i.endpoint), err)
		return
	}

	if id < 1 || id > len(i.objLookup) {
		respondWithError(w, http.StatusNotFound, fmt.Sprintf("%s with ID %d doesn't exist. Max ID: %d", h.Capitalize(i.resourceType), id, len(i.objLookup)), nil)
		return
	}

	resource, err := i.getSingleFunc(r, i.endpoint, int32(id))
	if handleHTTPError(w, err) {
		return
	}

	respondWithJSON(w, http.StatusOK, resource)
}

func handleEndpointNameOrID[T h.HasID, R any, L IsAPIResourceList](cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInput[T, R, L], segments []string) {
	segment := segments[0]

	handleEndpointParameters(cfg, w, r, i, segment)

	parseRes, err := parseSingleSegmentResource(i.resourceType, segment, "", i.objLookup)
	if handleHTTPError(w, err) {
		return
	}

	if i.getMultipleFunc != nil && parseRes.Name != "" {
		resources, err := i.getMultipleFunc(r, i.endpoint, parseRes.Name)
		if handleHTTPError(w, err) {
			return
		}
		respondWithJSON(w, http.StatusMultipleChoices, resources)
		return
	}

	resource, err := i.getSingleFunc(r, i.endpoint, parseRes.ID)
	if handleHTTPError(w, err) {
		return
	}

	respondWithJSON(w, http.StatusOK, resource)
}

func handleEndpointNameVersion[T h.HasID, R any, L IsAPIResourceList](w http.ResponseWriter, r *http.Request, i handlerInput[T, R, L], segments []string) {
	name := segments[0]
	versionStr := segments[1]

	parseRes, err := parseNameVersionResource(i.resourceType, name, versionStr, "", i.objLookup)
	if handleHTTPError(w, err) {
		return
	}

	resource, err := i.getSingleFunc(r, i.endpoint, parseRes.ID)
	if handleHTTPError(w, err) {
		return
	}

	respondWithJSON(w, http.StatusOK, resource)
}

func handleEndpointParameters[T h.HasID, R any, L IsAPIResourceList](cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInput[T, R, L], segment string) {
	if segment == "parameters" {
		parameterList, err := cfg.getQueryParamList(r, i.queryLookup, i.endpoint)
		if handleHTTPError(w, err) {
			return
		}
		respondWithJSON(w, http.StatusOK, parameterList)
	}
}


func handleHierarchicalEndpoints[T h.HasID, R any, L IsAPIResourceList](subsection string, i handlerInput[T, R, L]) {
	fn, ok := i.subSections[subsection]
	if !ok {
		fmt.Printf("this should trigger an error: subsection %s is not supported. Supported sub-sections: connected, monsters, monster-formations, shops, treasures.\n", subsection)
		return
	}
	fn(subsection)
}
