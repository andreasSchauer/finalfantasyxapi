package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

func handleSubsection[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInput[T, R, A, L], segments []string) {
	idStr := segments[0]
	sectionName := segments[1]

	parseRes, err := parseID(idStr, i.resTypeSingle, len(i.objLookup))
	if handleHTTPError(w, err) {
		return
	}

	subsection, ok := i.subsections[SectionName(sectionName)]
	if !ok {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("subsection '%s' does not exist for endpoint /%s. supported subsections: %s.", sectionName, i.endpoint, formatSectionNames(i.subsections)), nil)
		return
	}

	q := r.URL.Query()
	if len(q) > 0 {
		respondWithError(w, http.StatusBadRequest, "query parameters can't be used in combination with subsections.", nil)
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

	list, err := newSimpleResourceList(cfg, r, i, parseRes.ID, sectionName, subsection)
	if handleHTTPError(w, err) {
		return
	}
	respondWithJSON(w, http.StatusOK, list)
}


func handleParameters(cfg *Config, w http.ResponseWriter, r *http.Request, endpoint EndpointName, queryLookup map[QueryParamName]QueryParam) {
	segment := string(snParameters)
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

func handleSections[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInput[T, R, A, L]) {
	segment := string(snSections)
	err := verifyQueryParamsAltList(cfg, r, i.endpoint, i.queryLookup, &segment)
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

func handleBody(w http.ResponseWriter, fields []FieldDoc) {
	respondWithJSON(w, http.StatusOK, fields)
}

func handleSimple[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInput[T, R, A, L]) {
	segment := string(snSimple)

	_, ok := i.subsections[snSimple]
	if !ok {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("simple view is not available for endpoint /%s.", i.endpoint), nil)
		return
	}

	ids, err := getIDsQuery(cfg, r, i, &segment)
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

func getIDsQuery[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], segment *string) ([]int32, error) {
	err := verifyQueryParams[any](r, i.endpoint, i.queryLookup, nil, segment)
	if err != nil {
		return nil, err
	}

	// ifs ids param is empty, just retrieve all ids
	queryParamIDs := i.queryLookup[qpnIDs]
	_, err = checkEmptyQuery(r, queryParamIDs)
	if queryIsEmpty(err) {
		return i.retrieveFunc(r, i)
	}

	ids, err := parseIdListQuery(cfg, r, queryParamIDs, i.objLookup)
	if err != nil {
		return nil, err
	}

	setLimitMax(cfg, r, r.URL.Query())

	return ids, nil
}
