package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type expSong struct {
	testGeneral
	expUnique
	composer          *string
	arranger          *string
	durationInSeconds int32
	canLoop           bool
	backgroundMusic   []testBackgroundMusic
	cues              []testCue
	bossFights        []int32
	fmvs              []int32
}

func (e expSong) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareSongs(test test, exp expSong, got apitypes.Song) {
	test.t.Helper()
	test.compareExpUnique(exp.expUnique, got.ID, got.Name)
	test.compare("composer", exp.composer, got.Composer)
	test.compare("arranger", exp.arranger, got.Arranger)
	test.compare("duration in seconds", exp.durationInSeconds, got.DurationInSeconds)
	test.compare("can loop", exp.canLoop, got.CanLoop)
	test.checkTestStructsInSlice("background music", exp.backgroundMusic, got.BackgroundMusic, compareBackgroundMusic)
	test.checkTestStructsInSlice("cues", exp.cues, got.Cues, compareCues)
	test.checkResIDsInSlice("boss fights", apilookups.E.MonsterFormations.Endpoint, exp.bossFights, got.BossFights)
	test.checkResIDsInSlice("fmvs", apilookups.E.FMVs.Endpoint, exp.fmvs, got.FMVs)
}

type testBackgroundMusic struct {
	index                  int
	replacesEncounterMusic bool
	areas                  []int32
}

func (t testBackgroundMusic) GetIndex() int {
	return t.index
}

func compareBackgroundMusic(test test, _ string, exp testBackgroundMusic, got apitypes.BackgroundMusic) {
	test.t.Helper()
	test.compare("bg music - replaces encounter music", exp.replacesEncounterMusic, got.ReplacesEncounterMusic)
	test.checkResIDsInSlice("bg music - areas", apilookups.E.Areas.Endpoint, exp.areas, got.Areas)
}

type testCue struct {
	index                  int
	triggerArea            *int32
	includedAreas          []int32
	replacesEncounterMusic bool
	replacesBGMusic        *string
}

func (t testCue) GetIndex() int {
	return t.index
}

func compareCues(test test, fieldName string, exp testCue, got apitypes.Cue) {
	test.t.Helper()
	test.compIDAPIResourcePtrs(fieldName+" - trigger area", apilookups.E.Areas.Endpoint, exp.triggerArea, got.TriggerArea)
	test.checkResIDsInSlice(fieldName+" - included areas", apilookups.E.Areas.Endpoint, exp.includedAreas, got.IncludedAreas)
	test.compare(fieldName+" - replaces encounter music", exp.replacesEncounterMusic, got.ReplacesEncounterMusic)
	test.compare(fieldName+" - replaces bg music", exp.replacesBGMusic, got.ReplacesBGMusic)
}
