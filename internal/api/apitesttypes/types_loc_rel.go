package apitesttypes

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
	Music      *TestLocMusic
}

func (e LocRel) Compare(test Test, got apitypes.LocRel) {
	test.T.Helper()
	test.compTestStructPtrs("music", e.Music, got.Music)
	test.CheckResIDsInSlice("characters", apilookups.E.Characters.Endpoint, e.Characters, got.Characters)
	test.CheckResIDsInSlice("aeons", apilookups.E.Aeons.Endpoint, e.Aeons, got.Aeons)
	test.CheckResIDsInSlice("shops", apilookups.E.Shops.Endpoint, e.Shops, got.Shops)
	test.CheckResIDsInSlice("treasures", apilookups.E.Treasures.Endpoint, e.Treasures, got.Treasures)
	test.CheckResIDsInSlice("monsters", apilookups.E.Monsters.Endpoint, e.Monsters, got.Monsters)
	test.CheckResIDsInSlice("formations", apilookups.E.MonsterFormations.Endpoint, e.Formations, got.Formations)
	test.CheckResIDsInSlice("quests", apilookups.E.Quests.Endpoint, e.Quests, got.Quests)
	test.CheckResIDsInSlice("fmvs", apilookups.E.FMVs.Endpoint, e.Fmvs, got.FMVs)
}

type TestLocMusic struct {
	BgMusic   []int32
	CuesMusic []int32
	FmvsMusic []int32
	BossMusic []int32
}

func (e TestLocMusic) Compare(test Test, _ string, got apitypes.LocBasedMusic) {
	test.T.Helper()
	songsEndpoint := apilookups.E.Songs.Endpoint

	test.CheckResIDsInSlice("bg music", songsEndpoint, e.BgMusic, got.BackgroundMusic)
	test.CheckResIDsInSlice("cues music", songsEndpoint, e.CuesMusic, got.Cues)
	test.CheckResIDsInSlice("fmvs music", songsEndpoint, e.FmvsMusic, got.FMVs)
	test.CheckResIDsInSlice("boss music", songsEndpoint, e.BossMusic, got.BossMusic)
}
