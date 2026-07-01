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

	// this is for the simple subsection /endpoint/{id}/simple,
	// also used when aspects of the resource itself need to be simplified (like /aeons/{id}/stats)
	// the resource fetches itself and doesn't need a db query
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

func handleParameters[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInput[T, R, A, L]) {
	segment := string(snParameters)
	err := verifyQueryParamsAltListID(cfg, r, i.endpoint, &segment)
	if handleHTTPError(w, err) {
		return
	}

	setLimitMax(cfg, r, r.URL.Query())

	parameterList, err := getQueryParamList(cfg, r, i.endpoint, i.queryLookup)
	if handleHTTPError(w, err) {
		return
	}
	respondWithJSON(w, http.StatusOK, parameterList)
}

func handleParametersEnums(cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInputEnums) {
	segment := string(snParameters)
	err := verifyQueryParamsAltListKey(cfg, r, i.endpoint, &segment)
	if handleHTTPError(w, err) {
		return
	}
	
	setLimitMax(cfg, r, r.URL.Query())

	parameterList, err := getQueryParamList(cfg, r, i.endpoint, i.queryLookup)
	if handleHTTPError(w, err) {
		return
	}
	respondWithJSON(w, http.StatusOK, parameterList)
}

func handleSections[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInput[T, R, A, L]) {
	segment := string(snSections)
	err := verifyQueryParamsAltListID(cfg, r, i.endpoint, &segment)
	if handleHTTPError(w, err) {
		return
	}
	
	setLimitMax(cfg, r, r.URL.Query())

	sectionList, err := getSectionList(cfg, r, i)
	if handleHTTPError(w, err) {
		return
	}
	respondWithJSON(w, http.StatusOK, sectionList)
}

func handleSimple[T seeding.Lookupable, R any, A APIResource, L APIResourceList](cfg *Config, w http.ResponseWriter, r *http.Request, i handlerInput[T, R, A, L]) {
	segment := string(snSimple)

	_, ok := i.subsections[snSimple]
	if !ok {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("simple view is not available for endpoint /%s.", i.endpoint), nil)
		return
	}

	var ids []int32

	err := verifyQueryParamsID(r, i.endpoint, i.queryLookup, nil, &segment)
	if handleHTTPError(w, err) {
		return
	}

	queryParamIDs := i.queryLookup[qpnIDs]
	_, err = checkEmptyQuery(r, queryParamIDs)
	if queryIsEmpty(err) {
		ids, err = i.retrieveFunc(r, i)
		if handleHTTPError(w, err) {
			return
		}
	} else {
		ids, err = parseIdListQuery(cfg, r, queryParamIDs, i.objLookup)
		if handleHTTPError(w, err) {
			return
		}

		setLimitMax(cfg, r, r.URL.Query())
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
