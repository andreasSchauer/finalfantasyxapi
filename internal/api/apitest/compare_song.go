package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

func compareSongs(test Test, exp Song, got apitypes.Song) {
	test.t.Helper()
	test.compareExpUnique(exp.Unique, got.ID, got.Name)
	test.compare("composer", exp.composer, got.Composer)
	test.compare("arranger", exp.arranger, got.Arranger)
	test.compare("duration in seconds", exp.durationInSeconds, got.DurationInSeconds)
	test.compare("can loop", exp.canLoop, got.CanLoop)
	test.checkTestStructsInSlice("background music", exp.backgroundMusic, got.BackgroundMusic, compareBackgroundMusic)
	test.checkTestStructsInSlice("cues", exp.cues, got.Cues, compareCues)
	test.checkResIDsInSlice("boss fights", apilookups.E.MonsterFormations.Endpoint, exp.bossFights, got.BossFights)
	test.checkResIDsInSlice("fmvs", apilookups.E.FMVs.Endpoint, exp.fmvs, got.FMVs)
}

func compareBackgroundMusic(test Test, _ string, exp TestBackgroundMusic, got apitypes.BackgroundMusic) {
	test.t.Helper()
	test.compare("bg music - replaces encounter music", exp.replacesEncounterMusic, got.ReplacesEncounterMusic)
	test.checkResIDsInSlice("bg music - areas", apilookups.E.Areas.Endpoint, exp.areas, got.Areas)
}

func compareCues(test Test, fieldName string, exp TestCue, got apitypes.Cue) {
	test.t.Helper()
	test.compIDAPIResourcePtrs(fieldName+" - trigger area", apilookups.E.Areas.Endpoint, exp.triggerArea, got.TriggerArea)
	test.checkResIDsInSlice(fieldName+" - included areas", apilookups.E.Areas.Endpoint, exp.includedAreas, got.IncludedAreas)
	test.compare(fieldName+" - replaces encounter music", exp.replacesEncounterMusic, got.ReplacesEncounterMusic)
	test.compare(fieldName+" - replaces bg music", exp.replacesBGMusic, got.ReplacesBGMusic)
}
