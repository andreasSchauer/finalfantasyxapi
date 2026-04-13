package api

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type SimpleResourceList struct {
	ListParams
	ParentResource APIResource      `json:"parent_resource,omitempty"`
	Results        []SimpleResource `json:"results"`
}

func (l SimpleResourceList) getListParams() ListParams {
	return l.ListParams
}

func (l SimpleResourceList) getResults() []SimpleResource {
	return l.Results
}


func newSimpleResourceList[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, i handlerInput[T, R, A, L], id int32, sectionName string, subsection Subsection) (SimpleResourceList, error) {
	dbIDs, err := subsection.dbQuery(r.Context(), id)
	if err != nil {
		return SimpleResourceList{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't retrieve %s of %s with id '%d'", sectionName, i.resourceType, id), err)
	}

	results, err := createSimpleResources(cfg, r, dbIDs, subsection)
	if err != nil {
		return SimpleResourceList{}, err
	}

	listParams, shownResults, err := createPaginatedList(cfg, r, results)
	if err != nil {
		return SimpleResourceList{}, err
	}

	subResList := SimpleResourceList{
		ListParams:     listParams,
		ParentResource: i.idToResFunc(cfg, i, id),
		Results:        shownResults,
	}

	return subResList, nil
}

func createSimpleResources(cfg *Config, r *http.Request, dbIDs []int32, subsection Subsection) ([]SimpleResource, error) {
	subs := []SimpleResource{}
	
	if subsection.relationsFn != nil {
		var err error
		subsection.relations, err = subsection.relationsFn(cfg, r, dbIDs)
		if err != nil {
			return nil, err
		}
	}

	for _, id := range dbIDs {
		subRes, err := subsection.createSubFn(cfg, r, id, subsection)
		if err != nil {
			return nil, err
		}

		subs = append(subs, subRes)
	}

	return subs, nil
}








