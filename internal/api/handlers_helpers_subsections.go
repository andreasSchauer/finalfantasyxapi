package api

import (
	"fmt"
	"net/http"
)

func (i handlerInput[T, R, A]) handleSubsection(cfg *Config, w http.ResponseWriter, r *http.Request, segments []string) {
	idStr := segments[0]
	sectionName := segments[1]

	parseRes, err := parseID(idStr, i.resTypeSingle, len(i.objLookup))
	if handleHTTPError(w, err) {
		return
	}

	subsection, ok := i.subsections[SectionName(sectionName)]
	if !ok {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("subsection '%s' does not exist for endpoint /%s. supported subsections: %s", sectionName, i.Endpoint, formatSectionNames(i.subsections)), nil)
		return
	}

	q := r.URL.Query()
	if len(q) > 0 {
		respondWithError(w, http.StatusBadRequest, "query parameters can't be used in combination with subsections", nil)
		return
	}

	setLimitMax(cfg, r, q)

	// if a subsection refers to the resource itself (/endpoint/{id}/simple, or /aeons/{id}/stats),
	// it doesn't need a db query and can create a SimpleResource with the given id,
	// instead of querying the db for relations of the given id
	if subsection.dbQuery == nil {
		if subsection.relationsFn != nil {
			var err error
			subsection.relations, err = subsection.relationsFn(cfg, r, []int32{parseRes.ID})
			if handleHTTPError(w, err) {
				return
			}
		}

		simpleRes, err := subsection.createSubFn(cfg, r, parseRes.ID, subsection)
		if handleHTTPError(w, err) {
			return
		}
		respondWithJSON(w, http.StatusOK, simpleRes)
		return
	}

	list, err := i.toSimpleResourceList(cfg, r, parseRes.ID, sectionName, subsection)
	if handleHTTPError(w, err) {
		return
	}
	respondWithJSON(w, http.StatusOK, list)
}

func (i handlerInput[T, R, A]) handleParameters(cfg *Config, w http.ResponseWriter, r *http.Request) {
	handleParametersWrapper(cfg, w, r, i.Endpoint, i.queryLookup)
}

func (i handlerInputEnums) handleParameters(cfg *Config, w http.ResponseWriter, r *http.Request) {
	handleParametersWrapper(cfg, w, r, i.Endpoint, i.queryLookup)
}

func (i handlerInputEndpoints) handleParameters(cfg *Config, w http.ResponseWriter, r *http.Request) {
	handleParametersWrapper(cfg, w, r, i.Endpoint, i.queryLookup)
}

func handleParametersWrapper(cfg *Config, w http.ResponseWriter, r *http.Request, endpoint EndpointName, queryLookup map[QueryParamName]QueryParam) {
	segment := string(SnParameters)
	err := verifyQueryParamsAltList(cfg, r, endpoint, queryLookup, &segment)
	if handleHTTPError(w, err) {
		return
	}

	setLimitMax(cfg, r, r.URL.Query())

	parameterList, err := getQueryParamList(cfg, r, endpoint, queryLookup)
	if handleHTTPError(w, err) {
		return
	}
	respondWithJSON(w, http.StatusOK, parameterList)
}

func (i handlerInput[T, R, A]) handleSections(cfg *Config, w http.ResponseWriter, r *http.Request) {
	segment := string(SnSections)
	err := verifyQueryParamsAltList(cfg, r, i.Endpoint, i.queryLookup, &segment)
	if handleHTTPError(w, err) {
		return
	}

	setLimitMax(cfg, r, r.URL.Query())

	sectionList, err := getSectionList(cfg, r, i.subsections)
	if handleHTTPError(w, err) {
		return
	}
	respondWithJSON(w, http.StatusOK, sectionList)
}

func (i handlerInput[T, R, A]) handleSimple(cfg *Config, w http.ResponseWriter, r *http.Request) {
	segment := string(SnSimple)

	_, ok := i.subsections[SnSimple]
	if !ok {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("simple view is not available for endpoint /%s", i.Endpoint), nil)
		return
	}

	ids, err := i.queryIDsParam(cfg, r, &segment)
	if handleHTTPError(w, err) {
		return
	}

	resources, err := createSimpleResources(cfg, r, ids, i.subsections[SectionName(segment)])
	if handleHTTPError(w, err) {
		return
	}

	listParams, shownResources, err := createPaginatedList(cfg, r, resources)
	if handleHTTPError(w, err) {
		return
	}

	subResList := SimpleResourceList{
		ListParams: listParams,
		Results:    shownResources,
	}

	respondWithJSON(w, http.StatusOK, subResList)
}

func (i handlerInput[T, R, A]) queryIDsParam(cfg *Config, r *http.Request, segment *string) ([]int32, error) {
	err := verifyQueryParams[any](r, i.Endpoint, i.queryLookup, nil, segment)
	if err != nil {
		return nil, err
	}

	// ifs ids param is empty, just retrieve all ids
	queryParamIDs := i.queryLookup[QpnIDs]
	_, err = queryParamIDs.checkEmpty(r)
	if queryIsEmpty(err) {
		return i.retrieveFunc(r, i)
	}

	ids, err := queryParamIDs.parseIDList(cfg, r, i.objLookup)
	if err != nil {
		return nil, err
	}

	setLimitMax(cfg, r, r.URL.Query())

	return ids, nil
}
