package api

import "net/http"

type Subsection struct {
	dbQuery     	DbQueryIntMany
	createSubFn 	func(*Config, *http.Request, int32, Subsection) (SimpleResource, error)
	relationsFn		func(*Config, *http.Request, []int32) (map[int32]map[Relation][]int32, error)
	relations		map[int32]map[Relation][]int32
}


type Relation string

const (
	RelationAreas 		Relation = "areas"
	RelationTreasures	Relation = "treasures"
	RelationShops		Relation = "shops"
	RelationMonsters	Relation = "monsters"
	RelationRanks		Relation = "ranks"
)