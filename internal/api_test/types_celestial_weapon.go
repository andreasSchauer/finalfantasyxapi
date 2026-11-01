package apitest

import "github.com/andreasSchauer/finalfantasyxapi/internal/api"

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

func compareCelestialWeapons(test test, exp expCelestialWeapon, got api.CelestialWeapon) {
	test.t.Helper()
	test.compareExpUnique(exp.expUnique, got.ID, got.Name)
	test.compIDAPIResource("equipment", test.cfg.E.Equipment.Endpoint, exp.equipment, got.Equipment)
	test.checkResIDsInSlice("auto-abilities", test.cfg.E.AutoAbilities.Endpoint, exp.autoAbilities, got.AutoAbilities)
	test.compIDAPIResource("crest", test.cfg.E.KeyItems.Endpoint, exp.crest, got.Crest)
	test.compIDAPIResource("sigil", test.cfg.E.KeyItems.Endpoint, exp.sigil, got.Sigil)
	test.compIDAPIResource("weapon treasure", test.cfg.E.Treasures.Endpoint, exp.wpnTreasure, got.WpnTreasure)
	test.compIDAPIResource("crest treasure", test.cfg.E.Treasures.Endpoint, exp.crestTreasure, got.CrestTreasure)
	test.compIDAPIResource("sigil quest", test.cfg.E.Quests.Endpoint, exp.sigilQuest, got.SigilQuest)
}
