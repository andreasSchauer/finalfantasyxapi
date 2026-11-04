package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

type handlerInput[T seedtypes.Lookupable, R any, A APIResource] struct {
	Endpoint         ts.EndpointName
	resTypeSingle    ts.ResTypeSingle
	resTypePlural    ts.ResTypePlural
	usage            []string
	objLookup        map[string]T
	objLookupID      map[int32]T
	queryLookup      map[ts.QueryParamName]QueryParam
	getMultipleQuery DBQueryStringMany
	retrieveQuery    DBQueryNoInput
	idToResFunc      func(*Config, handlerInput[T, R, A], int32) A
	getSingleFunc    func(*http.Request, handlerInput[T, R, A], int32) (R, error)
	retrieveFunc     func(*http.Request, handlerInput[T, R, A]) ([]int32, error)
	avlFunc          func(*Config, *http.Request, []int32) ([]int32, error)
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
	slice       []ts.EndpointName
	queryLookup map[ts.QueryParamName]QueryParam
}

type handlerInputService[P serviceParams, R serviceResponse] struct {
	Endpoint  ts.EndpointName
	usage     []string
	verifyFn  func(paramsVerifyer, P) (P, error)
	executeFn func(*Config, P) (R, error)
	docsFn    func(*Config) ParamsDoc
	samplesFn func() []SampleParams[P]
}
