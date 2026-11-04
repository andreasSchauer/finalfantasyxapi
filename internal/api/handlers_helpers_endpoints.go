package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (i handlerInput[T, R, A]) handleList(server *Server, w http.ResponseWriter, r *http.Request) {
	ids, err := i.retrieveFunc(r, i)
	if handleHTTPError(w, err) {
		return
	}
	resources := i.toResources(server, ids)

	resourceList, err := newAPIResourceList(server, r, resources)
	if handleHTTPError(w, err) {
		return
	}
	respondWithJSON(w, http.StatusOK, resourceList)
}

func (i handlerInput[T, R, A]) handleIDOnly(server *Server, w http.ResponseWriter, r *http.Request, segments []string) {
	segment := segments[0]

	if segment == string(ts.SnParameters) {
		i.handleParameters(server, w, r)
		return
	}

	if segment == string(ts.SnSections) {
		i.handleSections(server, w, r)
		return
	}

	if segment == string(ts.SnSimple) {
		i.handleSimple(server, w, r)
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

func (i handlerInput[T, R, A]) handleNameID(server *Server, w http.ResponseWriter, r *http.Request, segments []string) {
	segment := segments[0]

	if segment == string(ts.SnParameters) {
		i.handleParameters(server, w, r)
		return
	}

	if segment == string(ts.SnSections) {
		i.handleSections(server, w, r)
		return
	}

	if segment == string(ts.SnSimple) {
		i.handleSimple(server, w, r)
		return
	}

	parseRes, err := i.parseSingleSegmentResource(segment)
	if handleHTTPError(w, err) {
		return
	}

	if i.getMultipleQuery != nil && parseRes.Name != "" {
		resources, err := i.handleMultiple(server, r, parseRes.Name)
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

func (i handlerInput[T, R, A]) handleNameVersion(w http.ResponseWriter, r *http.Request, segments []string) {
	name := segments[0]
	versionStr := segments[1]

	parseRes, err := i.parseNameVersionResource(name, versionStr)
	if handleHTTPError(w, err) {
		return
	}

	resource, err := i.getSingleFunc(r, i, parseRes.ID)
	if handleHTTPError(w, err) {
		return
	}

	respondWithJSON(w, http.StatusOK, resource)
}

func (i handlerInput[T, R, A]) handleSubsections(server *Server, w http.ResponseWriter, r *http.Request, segments []string) {
	posIDStr := segments[0]
	idIsValid := isValidInt(posIDStr)
	posSection := segments[1]
	sectionIsInt := isValidInt(posSection)

	switch {
	// /ep/a/2 + /ep/a/a
	case !idIsValid:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", i.compileUsage()), nil)
		return

	// /ep/2/2
	case sectionIsInt:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("invalid subsection '%s'. subsection can't be an integer. use /api/%s/sections for valid subsections", posSection, i.Endpoint), nil)
		return

	// /ep/2/a (no subsections)
	case i.subsections == nil:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("endpoint /%s doesn't have any subsections", i.Endpoint), nil)
		return

	// /ep/2/a (subsections)
	case i.subsections != nil:
		i.handleSubsection(server, w, r, segments)
		return
	}
}

func (i handlerInput[T, R, A]) handleSubOrNameVer(server *Server, w http.ResponseWriter, r *http.Request, segments []string) {
	isSubsection, isNameVersion, subsectionIsInt := getSegmentCases(segments)

	switch {
	case isSubsection:
		i.handleSubsection(server, w, r, segments)
		return

	case isNameVersion:
		i.handleNameVersion(w, r, segments)
		return

	case subsectionIsInt:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("invalid subsection '%s'. subsection can't be an integer. use /api/%s/sections for available subsections", segments[1], i.Endpoint), nil)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", i.compileUsage()), nil)
		return
	}
}

// if len(segments) is 0, this function crafts the response for /enums
func (i handlerInputEnums) handleList(server *Server, w http.ResponseWriter, r *http.Request) {
	err := verifyQueryParams[string](r, i.Endpoint, i.queryLookup, nil, nil)
	if handleHTTPError(w, err) {
		return
	}

	enums := enums.EnumLookupToSlice(i.lookup)
	resources := enumsToNamedAPIResources(server, enums)

	resourceList, err := newAPIResourceList(server, r, resources)
	if handleHTTPError(w, err) {
		return
	}
	respondWithJSON(w, http.StatusOK, resourceList)
}

// if len(segments) is 1, this function crafts the response for /enums/{enum_name} and /enums/parameters
func (i handlerInputEnums) handleSingle(server *Server, w http.ResponseWriter, r *http.Request, segments []string) {
	segment := segments[0]

	if segment == string(ts.SnParameters) {
		i.handleParameters(server, w, r)
		return
	}

	enum, key, err := i.parseSegment(server, segment)
	if handleHTTPError(w, err) {
		return
	}

	// happens after enum parsing, just for error priority reasons. if I prioritize wrong query param errors over wrong enum errors, I can simply swap the two functions
	err = verifyQueryParams(r, i.Endpoint, i.queryLookup, &key, nil)
	if handleHTTPError(w, err) {
		return
	}

	respondWithJSON(w, http.StatusOK, enum)
}

// if len(segments) is 0, this function crafts the response for /endpoints
func (i handlerInputEndpoints) handleList(server *Server, w http.ResponseWriter, r *http.Request) {
	err := verifyQueryParams[string](r, i.Endpoint, i.queryLookup, nil, nil)
	if handleHTTPError(w, err) {
		return
	}

	endpoints := i.slice
	resources := endpointsToNamedAPIResources(server, endpoints)
	setLimitMax(r, r.URL.Query())

	resourceList, err := newAPIResourceList(server, r, resources)
	if handleHTTPError(w, err) {
		return
	}
	respondWithJSON(w, http.StatusOK, resourceList)
}

// if len(segments) is 1, this function checks if /endpoints was combined with 'parameters'. returns an error, if it wasn't
func (i handlerInputEndpoints) handleSections(server *Server, w http.ResponseWriter, r *http.Request, segments []string) {
	segment := segments[0]

	if segment == string(ts.SnParameters) {
		i.handleParameters(server, w, r)
		return
	}

	respondWithError(w, http.StatusBadRequest, "wrong format. '/endpoints' doesn't support single-resource requests", nil)
}

func (i handlerInputService[P, R]) handleBody(server *Server, w http.ResponseWriter, r *http.Request, segments []string) {
	err := verifyQueryParamsService(r)
	if handleHTTPError(w, err) {
		return
	}

	segment := segments[0]

	if segment == string(ts.SnBody) {
		respondWithJSON(w, http.StatusOK, getBody(server, i.Endpoint))
		return
	}

	respondWithError(w, http.StatusBadRequest, fmt.Sprintf("'GET api/%s' only supports '/body'", i.Endpoint), nil)
}

func (i handlerInputService[P, R]) handleBodySections(server *Server, w http.ResponseWriter, r *http.Request, segments []string) {
	err := verifyQueryParamsService(r)
	if handleHTTPError(w, err) {
		return
	}

	if segments[0] != string(ts.SnBody) {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("'GET api/%s' only supports '/body'", i.Endpoint), nil)
		return
	}

	switch segments[1] {
	case string(ts.SnDocs):
		respondWithJSON(w, http.StatusOK, i.docsFn(server))
		return

	case string(ts.SnSamples):
		respondWithJSON(w, http.StatusOK, i.samplesFn())
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("'GET api/%s/body' only supports the following sections: 'docs', 'samples'", i.Endpoint), nil)
		return
	}
}

func (i handlerInputService[P, R]) handlePost(server *Server, w http.ResponseWriter, r *http.Request) {
	err := verifyQueryParamsService(r)
	if handleHTTPError(w, err) {
		return
	}

	response, err := servicePost(server, r, i)
	if handleHTTPError(w, err) {
		return
	}

	respondWithJSON(w, http.StatusOK, response)
}
