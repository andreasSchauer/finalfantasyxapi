package api

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type SubResourceList struct {
	ListParams
	ParentResource APIResource   `json:"parent_resource,omitempty"`
	Results        []SubResource `json:"results"`
}

func (l SubResourceList) getListParams() ListParams {
	return l.ListParams
}

func (l SubResourceList) getResults() []SubResource {
	return l.Results
}

func newSubResourceList[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], id int32, sectionName string, fns SubSectionFns) (SubResourceList, error) {
	dbIDs, err := fns.dbQuery(r.Context(), id)
	if err != nil {
		return SubResourceList{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't retrieve %s of %s with id '%d'", sectionName, i.resourceType, id), err)
	}

	results, err := createSubResources(cfg, r, dbIDs, fns.createSubFn)
	if err != nil {
		return SubResourceList{}, err
	}

	listParams, shownResults, err := createPaginatedList(cfg, r, results)
	if err != nil {
		return SubResourceList{}, err
	}

	subResList := SubResourceList{
		ListParams:     listParams,
		ParentResource: i.idToResFunc(cfg, i, id),
		Results:        shownResults,
	}

	return subResList, nil
}



func createSubResources(cfg *Config, r *http.Request, dbIDs []int32, createFn func(*Config, *http.Request, int32) (SubResource, error)) ([]SubResource, error) {
	subs := []SubResource{}
	
	for _, id := range dbIDs {
		subRes, err := createFn(cfg, r, id)
		if err != nil {
			return nil, err
		}

		subs = append(subs, subRes)
	}

	return subs, nil
}