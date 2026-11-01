package apitest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api"

type expLocRel struct {
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

func compareLocRel(test test, exp expLocRel, got api.LocRel) {
	test.t.Helper()
	test.compTestStructPtrs("music", exp.music, got.Music, compareLocMusic)
	test.checkResIDsInSlice("characters", test.cfg.E.Characters.Endpoint, exp.characters, got.Characters)
	test.checkResIDsInSlice("aeons", test.cfg.E.Aeons.Endpoint, exp.aeons, got.Aeons)
	test.checkResIDsInSlice("shops", test.cfg.E.Shops.Endpoint, exp.shops, got.Shops)
	test.checkResIDsInSlice("treasures", test.cfg.E.Treasures.Endpoint, exp.treasures, got.Treasures)
	test.checkResIDsInSlice("monsters", test.cfg.E.Monsters.Endpoint, exp.monsters, got.Monsters)
	test.checkResIDsInSlice("formations", test.cfg.E.MonsterFormations.Endpoint, exp.formations, got.Formations)
	test.checkResIDsInSlice("quests", test.cfg.E.Quests.Endpoint, exp.quests, got.Quests)
	test.checkResIDsInSlice("fmvs", test.cfg.E.FMVs.Endpoint, exp.fmvs, got.FMVs)
}

type testLocMusic struct {
	bgMusic   []int32
	cuesMusic []int32
	fmvsMusic []int32
	bossMusic []int32
}

func compareLocMusic(test test, _ string, exp testLocMusic, got api.LocBasedMusic) {
	test.t.Helper()
	songsEndpoint := test.cfg.E.Songs.Endpoint

	test.checkResIDsInSlice("bg music", songsEndpoint, exp.bgMusic, got.BackgroundMusic)
	test.checkResIDsInSlice("cues music", songsEndpoint, exp.cuesMusic, got.Cues)
	test.checkResIDsInSlice("fmvs music", songsEndpoint, exp.fmvsMusic, got.FMVs)
	test.checkResIDsInSlice("boss music", songsEndpoint, exp.bossMusic, got.BossMusic)
}
