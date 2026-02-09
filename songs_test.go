package main

import (
	//"net/http"
	"testing"

	//h "github.com/andreasSchauer/finalfantasyxapi/internal/helpers"
)

type expSong struct {
	testGeneral
	expUnique
	composer			*string
	arranger			*string
	durationInSeconds	int32
	canLoop				bool
	backgroundMusic		[]testBackgroundMusic
	cues				[]testCue
	bossFights			[]int32
	fmvs				[]int32
}

func (e expSong) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareSongs(test test, exp expSong, got Song) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compare(test, "composer", exp.composer, got.Composer)
	compare(test, "arranger", exp.arranger, got.Arranger)
	compare(test, "duration in seconds", exp.durationInSeconds, got.DurationInSeconds)
	compare(test, "can loop", exp.canLoop, got.CanLoop)
	compTestStructSlices(test, "background music", exp.backgroundMusic, got.BackgroundMusic, compareBackgroundMusic)
	compTestStructSlices(test, "cues", exp.cues, got.Cues, compareCues)
	compareResListTest(test, rltIDs("boss fights", test.cfg.e.monsterFormations.endpoint, exp.bossFights, got.BossFights))
	compareResListTest(test, rltIDs("fmvs", test.cfg.e.fmvs.endpoint, exp.fmvs, got.FMVs))
}

type testBackgroundMusic struct {
	index					int
	replacesEncounterMusic	bool
	areas					[]int32
}

func (t testBackgroundMusic) GetIndex() int {
	return t.index
}

func compareBackgroundMusic(test test, exp testBackgroundMusic, got BackgroundMusic) {
	compare(test, "bg music - replaces encounter music", exp.replacesEncounterMusic, got.ReplacesEncounterMusic)
	compareResListTest(test, rltIDs("bg music - areas", test.cfg.e.areas.endpoint, exp.areas, got.Areas))
}

type testCue struct {
	index					int
	triggerArea				*int32
	includedAreas			[]int32
	replacesEncounterMusic	bool
	replacesBGMusic			*string
}

func (t testCue) GetIndex() int {
	return t.index
}

func compareCues(test test, exp testCue, got Cue) {
	compIdApiResourcePtrs(test, "cue - trigger area", test.cfg.e.areas.endpoint, exp.triggerArea, got.TriggerArea)
	compareResListTest(test, rltIDs("cue - included areas", test.cfg.e.areas.endpoint, exp.includedAreas, got.IncludedAreas))
	compare(test, "cue - replaces encounter music", exp.replacesEncounterMusic, got.ReplacesEncounterMusic)
	compare(test, "cue - replaces bg music", exp.replacesBGMusic, got.ReplacesBGMusic)
}

func TestGetSong(t *testing.T) {
	tests := []expSong{
		
	}

	testSingleResources(t, tests, "GetSongs", testCfg.HandleSongs, compareSongs)
}

func TestRetrieveSongs(t *testing.T) {
	tests := []expListIDs{
		
	}

	testIdList(t, tests, testCfg.e.songs.endpoint, "RetrieveSongs", testCfg.HandleSongs, compareAPIResourceLists[NamedApiResourceList])
}