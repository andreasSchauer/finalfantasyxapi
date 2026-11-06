package api

import (
	"fmt"
	"net/http"

	ah "github.com/andreasSchauer/finalfantasyxapi/internal/apihelpers"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/queryparam"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (i handlerInput[T, R, A]) handleSubsection(w http.ResponseWriter, r *http.Request, segments []string) {
	idStr := segments[0]
	sectionName := segments[1]

	parseRes, err := ah.ParseID(idStr, i.resTypeSingle, len(i.objLookup))
	if HandleHTTPError(w, err) {
		return
	}

	subsection, ok := i.subsections[ts.SectionName(sectionName)]
	if !ok {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("subsection '%s' does not exist for endpoint /%s. supported subsections: %s", sectionName, i.Endpoint, formatSectionNames(i.subsections)), nil)
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
	if subsection.dbQuery == nil {
		if subsection.relationsFn != nil {
			var err error
			subsection.relations, err = subsection.relationsFn(r, []int32{parseRes.ID})
			if HandleHTTPError(w, err) {
				return
			}
		}

		simpleRes, err := subsection.createSubFn(r, parseRes.ID, subsection)
		if HandleHTTPError(w, err) {
			return
		}
		respondWithJSON(w, http.StatusOK, simpleRes)
		return
	}

	list, err := i.toSimpleResourceList(r, parseRes.ID, sectionName, subsection)
	if HandleHTTPError(w, err) {
		return
	}
	respondWithJSON(w, http.StatusOK, list)
}

func (i handlerInput[T, R, A]) handleParameters(w http.ResponseWriter, r *http.Request) {
	handleParametersWrapper(w, r, i.Endpoint, i.queryLookup)
}

func (i handlerInputEnums) handleParameters(w http.ResponseWriter, r *http.Request) {
	handleParametersWrapper(w, r, i.Endpoint, i.queryLookup)
}

func (i handlerInputEndpoints) handleParameters(w http.ResponseWriter, r *http.Request) {
	handleParametersWrapper(w, r, i.Endpoint, i.queryLookup)
}

func handleParametersWrapper(w http.ResponseWriter, r *http.Request, endpoint ts.EndpointName, queryLookup map[ts.QueryParamName]queryparam.QueryParam) {
	segment := string(ts.SnParameters)
	err := verifyQueryParamsAltList(r, endpoint, queryLookup, &segment)
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

func (i handlerInput[T, R, A]) handleSections(w http.ResponseWriter, r *http.Request) {
	segment := string(ts.SnSections)
	err := verifyQueryParamsAltList(r, i.Endpoint, i.queryLookup, &segment)
	if HandleHTTPError(w, err) {
		return
	}

	queryparam.SetLimitMax(r, r.URL.Query())

	sectionList, err := getSectionList(r, i.subsections)
	if HandleHTTPError(w, err) {
		return
	}
	respondWithJSON(w, http.StatusOK, sectionList)
}

func (i handlerInput[T, R, A]) handleSimple(w http.ResponseWriter, r *http.Request) {
	segment := string(ts.SnSimple)

	_, ok := i.subsections[ts.SnSimple]
	if !ok {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("simple view is not available for endpoint /%s", i.Endpoint), nil)
		return
	}

	ids, err := i.queryIDsParam(r, &segment)
	if HandleHTTPError(w, err) {
		return
	}

	resources, err := createSimpleResources(r, ids, i.subsections[ts.SectionName(segment)])
	if HandleHTTPError(w, err) {
		return
	}

	listParams, shownResources, err := CreatePaginatedList(r, resources)
	if HandleHTTPError(w, err) {
		return
	}

	subResList := apiresource.SimpleResourceList{
		ListParams: listParams,
		Results:    shownResources,
	}

	respondWithJSON(w, http.StatusOK, subResList)
}

func (i handlerInput[T, R, A]) queryIDsParam(r *http.Request, segment *string) ([]int32, error) {
	err := verifyQueryParams[any](r, i.Endpoint, i.queryLookup, nil, segment)
	if err != nil {
		return nil, err
	}

	// ifs ids param is empty, just retrieve all ids
	queryParamIDs := i.queryLookup[ts.QpnIDs]
	_, err = queryParamIDs.CheckEmpty(r)
	if errs.QueryIsEmpty(err) {
		return i.retrieveFunc(r, i)
	}

	ids, err := queryParamIDs.ParseIDList(r, i.objLookup)
	if err != nil {
		return nil, err
	}

	queryparam.SetLimitMax(r, r.URL.Query())

	return ids, nil
}
