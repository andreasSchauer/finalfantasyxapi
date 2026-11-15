package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/sections/sectiontypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetTreasure(t *testing.T) {
	t.Parallel()
	tests := []Treasure{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/treasures/a",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid id 'a'",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/treasures/343",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "treasure with provided id '343' doesn't exist. max id: 342",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/treasures/3/a",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "endpoint /treasures doesn't have any subsections",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/treasures/a/3",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "wrong format. usage: '/api/treasures', '/api/treasures/{id}'",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/treasures/2/2",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid subsection '2'. subsection can't be an integer. use /api/treasures/sections for valid subsections",
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/treasures/199",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths:     map[string]int{},
			},
			IDOnly:          newExpIDOnly(199),
			area:            158,
			availability:    database.AvailabilityTypePreStory,
			isAnimaTreasure: false,
			treasureType:    database.TreasureTypeChest,
			lootType:        database.LootTypeItem,
			gilAmount:       nil,
			items: []testResAmount[apiresource.Typed]{
				newTestResAmount[apiresource.Typed](82, 1),
			},
			equipment: nil,
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/treasures/13",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths:     map[string]int{},
			},
			IDOnly:          newExpIDOnly(13),
			area:            15,
			availability:    database.AvailabilityTypePreStory,
			isAnimaTreasure: false,
			treasureType:    database.TreasureTypeObject,
			lootType:        database.LootTypeItem,
			gilAmount:       nil,
			items: []testResAmount[apiresource.Typed]{
				newTestResAmount[apiresource.Typed](147, 1),
			},
			equipment: nil,
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/treasures/62",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths:     map[string]int{},
			},
			IDOnly:          newExpIDOnly(62),
			area:            60,
			availability:    database.AvailabilityTypeAlways,
			isAnimaTreasure: true,
			treasureType:    database.TreasureTypeChest,
			lootType:        database.LootTypeEquipment,
			gilAmount:       nil,
			items:           nil,
			equipment: &testFoundEquipment{
				equipmentName:    877,
				abilities:        []int32{55, 64, 58},
				emptySlotsAmount: 0,
			},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/treasures/284",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths:     map[string]int{},
			},
			IDOnly:          newExpIDOnly(284),
			area:            212,
			availability:    database.AvailabilityTypeAlways,
			isAnimaTreasure: false,
			treasureType:    database.TreasureTypeChest,
			lootType:        database.LootTypeGil,
			gilAmount:       h.GetInt32Ptr(20000),
			items:           nil,
			equipment:       nil,
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/treasures/45",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths:     map[string]int{},
			},
			IDOnly:          newExpIDOnly(45),
			area:            42,
			availability:    database.AvailabilityTypePost,
			isAnimaTreasure: false,
			treasureType:    database.TreasureTypeGift,
			lootType:        database.LootTypeItem,
			gilAmount:       nil,
			items: []testResAmount[apiresource.Typed]{
				newTestResAmount[apiresource.Typed](97, 1),
			},
			equipment: nil,
		},
	}

	testSingleResources(t, tests, "GetTreasure", testSrv.HandleTreasures)
}

func TestRetrieveTreasures(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/treasures?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   342,
			Results: []int32{1, 28, 53, 84, 123, 127, 167, 191, 214, 286, 300, 342},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/treasures?location=4&loot_type=item&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   21,
			Results: []int32{15, 18, 20, 23, 35, 37, 41, 44},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/treasures?sublocation=43&treasure_type=chest", // omega ruins
				ExpectedStatus: http.StatusOK,
			},
			Count:   15,
			Results: []int32{328, 330, 335, 337, 339, 342},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/treasures?area=212", // gagazet - mountain trail
				ExpectedStatus: http.StatusOK,
			},
			Count:   6,
			Results: []int32{284, 285, 286, 287, 288, 289},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/treasures?anima=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   6,
			Results: []int32{39, 62, 150, 210, 253, 302},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/treasures?availability=pre-story&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   99,
			Results: []int32{1, 23, 91, 119, 139, 206, 239, 261, 263},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/treasures?availability=post&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   59,
			Results: []int32{5, 30, 43, 64, 81, 109, 153, 211, 308, 342},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/treasures?availability=always&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   174,
			Results: []int32{2, 54, 80, 108, 147, 179, 214, 265, 295, 303},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/treasures?availability=pre-story&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   99,
			Results: []int32{1, 55, 112, 206, 248, 267},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/treasures?availability=post-story",
				ExpectedStatus: http.StatusOK,
			},
			Count:   10,
			Results: []int32{317, 321, 324, 326},
		},
	}

	testIDList(t, tests, apilookups.E.Treasures.Endpoint, "RetrieveTreasures", testSrv.HandleTreasures, compareAPIResourceLists[apiresource.Unnamed])
}

func TestSubsectionTreasures(t *testing.T) {
	t.Parallel()
	tests := []ListIDs{
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/locations/6/treasures",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleLocations,
			},
			Count:          17,
			ParentResource: h.GetStrPtr("/locations/6"),
			Results:        []int32{50, 53, 55, 59, 62, 65, 66},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/sublocations/21/treasures",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleSublocations,
			},
			Count:          5,
			ParentResource: h.GetStrPtr("/sublocations/21"),
			Results:        []int32{167, 168, 169, 170, 171},
		},
		{
			TestGeneral: TestGeneral{
				RequestURL:     "/api/areas/223/treasures",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleAreas,
			},
			Count:          3,
			ParentResource: h.GetStrPtr("/areas/223"),
			Results:        []int32{298, 299, 300},
		},
	}

	testIDList(t, tests, apilookups.E.Treasures.Endpoint, "SubsectionTreasures", nil, compareSimpleResourceLists[apiresource.Unnamed, sectiontypes.TreasureSimple])
}
