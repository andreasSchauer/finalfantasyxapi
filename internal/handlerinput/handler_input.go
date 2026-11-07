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
	ResTypeSingle    ts.ResTypeSingle
	ResTypePlural    ts.ResTypePlural
	Usage            []string
	ObjLookup        map[string]T
	ObjLookupID      map[int32]T
	QueryLookup      map[ts.QueryParamName]queryparam.QueryParam
	GetMultipleQuery DBQueryStringMany
	RetrieveQuery    DBQueryNoInput
	IDToResFunc      func(HandlerInput[T, R, A], int32) A
	GetSingleFunc    func(*http.Request, HandlerInput[T, R, A], int32) (R, error)
	RetrieveFunc     func(*http.Request, HandlerInput[T, R, A]) ([]int32, error)
	AvlFunc          func(*http.Request, []int32) ([]int32, error)
	Subsections      map[ts.SectionName]Subsection
}

type HandlerInputEnums struct {
	Endpoint    ts.EndpointName
	Usage       []string
	Lookup      map[string]enums.EnumResponse
	QueryLookup map[ts.QueryParamName]queryparam.QueryParam
}

type HandlerInputEndpoints struct {
	Endpoint    ts.EndpointName
	Usage       []string
	Slice       []ts.EndpointName
	QueryLookup map[ts.QueryParamName]queryparam.QueryParam
}
