package api

import (
	"fmt"
	"net/http"
)

type SimpleResourceList struct {
	ListParams
	ParentResource APIResource      `json:"parent_resource,omitempty"`
	Results        []SimpleResource `json:"results"`
}

func (l SimpleResourceList) GetListParams() ListParams {
	return l.ListParams
}

func (l SimpleResourceList) GetResults() []SimpleResource {
	return l.Results
}

func (i handlerInput[T, R, A]) toSimpleResourceList(srv *Server, r *http.Request, id int32, sectionName string, subsection Subsection) (SimpleResourceList, error) {
	dbIDs, err := subsection.dbQuery(r.Context(), id)
	if err != nil {
		return SimpleResourceList{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't retrieve %s of %s with id '%d'", sectionName, i.resTypeSingle, id), err)
	}

	results, err := createSimpleResources(srv, r, dbIDs, subsection)
	if err != nil {
		return SimpleResourceList{}, err
	}

	listParams, shownResults, err := createPaginatedList(srv, r, results)
	if err != nil {
		return SimpleResourceList{}, err
	}

	subResList := SimpleResourceList{
		ListParams:     listParams,
		ParentResource: i.idToResFunc(srv, i, id),
		Results:        shownResults,
	}

	return subResList, nil
}

func createSimpleResources(srv *Server, r *http.Request, dbIDs []int32, subsection Subsection) ([]SimpleResource, error) {
	subs := []SimpleResource{}

	if subsection.relationsFn != nil {
		var err error
		subsection.relations, err = subsection.relationsFn(srv, r, dbIDs)
		if err != nil {
			return nil, err
		}
	}

	for _, id := range dbIDs {
		subRes, err := subsection.createSubFn(srv, r, id, subsection)
		if err != nil {
			return nil, err
		}

		subs = append(subs, subRes)
	}

	return subs, nil
}
