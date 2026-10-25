package apiTest

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
	compTestStructPtrs(test, "music", exp.music, got.Music, compareLocMusic)
	checkResIDsInSlice(test, "characters", test.cfg.E.Characters.Endpoint, exp.characters, got.Characters)
	checkResIDsInSlice(test, "aeons", test.cfg.E.Aeons.Endpoint, exp.aeons, got.Aeons)
	checkResIDsInSlice(test, "shops", test.cfg.E.Shops.Endpoint, exp.shops, got.Shops)
	checkResIDsInSlice(test, "treasures", test.cfg.E.Treasures.Endpoint, exp.treasures, got.Treasures)
	checkResIDsInSlice(test, "monsters", test.cfg.E.Monsters.Endpoint, exp.monsters, got.Monsters)
	checkResIDsInSlice(test, "formations", test.cfg.E.MonsterFormations.Endpoint, exp.formations, got.Formations)
	checkResIDsInSlice(test, "quests", test.cfg.E.Quests.Endpoint, exp.quests, got.Quests)
	checkResIDsInSlice(test, "fmvs", test.cfg.E.FMVs.Endpoint, exp.fmvs, got.FMVs)
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

	checkResIDsInSlice(test, "bg music", songsEndpoint, exp.bgMusic, got.BackgroundMusic)
	checkResIDsInSlice(test, "cues music", songsEndpoint, exp.cuesMusic, got.Cues)
	checkResIDsInSlice(test, "fmvs music", songsEndpoint, exp.fmvsMusic, got.FMVs)
	checkResIDsInSlice(test, "boss music", songsEndpoint, exp.bossMusic, got.BossMusic)
}
