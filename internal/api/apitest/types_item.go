package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/database"
)

type Item struct {
	TestGeneral
	Unique
	untypedItem        int32
	category           database.ItemCategory
	monsters           []TestMonItemAmts
	treasures          map[int32]int32
	shops              []int32
	quests             map[int32]int32
	blitzballPrizes    map[int32]int32
	aeonLearnAbilities map[int32]int32
	autoAbilities      map[string]int32
	mixes              []int32
}

func (e Item) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

type TestMonItemAmts struct {
	index               int
	monster             int32
	stealCommon         int32
	stealRare           int32
	dropCommon          int32
	dropRare            int32
	secondaryDropCommon int32
	secondaryDropRare   int32
	bribe               int32
	other               int32
}

func (t TestMonItemAmts) GetIndex() int {
	return t.index
}
