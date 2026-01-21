package main

import (
	"fmt"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type SubResource interface {
	GetSectionName() string
}

type SubResourceList struct {
	ListParams
	ParentResource 	APIResource 	`json:"parent_resource,omitempty"`
	Results			[]SubResource	`json:"results"`
}


func newSubResourceList[T h.HasID, R any, A APIResource, L APIResourceList](cfg *Config, r *http.Request, iParent handlerInput[T, R, A, L], id int32, sectionName string, fns SubSectionFns) (SubResourceList, error) {
	dbIDs, err := fns.dbQuery(r.Context(), id)
	if err != nil {
		return SubResourceList{}, newHTTPError(http.StatusInternalServerError, fmt.Sprintf("couldn't retrieve %s of %s with id '%d'", sectionName, iParent.resourceType, id), err)
	}

	results := fns.getResultsFn(cfg, dbIDs)
	listParams, shownResults, err := createPaginatedList(cfg, r, results)
	if err != nil {
		return SubResourceList{}, err
	}

	subResList := SubResourceList{
		ListParams: listParams,
		ParentResource: iParent.idToResFunc(cfg, iParent, id),
		Results: shownResults,
	}

	return subResList, nil
}




func toSubResourceSlice[T SubResource](s []T) []SubResource {
	out := make([]SubResource, len(s))
	for i, v := range s {
		out[i] = v
	}
	return out
}
