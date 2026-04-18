package api

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

func compareCelestialWeapons(test test, exp expCelestialWeapon, got CelestialWeapon) {
	compareExpUnique(test, exp.expUnique, got.ID, got.Name)
	compIdApiResource(test, "equipment", test.cfg.e.equipment.endpoint, exp.equipment, got.Equipment)
	checkResIDsInSlice(test, "auto-abilities", test.cfg.e.autoAbilities.endpoint, exp.autoAbilities, got.AutoAbilities)
	compIdApiResource(test, "crest", test.cfg.e.keyItems.endpoint, exp.crest, got.Crest)
	compIdApiResource(test, "sigil", test.cfg.e.keyItems.endpoint, exp.sigil, got.Sigil)
	compIdApiResource(test, "weapon treasure", test.cfg.e.treasures.endpoint, exp.wpnTreasure, got.WpnTreasure)
	compIdApiResource(test, "crest treasure", test.cfg.e.treasures.endpoint, exp.crestTreasure, got.CrestTreasure)
	compIdApiResource(test, "sigil quest", test.cfg.e.quests.endpoint, exp.sigilQuest, got.SigilQuest)
}
