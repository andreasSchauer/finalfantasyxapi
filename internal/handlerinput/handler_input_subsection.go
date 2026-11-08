package handlerinput

import (
	"net/http"

	"github.com/andreasSchauer/finalfantasyxapi/internal/apiresource"
)

// might get moved to another package

type Subsection struct {
	DBQuery     DBQueryIntMany
	CreateSubFn func(*http.Request, int32, Subsection) (apiresource.SimpleResource, error)
	RelationsFn func(*http.Request, []int32) (map[int32]map[Relation][]int32, error)
	Relations   map[int32]map[Relation][]int32
}

type Relation string

const (
	RelationAreas     Relation = "areas"
	RelationTreasures Relation = "treasures"
	RelationShops     Relation = "shops"
	RelationMonsters  Relation = "monsters"
)
