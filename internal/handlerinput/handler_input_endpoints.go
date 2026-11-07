package handlerinput

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/queryparam"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

type HandlerInputEndpoints struct {
	Endpoint    ts.EndpointName
	Usage       []string
	Slice       []ts.EndpointName
	QueryLookup map[ts.QueryParamName]queryparam.QueryParam
}

func (i HandlerInputEndpoints) Retrieve(r *http.Request) (apiresource.APIResourceList[apiresource.NamedAPIResource], error) {
	resources := EndpointsToNamedAPIResources(i.Slice)
	queryparam.SetLimitMax(r, r.URL.Query())

	return NewAPIResourceList(r, resources)
}

func (i HandlerInputEndpoints) VerifyQueries(r *http.Request) error {
	return verifyQueryParams[string](r, i.Endpoint, i.QueryLookup, nil, nil)
}

func (i HandlerInputEndpoints) VerifyQueriesAltList(r *http.Request, sectionName ts.SectionName) error {
	sn := string(sectionName)
	return verifyQueryParamsAltList(r, i.Endpoint, i.QueryLookup, &sn)
}