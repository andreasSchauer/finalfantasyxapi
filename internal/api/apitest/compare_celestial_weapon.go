package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

func compareCelestialWeapons(test Test, exp CelestialWeapon, got apitypes.CelestialWeapon) {
	test.t.Helper()
	test.compareExpUnique(exp.Unique, got.ID, got.Name)
	test.compIDAPIResource("equipment", apilookups.E.Equipment.Endpoint, exp.equipment, got.Equipment)
	test.checkResIDsInSlice("auto-abilities", apilookups.E.AutoAbilities.Endpoint, exp.autoAbilities, got.AutoAbilities)
	test.compIDAPIResource("crest", apilookups.E.KeyItems.Endpoint, exp.crest, got.Crest)
	test.compIDAPIResource("sigil", apilookups.E.KeyItems.Endpoint, exp.sigil, got.Sigil)
	test.compIDAPIResource("weapon treasure", apilookups.E.Treasures.Endpoint, exp.wpnTreasure, got.WpnTreasure)
	test.compIDAPIResource("crest treasure", apilookups.E.Treasures.Endpoint, exp.crestTreasure, got.CrestTreasure)
	test.compIDAPIResource("sigil quest", apilookups.E.Quests.Endpoint, exp.sigilQuest, got.SigilQuest)
}
