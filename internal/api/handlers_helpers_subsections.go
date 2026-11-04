package api

import (
	"fmt"
	"net/http"

	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

func (i handlerInput[T, R, A]) handleSubsection(server *Server, w http.ResponseWriter, r *http.Request, segments []string) {
	idStr := segments[0]
	sectionName := segments[1]

	parseRes, err := parseID(idStr, i.resTypeSingle, len(i.objLookup))
	if handleHTTPError(w, err) {
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

	setLimitMax(r, q)

	// if a subsection refers to the resource itself (/endpoint/{id}/simple, or /aeons/{id}/stats),
	// it doesn't need a db query and can create a SimpleResource with the given id,
	// instead of querying the db for relations of the given id
	if subsection.dbQuery == nil {
		if subsection.relationsFn != nil {
			var err error
			subsection.relations, err = subsection.relationsFn(server, r, []int32{parseRes.ID})
			if handleHTTPError(w, err) {
				return
			}
		}

		simpleRes, err := subsection.createSubFn(server, r, parseRes.ID, subsection)
		if handleHTTPError(w, err) {
			return
		}
		respondWithJSON(w, http.StatusOK, simpleRes)
		return
	}

	list, err := i.toSimpleResourceList(server, r, parseRes.ID, sectionName, subsection)
	if handleHTTPError(w, err) {
		return
	}
	respondWithJSON(w, http.StatusOK, list)
}

func (i handlerInput[T, R, A]) handleParameters(server *Server, w http.ResponseWriter, r *http.Request) {
	handleParametersWrapper(server, w, r, i.Endpoint, i.queryLookup)
}

func (i handlerInputEnums) handleParameters(server *Server, w http.ResponseWriter, r *http.Request) {
	handleParametersWrapper(server, w, r, i.Endpoint, i.queryLookup)
}

func (i handlerInputEndpoints) handleParameters(server *Server, w http.ResponseWriter, r *http.Request) {
	handleParametersWrapper(server, w, r, i.Endpoint, i.queryLookup)
}

func handleParametersWrapper(server *Server, w http.ResponseWriter, r *http.Request, endpoint ts.EndpointName, queryLookup map[ts.QueryParamName]QueryParam) {
	segment := string(ts.SnParameters)
	err := verifyQueryParamsAltList(r, endpoint, queryLookup, &segment)
	if handleHTTPError(w, err) {
		return
	}

	setLimitMax(r, r.URL.Query())

	parameterList, err := getQueryParamList(server, r, endpoint, queryLookup)
	if handleHTTPError(w, err) {
		return
	}
	respondWithJSON(w, http.StatusOK, parameterList)
}

func (i handlerInput[T, R, A]) handleSections(server *Server, w http.ResponseWriter, r *http.Request) {
	segment := string(ts.SnSections)
	err := verifyQueryParamsAltList(r, i.Endpoint, i.queryLookup, &segment)
	if handleHTTPError(w, err) {
		return
	}

	setLimitMax(r, r.URL.Query())

	sectionList, err := getSectionList(server, r, i.subsections)
	if handleHTTPError(w, err) {
		return
	}
	respondWithJSON(w, http.StatusOK, sectionList)
}

func (i handlerInput[T, R, A]) handleSimple(server *Server, w http.ResponseWriter, r *http.Request) {
	segment := string(ts.SnSimple)

	_, ok := i.subsections[ts.SnSimple]
	if !ok {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("simple view is not available for endpoint /%s", i.Endpoint), nil)
		return
	}

	ids, err := i.queryIDsParam(server, r, &segment)
	if handleHTTPError(w, err) {
		return
	}

	resources, err := createSimpleResources(server, r, ids, i.subsections[ts.SectionName(segment)])
	if handleHTTPError(w, err) {
		return
	}

	listParams, shownResources, err := createPaginatedList(server, r, resources)
	if handleHTTPError(w, err) {
		return
	}

	subResList := SimpleResourceList{
		ListParams: listParams,
		Results:    shownResources,
	}

	respondWithJSON(w, http.StatusOK, subResList)
}

func (i handlerInput[T, R, A]) queryIDsParam(server *Server, r *http.Request, segment *string) ([]int32, error) {
	err := verifyQueryParams[any](r, i.Endpoint, i.queryLookup, nil, segment)
	if err != nil {
		return nil, err
	}

	// ifs ids param is empty, just retrieve all ids
	queryParamIDs := i.queryLookup[ts.QpnIDs]
	_, err = queryParamIDs.checkEmpty(r)
	if queryIsEmpty(err) {
		return i.retrieveFunc(r, i)
	}

	ids, err := queryParamIDs.parseIDList(server, r, i.objLookup)
	if err != nil {
		return nil, err
	}

	setLimitMax(r, r.URL.Query())

	return ids, nil
}
