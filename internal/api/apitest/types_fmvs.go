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

func compareFMVs(test test, exp FMV, got apitypes.FMV) {
	test.t.Helper()
	test.compareExpUnique(exp.Unique, got.ID, got.Name)
	test.compIDAPIResource("area", apilookups.E.Areas.Endpoint, exp.area, got.Area)
	test.compIDAPIResourcePtrs("song", apilookups.E.Songs.Endpoint, exp.song, got.Song)
}
