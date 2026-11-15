package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type Element struct {
	TestGeneral
	Unique
	StatusProtection   *int32
	AutoAbilities      []int32
	PlayerAbilities    []int32
	OverdriveAbilities []int32
	ItemAbilities      []int32
	EnemyAbilities     []int32
	MonstersWeak       []int32
	MonstersHalved     []int32
	MonstersImmune     []int32
	MonstersAbsorb     []int32
}

func (e Element) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e Element) Compare(test Test, got apitypes.Element) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compIDAPIResourcePtrs("status protection", apilookups.E.StatusConditions.Endpoint, e.StatusProtection, got.StatusProtection)
	test.checkResIDsInSlice("auto-abilities", apilookups.E.AutoAbilities.Endpoint, e.AutoAbilities, got.AutoAbilities)
	test.checkResIDsInSlice("player abilities", apilookups.E.PlayerAbilities.Endpoint, e.PlayerAbilities, got.PlayerAbilities)
	test.checkResIDsInSlice("overdrive abilities", apilookups.E.OverdriveAbilities.Endpoint, e.OverdriveAbilities, got.OverdriveAbilities)
	test.checkResIDsInSlice("item abilities", apilookups.E.ItemAbilities.Endpoint, e.ItemAbilities, got.ItemAbilities)
	test.checkResIDsInSlice("enemy abilities", apilookups.E.EnemyAbilities.Endpoint, e.EnemyAbilities, got.EnemyAbilities)
	test.checkResIDsInSlice("monsters weak", apilookups.E.Monsters.Endpoint, e.MonstersWeak, got.MonstersWeak)
	test.checkResIDsInSlice("monsters halved", apilookups.E.Monsters.Endpoint, e.MonstersHalved, got.MonstersHalved)
	test.checkResIDsInSlice("monsters immune", apilookups.E.Monsters.Endpoint, e.MonstersImmune, got.MonstersImmune)
	test.checkResIDsInSlice("monsters absorb", apilookups.E.Monsters.Endpoint, e.MonstersAbsorb, got.MonstersAbsorb)
}
