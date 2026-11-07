package handlerinput

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/queryparam"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

type HandlerInputEnums struct {
	Endpoint    ts.EndpointName
	Usage       []string
	Lookup      map[string]enums.EnumResponse
	QueryLookup map[ts.QueryParamName]queryparam.QueryParam
}

func (i HandlerInputEnums) Retrieve(r *http.Request) (apiresource.APIResourceList[apiresource.NamedAPIResource], error) {
	enums := enums.EnumLookupToSlice(i.Lookup)
	resources := enumsToNamedAPIResources(enums)

	return NewAPIResourceList(r, resources)
}

func (i HandlerInputEnums) VerifyQueriesGet(r *http.Request, key string) error {
	return verifyQueryParams(r, i.Endpoint, i.QueryLookup, &key, nil)
}

func (i HandlerInputEnums) VerifyQueriesRetrieve(r *http.Request) error {
	return verifyQueryParams[string](r, i.Endpoint, i.QueryLookup, nil, nil)
}

func (i HandlerInputEnums) VerifyQueriesAltList(r *http.Request, sectionName ts.SectionName) error {
	sn := string(sectionName)
	return verifyQueryParamsAltList(r, i.Endpoint, i.QueryLookup, &sn)
}
