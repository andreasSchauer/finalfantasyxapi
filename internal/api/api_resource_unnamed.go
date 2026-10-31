package api

import (
	"fmt"
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding"
)

type UnnamedAPIResourceList struct {
	ListParams
	Results []UnnamedAPIResource `json:"results"`
}

func (l UnnamedAPIResourceList) GetListParams() ListParams {
	return l.ListParams
}

func (l UnnamedAPIResourceList) GetResults() []HasAPIResource {
	return toHasAPIResSlice(l.Results)
}

type UnnamedAPIResource struct {
	ID  int32  `json:"id"`
	URL string `json:"url"`
}

func (r UnnamedAPIResource) IsZero() bool {
	return r.ID == 0
}

func (r UnnamedAPIResource) GetID() int32 {
	return r.ID
}

func (r UnnamedAPIResource) GetURL() string {
	return r.URL
}

func (r UnnamedAPIResource) GetTestKey() string {
	return fmt.Sprint(r.ID)
}

func (r UnnamedAPIResource) String() string {
	return fmt.Sprintf("unnamed api resource with url: %s", r.URL)
}

func (r UnnamedAPIResource) GetAPIResource() APIResource {
	return r
}

func idToUnnamedAPIResource[T seeding.IsUnnamed, R any, A APIResource, L APIResourceList](cfg *Config, i handlerInput[T, R, A, L], id int32) UnnamedAPIResource {
	res, _ := seeding.GetResourceByID(id, i.objLookupID) // no error needed, because everything was verified through seeding
	params := res.GetResParamsUnnamed()

	return UnnamedAPIResource{
		ID:  params.ID,
		URL: CreateResourceURL(cfg, i.Endpoint, params.ID),
	}
}

func newUnnamedAPIResourceList(cfg *Config, r *http.Request, resources []UnnamedAPIResource) (UnnamedAPIResourceList, error) {
	listParams, shownResources, err := createPaginatedList(cfg, r, resources)
	if err != nil {
		return UnnamedAPIResourceList{}, err
	}

	list := UnnamedAPIResourceList{
		ListParams: listParams,
		Results:    shownResources,
	}

	return list, nil
}
