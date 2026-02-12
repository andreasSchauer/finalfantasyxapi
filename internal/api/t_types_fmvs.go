package api

type expFMV struct {
	testGeneral
	expUnique
	area int32
	song *int32
}

func (e expFMV) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareFMVs(test test, exp expFMV, got FMV) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIdApiResource(test, "area", test.cfg.e.areas.endpoint, exp.area, got.Area)
	compIdApiResourcePtrs(test, "song", test.cfg.e.songs.endpoint, exp.song, got.Song)
}
