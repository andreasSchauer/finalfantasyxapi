package api

import (
	"fmt"

	"github.com/andreasSchauer/finalfantasyxapi/internal/lookup"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

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

func idToUnnamedAPIResource[T seedtypes.IsUnnamed, R any, A APIResource](cfg *Config, i handlerInput[T, R, A], id int32) UnnamedAPIResource {
	res, _ := lookup.GetResourceByID(id, i.objLookupID) // no error needed, because everything was verified through seedtypes
	params := res.GetResParamsUnnamed()

	return UnnamedAPIResource{
		ID:  params.ID,
		URL: CreateResourceURL(cfg, i.Endpoint, params.ID),
	}
}
