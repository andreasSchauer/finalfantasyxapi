package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

func compareLocRel(test Test, exp LocRel, got apitypes.LocRel) {
	test.t.Helper()
	test.compTestStructPtrs("music", exp.music, got.Music, compareLocMusic)
	test.checkResIDsInSlice("characters", apilookups.E.Characters.Endpoint, exp.characters, got.Characters)
	test.checkResIDsInSlice("aeons", apilookups.E.Aeons.Endpoint, exp.aeons, got.Aeons)
	test.checkResIDsInSlice("shops", apilookups.E.Shops.Endpoint, exp.shops, got.Shops)
	test.checkResIDsInSlice("treasures", apilookups.E.Treasures.Endpoint, exp.treasures, got.Treasures)
	test.checkResIDsInSlice("monsters", apilookups.E.Monsters.Endpoint, exp.monsters, got.Monsters)
	test.checkResIDsInSlice("formations", apilookups.E.MonsterFormations.Endpoint, exp.formations, got.Formations)
	test.checkResIDsInSlice("quests", apilookups.E.Quests.Endpoint, exp.quests, got.Quests)
	test.checkResIDsInSlice("fmvs", apilookups.E.FMVs.Endpoint, exp.fmvs, got.FMVs)
}

func compareLocMusic(test Test, _ string, exp TestLocMusic, got apitypes.LocBasedMusic) {
	test.t.Helper()
	songsEndpoint := apilookups.E.Songs.Endpoint

	test.checkResIDsInSlice("bg music", songsEndpoint, exp.bgMusic, got.BackgroundMusic)
	test.checkResIDsInSlice("cues music", songsEndpoint, exp.cuesMusic, got.Cues)
	test.checkResIDsInSlice("fmvs music", songsEndpoint, exp.fmvsMusic, got.FMVs)
	test.checkResIDsInSlice("boss music", songsEndpoint, exp.bossMusic, got.BossMusic)
}
