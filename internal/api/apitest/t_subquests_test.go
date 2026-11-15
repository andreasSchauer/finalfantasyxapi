package apitest

import (
	"net/http"
	"testing"

	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apiresource"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apitesttypes"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/sections/sectiontypes"
	h "github.com/andreasSchauer/finalfantasyxapi/internal/general/helpers"
)

func TestGetSubquest(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.Subquest{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/subquests/89",
				ExpectedStatus: http.StatusNotFound,
				ExpectedErr:    "subquest with provided id '89' doesn't exist. max id: 88",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/subquests/10/a",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "endpoint /subquests doesn't have any subsections",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/subquests/a",
				ExpectedStatus: http.StatusBadRequest,
				ExpectedErr:    "invalid id 'a'",
			},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/subquests/5",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"completions": 1,
				},
			},
			Unique:          apitesttypes.NewExpUnique(5, "jormungand"),
			UntypedQuest:    15,
			ParentSidequest: 1,
			Completion: apitesttypes.TestQuestCompletion{
				Areas:  []int32{203},
				Reward: apitesttypes.NewTestResAmount[apiresource.Typed](50, 99),
			},
			ArenaCreation: h.GetInt32Ptr(5),
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/subquests/39",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"completions": 2,
				},
			},
			Unique:          apitesttypes.NewExpUnique(39, "ifrit - first win"),
			UntypedQuest:    49,
			ParentSidequest: 2,
			Completion: apitesttypes.TestQuestCompletion{
				Areas:  []int32{207},
				Reward: apitesttypes.NewTestResAmount[apiresource.Typed](3, 30),
			},
			ArenaCreation: nil,
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/subquests/76",
				ExpectedStatus: http.StatusOK,
				DontCheck:      map[string]bool{},
				ExpLengths: map[string]int{
					"completions": 2,
				},
			},
			Unique:          apitesttypes.NewExpUnique(76, "first hunt after acquiring the airship"),
			UntypedQuest:    86,
			ParentSidequest: 6,
			Completion: apitesttypes.TestQuestCompletion{
				Areas:  []int32{142, 143},
				Reward: apitesttypes.NewTestResAmount[apiresource.Typed](98, 1),
			},
			ArenaCreation: nil,
		},
	}

	testSingleResources(t, tests, "GetSubquest", testSrv.HandleSubquests)
}

func TestRetrieveSubquests(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ListIDs{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/subquests?limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   88,
			Results: []int32{1, 5, 17, 28, 30, 34, 44, 51, 56, 66, 72, 79, 88},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/subquests?availability=post&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   39,
			Results: []int32{1, 8, 12, 23, 33, 52, 68, 71, 76},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/subquests?repeatable=true&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   12,
			Results: []int32{38, 42, 44, 50, 52, 60, 64, 66},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/subquests?availability=always&repeatable=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   10,
			Results: []int32{38, 40, 50, 60, 64, 66},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/subquests?availability=post&repeatable=false&limit=max",
				ExpectedStatus: http.StatusOK,
			},
			Count:   37,
			Results: []int32{1, 14, 24, 35, 70, 77},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/subquests?availability=pre-story",
				ExpectedStatus: http.StatusOK,
			},
			Count:   1,
			Results: []int32{87},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/subquests?availability=post&repeatable=true",
				ExpectedStatus: http.StatusOK,
			},
			Count:   2,
			Results: []int32{48, 52},
		},
	}

	testIDList(t, tests, apilookups.E.Subquests.Endpoint, "RetrieveSubquests", testSrv.HandleSubquests, compareAPIResourceLists[apiresource.Named])
}

func TestSubsectionSubquests(t *testing.T) {
	t.Parallel()
	tests := []apitesttypes.ListIDs{
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/sidequests/1/subquests",
				ExpectedStatus: http.StatusOK,
			},
			Count:          36,
			ParentResource: h.GetStrPtr("/sidequests/1"),
			Results:        []int32{1, 8, 12, 14, 18, 23, 31, 36},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/sidequests/2/subquests",
				ExpectedStatus: http.StatusOK,
			},
			Count:          16,
			ParentResource: h.GetStrPtr("/sidequests/2"),
			Results:        []int32{37, 40, 41, 43, 44, 46, 49, 52},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/sidequests/4/subquests",
				ExpectedStatus: http.StatusOK,
			},
			Count:          9,
			ParentResource: h.GetStrPtr("/sidequests/4"),
			Results:        []int32{59, 62, 66, 67},
		},
		{
			TestGeneral: apitesttypes.TestGeneral{
				RequestURL:     "/api/sidequests/9/subquests",
				ExpectedStatus: http.StatusOK,
			},
			Count:          0,
			ParentResource: h.GetStrPtr("/sidequests/9"),
			Results:        []int32{},
		},
	}

	testIDList(t, tests, apilookups.E.Subquests.Endpoint, "SubsectionSubquests", testSrv.HandleSidequests, compareSimpleResourceLists[apiresource.Named, sectiontypes.SubquestSimple])
}
