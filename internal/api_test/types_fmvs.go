package apiTest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api"

type expFMV struct {
	testGeneral
	expUnique
	area int32
	song *int32
}

func (e expFMV) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareFMVs(test test, exp expFMV, got api.FMV) {
	test.t.Helper()
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIdApiResource(test, "area", test.cfg.E.Areas.Endpoint, exp.area, got.Area)
	compIdApiResourcePtrs(test, "song", test.cfg.E.Songs.Endpoint, exp.song, got.Song)
}
