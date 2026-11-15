package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
)

func TestGetCharacterClass(t *testing.T) {
	t.Parallel()
	tests := []CharacterClass{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/character-classes/23",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "character class with provided id '23' doesn't exist. max id: 22",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/character-classes/1",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"members":              7,
					"default abilities":    5,
					"learnable abilities":  85,
					"default overdrives":   0,
					"learnable overdrives": 0,
					"submenus":             11,
				},
			},
			Unique: Unique{
				ID:   1,
				Name: "characters",
			},
			category:            "group",
			members:             []int32{1, 2, 3, 4, 5, 6, 7},
			defaultAbilities:    []int32{374, 376, 377, 378, 379},
			learnableAbilities:  []int32{1, 12, 27, 29, 35, 51, 69, 87},
			defaultOverdrives:   []int32{},
			learnableOverdrives: []int32{},
			submenus:            []int32{2, 3, 4, 5, 6, 7, 8, 9, 11, 13, 15},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/character-classes/3",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"members":              6,
					"default abilities":    3,
					"learnable abilities":  67,
					"default overdrives":   0,
					"learnable overdrives": 0,
					"submenus":             4,
				},
			},
			Unique: Unique{
				ID:   3,
				Name: "standard-aeons",
			},
			category:            "group",
			members:             []int32{9, 10, 11, 12, 13, 14},
			defaultAbilities:    []int32{375, 380, 381},
			learnableAbilities:  []int32{3, 14, 19, 29, 50, 57, 64, 85},
			defaultOverdrives:   []int32{},
			learnableOverdrives: []int32{},
			submenus:            []int32{2, 3, 4, 5},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/character-classes/4",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"members":              3,
					"default abilities":    2,
					"learnable abilities":  0,
					"default overdrives":   1,
					"learnable overdrives": 0,
					"submenus":             0,
				},
			},
			Unique: Unique{
				ID:   4,
				Name: "magus-sisters",
			},
			category:            "group",
			members:             []int32{16, 17, 18},
			defaultAbilities:    []int32{102, 375},
			learnableAbilities:  []int32{},
			defaultOverdrives:   []int32{124},
			learnableOverdrives: []int32{},
			submenus:            []int32{},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/character-classes/6",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"members":              1,
					"default abilities":    9,
					"learnable abilities":  7,
					"default overdrives":   1,
					"learnable overdrives": 7,
					"submenus":             1,
				},
			},
			Unique: Unique{
				ID:   6,
				Name: "yuna",
			},
			category:            "character",
			members:             []int32{2},
			defaultAbilities:    []int32{45, 53, 383, 385, 386, 387, 388, 389, 390},
			learnableAbilities:  []int32{391, 392, 393, 394, 395, 396, 397},
			defaultOverdrives:   []int32{5},
			learnableOverdrives: []int32{6, 7, 8, 9, 10, 11, 12},
			submenus:            []int32{1},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/character-classes/8",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"members":              1,
					"default abilities":    10,
					"learnable abilities":  0,
					"default overdrives":   4,
					"learnable overdrives": 15,
					"submenus":             1,
				},
			},
			Unique: Unique{
				ID:   8,
				Name: "lulu",
			},
			category:            "character",
			members:             []int32{4},
			defaultAbilities:    []int32{69, 70, 71, 72, 383, 384, 385, 387, 388, 389},
			learnableAbilities:  []int32{},
			defaultOverdrives:   []int32{17, 18, 19, 20},
			learnableOverdrives: []int32{21, 26, 30, 32, 34, 35},
			submenus:            []int32{14},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/character-classes/19",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"members":              1,
					"default abilities":    5,
					"learnable abilities":  0,
					"default overdrives":   0,
					"learnable overdrives": 0,
					"submenus":             1,
				},
			},
			Unique: Unique{
				ID:   19,
				Name: "yojimbo",
			},
			category:            "aeon",
			members:             []int32{15},
			defaultAbilities:    []int32{94, 95, 96, 97, 98},
			learnableAbilities:  []int32{},
			defaultOverdrives:   []int32{},
			learnableOverdrives: []int32{},
			submenus:            []int32{12},
		},
	}

	testSingleResources(t, tests, "GetCharacterClass", testSrv.HandleCharacterClasses)
}

func TestRetrieveCharacterClasses(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/character-classes?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   22,
			Results: []int32{1, 5, 8, 12, 13, 17, 22},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/character-classes?category=1",
				ExpectedStatus: http.StatusOK,
			},
			Count:   4,
			Results: []int32{1, 2, 3, 4},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/character-classes?category=character",
				ExpectedStatus: http.StatusOK,
			},
			Count:   8,
			Results: []int32{5, 7, 8, 9, 10, 12},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/character-classes?category=aeon",
				ExpectedStatus: http.StatusOK,
			},
			Count:   10,
			Results: []int32{13, 15, 18, 20, 21, 22},
		},
	}

	testIDList(t, tests, apilookups.E.CharacterClasses.Endpoint, "RetrieveCharacterClasses", testSrv.HandleCharacterClasses, compareAPIResourceLists[apiresource.Named])
}
