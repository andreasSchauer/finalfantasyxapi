package main

import (
	"context"
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type handlerInput[T h.HasID, R any, A APIResource, L APIResourceList] struct {
	endpoint        string
	resourceType    string
	objLookup       map[string]T
	objLookupID     map[int32]T
	queryLookup     map[string]QueryType
	retrieveQuery   func(context.Context) ([]int32, error)
	idToResFunc     func(*Config, handlerInput[T, R, A, L], int32) A
	resToListFunc	func(*Config, *http.Request, handlerInput[T, R, A, L], []A) (L, error)
	getSingleFunc   func(*http.Request, int32) (R, error)
	getMultipleFunc func(*http.Request, string) (L, error)
	retrieveFunc    func(*http.Request) (L, error)
	subsections     map[string]func(string) (APIResourceList, error)
}

// use handlerView, if the type of the handlerInput must be determined first (like with abilities)
type handlerView interface {
	Endpoint() string
	ResourceType() string
	ObjLookup() map[string]h.HasID
	ObjLookupID() map[int32]h.HasID
	QueryLookup() map[string]QueryType
	GetSingleFunc() func(*http.Request, int32) (any, error)
	GetMultipleFunc() func(*http.Request, string) (APIResourceList, error)
	RetrieveFunc() func(*http.Request) (APIResourceList, error)
	Subsections() map[string]func(string) (APIResourceList, error)
}

func (i handlerInput[T, R, A, L]) Endpoint() string {
	return i.endpoint
}

func (i handlerInput[T, R, A, L]) ResourceType() string {
	return i.resourceType
}

func (i handlerInput[T, R, A, L]) ObjLookup() map[string]h.HasID {
	out := make(map[string]h.HasID, len(i.objLookup))

	for k, v := range i.objLookup {
		out[k] = v
	}

	return out
}

func (i handlerInput[T, R, A, L]) ObjLookupID() map[int32]h.HasID {
	out := make(map[int32]h.HasID, len(i.objLookup))

	for k, v := range i.objLookupID {
		out[k] = v
	}

	return out
}

func (i handlerInput[T, R, A, L]) QueryLookup() map[string]QueryType {
	return i.queryLookup
}

func (i handlerInput[T, R, A, L]) GetSingleFunc() func(*http.Request, int32) (any, error) {
	if i.getSingleFunc == nil {
		return nil
	}
	return func(r *http.Request, id int32) (any, error) {
		return i.getSingleFunc(r, id)
	}
}

func (i handlerInput[T, R, A, L]) GetMultipleFunc() func(*http.Request, string) (APIResourceList, error) {
	if i.retrieveFunc == nil {
		return nil
	}
	return func(r *http.Request, name string) (APIResourceList, error) {
		return i.getMultipleFunc(r, name)
	}
}

func (i handlerInput[T, R, A, L]) RetrieveFunc() func(*http.Request) (APIResourceList, error) {
	if i.getMultipleFunc == nil {
		return nil
	}
	return func(r *http.Request) (APIResourceList, error) {
		return i.retrieveFunc(r)
	}
}

func (i handlerInput[T, R, A, L]) Subsections() map[string]func(string) (APIResourceList, error) {
	return i.subsections
}
