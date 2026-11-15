package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type Song struct {
	TestGeneral
	Unique
	composer          *string
	arranger          *string
	durationInSeconds int32
	canLoop           bool
	backgroundMusic   []testBackgroundMusic
	cues              []testCue
	bossFights        []int32
	fmvs              []int32
}

func (e Song) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e Song) Compare(test Test, got apitypes.Song) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compare("composer", e.composer, got.Composer)
	test.compare("arranger", e.arranger, got.Arranger)
	test.compare("duration in seconds", e.durationInSeconds, got.DurationInSeconds)
	test.compare("can loop", e.canLoop, got.CanLoop)
	test.checkTestStructsInSlice("background music", e.backgroundMusic, got.BackgroundMusic)
	test.checkTestStructsInSlice("cues", e.cues, got.Cues)
	test.checkResIDsInSlice("boss fights", apilookups.E.MonsterFormations.Endpoint, e.bossFights, got.BossFights)
	test.checkResIDsInSlice("fmvs", apilookups.E.FMVs.Endpoint, e.fmvs, got.FMVs)
}

type testBackgroundMusic struct {
	index                  int
	replacesEncounterMusic bool
	areas                  []int32
}

func (e testBackgroundMusic) GetIndex() int {
	return e.index
}

func (e testBackgroundMusic) Compare(test Test, _ string, got apitypes.BackgroundMusic) {
	test.T.Helper()
	test.compare("bg music - replaces encounter music", e.replacesEncounterMusic, got.ReplacesEncounterMusic)
	test.checkResIDsInSlice("bg music - areas", apilookups.E.Areas.Endpoint, e.areas, got.Areas)
}

type testCue struct {
	index                  int
	triggerArea            *int32
	includedAreas          []int32
	replacesEncounterMusic bool
	replacesBGMusic        *string
}

func (e testCue) GetIndex() int {
	return e.index
}

func (e testCue) Compare(test Test, fieldName string, got apitypes.Cue) {
	test.T.Helper()
	test.compIDAPIResourcePtrs(fieldName+" - trigger area", apilookups.E.Areas.Endpoint, e.triggerArea, got.TriggerArea)
	test.checkResIDsInSlice(fieldName+" - included areas", apilookups.E.Areas.Endpoint, e.includedAreas, got.IncludedAreas)
	test.compare(fieldName+" - replaces encounter music", e.replacesEncounterMusic, got.ReplacesEncounterMusic)
	test.compare(fieldName+" - replaces bg music", e.replacesBGMusic, got.ReplacesBGMusic)
}
