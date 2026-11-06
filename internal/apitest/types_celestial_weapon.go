package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api"
	"github.com/andreasSchauer/finalfantasyxapi/internal/apitypes"
)

type expCelestialWeapon struct {
	testGeneral
	expUnique
	equipment     int32
	autoAbilities []int32
	crest         int32
	sigil         int32
	wpnTreasure   int32
	crestTreasure int32
	sigilQuest    int32
}

func (e expCelestialWeapon) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func compareCelestialWeapons(test test, exp expCelestialWeapon, got apitypes.CelestialWeapon) {
	test.t.Helper()
	test.compareExpUnique(exp.expUnique, got.ID, got.Name)
	test.compIDAPIResource("equipment", api.E.Equipment.Endpoint, exp.equipment, got.Equipment)
	test.checkResIDsInSlice("auto-abilities", api.E.AutoAbilities.Endpoint, exp.autoAbilities, got.AutoAbilities)
	test.compIDAPIResource("crest", api.E.KeyItems.Endpoint, exp.crest, got.Crest)
	test.compIDAPIResource("sigil", api.E.KeyItems.Endpoint, exp.sigil, got.Sigil)
	test.compIDAPIResource("weapon treasure", api.E.Treasures.Endpoint, exp.wpnTreasure, got.WpnTreasure)
	test.compIDAPIResource("crest treasure", api.E.Treasures.Endpoint, exp.crestTreasure, got.CrestTreasure)
	test.compIDAPIResource("sigil quest", api.E.Quests.Endpoint, exp.sigilQuest, got.SigilQuest)
}
