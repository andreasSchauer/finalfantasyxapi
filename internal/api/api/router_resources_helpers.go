package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func handleList[T seedtypes.Lookupable, R any, A apiresource.APIResource](w http.ResponseWriter, r *http.Request, i handlerinput.Main[T, R, A]) {
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

func handleIDOnly[T seedtypes.Lookupable, R any, A apiresource.APIResource](w http.ResponseWriter, r *http.Request, i handlerinput.Main[T, R, A], segment string) {
	if segment == string(ts.SnParameters) {
		handleParameters(w, r, i)
		return
	}

	if segment == string(ts.SnSections) {
		handleSections(w, r, i)
		return
	}

	if segment == string(ts.SnSimple) {
		handleSimple(w, r, i)
		return
	}

	parseRes, err := parseID(segment, i.ResTypeSingle, len(i.ObjLookupID))
	if HandleHTTPError(w, err) {
		return
	}

	resource, err := i.GetSingleFunc(r, i, parseRes.ID)
	if HandleHTTPError(w, err) {
		return
	}

	respondWithJSON(w, http.StatusOK, resource)
}

func handleNameID[T seedtypes.Lookupable, R any, A apiresource.APIResource](w http.ResponseWriter, r *http.Request, i handlerinput.Main[T, R, A], segment string) {
	if segment == string(ts.SnParameters) {
		handleParameters(w, r, i)
		return
	}

	if segment == string(ts.SnSections) {
		handleSections(w, r, i)
		return
	}

	if segment == string(ts.SnSimple) {
		handleSimple(w, r, i)
		return
	}

	parseRes, err := parseSingleSegmentResource(i, segment)
	if HandleHTTPError(w, err) {
		return
	}

	if i.GetMultipleQuery != nil && parseRes.Name != "" {
		resources, err := handleMultiple(r, i, parseRes.Name)
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

func handleMultiple[T seedtypes.Lookupable, R any, A apiresource.APIResource](r *http.Request, i handlerinput.Main[T, R, A], name string) (apiresource.APIResourceList[A], error) {
	dbIDs, err := i.GetMultipleQuery(r.Context(), name)
	if err != nil {
		return apiresource.APIResourceList[A]{}, errs.NewHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't get multiple %s with name '%s'", i.ResTypeSingle, name), err)
	}

	return i.ToAPIResList(r, dbIDs)
}

func handleNameVersion[T seedtypes.Lookupable, R any, A apiresource.APIResource](w http.ResponseWriter, r *http.Request, i handlerinput.Main[T, R, A], name, versionStr string) {
	parseRes, err := parseNameVersionResource(i, name, versionStr)
	if HandleHTTPError(w, err) {
		return
	}

	resource, err := i.GetSingleFunc(r, i, parseRes.ID)
	if HandleHTTPError(w, err) {
		return
	}

	respondWithJSON(w, http.StatusOK, resource)
}

func handleSubsections[T seedtypes.Lookupable, R any, A apiresource.APIResource](w http.ResponseWriter, r *http.Request, i handlerinput.Main[T, R, A], posIDStr, posSection string) {
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
		handleSubsection(w, r, i, posIDStr, posSection)
		return
	}
}

func handleSubOrNameVer[T seedtypes.Lookupable, R any, A apiresource.APIResource](w http.ResponseWriter, r *http.Request, i handlerinput.Main[T, R, A], segment1, segment2 string) {
	isSubsection, isNameVersion, subsectionIsInt := getSegmentCases(segment1, segment2)

	switch {
	case isSubsection:
		handleSubsection(w, r, i, segment1, segment2)
		return

	case isNameVersion:
		handleNameVersion(w, r, i, segment1, segment2)
		return

	case subsectionIsInt:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("invalid subsection '%s'. subsection can't be an integer. use /api/%s/sections for available subsections", segment2, i.Endpoint), nil)
		return

	default:
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", i.CompileUsage()), nil)
		return
	}
}

func formatError[T seedtypes.Lookupable, R any, A apiresource.APIResource](w http.ResponseWriter, i handlerinput.Main[T, R, A]) {
	respondWithError(w, http.StatusBadRequest, fmt.Sprintf("wrong format. usage: %s", i.CompileUsage()), nil)
}
