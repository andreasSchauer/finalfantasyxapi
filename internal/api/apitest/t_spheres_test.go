package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

func TestGetSphere(t *testing.T) {
	t.Parallel()
	tests := []Sphere{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/spheres/31",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "sphere with provided id '31' doesn't exist. max id: 30",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/spheres/4?rel_availability=pre-story&rel_repeatable=true",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"monsters":         10,
					"treasures":        0,
					"shops":            0,
					"quests":           0,
					"blitzball prizes": 0,
				},
			},
			Unique:      newExpUnique(4, "ability sphere"),
			Item:        73,
			CreatedNode: nil,
			Monsters: []TestMonItemAmts{
				{
					Index:      3,
					Monster:    96,
					DropCommon: 1,
					DropRare:   1,
				},
				{
					Index:      9,
					Monster:    123,
					DropCommon: 1,
					DropRare:   1,
				},
			},
			Treasures:       nil,
			Shops:           []int32{},
			Quests:          nil,
			BlitzballPrizes: nil,
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/spheres/16",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"monsters":         3,
					"treasures":        4,
					"shops":            0,
					"quests":           1,
					"blitzball prizes": 0,
				},
			},
			Unique: newExpUnique(16, "hp sphere"),
			Item:   85,
			CreatedNode: &apitypes.CreatedNode{
				Node:  "hp",
				Value: 300,
			},
			Monsters: []TestMonItemAmts{
				{
					Index:    0,
					Monster:  28,
					DropRare: 1,
				},
				{
					Index:      1,
					Monster:    196,
					DropCommon: 1,
					DropRare:   1,
				},
				{
					Index:      2,
					Monster:    282,
					DropCommon: 1,
				},
			},
			Treasures: map[int32]int32{
				78:  1,
				253: 1,
				288: 1,
				313: 1,
			},
			Shops: []int32{},
			Quests: map[int32]int32{
				92: 3,
			},
			BlitzballPrizes: nil,
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/spheres/16?rel_availability=post-game&rel_repeatable=false",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"monsters":         1,
					"treasures":        3,
					"shops":            0,
					"quests":           1,
					"blitzball prizes": 0,
				},
			},
			Unique: newExpUnique(16, "hp sphere"),
			Item:   85,
			CreatedNode: &apitypes.CreatedNode{
				Node:  "hp",
				Value: 300,
			},
			Monsters: []TestMonItemAmts{
				{
					Index:    0,
					Monster:  28,
					DropRare: 1,
				},
			},
			Treasures: map[int32]int32{
				78:  1,
				288: 1,
				313: 1,
			},
			Shops: []int32{},
			Quests: map[int32]int32{
				92: 3,
			},
			BlitzballPrizes: nil,
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/spheres/26?rel_availability=post&rel_repeatable=true",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"monsters":         0,
					"treasures":        0,
					"shops":            1,
					"quests":           0,
					"blitzball prizes": 0,
				},
			},
			Unique:          newExpUnique(26, "clear sphere"),
			Item:            95,
			CreatedNode:     nil,
			Monsters:        []TestMonItemAmts{},
			Treasures:       nil,
			Shops:           []int32{33},
			Quests:          nil,
			BlitzballPrizes: nil,
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/spheres/26",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"monsters":         0,
					"treasures":        0,
					"shops":            1,
					"quests":           0,
					"blitzball prizes": 0,
				},
			},
			Unique:          newExpUnique(26, "clear sphere"),
			Item:            95,
			CreatedNode:     nil,
			Monsters:        []TestMonItemAmts{},
			Treasures:       nil,
			Shops:           []int32{33},
			Quests:          nil,
			BlitzballPrizes: nil,
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/spheres/2?rel_availability=always&rel_repeatable=false",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"monsters":         1,
					"treasures":        1,
					"shops":            0,
					"quests":           0,
					"blitzball prizes": 0,
				},
			},
			Unique:      newExpUnique(2, "mana sphere"),
			Item:        71,
			CreatedNode: nil,
			Monsters: []TestMonItemAmts{
				{
					Index:      0,
					Monster:    107,
					DropCommon: 1,
					DropRare:   1,
				},
			},
			Treasures: map[int32]int32{
				53: 2,
			},
			Shops:           []int32{},
			Quests:          nil,
			BlitzballPrizes: nil,
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/spheres/2?rel_availability=post&rel_repeatable=true",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"monsters":         10,
					"treasures":        0,
					"shops":            0,
					"quests":           1,
					"blitzball prizes": 0,
				},
			},
			Unique:      newExpUnique(2, "mana sphere"),
			Item:        71,
			CreatedNode: nil,
			Monsters: []TestMonItemAmts{
				{
					Index:      1,
					Monster:    205,
					DropCommon: 1,
					DropRare:   1,
				},
				{
					Index:      5,
					Monster:    240,
					DropCommon: 2,
					DropRare:   3,
				},
			},
			Treasures: nil,
			Shops:     []int32{},
			Quests: map[int32]int32{
				58: 10,
			},
			BlitzballPrizes: nil,
		},
	}

	testSingleResources(t, tests, "GetSphere", testSrv.HandleSpheres)
}

func TestRetrieveSpheres(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/spheres?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   30,
			Results: []int32{1, 13, 15, 19, 23, 27, 30},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/spheres?color=yellow&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   6,
			Results: []int32{6, 7, 8, 9, 10, 11},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/spheres?availability=always&location=12&repeatable=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   3,
			Results: []int32{1, 2, 3},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/spheres?availability=post&repeatable=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   14,
			Results: []int32{7, 10, 16, 21, 24, 30},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/spheres?availability=post-game&repeatable=true&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   30,
			Results: []int32{1, 5, 10, 20, 25, 29, 30},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/spheres?availability=always&repeatable=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   16,
			Results: []int32{1, 5, 11, 15, 27, 29},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/spheres?availability=always&repeatable=false",
				ExpectedStatus: http.StatusOK,
			},
			Count:   6,
			Results: []int32{16, 17, 18, 20, 21, 25},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/spheres?availability=always&repeatable=false&methods=treasure",
				ExpectedStatus: http.StatusOK,
			},
			Count:   15,
			Results: []int32{2, 4, 12, 15, 21, 25, 29},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/spheres?availability=always&repeatable=false&methods=monster",
				ExpectedStatus: http.StatusOK,
			},
			Count:   2,
			Results: []int32{16, 17},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/spheres?availability=always&repeatable=false&methods=monster,treasure",
				ExpectedStatus: http.StatusOK,
			},
			Count:   6,
			Results: []int32{16, 17, 20, 21, 25, 29},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/spheres?availability=story&sublocation=28", // macalania
				ExpectedStatus: http.StatusOK,
			},
			Count:   3,
			Results: []int32{4, 7, 10},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/spheres?availability=story&repeatable=true&sublocation=28", // macalania
				ExpectedStatus: http.StatusOK,
			},
			Count:   0,
			Results: []int32{},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/spheres?availability=post&repeatable=true&area=203", // calm lands - arena
				ExpectedStatus: http.StatusOK,
			},
			Count:   17,
			Results: []int32{13, 15, 17, 19, 22, 25, 26, 30},
		},
	}

	testIDList(t, tests, apilookups.E.Spheres.Endpoint, "RetrieveSpheres", testSrv.HandleSpheres, compareAPIResourceLists[apiresource.Named])
}
