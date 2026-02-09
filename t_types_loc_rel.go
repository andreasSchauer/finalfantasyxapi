package main


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