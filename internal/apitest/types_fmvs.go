package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
)

type expFMV struct {
	testGeneral
	expUnique
	area int32
	song *int32
}

func (e expFMV) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareFMVs(test test, exp expFMV, got apitypes.FMV) {
	test.t.Helper()
	test.compareExpUnique(exp.expUnique, got.ID, got.Name)
	test.compIDAPIResource("area", api.E.Areas.Endpoint, exp.area, got.Area)
	test.compIDAPIResourcePtrs("song", api.E.Songs.Endpoint, exp.song, got.Song)
}
