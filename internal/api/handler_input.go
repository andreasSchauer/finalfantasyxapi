package api

import (
	"net/http"

	h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type handlerInput[T h.HasID, R any, A APIResource, L APIResourceList] struct {
	endpoint         string
	resourceType     string
	usage            []string
	objLookup        map[string]T
	objLookupID      map[int32]T
	queryLookup      map[string]QueryParam
	getMultipleQuery DbQueryStringMany
	retrieveQuery    DbQueryNoInput
	idToResFunc      func(*Config, handlerInput[T, R, A, L], int32) A
	resToListFunc    func(*Config, *http.Request, []A) (L, error)
	getSingleFunc    func(*http.Request, handlerInput[T, R, A, L], int32) (R, error)
	retrieveFunc     func(*http.Request, handlerInput[T, R, A, L]) (L, error)
	subsections      map[string]Subsection
}

type Subsection struct {
	dbQuery     	DbQueryIntMany
	createSubFn 	func(*Config, *http.Request, int32) (SimpleResource, error)
	relationsFn		func(*Config, *http.Request, []int32) (map[int32]map[Relation][]int32, error)
	relations		map[int32]map[Relation][]int32
}


type Relation string

const (
	RelationAreas 		Relation = "areas"
	RelationTreasures	Relation = "treasures"
	RelationShops		Relation = "shops"
	RelationMonsters	Relation = "monsters"
)