package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/sections/sectiontypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/tests/apitesttypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetTreasure(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.Treasure{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/treasures/a",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid id 'a'",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/treasures/343",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "treasure with provided id '343' doesn't exist. max id: 342",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/treasures/3/a",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "endpoint /treasures doesn't have any subsections",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/treasures/a/3",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "wrong format. usage: '/api/treasures', '/api/treasures/{id}'",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/treasures/2/2",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid subsection '2'. subsection can't be an integer. use /api/treasures/sections for valid subsections",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/treasures/199",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths:     map[string]int{},
			},
			IDOnly:          apitesttypes.NewExpIDOnly(199),
			Area:            158,
			Availability:    database.AvailabilityTypePreStory,
			IsAnimaTreasure: false,
			TreasureType:    database.TreasureTypeChest,
			LootType:        database.LootTypeItem,
			GilAmount:       nil,
			Items: []apitesttypes.TestResAmount[apiresource.Typed]{
				apitesttypes.NewTestResAmount[apiresource.Typed](82, 1),
			},
			Equipment: nil,
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/treasures/13",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths:     map[string]int{},
			},
			IDOnly:          apitesttypes.NewExpIDOnly(13),
			Area:            15,
			Availability:    database.AvailabilityTypePreStory,
			IsAnimaTreasure: false,
			TreasureType:    database.TreasureTypeObject,
			LootType:        database.LootTypeItem,
			GilAmount:       nil,
			Items: []apitesttypes.TestResAmount[apiresource.Typed]{
				apitesttypes.NewTestResAmount[apiresource.Typed](147, 1),
			},
			Equipment: nil,
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/treasures/62",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths:     map[string]int{},
			},
			IDOnly:          apitesttypes.NewExpIDOnly(62),
			Area:            60,
			Availability:    database.AvailabilityTypeAlways,
			IsAnimaTreasure: true,
			TreasureType:    database.TreasureTypeChest,
			LootType:        database.LootTypeEquipment,
			GilAmount:       nil,
			Items:           nil,
			Equipment: &apitesttypes.TestFoundEquipment{
				EquipmentName:    877,
				Abilities:        []int32{55, 64, 58},
				EmptySlotsAmount: 0,
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/treasures/284",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths:     map[string]int{},
			},
			IDOnly:          apitesttypes.NewExpIDOnly(284),
			Area:            212,
			Availability:    database.AvailabilityTypeAlways,
			IsAnimaTreasure: false,
			TreasureType:    database.TreasureTypeChest,
			LootType:        database.LootTypeGil,
			GilAmount:       h.GetInt32Ptr(20000),
			Items:           nil,
			Equipment:       nil,
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/treasures/45",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths:     map[string]int{},
			},
			IDOnly:          apitesttypes.NewExpIDOnly(45),
			Area:            42,
			Availability:    database.AvailabilityTypePost,
			IsAnimaTreasure: false,
			TreasureType:    database.TreasureTypeGift,
			LootType:        database.LootTypeItem,
			GilAmount:       nil,
			Items: []apitesttypes.TestResAmount[apiresource.Typed]{
				apitesttypes.NewTestResAmount[apiresource.Typed](97, 1),
			},
			Equipment: nil,
		},
	}

	testSingleResources(t, tests, "GetTreasure", testSrv.HandleTreasures)
}

func TestRetrieveTreasures(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ListIDs{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/treasures?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   342,
			Results: []int32{1, 28, 53, 84, 123, 127, 167, 191, 214, 286, 300, 342},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/treasures?location=4&loot_type=item&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   21,
			Results: []int32{15, 18, 20, 23, 35, 37, 41, 44},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/treasures?sublocation=43&treasure_type=chest", // omega ruins
				ExpectedStatus: http.StatusOK,
			},
			Count:   15,
			Results: []int32{328, 330, 335, 337, 339, 342},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/treasures?area=212", // gagazet - mountain trail
				ExpectedStatus: http.StatusOK,
			},
			Count:   6,
			Results: []int32{284, 285, 286, 287, 288, 289},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/treasures?anima=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   6,
			Results: []int32{39, 62, 150, 210, 253, 302},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/treasures?availability=pre-story&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   99,
			Results: []int32{1, 23, 91, 119, 139, 206, 239, 261, 263},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/treasures?availability=post&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   59,
			Results: []int32{5, 30, 43, 64, 81, 109, 153, 211, 308, 342},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/treasures?availability=always&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   174,
			Results: []int32{2, 54, 80, 108, 147, 179, 214, 265, 295, 303},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/treasures?availability=pre-story&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   99,
			Results: []int32{1, 55, 112, 206, 248, 267},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
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
	tests := []apitesttypes.ListIDs{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/locations/6/treasures",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleLocations,
			},
			Count:          17,
			ParentResource: h.GetStrPtr("/locations/6"),
			Results:        []int32{50, 53, 55, 59, 62, 65, 66},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/sublocations/21/treasures",
				ExpectedStatus: http.StatusOK,
				Handler:        testSrv.HandleSublocations,
			},
			Count:          5,
			ParentResource: h.GetStrPtr("/sublocations/21"),
			Results:        []int32{167, 168, 169, 170, 171},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
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
