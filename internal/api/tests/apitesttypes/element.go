package apitesttypes

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
	test.CheckResIDsInSlice("auto-abilities", apilookups.E.AutoAbilities.Endpoint, e.AutoAbilities, got.AutoAbilities)
	test.CheckResIDsInSlice("player abilities", apilookups.E.PlayerAbilities.Endpoint, e.PlayerAbilities, got.PlayerAbilities)
	test.CheckResIDsInSlice("overdrive abilities", apilookups.E.OverdriveAbilities.Endpoint, e.OverdriveAbilities, got.OverdriveAbilities)
	test.CheckResIDsInSlice("item abilities", apilookups.E.ItemAbilities.Endpoint, e.ItemAbilities, got.ItemAbilities)
	test.CheckResIDsInSlice("enemy abilities", apilookups.E.EnemyAbilities.Endpoint, e.EnemyAbilities, got.EnemyAbilities)
	test.CheckResIDsInSlice("monsters weak", apilookups.E.Monsters.Endpoint, e.MonstersWeak, got.MonstersWeak)
	test.CheckResIDsInSlice("monsters halved", apilookups.E.Monsters.Endpoint, e.MonstersHalved, got.MonstersHalved)
	test.CheckResIDsInSlice("monsters immune", apilookups.E.Monsters.Endpoint, e.MonstersImmune, got.MonstersImmune)
	test.CheckResIDsInSlice("monsters absorb", apilookups.E.Monsters.Endpoint, e.MonstersAbsorb, got.MonstersAbsorb)
}
