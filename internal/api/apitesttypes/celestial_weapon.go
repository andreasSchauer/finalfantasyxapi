package apitesttypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type CelestialWeapon struct {
	TestGeneral
	Unique
	Equipment     int32
	AutoAbilities []int32
	Crest         int32
	Sigil         int32
	WpnTreasure   int32
	CrestTreasure int32
	SigilQuest    int32
}

func (e CelestialWeapon) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e CelestialWeapon) Compare(test Test, got apitypes.CelestialWeapon) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compIDAPIResource("equipment", apilookups.E.Equipment.Endpoint, e.Equipment, got.Equipment)
	test.CheckResIDsInSlice("auto-abilities", apilookups.E.AutoAbilities.Endpoint, e.AutoAbilities, got.AutoAbilities)
	test.compIDAPIResource("crest", apilookups.E.KeyItems.Endpoint, e.Crest, got.Crest)
	test.compIDAPIResource("sigil", apilookups.E.KeyItems.Endpoint, e.Sigil, got.Sigil)
	test.compIDAPIResource("weapon treasure", apilookups.E.Treasures.Endpoint, e.WpnTreasure, got.WpnTreasure)
	test.compIDAPIResource("crest treasure", apilookups.E.Treasures.Endpoint, e.CrestTreasure, got.CrestTreasure)
	test.compIDAPIResource("sigil quest", apilookups.E.Quests.Endpoint, e.SigilQuest, got.SigilQuest)
}
