package api

import (
	"fmt"
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/queryparam"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (i HandlerInput[T, R, A]) HandleList(w http.ResponseWriter, r *http.Request) {
	ids, err := i.retrieveFunc(r, i)
	if HandleHTTPError(w, err) {
		return
	}
	resources := i.ToAPIResources(ids)

	resourceList, err := newAPIResourceList(r, resources)
	if HandleHTTPError(w, err) {
		return
	}
	respondWithJSON(w, http.StatusOK, resourceList)
}

func (i HandlerInput[T, R, A]) HandleIDOnly(w http.ResponseWriter, r *http.Request, segments []string) {
	segment := segments[0]

	if segment == string(ts.SnParameters) {
		i.HandleParameters(w, r)
		return
	}

	if segment == string(ts.SnSections) {
		i.HandleSections(w, r)
		return
	}

	if segment == string(ts.SnSimple) {
		i.HandleSimple(w, r)
		return
	}

	parseRes, err := ah.ParseID(segment, i.resTypeSingle, len(i.objLookupID))
	if HandleHTTPError(w, err) {
		return
	}

	resource, err := i.getSingleFunc(r, i, parseRes.ID)
	if HandleHTTPError(w, err) {
		return
	}

	respondWithJSON(w, http.StatusOK, resource)
}

func (i HandlerInput[T, R, A]) HandleNameID(w http.ResponseWriter, r *http.Request, segments []string) {
	segment := segments[0]

	if segment == string(ts.SnParameters) {
		i.HandleParameters(w, r)
		return
	}

	if segment == string(ts.SnSections) {
		i.HandleSections(w, r)
		return
	}

	if segment == string(ts.SnSimple) {
		i.HandleSimple(w, r)
		return
	}

	parseRes, err := i.ParseSingleSegmentResource(segment)
	if HandleHTTPError(w, err) {
		return
	}

	if i.getMultipleQuery != nil && parseRes.Name != "" {
		resources, err := i.HandleMultiple(r, parseRes.Name)
		if HandleHTTPError(w, err) {
			return
		}
		respondWithJSON(w, http.StatusMultipleChoices, resources)
		return
	}

	resource, err := i.getSingleFunc(r, i, parseRes.ID)
	if HandleHTTPError(w, err) {
		return
	}

	respondWithJSON(w, http.StatusOK, resource)
}

func (i HandlerInput[T, R, A]) HandleNameVersion(w http.ResponseWriter, r *http.Request, segments []string) {
	name := segments[0]
	versionStr := segments[1]

	parseRes, err := i.ParseNameVersionResource(name, versionStr)
	if HandleHTTPError(w, err) {
		return
	}

	resource, err := i.getSingleFunc(r, i, parseRes.ID)
	if HandleHTTPError(w, err) {
		return
	}

	respondWithJSON(w, http.StatusOK, resource)
}

func (i HandlerInput[T, R, A]) HandleSubsections(w http.ResponseWriter, r *http.Request, segments []string) {
	posIDStr := segments[0]
	idIsValid := isValidInt(posIDStr)
	posSection := segments[1]
	sectionIsInt := isValidInt(posSection)

	switch {
	// /ep/a/2 + /ep/a/a
	case !idIsValid:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", i.CompileUsage()), nil)
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
		i.HandleSubsection(w, r, segments)
		return
	}
}

func (i HandlerInput[T, R, A]) HandleSubOrNameVer(w http.ResponseWriter, r *http.Request, segments []string) {
	isSubsection, isNameVersion, subsectionIsInt := getSegmentCases(segments)

	switch {
	case isSubsection:
		i.HandleSubsection(w, r, segments)
		return

	case isNameVersion:
		i.HandleNameVersion(w, r, segments)
		return

	case subsectionIsInt:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("invalid subsection '%s'. subsection can't be an integer. use /api/%s/sections for available subsections", segments[1], i.Endpoint), nil)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", i.CompileUsage()), nil)
		return
	}
}

// if len(segments) is 0, this function crafts the response for /enums

func (i HandlerInputEnums) HandleList(w http.ResponseWriter, r *http.Request) {
	err := verifyQueryParams[string](r, i.Endpoint, i.queryLookup, nil, nil)
	if HandleHTTPError(w, err) {
		return
	}

	enums := enums.EnumLookupToSlice(i.lookup)
	resources := enumsToNamedAPIResources(enums)

	resourceList, err := newAPIResourceList(r, resources)
	if HandleHTTPError(w, err) {
		return
	}
	respondWithJSON(w, http.StatusOK, resourceList)
}

// if len(segments) is 1, this function crafts the response for /enums/{enum_name} and /enums/parameters

func (i HandlerInputEnums) HandleSingle(w http.ResponseWriter, r *http.Request, segments []string) {
	segment := segments[0]

	if segment == string(ts.SnParameters) {
		i.HandleParameters(w, r)
		return
	}

	enum, key, err := i.ParseSegment(segment)
	if HandleHTTPError(w, err) {
		return
	}

	// happens after enum parsing, just for error priority reasons. if I prioritize wrong query param errors over wrong enum errors, I can simply swap the two functions
	err = verifyQueryParams(r, i.Endpoint, i.queryLookup, &key, nil)
	if HandleHTTPError(w, err) {
		return
	}

	respondWithJSON(w, http.StatusOK, enum)
}

// HandleList crafts the response for /endpoints, if len(segments) is 0
func (i HandlerInputEndpoints) HandleList(w http.ResponseWriter, r *http.Request) {
	err := verifyQueryParams[string](r, i.Endpoint, i.queryLookup, nil, nil)
	if HandleHTTPError(w, err) {
		return
	}

	endpoints := i.Slice
	resources := endpointsToNamedAPIResources(endpoints)
	queryparam.SetLimitMax(r, r.URL.Query())

	resourceList, err := newAPIResourceList(r, resources)
	if HandleHTTPError(w, err) {
		return
	}
	respondWithJSON(w, http.StatusOK, resourceList)
}

// HandleSections checks if /endpoints was combined with 'parameters' if len(segments) is 1. returns an error, if it wasn't
func (i HandlerInputEndpoints) HandleSections(w http.ResponseWriter, r *http.Request, segments []string) {
	segment := segments[0]

	if segment == string(ts.SnParameters) {
		i.HandleParameters(w, r)
		return
	}

	respondWithError(w, http.StatusBadRequest, "wrong format. '/endpoints' doesn't support single-resource requests", nil)
}

func (i HandlerInputService[P, R]) HandleBody(w http.ResponseWriter, r *http.Request, segments []string) {
	err := verifyQueryParamsService(r)
	if HandleHTTPError(w, err) {
		return
	}

	segment := segments[0]

	if segment == string(ts.SnBody) {
		respondWithJSON(w, http.StatusOK, getBody(i.Endpoint))
		return
	}

	respondWithError(w, http.StatusBadRequest, fmt.Sprintf("'GET api/%s' only supports '/body'", i.Endpoint), nil)
}

func (i HandlerInputService[P, R]) HandleBodySections(w http.ResponseWriter, r *http.Request, segments []string) {
	err := verifyQueryParamsService(r)
	if HandleHTTPError(w, err) {
		return
	}

	if segments[0] != string(ts.SnBody) {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("'GET api/%s' only supports '/body'", i.Endpoint), nil)
		return
	}

	switch segments[1] {
	case string(ts.SnDocs):
		respondWithJSON(w, http.StatusOK, i.docsFn())
		return

	case string(ts.SnSamples):
		respondWithJSON(w, http.StatusOK, i.samplesFn())
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("'GET api/%s/body' only supports the following sections: 'docs', 'samples'", i.Endpoint), nil)
		return
	}
}

func (i HandlerInputService[P, R]) HandlePost(w http.ResponseWriter, r *http.Request) {
	err := verifyQueryParamsService(r)
	if HandleHTTPError(w, err) {
		return
	}

	response, err := servicePost(r, i)
	if HandleHTTPError(w, err) {
		return
	}

	respondWithJSON(w, http.StatusOK, response)
}
