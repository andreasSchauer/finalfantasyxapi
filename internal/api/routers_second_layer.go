package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func HandleList[T seedtypes.Lookupable, R any, A apiresource.APIResource](w http.ResponseWriter, r *http.Request, i handlerinput.HandlerInput[T, R, A]) {
	ids, err := i.RetrieveFunc(r, i)
	if HandleHTTPError(w, err) {
		return
	}
	resources := i.ToAPIResources(ids)

	resourceList, err := handlerinput.NewAPIResourceList(r, resources)
	if HandleHTTPError(w, err) {
		return
	}
	respondWithJSON(w, http.StatusOK, resourceList)
}

func HandleIDOnly[T seedtypes.Lookupable, R any, A apiresource.APIResource](w http.ResponseWriter, r *http.Request, i handlerinput.HandlerInput[T, R, A], segment string) {
	if segment == string(ts.SnParameters) {
		HandleParameters(w, r, i)
		return
	}

	if segment == string(ts.SnSections) {
		HandleSections(w, r, i)
		return
	}

	if segment == string(ts.SnSimple) {
		HandleSimple(w, r, i)
		return
	}

	parseRes, err := ParseID(segment, i.ResTypeSingle, len(i.ObjLookupID))
	if HandleHTTPError(w, err) {
		return
	}

	resource, err := i.GetSingleFunc(r, i, parseRes.ID)
	if HandleHTTPError(w, err) {
		return
	}

	respondWithJSON(w, http.StatusOK, resource)
}

func HandleNameID[T seedtypes.Lookupable, R any, A apiresource.APIResource](w http.ResponseWriter, r *http.Request, i handlerinput.HandlerInput[T, R, A], segment string) {
	if segment == string(ts.SnParameters) {
		HandleParameters(w, r, i)
		return
	}

	if segment == string(ts.SnSections) {
		HandleSections(w, r, i)
		return
	}

	if segment == string(ts.SnSimple) {
		HandleSimple(w, r, i)
		return
	}

	parseRes, err := ParseSingleSegmentResource(i, segment)
	if HandleHTTPError(w, err) {
		return
	}

	if i.GetMultipleQuery != nil && parseRes.Name != "" {
		resources, err := HandleMultiple(r, i, parseRes.Name)
		if HandleHTTPError(w, err) {
			return
		}
		respondWithJSON(w, http.StatusMultipleChoices, resources)
		return
	}

	resource, err := i.GetSingleFunc(r, i, parseRes.ID)
	if HandleHTTPError(w, err) {
		return
	}

	respondWithJSON(w, http.StatusOK, resource)
}

func HandleMultiple[T seedtypes.Lookupable, R any, A apiresource.APIResource](r *http.Request, i handlerinput.HandlerInput[T, R, A], name string) (apiresource.APIResourceList[A], error) {
	dbIDs, err := i.GetMultipleQuery(r.Context(), name)
	if err != nil {
		return apiresource.APIResourceList[A]{}, errs.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get multiple %s with name '%s'", i.ResTypeSingle, name), err)
	}

	return i.ToAPIResList(r, dbIDs)
}

func HandleNameVersion[T seedtypes.Lookupable, R any, A apiresource.APIResource](w http.ResponseWriter, r *http.Request, i handlerinput.HandlerInput[T, R, A], name, versionStr string) {
	parseRes, err := ParseNameVersionResource(i, name, versionStr)
	if HandleHTTPError(w, err) {
		return
	}

	resource, err := i.GetSingleFunc(r, i, parseRes.ID)
	if HandleHTTPError(w, err) {
		return
	}

	respondWithJSON(w, http.StatusOK, resource)
}

func HandleSubsections[T seedtypes.Lookupable, R any, A apiresource.APIResource](w http.ResponseWriter, r *http.Request, i handlerinput.HandlerInput[T, R, A], posIDStr, posSection string) {
	idIsValid := isValidInt(posIDStr)
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
	case i.Subsections == nil:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("endpoint /%s doesn't have any subsections", i.Endpoint), nil)
		return

	// /ep/2/a (subsections)
	case i.Subsections != nil:
		HandleSubsection(w, r, i, posIDStr, posSection)
		return
	}
}

func HandleSubOrNameVer[T seedtypes.Lookupable, R any, A apiresource.APIResource](w http.ResponseWriter, r *http.Request, i handlerinput.HandlerInput[T, R, A], segment1, segment2 string) {
	isSubsection, isNameVersion, subsectionIsInt := getSegmentCases(segment1, segment2)

	switch {
	case isSubsection:
		HandleSubsection(w, r, i, segment1, segment2)
		return

	case isNameVersion:
		HandleNameVersion(w, r, i, segment1, segment2)
		return

	case subsectionIsInt:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("invalid subsection '%s'. subsection can't be an integer. use /api/%s/sections for available subsections", segment2, i.Endpoint), nil)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", i.CompileUsage()), nil)
		return
	}
}

// if len(segments) is 0, this function crafts the response for /enums

func HandleListEnums(w http.ResponseWriter, r *http.Request, i handlerinput.HandlerInputEnums) {
	err := i.VerifyQueriesRetrieve(r)
	if HandleHTTPError(w, err) {
		return
	}

	resourceList, err := i.Retrieve(r)
	if HandleHTTPError(w, err) {
		return
	}

	respondWithJSON(w, http.StatusOK, resourceList)
}

// if len(segments) is 1, this function crafts the response for /enums/{enum_name} and /enums/parameters

func HandleSingleEnums(w http.ResponseWriter, r *http.Request, i handlerinput.HandlerInputEnums, segment string) {
	if segment == string(ts.SnParameters) {
		HandleParametersEnums(w, r, i)
		return
	}

	enum, key, err := ParseSegment(i, segment)
	if HandleHTTPError(w, err) {
		return
	}

	// happens after enum parsing, just for error priority reasons. if I prioritize wrong query param errors over wrong enum errors, I can simply swap the two functions
	err = i.VerifyQueriesGet(r, key)
	if HandleHTTPError(w, err) {
		return
	}

	respondWithJSON(w, http.StatusOK, enum)
}

// HandleListEndpoints crafts the response for /endpoints, if len(segments) is 0
func HandleListEndpoints(w http.ResponseWriter, r *http.Request, i handlerinput.HandlerInputEndpoints) {
	err := i.VerifyQueries(r)
	if HandleHTTPError(w, err) {
		return
	}

	resourceList, err := i.Retrieve(r)
	if HandleHTTPError(w, err) {
		return
	}
	respondWithJSON(w, http.StatusOK, resourceList)
}

// HandleSectionsEndpoints checks if /endpoints was combined with 'parameters' if len(segments) is 1. returns an error, if it wasn't
func HandleSectionsEndpoints(w http.ResponseWriter, r *http.Request, i handlerinput.HandlerInputEndpoints, segment string) {
	if segment == string(ts.SnParameters) {
		HandleParametersEndpoints(w, r, i)
		return
	}

	respondWithError(w, http.StatusBadRequest, "wrong format. '/endpoints' doesn't support single-resource requests", nil)
}

func HandleBody[P ServiceParams, R ServiceResponse](w http.ResponseWriter, r *http.Request, i HandlerInputService[P, R], segment string) {
	err := i.VerifyQueries(r)
	if HandleHTTPError(w, err) {
		return
	}

	if segment == string(ts.SnBody) {
		respondWithJSON(w, http.StatusOK, getBody(i.Endpoint))
		return
	}

	respondWithError(w, http.StatusBadRequest, fmt.Sprintf("'GET api/%s' only supports '/body'", i.Endpoint), nil)
}

func HandleBodySections[P ServiceParams, R ServiceResponse](w http.ResponseWriter, r *http.Request, i HandlerInputService[P, R], segments []string) {
	err := i.VerifyQueries(r)
	if HandleHTTPError(w, err) {
		return
	}

	if segments[0] != string(ts.SnBody) {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("'GET api/%s' only supports '/body'", i.Endpoint), nil)
		return
	}

	switch segments[1] {
	case string(ts.SnDocs):
		respondWithJSON(w, http.StatusOK, i.DocsFn())
		return

	case string(ts.SnSamples):
		respondWithJSON(w, http.StatusOK, i.Samples)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("'GET api/%s/body' only supports the following sections: 'docs', 'samples'", i.Endpoint), nil)
		return
	}
}

func HandlePost[P ServiceParams, R ServiceResponse](w http.ResponseWriter, r *http.Request, i HandlerInputService[P, R]) {
	err := i.VerifyQueries(r)
	if HandleHTTPError(w, err) {
		return
	}

	response, err := i.ServicePost(r)
	if HandleHTTPError(w, err) {
		return
	}

	respondWithJSON(w, http.StatusOK, response)
}
