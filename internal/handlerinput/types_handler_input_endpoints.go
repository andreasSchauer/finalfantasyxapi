package handlerinput

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/queryparam"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

type Endpoints struct {
	Endpoint    ts.EndpointName
	Usage       []string
	Slice       []ts.EndpointName
	QueryLookup map[ts.QueryParamName]queryparam.QueryParam
}

func (i Endpoints) Retrieve(r *http.Request) (apiresource.APIResourceList[apiresource.Named], error) {
	resources := EndpointsToNamedAPIResources(i.Slice)
	queryparam.SetLimitMax(r, r.URL.Query())

	return NewAPIResourceList(r, resources)
}

func (i Endpoints) VerifyQueries(r *http.Request) error {
	return verifyQueryParams[string](r, i.Endpoint, i.QueryLookup, nil, nil)
}

func (i Endpoints) VerifyQueriesAltList(r *http.Request, sectionName ts.SectionName) error {
	sn := string(sectionName)
	return verifyQueryParamsAltList(r, i.Endpoint, i.QueryLookup, &sn)
}
