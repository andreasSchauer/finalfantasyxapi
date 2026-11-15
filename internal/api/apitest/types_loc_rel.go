package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type LocRel struct {
	Characters []int32
	Aeons      []int32
	Shops      []int32
	Treasures  []int32
	Monsters   []int32
	Formations []int32
	Quests     []int32
	Fmvs       []int32
	Music      *testLocMusic
}

func (e LocRel) Compare(test Test, got apitypes.LocRel) {
	test.T.Helper()
	test.compTestStructPtrs("music", e.Music, got.Music)
	test.checkResIDsInSlice("characters", apilookups.E.Characters.Endpoint, e.Characters, got.Characters)
	test.checkResIDsInSlice("aeons", apilookups.E.Aeons.Endpoint, e.Aeons, got.Aeons)
	test.checkResIDsInSlice("shops", apilookups.E.Shops.Endpoint, e.Shops, got.Shops)
	test.checkResIDsInSlice("treasures", apilookups.E.Treasures.Endpoint, e.Treasures, got.Treasures)
	test.checkResIDsInSlice("monsters", apilookups.E.Monsters.Endpoint, e.Monsters, got.Monsters)
	test.checkResIDsInSlice("formations", apilookups.E.MonsterFormations.Endpoint, e.Formations, got.Formations)
	test.checkResIDsInSlice("quests", apilookups.E.Quests.Endpoint, e.Quests, got.Quests)
	test.checkResIDsInSlice("fmvs", apilookups.E.FMVs.Endpoint, e.Fmvs, got.FMVs)
}

type testLocMusic struct {
	BgMusic   []int32
	CuesMusic []int32
	FmvsMusic []int32
	BossMusic []int32
}

func (e testLocMusic) Compare(test Test, _ string, got apitypes.LocBasedMusic) {
	test.T.Helper()
	songsEndpoint := apilookups.E.Songs.Endpoint

	test.checkResIDsInSlice("bg music", songsEndpoint, e.BgMusic, got.BackgroundMusic)
	test.checkResIDsInSlice("cues music", songsEndpoint, e.CuesMusic, got.Cues)
	test.checkResIDsInSlice("fmvs music", songsEndpoint, e.FmvsMusic, got.FMVs)
	test.checkResIDsInSlice("boss music", songsEndpoint, e.BossMusic, got.BossMusic)
}
