package api

import (
	"context"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type handlerInput[T h.HasID, R any, A APIResource, L APIResourceList] struct {
	endpoint         string
	resourceType     string
	usage            []string
	objLookup        map[string]T
	objLookupID      map[int32]T
	queryLookup      map[string]QueryType
	getMultipleQuery func(context.Context, string) ([]int32, error)
	retrieveQuery    func(context.Context) ([]int32, error)
	idToResFunc      func(*Config, handlerInput[T, R, A, L], int32) A
	resToListFunc    func(*Config, *http.Request, []A) (L, error)
	getSingleFunc    func(*http.Request, handlerInput[T, R, A, L], int32) (R, error)
	retrieveFunc     func(*http.Request, handlerInput[T, R, A, L]) (L, error)
	subsections      map[string]SubSectionFns
}

type SubSectionFns struct {
	dbQuery     func(context.Context, int32) ([]int32, error)
	createSubFn func(*Config, *http.Request, int32) (SimpleResource, error)
}
