package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/errs"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/handlerinput"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/queryparam"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

func handleSubsection[T seedtypes.Lookupable, R any, A apiresource.APIResource](w http.ResponseWriter, r *http.Request, i handlerinput.Main[T, R, A], idStr, sectionName string) {
	parseRes, err := parseID(idStr, i.ResTypeSingle, len(i.ObjLookup))
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
	// it doesn't need a db query and can create a SimpleResource with the given id,
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

func handleParameters[T seedtypes.Lookupable, R any, A apiresource.APIResource](w http.ResponseWriter, r *http.Request, i handlerinput.Main[T, R, A]) {
	err := i.VerifyQueriesAltList(r, ts.SnParameters)
	if HandleHTTPError(w, err) {
		return
	}

	handleParametersWrapper(w, r, i.Endpoint, i.QueryLookup)
}

func handleParametersEnums(w http.ResponseWriter, r *http.Request, i handlerinput.Enums) {
	err := i.VerifyQueriesAltList(r, ts.SnParameters)
	if HandleHTTPError(w, err) {
		return
	}

	handleParametersWrapper(w, r, i.Endpoint, i.QueryLookup)
}

func handleParametersEndpoints(w http.ResponseWriter, r *http.Request, i handlerinput.Endpoints) {
	err := i.VerifyQueriesAltList(r, ts.SnParameters)
	if HandleHTTPError(w, err) {
		return
	}

	handleParametersWrapper(w, r, i.Endpoint, i.QueryLookup)
}

func handleParametersWrapper(w http.ResponseWriter, r *http.Request, endpoint ts.EndpointName, queryLookup map[ts.QueryParamName]queryparam.QueryParam) {
	queryparam.SetLimitMax(r, r.URL.Query())

	parameterList, err := handlerinput.GetQueryParamList(r, endpoint, queryLookup)
	if HandleHTTPError(w, err) {
		return
	}
	respondWithJSON(w, http.StatusOK, parameterList)
}

func handleSections[T seedtypes.Lookupable, R any, A apiresource.APIResource](w http.ResponseWriter, r *http.Request, i handlerinput.Main[T, R, A]) {
	err := i.VerifyQueriesAltList(r, ts.SnSections)
	if HandleHTTPError(w, err) {
		return
	}

	queryparam.SetLimitMax(r, r.URL.Query())

	sectionList, err := handlerinput.GetSectionList(r, i.Subsections)
	if HandleHTTPError(w, err) {
		return
	}
	respondWithJSON(w, http.StatusOK, sectionList)
}

func handleSimple[T seedtypes.Lookupable, R any, A apiresource.APIResource](w http.ResponseWriter, r *http.Request, i handlerinput.Main[T, R, A]) {
	_, ok := i.Subsections[ts.SnSimple]
	if !ok {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("simple view is not available for endpoint /%s", i.Endpoint), nil)
		return
	}

	ids, err := handleIDsQuery(r, i)
	if HandleHTTPError(w, err) {
		return
	}

	subsection := i.Subsections[ts.SnSimple]
	subResList, err := handlerinput.CreateSimpleResourceList(r, ids, subsection)
	if HandleHTTPError(w, err) {
		return
	}

	respondWithJSON(w, http.StatusOK, subResList)
}

func handleIDsQuery[T seedtypes.Lookupable, R any, A apiresource.APIResource](r *http.Request, i handlerinput.Main[T, R, A]) ([]int32, error) {
	err := i.VerifyIDsQuery(r)
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
