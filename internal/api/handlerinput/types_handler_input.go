// Package handlerinput ajsdioaüsdj
package handlerinput

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/queryparam"
	ts "github.com/andreasSchauer/finalfantasyxapi/internal/general/typestrings"
	"github.com/andreasSchauer/finalfantasyxapi/internal/seeding/seedtypes"
)

type Main[T seedtypes.Lookupable, R any, A apiresource.APIResource] struct {
	Endpoint         ts.EndpointName
	ResTypeSingle    ts.ResTypeSingle
	ResTypePlural    ts.ResTypePlural
	Usage            []string
	ObjLookup        map[string]T
	ObjLookupID      map[int32]T
	QueryLookup      map[ts.QueryParamName]queryparam.QueryParam
	GetMultipleQuery DBQueryStringMany
	RetrieveQuery    DBQueryNoInput
	IDToResFunc      func(Main[T, R, A], int32) A
	GetSingleFunc    func(*http.Request, Main[T, R, A], int32) (R, error)
	RetrieveFunc     func(*http.Request, Main[T, R, A]) ([]int32, error)
	AvlFunc          func(*http.Request, []int32) ([]int32, error)
	Subsections      map[ts.SectionName]Subsection
}
