package main

import (
	"testing"
)

/*
fields that need to be explicitly ignored with dontCheck:
- standard types and pointers (through compare),
- direct apiResource references or pointers
- structs and pointers to structs that are borrowed from the result
- basically anything that isn't a slice or a map and doesn't have nil checks in the test function body


fields that can be explicitly ignored with dontCheck:
- any field that is referenced explicitly as part of dont check in the function body


fields that are implicitly ignored by leaving them blank:
- slices of api resource references
- resourceAmount map references
- slices, structs, pointers to structs of some kind that have nil checks in the function body
*/

type testStructIdx interface {
	GetIndex() int
}

type test struct {
	t          *testing.T
	cfg        *Config
	name       string
	expLengths map[string]int
	dontCheck  map[string]bool
}

type testGeneral struct {
	requestURL     string
	expectedStatus int
	expectedErr    string
	dontCheck      map[string]bool
	expLengths     map[string]int
}

type expNameVer struct {
	id      int32
	name    string
	version *int32
}

func compareExpNameVer(test test, exp expNameVer, gotID int32, gotName string, gotVer *int32) {
	test.t.Helper()
	compare(test, "id", exp.id, gotID)
	compare(test, "name", exp.name, gotName)
	compare(test, "version", exp.version, gotVer)
}

type expUnique struct {
	id   int32
	name string
}

func compareExpUnique(test test, exp expUnique, gotID int32, gotName string) {
	test.t.Helper()
	compare(test, "id", exp.id, gotID)
	compare(test, "name", exp.name, gotName)
}

type expIdOnly struct {
	id int32
}

func compareExpIdOnly(test test, exp expIdOnly, gotID int32) {
	test.t.Helper()
	compare(test, "id", exp.id, gotID)
}

type expListIDs struct {
	testGeneral
	count          int
	previous       *string
	next           *string
	parentResource *string
	results        []int32
}

func (l expListIDs) getListParams() ListParams {
	return ListParams{
		Count:    l.count,
		Previous: l.previous,
		Next:     l.next,
	}
}

type expListNames struct {
	testGeneral
	count    int
	previous *string
	next     *string
	results  []string
}

func (l expListNames) getListParams() ListParams {
	return ListParams{
		Count:    l.count,
		Previous: l.previous,
		Next:     l.next,
	}
}

func compareListParams(test test, exp, got ListParams) {
	compare(test, "count", exp.Count, got.Count)
	compPageURL(test, "previous", exp.Previous, got.Previous)
	compPageURL(test, "next", exp.Next, got.Next)
}

type expLocRel struct {
	characters 	[]int32
	aeons      	[]int32
	shops      	[]int32
	treasures  	[]int32
	monsters   	[]int32
	formations 	[]int32
	sidequests 	[]int32
	fmvs       	[]int32
	music		*testLocMusic
}

func compareLocRel(test test, exp expLocRel, got LocRel) {
	compTestStructPtrs(test, "music", exp.music, got.Music, compareLocMusic)
	checkResIDsInSlice(test, "characters", test.cfg.e.characters.endpoint, exp.characters, got.Characters)
	checkResIDsInSlice(test, "aeons", test.cfg.e.aeons.endpoint, exp.aeons, got.Aeons)
	checkResIDsInSlice(test, "shops", test.cfg.e.shops.endpoint, exp.shops, got.Shops)
	checkResIDsInSlice(test, "treasures", test.cfg.e.treasures.endpoint, exp.treasures, got.Treasures)
	checkResIDsInSlice(test, "monsters", test.cfg.e.monsters.endpoint, exp.monsters, got.Monsters)
	checkResIDsInSlice(test, "formations", test.cfg.e.monsterFormations.endpoint, exp.formations, got.Formations)
	checkResIDsInSlice(test, "sidequests", test.cfg.e.sidequests.endpoint, exp.sidequests, got.Sidequests)
	checkResIDsInSlice(test, "fmvs", test.cfg.e.fmvs.endpoint, exp.fmvs, got.FMVs)
}

type testLocMusic struct {
	bgMusic    []int32
	cuesMusic  []int32
	fmvsMusic  []int32
	bossMusic  []int32
}

func compareLocMusic(test test, exp testLocMusic, got LocBasedMusic) {
	songsEndpoint := test.cfg.e.songs.endpoint

	checkResIDsInSlice(test, "bg music", songsEndpoint, exp.bgMusic, got.BackgroundMusic)
	checkResIDsInSlice(test, "cues music", songsEndpoint, exp.cuesMusic, got.Cues)
	checkResIDsInSlice(test, "fmvs music", songsEndpoint, exp.fmvsMusic, got.FMVs)
	checkResIDsInSlice(test, "boss music", songsEndpoint, exp.bossMusic, got.BossMusic)
}