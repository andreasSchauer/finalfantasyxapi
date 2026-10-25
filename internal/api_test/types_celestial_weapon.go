package apiTest

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
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIdApiResource(test, "equipment", test.cfg.E.Equipment.Endpoint, exp.equipment, got.Equipment)
	checkResIDsInSlice(test, "auto-abilities", test.cfg.E.AutoAbilities.Endpoint, exp.autoAbilities, got.AutoAbilities)
	compIdApiResource(test, "crest", test.cfg.E.KeyItems.Endpoint, exp.crest, got.Crest)
	compIdApiResource(test, "sigil", test.cfg.E.KeyItems.Endpoint, exp.sigil, got.Sigil)
	compIdApiResource(test, "weapon treasure", test.cfg.E.Treasures.Endpoint, exp.wpnTreasure, got.WpnTreasure)
	compIdApiResource(test, "crest treasure", test.cfg.E.Treasures.Endpoint, exp.crestTreasure, got.CrestTreasure)
	compIdApiResource(test, "sigil quest", test.cfg.E.Quests.Endpoint, exp.sigilQuest, got.SigilQuest)
}
