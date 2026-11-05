package api

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
)

type Subsection struct {
	dbQuery     DBQueryIntMany
	createSubFn func(*http.Request, int32, Subsection) (apiresource.SimpleResource, error)
	relationsFn func(*http.Request, []int32) (map[int32]map[Relation][]int32, error)
	relations   map[int32]map[Relation][]int32
}

type Relation string

const (
	RelationAreas     Relation = "areas"
	RelationTreasures Relation = "treasures"
	RelationShops     Relation = "shops"
	RelationMonsters  Relation = "monsters"
)
