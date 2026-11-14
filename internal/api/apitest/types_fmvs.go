package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type FMV struct {
	testGeneral
	Unique
	area int32
	song *int32
}

func (e FMV) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func (e FMV) Compare(test test, got apitypes.FMV) {
	test.t.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compIDAPIResource("area", apilookups.E.Areas.Endpoint, e.area, got.Area)
	test.compIDAPIResourcePtrs("song", apilookups.E.Songs.Endpoint, e.song, got.Song)
}
