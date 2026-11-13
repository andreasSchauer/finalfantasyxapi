package apitest

type LocRel struct {
	characters []int32
	aeons      []int32
	shops      []int32
	treasures  []int32
	monsters   []int32
	formations []int32
	quests     []int32
	fmvs       []int32
	music      *TestLocMusic
}

type TestLocMusic struct {
	bgMusic   []int32
	cuesMusic []int32
	fmvsMusic []int32
	bossMusic []int32
}
