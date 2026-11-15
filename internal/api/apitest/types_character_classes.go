package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type CharacterClass struct {
	TestGeneral
	Unique
	Category            string
	Members             []int32
	DefaultAbilities    []int32
	LearnableAbilities  []int32
	DefaultOverdrives   []int32
	LearnableOverdrives []int32
	Submenus            []int32
}

func (e CharacterClass) GetTestGeneral() TestGeneral {
	return e.TestGeneral
}

func (e CharacterClass) Compare(test Test, got apitypes.CharacterClass) {
	test.T.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compare("category", e.Category, got.Category)
	test.checkResIDsInSlice("members", apilookups.E.PlayerUnits.Endpoint, e.Members, got.Members)
	test.checkResIDsInSlice("default abilities", apilookups.E.Abilities.Endpoint, e.DefaultAbilities, got.DefaultAbilities)
	test.checkResIDsInSlice("learnable abilities", apilookups.E.Abilities.Endpoint, e.LearnableAbilities, got.LearnableAbilities)
	test.checkResIDsInSlice("default overdrives", apilookups.E.Overdrives.Endpoint, e.DefaultOverdrives, got.DefaultOverdrives)
	test.checkResIDsInSlice("learnable overdrives", apilookups.E.Overdrives.Endpoint, e.LearnableOverdrives, got.LearnableOverdrives)
	test.checkResIDsInSlice("submenus", apilookups.E.Submenus.Endpoint, e.Submenus, got.Submenus)
}
