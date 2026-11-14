package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type CelestialWeapon struct {
	testGeneral
	Unique
	equipment     int32
	autoAbilities []int32
	crest         int32
	sigil         int32
	wpnTreasure   int32
	crestTreasure int32
	sigilQuest    int32
}

func (e CelestialWeapon) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func (e CelestialWeapon) Compare(test test, got apitypes.CelestialWeapon) {
	test.t.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compIDAPIResource("equipment", apilookups.E.Equipment.Endpoint, e.equipment, got.Equipment)
	test.checkResIDsInSlice("auto-abilities", apilookups.E.AutoAbilities.Endpoint, e.autoAbilities, got.AutoAbilities)
	test.compIDAPIResource("crest", apilookups.E.KeyItems.Endpoint, e.crest, got.Crest)
	test.compIDAPIResource("sigil", apilookups.E.KeyItems.Endpoint, e.sigil, got.Sigil)
	test.compIDAPIResource("weapon treasure", apilookups.E.Treasures.Endpoint, e.wpnTreasure, got.WpnTreasure)
	test.compIDAPIResource("crest treasure", apilookups.E.Treasures.Endpoint, e.crestTreasure, got.CrestTreasure)
	test.compIDAPIResource("sigil quest", apilookups.E.Quests.Endpoint, e.sigilQuest, got.SigilQuest)
}
