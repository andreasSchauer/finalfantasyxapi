// Package handlerinput ajsdioaüsdj
package handlerinput

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/enums"
	"github.com/andreasSchauer/finalfantasyxapi/internal/queryparam"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seedtypes"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/typestrings"
)

type HandlerInput[T seedtypes.Lookupable, R any, A apiresource.APIResource] struct {
	Endpoint         ts.EndpointName
	resTypeSingle    ts.ResTypeSingle
	resTypePlural    ts.ResTypePlural
	usage            []string
	objLookup        map[string]T
	objLookupID      map[int32]T
	queryLookup      map[ts.QueryParamName]queryparam.QueryParam
	getMultipleQuery DBQueryStringMany
	retrieveQuery    DBQueryNoInput
	idToResFunc      func(HandlerInput[T, R, A], int32) A
	getSingleFunc    func(*http.Request, HandlerInput[T, R, A], int32) (R, error)
	retrieveFunc     func(*http.Request, HandlerInput[T, R, A]) ([]int32, error)
	avlFunc          func(*http.Request, []int32) ([]int32, error)
	subsections      map[ts.SectionName]Subsection
}

type HandlerInputEnums struct {
	Endpoint    ts.EndpointName
	usage       []string
	lookup      map[string]enums.EnumResponse
	queryLookup map[ts.QueryParamName]queryparam.QueryParam
}

type HandlerInputEndpoints struct {
	Endpoint    ts.EndpointName
	usage       []string
	Slice       []ts.EndpointName
	queryLookup map[ts.QueryParamName]queryparam.QueryParam
}
