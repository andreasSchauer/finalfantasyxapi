package apitesttypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type Song struct {
	TestGeneral
	Unique
	Composer          *string
	Arranger          *string
	DurationInSeconds int32
	CanLoop           bool
	BackgroundMusic   []TestBackgroundMusic
	Cues              []TestCue
	BossFights        []int32
	Fmvs              []int32
}

func (e Song) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e Song) Compare(test Test, got apitypes.Song) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compare("composer", e.Composer, got.Composer)
	test.compare("arranger", e.Arranger, got.Arranger)
	test.compare("duration in seconds", e.DurationInSeconds, got.DurationInSeconds)
	test.compare("can loop", e.CanLoop, got.CanLoop)
	test.checkTestStructsInSlice("background music", e.BackgroundMusic, got.BackgroundMusic)
	test.checkTestStructsInSlice("cues", e.Cues, got.Cues)
	test.CheckResIDsInSlice("boss fights", apilookups.E.MonsterFormations.Endpoint, e.BossFights, got.BossFights)
	test.CheckResIDsInSlice("fmvs", apilookups.E.FMVs.Endpoint, e.Fmvs, got.FMVs)
}

type TestBackgroundMusic struct {
	Index                  int
	ReplacesEncounterMusic bool
	Areas                  []int32
}

func (e TestBackgroundMusic) GetIndex() int {
	return e.Index
}

func (e TestBackgroundMusic) Compare(test Test, _ string, got apitypes.BackgroundMusic) {
	test.T.Helper()
	test.compare("bg music - replaces encounter music", e.ReplacesEncounterMusic, got.ReplacesEncounterMusic)
	test.CheckResIDsInSlice("bg music - areas", apilookups.E.Areas.Endpoint, e.Areas, got.Areas)
}

type TestCue struct {
	Index                  int
	TriggerArea            *int32
	IncludedAreas          []int32
	ReplacesEncounterMusic bool
	ReplacesBGMusic        *string
}

func (e TestCue) GetIndex() int {
	return e.Index
}

func (e TestCue) Compare(test Test, fn string, got apitypes.Cue) {
	test.T.Helper()
	test.compIDAPIResourcePtrs(childFn(fn, "trigger area"), apilookups.E.Areas.Endpoint, e.TriggerArea, got.TriggerArea)
	test.CheckResIDsInSlice(childFn(fn, "included areas"), apilookups.E.Areas.Endpoint, e.IncludedAreas, got.IncludedAreas)
	test.compare(childFn(fn, "replaces encounter music"), e.ReplacesEncounterMusic, got.ReplacesEncounterMusic)
	test.compare(childFn(fn, "replaces bg music"), e.ReplacesBGMusic, got.ReplacesBGMusic)
}
