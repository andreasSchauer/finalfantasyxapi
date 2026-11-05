package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type APIResource interface {
	HasAPIResource
	helpers.HasID
	HasURL
	GetTestKey() string
}

type HasAPIResource interface {
	GetAPIResource() APIResource
}

type HasURL interface {
	GetURL() string
}

type APIResourceList[A APIResource] struct {
	ListParams
	Results []A `json:"results"`
}

func (l APIResourceList[A]) GetListParams() ListParams {
	return l.ListParams
}

func (l APIResourceList[A]) GetResults() []A {
	return l.Results
}

func newAPIResourceList[A APIResource](r *http.Request, resources []A) (APIResourceList[A], error) {
	listParams, shownResources, err := createPaginatedList(r, resources)
	if err != nil {
		return APIResourceList[A]{}, err
	}

	list := APIResourceList[A]{
		ListParams: listParams,
		Results:    shownResources,
	}

	return list, nil
}

type SimpleResource interface {
	HasURL
}

type NameAmount interface {
	GetName() string
	GetVersion() *int32
	GetVal() int32
}

type ResourceAmountType[A APIResource] interface {
	ToResAmount() ResourceAmount[A]
	NameAmount
}
