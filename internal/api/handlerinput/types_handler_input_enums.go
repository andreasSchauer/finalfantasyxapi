package handlerinput

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/queryparam"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
)

type Enums struct {
	Endpoint    ts.EndpointName
	Usage       []string
	Lookup      map[string]enums.EnumResponse
	QueryLookup map[ts.QueryParamName]queryparam.QueryParam
}

func (i Enums) Retrieve(r *http.Request) (apiresource.APIResourceList[apiresource.Named], error) {
	enums := enums.EnumLookupToSlice(i.Lookup)
	resources := enumsToNamedAPIResources(enums)

	return NewAPIResourceList(r, resources)
}

func (i Enums) VerifyQueriesGet(r *http.Request, key string) error {
	return verifyQueryParams(r, i.Endpoint, i.QueryLookup, &key, nil)
}

func (i Enums) VerifyQueriesRetrieve(r *http.Request) error {
	return verifyQueryParams[string](r, i.Endpoint, i.QueryLookup, nil, nil)
}

func (i Enums) VerifyQueriesAltList(r *http.Request, sectionName ts.SectionName) error {
	sn := string(sectionName)
	return verifyQueryParamsAltList(r, i.Endpoint, i.QueryLookup, &sn)
}
