package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
)

type handlerInput[T seedtypes.Lookupable, R any, A APIResource] struct {
	Endpoint         EndpointName
	resTypeSingle    ResTypeSingle
	resTypePlural    ResTypePlural
	usage            []string
	objLookup        map[string]T
	objLookupID      map[int32]T
	queryLookup      map[QueryParamName]QueryParam
	getMultipleQuery DBQueryStringMany
	retrieveQuery    DBQueryNoInput
	idToResFunc      func(*Config, handlerInput[T, R, A], int32) A
	getSingleFunc    func(*http.Request, handlerInput[T, R, A], int32) (R, error)
	retrieveFunc     func(*http.Request, handlerInput[T, R, A]) ([]int32, error)
	avlFunc          func(*Config, *http.Request, []int32) ([]int32, error)
	subsections      map[SectionName]Subsection
}

type handlerInputEnums struct {
	Endpoint    EndpointName
	usage       []string
	lookup      map[string]EnumResponse
	queryLookup map[QueryParamName]QueryParam
}

type handlerInputEndpoints struct {
	Endpoint    EndpointName
	usage       []string
	slice       []EndpointName
	queryLookup map[QueryParamName]QueryParam
}

type handlerInputService[P serviceParams, R serviceResponse] struct {
	Endpoint  EndpointName
	usage     []string
	verifyFn  func(paramsVerifyer, P) (P, error)
	executeFn func(*Config, P) (R, error)
	docsFn    func(*Config) ParamsDoc
	samplesFn func() []SampleParams[P]
}
