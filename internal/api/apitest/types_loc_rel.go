package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type LocRel struct {
	characters []int32
	aeons      []int32
	shops      []int32
	treasures  []int32
	monsters   []int32
	formations []int32
	quests     []int32
	fmvs       []int32
	music      *testLocMusic
}

func (e LocRel) Compare(test test, got apitypes.LocRel) {
	test.t.Helper()
	test.compTestStructPtrs("music", e.music, got.Music)
	test.checkResIDsInSlice("characters", apilookups.E.Characters.Endpoint, e.characters, got.Characters)
	test.checkResIDsInSlice("aeons", apilookups.E.Aeons.Endpoint, e.aeons, got.Aeons)
	test.checkResIDsInSlice("shops", apilookups.E.Shops.Endpoint, e.shops, got.Shops)
	test.checkResIDsInSlice("treasures", apilookups.E.Treasures.Endpoint, e.treasures, got.Treasures)
	test.checkResIDsInSlice("monsters", apilookups.E.Monsters.Endpoint, e.monsters, got.Monsters)
	test.checkResIDsInSlice("formations", apilookups.E.MonsterFormations.Endpoint, e.formations, got.Formations)
	test.checkResIDsInSlice("quests", apilookups.E.Quests.Endpoint, e.quests, got.Quests)
	test.checkResIDsInSlice("fmvs", apilookups.E.FMVs.Endpoint, e.fmvs, got.FMVs)
}

type testLocMusic struct {
	bgMusic   []int32
	cuesMusic []int32
	fmvsMusic []int32
	bossMusic []int32
}

func (e testLocMusic) Compare(test test, _ string, got apitypes.LocBasedMusic) {
	test.t.Helper()
	songsEndpoint := apilookups.E.Songs.Endpoint

	test.checkResIDsInSlice("bg music", songsEndpoint, e.bgMusic, got.BackgroundMusic)
	test.checkResIDsInSlice("cues music", songsEndpoint, e.cuesMusic, got.Cues)
	test.checkResIDsInSlice("fmvs music", songsEndpoint, e.fmvsMusic, got.FMVs)
	test.checkResIDsInSlice("boss music", songsEndpoint, e.bossMusic, got.BossMusic)
}
