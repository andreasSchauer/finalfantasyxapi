package apitest

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
	BackgroundMusic   []testBackgroundMusic
	Cues              []testCue
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
	test.checkResIDsInSlice("boss fights", apilookups.E.MonsterFormations.Endpoint, e.BossFights, got.BossFights)
	test.checkResIDsInSlice("fmvs", apilookups.E.FMVs.Endpoint, e.Fmvs, got.FMVs)
}

type testBackgroundMusic struct {
	Index                  int
	ReplacesEncounterMusic bool
	Areas                  []int32
}

func (e testBackgroundMusic) GetIndex() int {
	return e.Index
}

func (e testBackgroundMusic) Compare(test Test, _ string, got apitypes.BackgroundMusic) {
	test.T.Helper()
	test.compare("bg music - replaces encounter music", e.ReplacesEncounterMusic, got.ReplacesEncounterMusic)
	test.checkResIDsInSlice("bg music - areas", apilookups.E.Areas.Endpoint, e.Areas, got.Areas)
}

type testCue struct {
	Index                  int
	TriggerArea            *int32
	IncludedAreas          []int32
	ReplacesEncounterMusic bool
	ReplacesBGMusic        *string
}

func (e testCue) GetIndex() int {
	return e.Index
}

func (e testCue) Compare(test Test, fieldName string, got apitypes.Cue) {
	test.T.Helper()
	test.compIDAPIResourcePtrs(fieldName+" - trigger area", apilookups.E.Areas.Endpoint, e.TriggerArea, got.TriggerArea)
	test.checkResIDsInSlice(fieldName+" - included areas", apilookups.E.Areas.Endpoint, e.IncludedAreas, got.IncludedAreas)
	test.compare(fieldName+" - replaces encounter music", e.ReplacesEncounterMusic, got.ReplacesEncounterMusic)
	test.compare(fieldName+" - replaces bg music", e.ReplacesBGMusic, got.ReplacesBGMusic)
}
