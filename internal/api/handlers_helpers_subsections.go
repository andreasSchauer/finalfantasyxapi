package api

import (
	"fmt"
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/queryparam"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func HandleSubsection[T seedtypes.Lookupable, R any, A apiresource.APIResource](w http.ResponseWriter, r *http.Request, i handlerinput.HandlerInput[T, R, A], segments []string) {
	idStr := segments[0]
	sectionName := segments[1]

	parseRes, err := ah.ParseID(idStr, i.ResTypeSingle, len(i.ObjLookup))
	if HandleHTTPError(w, err) {
		return
	}

	subsection, ok := i.Subsections[ts.SectionName(sectionName)]
	if !ok {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("subsection '%s' does not exist for endpoint /%s. supported subsections: %s", sectionName, i.Endpoint, handlerinput.FormatSectionNames(i.Subsections)), nil)
		return
	}

	q := r.URL.Query()
	if len(q) > 0 {
		respondWithError(w, http.StatusBadRequest, "query parameters can't be used in combination with subsections", nil)
		return
	}

	queryparam.SetLimitMax(r, q)

	// if a subsection refers to the resource itself (/endpoint/{id}/simple, or /aeons/{id}/stats),
	// it doesn't need a db query and can create a apiresource.SimpleResource with the given id,
	// instead of querying the db for relations of the given id
	if subsection.DBQuery == nil {
		if subsection.RelationsFn != nil {
			var err error
			subsection.Relations, err = subsection.RelationsFn(r, []int32{parseRes.ID})
			if HandleHTTPError(w, err) {
				return
			}
		}

		simpleRes, err := subsection.CreateSubFn(r, parseRes.ID, subsection)
		if HandleHTTPError(w, err) {
			return
		}
		respondWithJSON(w, http.StatusOK, simpleRes)
		return
	}

	list, err := i.ToSimpleResourceList(r, parseRes.ID, sectionName, subsection)
	if HandleHTTPError(w, err) {
		return
	}
	respondWithJSON(w, http.StatusOK, list)
}

func HandleParameters[T seedtypes.Lookupable, R any, A apiresource.APIResource](w http.ResponseWriter, r *http.Request, i handlerinput.HandlerInput[T, R, A]) {
	handleParametersWrapper(w, r, i.Endpoint, i.QueryLookup)
}

func HandleParametersEnums(w http.ResponseWriter, r *http.Request, i handlerinput.HandlerInputEnums) {
	handleParametersWrapper(w, r, i.Endpoint, i.QueryLookup)
}

func HandleParametersEndpoints(w http.ResponseWriter, r *http.Request, i handlerinput.HandlerInputEndpoints) {
	handleParametersWrapper(w, r, i.Endpoint, i.QueryLookup)
}

func handleParametersWrapper(w http.ResponseWriter, r *http.Request, endpoint ts.EndpointName, queryLookup map[ts.QueryParamName]queryparam.QueryParam) {
	segment := string(ts.SnParameters)
	err := handlerinput.VerifyQueryParamsAltList(r, endpoint, queryLookup, &segment)
	if HandleHTTPError(w, err) {
		return
	}

	queryparam.SetLimitMax(r, r.URL.Query())

	parameterList, err := getQueryParamList(r, endpoint, queryLookup)
	if HandleHTTPError(w, err) {
		return
	}
	respondWithJSON(w, http.StatusOK, parameterList)
}

func HandleSections[T seedtypes.Lookupable, R any, A apiresource.APIResource](w http.ResponseWriter, r *http.Request, i handlerinput.HandlerInput[T, R, A]) {
	segment := string(ts.SnSections)
	err := handlerinput.VerifyQueryParamsAltList(r, i.Endpoint, i.QueryLookup, &segment)
	if HandleHTTPError(w, err) {
		return
	}

	queryparam.SetLimitMax(r, r.URL.Query())

	sectionList, err := getSectionList(r, i.Subsections)
	if HandleHTTPError(w, err) {
		return
	}
	respondWithJSON(w, http.StatusOK, sectionList)
}

func HandleSimple[T seedtypes.Lookupable, R any, A apiresource.APIResource](w http.ResponseWriter, r *http.Request, i handlerinput.HandlerInput[T, R, A]) {
	segment := string(ts.SnSimple)

	_, ok := i.Subsections[ts.SnSimple]
	if !ok {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("simple view is not available for endpoint /%s", i.Endpoint), nil)
		return
	}

	ids, err := HandleIDsQuery(r, i, &segment)
	if HandleHTTPError(w, err) {
		return
	}

	resources, err := handlerinput.CreateSimpleResources(r, ids, i.Subsections[ts.SectionName(segment)])
	if HandleHTTPError(w, err) {
		return
	}

	listParams, shownResources, err := handlerinput.CreatePaginatedList(r, resources)
	if HandleHTTPError(w, err) {
		return
	}

	subResList := apiresource.SimpleResourceList{
		ListParams: listParams,
		Results:    shownResources,
	}

	respondWithJSON(w, http.StatusOK, subResList)
}

func HandleIDsQuery[T seedtypes.Lookupable, R any, A apiresource.APIResource](r *http.Request, i handlerinput.HandlerInput[T, R, A], segment *string) ([]int32, error) {
	err := handlerinput.VerifyQueryParams[any](r, i.Endpoint, i.QueryLookup, nil, segment)
	if err != nil {
		return nil, err
	}

	// ifs ids param is empty, just retrieve all ids
	queryParamIDs := i.QueryLookup[ts.QpnIDs]
	_, err = queryParamIDs.CheckEmpty(r)
	if errs.QueryIsEmpty(err) {
		return i.RetrieveFunc(r, i)
	}

	ids, err := queryParamIDs.ParseIDList(r, i.ObjLookup)
	if err != nil {
		return nil, err
	}

	queryparam.SetLimitMax(r, r.URL.Query())

	return ids, nil
}
