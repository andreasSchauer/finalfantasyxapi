package apitesttypes

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type Property struct {
	TestGeneral
	Unique
	AutoAbilities []int32
	Monsters      []int32
}

func (e Property) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e Property) Compare(test Test, got apitypes.Property) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.CheckResIDsInSlice("auto-abilities", apilookups.E.AutoAbilities.Endpoint, e.AutoAbilities, got.AutoAbilities)
	test.CheckResIDsInSlice("monsters", apilookups.E.Monsters.Endpoint, e.Monsters, got.Monsters)
}
