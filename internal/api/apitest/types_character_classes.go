package apitest

import (
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/apilookups"
	"github.com/andreasSchauer/finalfantasyxapi/internal/api/endpoints/apitypes"
)

type CharacterClass struct {
	testGeneral
	Unique
	category            string
	members             []int32
	defaultAbilities    []int32
	learnableAbilities  []int32
	defaultOverdrives   []int32
	learnableOverdrives []int32
	submenus            []int32
}

func (e CharacterClass) GetTestGeneral() testGeneral {
	return e.testGeneral
}

func (e CharacterClass) Compare(test test, got apitypes.CharacterClass) {
	test.t.Helper()
	test.compareExpUnique(e.Unique, got.ID, got.Name)
	test.compare("category", e.category, got.Category)
	test.checkResIDsInSlice("members", apilookups.E.PlayerUnits.Endpoint, e.members, got.Members)
	test.checkResIDsInSlice("default abilities", apilookups.E.Abilities.Endpoint, e.defaultAbilities, got.DefaultAbilities)
	test.checkResIDsInSlice("learnable abilities", apilookups.E.Abilities.Endpoint, e.learnableAbilities, got.LearnableAbilities)
	test.checkResIDsInSlice("default overdrives", apilookups.E.Overdrives.Endpoint, e.defaultOverdrives, got.DefaultOverdrives)
	test.checkResIDsInSlice("learnable overdrives", apilookups.E.Overdrives.Endpoint, e.learnableOverdrives, got.LearnableOverdrives)
	test.checkResIDsInSlice("submenus", apilookups.E.Submenus.Endpoint, e.submenus, got.Submenus)
}
