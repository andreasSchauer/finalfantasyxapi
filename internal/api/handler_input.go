package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

type handlerInput[T seedtypes.Lookupable, R any, A apiresource.APIResource] struct {
	Endpoint         ts.EndpointName
	resTypeSingle    ts.ResTypeSingle
	resTypePlural    ts.ResTypePlural
	usage            []string
	objLookup        map[string]T
	objLookupID      map[int32]T
	queryLookup      map[ts.QueryParamName]QueryParam
	getMultipleQuery DBQueryStringMany
	retrieveQuery    DBQueryNoInput
	idToResFunc      func(handlerInput[T, R, A], int32) A
	getSingleFunc    func(*http.Request, handlerInput[T, R, A], int32) (R, error)
	retrieveFunc     func(*http.Request, handlerInput[T, R, A]) ([]int32, error)
	avlFunc          func(*http.Request, []int32) ([]int32, error)
	subsections      map[ts.SectionName]Subsection
}

type handlerInputEnums struct {
	Endpoint    ts.EndpointName
	usage       []string
	lookup      map[string]enums.EnumResponse
	queryLookup map[ts.QueryParamName]QueryParam
}

type handlerInputEndpoints struct {
	Endpoint    ts.EndpointName
	usage       []string
	Slice       []ts.EndpointName
	queryLookup map[ts.QueryParamName]QueryParam
}

type handlerInputService[P ServiceParams, R ServiceResponse] struct {
	Endpoint  ts.EndpointName
	usage     []string
	verifyFn  func(paramsVerifyer, P) (P, error)
	executeFn func(P) (R, error)
	docsFn    func() ParamsDoc
	samplesFn func() []SampleParams[P]
}
