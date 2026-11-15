package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type Sublocation struct {
	TestGeneral
	Unique
	ParentLocation        int32
	ConnectedSublocations []int32
	Areas                 []int32
	LocRel
}

func (e Sublocation) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e Sublocation) Compare(test Test, got apitypes.Sublocation) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compIDAPIResource("location", apilookups.E.Locations.Endpoint, e.ParentLocation, got.ParentLocation)
	test.compTestStructPtrs("music", e.Music, got.Music)
	test.checkResIDsInSlice("connected sublocations", apilookups.E.Sublocations.Endpoint, e.ConnectedSublocations, got.ConnectedSublocations)
	test.checkResIDsInSlice("areas", apilookups.E.Areas.Endpoint, e.Areas, got.Areas)
	e.LocRel.Compare(test, got.LocRel)
}
